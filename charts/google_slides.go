package charts

import (
	"fmt"
	"strings"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/v2/charts/wchart"
	"github.com/grokify/gocharts/v2/charts/wchart/sts2wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/gocharts/v2/data/timeseries/interval"
	"github.com/grokify/gogoogle/slidesutil/v1"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/math/ratio"
	"github.com/grokify/mogo/net/urlutil"
	"github.com/grokify/mogo/strconv/strconvutil"
	"github.com/grokify/mogo/time/timeutil"
)

const DefaultXAxisTimeFormat = "Jan '06"

type KpiSlideOpts struct {
	SlideType         string
	KpiID             uint64
	KpiTypeAbbr       string
	ImageBaseURL      string
	ImageHeight       uint64
	ImageWidth        uint64
	ImageRatio        float64
	Title             string
	Reference         string
	Verbose           bool
	ValueToString     func(int64) string
	XAxisTimeToString func(time.Time) string
	SlideBuildExec    bool
}

func KpiTypeAbbrIsDollars(abbr string) bool {
	abbr = strings.ToUpper(strings.TrimSpace(abbr))
	if abbr == "MRR" || abbr == "ARR" {
		return true
	}
	return false
}

func KpiSlideOptsDefaultify(opts KpiSlideOpts) KpiSlideOpts {
	if opts.ValueToString == nil {
		if KpiTypeAbbrIsDollars(opts.KpiTypeAbbr) {
			opts.ValueToString = func(val int64) string {
				return "$" + strconvutil.Commify(val)
			}
		} else {
			opts.ValueToString = func(val int64) string {
				return strconvutil.Commify(val)
			}
		}
	}
	return opts
}

func KpiSlideOptsSize2Col(opts KpiSlideOpts) KpiSlideOpts {
	opts.ImageRatio = ratio.RatioAcademy
	if opts.ImageHeight == 0 && opts.ImageWidth == 0 {
		opts.ImageHeight = 600
	}
	return opts
}

func CreateKPISlide(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, opts KpiSlideOpts) (timeseries.TimeSeries, error) {
	ds, err := GetKpiAsDataSeries(skClient, opts.KpiID, timeutil.TimeZeroRFC3339(), time.Now())
	if err != nil {
		return ds, err
	}
	if opts.Verbose {
		err := fmtutil.PrintJSON(ds)
		if err != nil {
			return ds, err
		}
	}
	if ds.Interval == timeutil.IntervalMonth {
		itemLast, err := ds.Last()
		if err == nil {
			itemLastMonthStart := timeutil.NewTimeMore(itemLast.Time, 0).MonthStart()
			nowMonthStart := timeutil.NewTimeMore(time.Now(), 0).MonthStart()
			if itemLastMonthStart.Equal(nowMonthStart) {
				_, err := ds.Pop()
				if err != nil {
					return ds, err
				}
			}
		}
	} else if ds.Interval == timeutil.IntervalQuarter {
		itemLast, err := ds.Last()
		if err == nil {
			itemLastQtrStart := timeutil.NewTimeMore(itemLast.Time, 0).QuarterStart()
			nowQtrStart := timeutil.NewTimeMore(time.Now(), 0).QuarterStart()
			if itemLastQtrStart.Equal(nowQtrStart) {
				_, err := ds.Pop()
				if err != nil {
					return ds, err
				}
			}
		}
	}

	if len(opts.Title) > 0 {
		ds.SeriesName = opts.Title
	}
	if opts.XAxisTimeToString == nil {
		opts.XAxisTimeToString = func(dt time.Time) string {
			return dt.Format(DefaultXAxisTimeFormat)
		}
	}
	graph, err := sts2wchart.TimeSeriesToLineChart(ds, &sts2wchart.LineChartOpts{
		TitleSuffixCurrentValue:     true,
		TitleSuffixCurrentValueFunc: opts.ValueToString,
		TitleSuffixCurrentDateFunc: func(dt time.Time) string {
			if ds.Interval == timeutil.IntervalQuarter {
				lastQuarter, err := ds.Last()
				if err != nil {
					return ""
				}
				return timeutil.FormatQuarterYYYYQ(lastQuarter.Time)
			}
			monthAgo := timeutil.NewTimeMore(dt, 0).MonthStart()
			return monthAgo.Format("Jan '06")
		},
		Height:           opts.ImageHeight,
		Width:            opts.ImageWidth,
		AspectRatio:      opts.ImageRatio,
		Legend:           true,
		RegressionDegree: 1,
		Interval:         ds.Interval,
		QAgoAnnotation:   true,
		YAgoAnnotation:   true,
		AgoAnnotationPct: true,
		YAxisLeft:        true,
		XAxisTickFunc:    opts.XAxisTimeToString,
	})
	if err != nil {
		return ds, err
	}

	localChartFilename := fmt.Sprintf("_output_line_%d.png", opts.KpiID)
	err = wchart.WritePNGFile(localChartFilename, graph)
	if err != nil {
		return ds, err
	}
	fmt.Printf("WROTE [%s]\n", localChartFilename)

	if pc != nil && opts.SlideBuildExec {
		opts.ImageBaseURL = strings.TrimSpace(opts.ImageBaseURL)
		if len(opts.ImageBaseURL) > 0 {
			imageURL := urlutil.JoinAbsolute(opts.ImageBaseURL, localChartFilename)

			xoxString, err := getXoxString(ds, opts.KpiID, opts.KpiTypeAbbr, opts.Reference, opts.ValueToString, opts.Verbose)
			if err != nil {
				return ds, err
			}

			err = pc.CreateSlideImageSidebarRight(ds.SeriesName, "", imageURL, xoxString)
			if err != nil {
				return ds, err
			}
		}
	}
	return ds, nil
}

func getXoxString(ds timeseries.TimeSeries, kpiID uint64, kpiTypeAbbr, sourceString string, fmtValue func(int64) string, verbose bool) (string, error) {
	xoxString := ""
	xox, err := interval.NewXoXTimeSeries(ds)
	if err != nil {
		return "", err
	}
	xoxLast := xox.Last()

	xoxLines := []string{
		fmt.Sprintf("%s: %s\n", kpiTypeAbbr, fmtValue(xoxLast.Value)),
		fmt.Sprintf("MoM: %.1f%%", xoxLast.MoM),
		fmt.Sprintf("%s: %s\n", kpiTypeAbbr, fmtValue(xoxLast.MMAgoValue)),
		fmt.Sprintf("QoQ: %.1f%%", xoxLast.QoQ),
		fmt.Sprintf("%s: %s\n", kpiTypeAbbr, fmtValue(xoxLast.MQAgoValue)),
		fmt.Sprintf("YoY: %.1f%%", xoxLast.YoY),
		fmt.Sprintf("%s: %s\n", kpiTypeAbbr, fmtValue(xoxLast.MYAgoValue))}
	if len(strings.TrimSpace(sourceString)) > 0 {
		xoxLines = append(xoxLines, sourceString)
	}
	xoxString = strings.Join(xoxLines, "\n")
	if verbose {
		fmtutil.MustPrintJSON(xox)
		fmtutil.MustPrintJSON(xoxLast)
		fmt.Println(xoxString)
	}
	return xoxString, nil
}

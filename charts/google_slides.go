package charts

import (
	"fmt"
	"strings"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/charts/wchart"
	"github.com/grokify/gocharts/charts/wchart/sts2wchart"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/gotilla/fmt/fmtutil"
	"github.com/grokify/gotilla/net/urlutil"
	"github.com/grokify/gotilla/time/month"
	"github.com/grokify/gotilla/time/timeutil"
)

type KpiSlideOpts struct {
	SlideType     string
	KpiID         uint64
	KpiTypeAbbr   string
	ImageBaseURL  string
	Title         string
	Reference     string
	Verbose       bool
	ValueToString func(int64) string
}

// func CreateKPISlide(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, kpiID uint64, imageServerURL string, sourceString string, verbose bool) error {

func CreateKPISlide(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, opts KpiSlideOpts) error {
	ds, err := GetKpiAsDataSeries(skClient, opts.KpiID, timeutil.TimeZeroRFC3339(), time.Now())
	if err != nil {
		return err
	}
	if ds.Interval == timeutil.Month {
		itemLast, err := ds.Last()
		if err == nil {
			itemLastMonthStart := timeutil.MonthStart(itemLast.Time)
			nowMonthStart := timeutil.MonthStart(time.Now())
			if itemLastMonthStart.Equal(nowMonthStart) {
				ds.Pop()
			}
		}
	} else if ds.Interval == timeutil.Quarter {
		itemLast, err := ds.Last()
		if err == nil {
			itemLastQtrStart := timeutil.QuarterStart(itemLast.Time)
			nowQtrStart := timeutil.QuarterStart(time.Now())
			if itemLastQtrStart.Equal(nowQtrStart) {
				ds.Pop()
			}
		}
	}

	if len(opts.Title) > 0 {
		ds.SeriesName = opts.Title
	}
	graph := sts2wchart.DataSeriesMonthToLineChart(ds, sts2wchart.LineChartOpts{
		TitleSuffixCurrentValue:     true,
		TitleSuffixCurrentValueFunc: opts.ValueToString,
		TitleSuffixCurrentDateFunc: func(dt time.Time) string {
			if ds.Interval == timeutil.Quarter {
				lastQuarter, err := ds.Last()
				if err != nil {
					return ""
				}
				return timeutil.FormatQuarterYYYYQ(lastQuarter.Time)
			}
			monthAgo := month.MonthBegin(dt, 0)
			return monthAgo.Format("Jan '06")
		},
		Legend:           true,
		RegressionDegree: 1,
		Interval:         ds.Interval,
		QAgoAnnotation:   true,
		YAgoAnnotation:   true,
		AgoAnnotationPct: true})

	localChartFilename := fmt.Sprintf("_output_line_%d.png", opts.KpiID)
	err = wchart.WritePNG(localChartFilename, graph)
	if err != nil {
		return err
	}

	if pc != nil {
		opts.ImageBaseURL = strings.TrimSpace(opts.ImageBaseURL)
		if len(opts.ImageBaseURL) > 0 {
			imageURL := urlutil.JoinAbsolute(opts.ImageBaseURL, localChartFilename)

			xoxString, err := getXoxString(ds, opts.KpiID, opts.KpiTypeAbbr, opts.Reference, opts.ValueToString, opts.Verbose)
			if err != nil {
				return err
			}

			err = pc.CreateSlideImageSidebarRight(ds.SeriesName, "", imageURL, xoxString)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func getXoxString(ds statictimeseries.DataSeries, kpiID uint64, kpiTypeAbbr, sourceString string, fmtValue func(int64) string, verbose bool) (string, error) {
	xoxString := ""
	xox, err := statictimeseries.NewXoXDataSeries(ds)
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
		fmtutil.PrintJSON(xox)
		fmtutil.PrintJSON(xoxLast)
		fmt.Println(xoxString)
	}
	return xoxString, nil
}

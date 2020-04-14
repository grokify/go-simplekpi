package sk2gocharts

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
	"github.com/grokify/gotilla/strconv/strconvutil"
	"github.com/grokify/gotilla/time/month"
	"github.com/grokify/gotilla/time/timeutil"
)

func CreateKPISlide(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, kpiID uint64, imageServerURL string, verbose bool) error {
	ds, err := GetKpiAsDataSeries(skClient, kpiID, timeutil.TimeZeroRFC3339(), time.Now())
	if err != nil {
		return err
	}

	ds.Pop()
	graph := sts2wchart.DataSeriesMonthToLineChart(ds, sts2wchart.LineChartMonthOpts{
		TitleSuffixCurrentValue: true,
		TitleSuffixCurrentDateFunc: func(dt time.Time) string {
			monthAgo := month.MonthBegin(dt, 0)
			return monthAgo.Format("Jan '06")
		},
		RegressionDegree: 3,
		QAgoAnnotation:   true,
		YAgoAnnotation:   true,
		AgoAnnotationPct: true})
	localChartFilename := fmt.Sprintf("_output_line_%d.png", kpiID)
	err = wchart.WritePNG(localChartFilename, graph)
	if err != nil {
		return err
	}

	xoxString := ""
	if 1 == 1 {
		xox, err := statictimeseries.NewXoXDSMonth(ds)
		if err != nil {
			return err
		}
		if verbose {
			fmtutil.PrintJSON(xox)
		}
		xoxLast := xox.Last()
		if verbose {
			fmtutil.PrintJSON(xoxLast)
		}
		xoxLines := []string{}
		xoxLines = append(xoxLines, fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.Value)))
		xoxLines = append(xoxLines, fmt.Sprintf("MoM: %.1f%%", xoxLast.MoM))
		xoxLines = append(xoxLines, fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MMAgoValue)))
		xoxLines = append(xoxLines, fmt.Sprintf("QoQ: %.1f%%", xoxLast.QoQ))
		xoxLines = append(xoxLines, fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MQAgoValue)))
		xoxLines = append(xoxLines, fmt.Sprintf("YoY: %.1f%%", xoxLast.YoY))
		xoxLines = append(xoxLines, fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MYAgoValue)))
		xoxLines = append(xoxLines, fmt.Sprintf("Source: AGW Logs via\nMetabase &\nSimpleKPI #%d", kpiID))
		xoxString = strings.Join(xoxLines, "\n")
		fmt.Println(xoxString)
	}

	if pc != nil {
		imageServerURL = strings.TrimSpace(imageServerURL)
		if len(imageServerURL) > 0 {
			imageURL := urlutil.JoinAbsolute(imageServerURL, localChartFilename)
			err = pc.CreateSlideImageSidebarRight(ds.SeriesName, "", imageURL, xoxString)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

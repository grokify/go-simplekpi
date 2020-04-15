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
	"github.com/grokify/gotilla/strconv/strconvutil"
	"github.com/grokify/gotilla/time/month"
	"github.com/grokify/gotilla/time/timeutil"
)

func CreateKPISlide(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, kpiID uint64, imageServerURL string, sourceString string, verbose bool) error {
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
		Legend:           true,
		RegressionDegree: 3,
		QAgoAnnotation:   true,
		YAgoAnnotation:   true,
		AgoAnnotationPct: true})

	localChartFilename := fmt.Sprintf("_output_line_%d.png", kpiID)
	err = wchart.WritePNG(localChartFilename, graph)
	if err != nil {
		return err
	}

	if pc != nil {
		imageServerURL = strings.TrimSpace(imageServerURL)
		if len(imageServerURL) > 0 {
			imageURL := urlutil.JoinAbsolute(imageServerURL, localChartFilename)

			xoxString, err := getXoxString(ds, kpiID, sourceString, verbose)
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

func getXoxString(ds statictimeseries.DataSeries, kpiID uint64, sourceString string, verbose bool) (string, error) {
	xoxString := ""
	xox, err := statictimeseries.NewXoXDSMonth(ds)
	if err != nil {
		return "", err
	}
	xoxLast := xox.Last()

	xoxLines := []string{
		fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.Value)),
		fmt.Sprintf("MoM: %.1f%%", xoxLast.MoM),
		fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MMAgoValue)),
		fmt.Sprintf("QoQ: %.1f%%", xoxLast.QoQ),
		fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MQAgoValue)),
		fmt.Sprintf("YoY: %.1f%%", xoxLast.YoY),
		fmt.Sprintf("MAU: %s\n", strconvutil.Commify(xoxLast.MYAgoValue))}
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

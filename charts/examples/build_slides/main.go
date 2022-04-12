package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/grokify/go-simplekpi/charts"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/goauth/google"
	"github.com/grokify/gocharts/v2/charts/wchart"
	"github.com/grokify/gocharts/v2/charts/wchart/sts2wchart"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
	"github.com/grokify/mogo/math/ratio"
	"github.com/grokify/mogo/strconv/strconvutil"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath  string `short:"e" long:"envpath" description:".env Filepath" required:"false"`
	Site     string `short:"s" long:"site" description:"Your site" required:"false"`
	Username string `short:"u" long:"username" description:"Your username" required:"false"`
	Password string `short:"p" long:"password" description:"Your password" required:"false"`
	Kpiid    int32  `short:"k" long:"kpiid" description:"KPI ID" required:"false"`
}

func BuildSingleImages(skAPIClient *simplekpi.APIClient, imageBaseURL string, kpis []uint64) error {
	for _, kpiID := range kpis {
		opts := charts.KpiSlideOpts{
			KpiID: kpiID,
			Title: "Active MRR",
			ValueToString: func(v int64) string {
				return "$" + strconvutil.Commify(v)
			},
			KpiTypeAbbr:  "MRR",
			ImageBaseURL: imageBaseURL,
			Reference:    fmt.Sprintf("Source: SimpleKPI #%d", kpiID),
			Verbose:      true}
		_, err := charts.CreateKPISlide(skAPIClient, nil, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	imageBaseURL := "https://7e388a1e.ngrok.io"

	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Site) == 0 {
		_, err := config.LoadDotEnv(".env", os.Getenv("ENV_PATH"), opts.EnvPath)
		if err != nil {
			log.Fatal(err)
		}
		opts.Site = os.Getenv("SIMPLEKPI_SITE")
		opts.Username = os.Getenv("SIMPLEKPI_USERNAME")
		opts.Password = os.Getenv("SIMPLEKPI_TOKEN")
	}

	skAPIClient, err := simplekpiutil.NewApiClient(opts.Site, opts.Username, opts.Password)
	if err != nil {
		log.Fatal(err)
	}

	err1 := BuildSingleImages(skAPIClient, imageBaseURL, []uint64{157})
	if err1 != nil {
		log.Fatal(err1)
	}
	panic("Z")

	if 1 == 0 {
		//t0 := timeutil.TimeZeroRFC3339()
		t0, err := time.Parse(timeutil.RFC3339FullDate, "2017-01-01")
		if err != nil {
			log.Fatal(err)
		}

		ds1, ds2, ds3, err := charts.PercentTwoKPIs(
			skAPIClient, 159, 158, t0, time.Now())
		if err != nil {
			log.Fatal(err)
		}
		ds1.SetSeriesName("MAA Platform Apps")
		ds2.SetSeriesName("MAA RC Native Apps")
		ds3.SetSeriesName("MAA % Adoption Platform vs. Native")

		fmtutil.PrintJSON(ds3)

		if 1 == 1 {
			tss := timeseries.NewTimeSeriesSet("")
			tss.Name = "Adoption"
			tss.AddSeries(ds1)
			tss.AddSeries(ds2)
			tss.AddSeries(ds3)
			xlsx := "_Adoption.xlsx"
			err = tss.WriteXLSX(xlsx,
				&timeseries.TimeSeriesSetTableOpts{
					FuncFormatTime: timeutil.FormatTimeToString("2006-01"),
				})
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("WROTE [%v]\n", xlsx)
		}

		ds1.Pop()
		ds2.Pop()
		ds3.Pop()

		opts := sts2wchart.LineChartOpts{
			RegressionDegree: 1,
			QAgoAnnotation:   false,
			YAgoAnnotation:   false,
			AgoAnnotationPct: false,
			Height:           600,
			AspectRatio:      ratio.RatioAcademy,
			Interval:         timeutil.Month}

		if 1 == 1 {
			opts.QAgoAnnotation = true
			opts.YAgoAnnotation = true
			opts.AgoAnnotationPct = true
			graph1, err := sts2wchart.TimeSeriesToLineChart(ds1, &opts)
			if err != nil {
				log.Fatal(err)
			}
			file1 := "_MAA Platform.png"
			err = wchart.WritePNG(file1, graph1)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("WROTE [%s]\n", file1)

			graph2, err := sts2wchart.TimeSeriesToLineChart(ds2, &opts)
			if err != nil {
				log.Fatal(err)
			}
			file2 := "_MAA RC Native.png"
			err = wchart.WritePNG(file2, graph2)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("WROTE [%s]\n", file2)
		}
		if 1 == 1 {
			opts.QAgoAnnotation = false
			opts.YAgoAnnotation = false
			opts.AgoAnnotationPct = false
			graph3, err := sts2wchart.TimeSeriesToLineChart(ds3, &opts)
			if err != nil {
				log.Fatal(err)
			}
			file3 := "_MAA Adoption Rate.png"
			err = wchart.WritePNG(file3, graph3)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("WROTE [%s]\n", file3)
		}

		panic("ZZ")
	}

	googHTTPClient, err := google.NewClientFileStoreWithDefaultsCliEnv("", "")
	if err != nil {
		log.Fatal(err)
	}

	pc, err := slidesutil.NewPresentationCreator(googHTTPClient)
	if err != nil {
		log.Fatal(err)
	}

	_, err = pc.Create(
		"Platform Analytics - "+time.Now().Format(time.RFC3339),
		"Platform Analytics",
		"Developer Program")
	if err != nil {
		log.Fatal(err)
	}

	kpis := []uint64{158, 159}

	for _, kpiID := range kpis {
		opts := charts.KpiSlideOpts{
			KpiID:          kpiID,
			ImageBaseURL:   imageBaseURL,
			Reference:      fmt.Sprintf("Source: Metabase &\nSimpleKPI #%d", kpiID),
			SlideBuildExec: true,
			Verbose:        true}

		opts = charts.KpiSlideOptsDefaultify(opts)
		opts = charts.KpiSlideOptsSize2Col(opts)

		_, err = charts.CreateKPISlide(skAPIClient, pc, opts)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("DONE")
}

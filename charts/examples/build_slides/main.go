package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/grokify/go-simplekpi/charts"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/googleutil/slidesutil/v1"
	"github.com/grokify/gotilla/config"
	"github.com/grokify/oauth2more/google"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath  string `short:"e" long:"envpath" description:".env Filepath" required:"false"`
	Site     string `short:"s" long:"site" description:"Your site" required:"false"`
	Username string `short:"u" long:"username" description:"Your username" required:"false"`
	Password string `short:"p" long:"password" description:"Your password" required:"false"`
	Kpiid    int32  `short:"k" long:"kpiid" description:"KPI ID" required:"false"`
}

func main() {
	imageBaseURL := "https://06afcd6b.ngrok.io/"

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

	kpis := []uint64{94, 92, 93}

	for _, kpiID := range kpis {
		err = charts.CreateKPISlide(skAPIClient, pc,
			kpiID, imageBaseURL,
			fmt.Sprintf("Source: Metabase &\nSimpleKPI #%d", kpiID),
			true)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Println("DONE")
}

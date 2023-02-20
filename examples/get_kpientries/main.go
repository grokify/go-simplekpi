package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/mogo/config"
	"github.com/grokify/mogo/fmt/fmtutil"
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
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Site) == 0 {
		_, err := config.LoadDotEnv([]string{opts.EnvPath, ".env", os.Getenv("ENV_PATH")}, 1)
		if err != nil {
			log.Fatal(err)
		}
		opts.Site = os.Getenv("SIMPLEKPI_SITE")
		opts.Username = os.Getenv("SIMPLEKPI_USERNAME")
		opts.Password = os.Getenv("SIMPLEKPI_TOKEN")
	}

	client, err := simplekpiutil.NewApiClient(opts.Site, opts.Username, opts.Password)
	if err != nil {
		log.Fatal(err)
	}

	params := &simplekpi.GetAllKPIEntriesOpts{}
	if opts.Kpiid > 0 {
		params.Kpiid = optional.NewInt32(opts.Kpiid)
	}

	kpi, resp, err := client.KPIsApi.GetKPI(
		context.Background(),
		int64(opts.Kpiid))
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode > 299 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(kpi)

	kpientries, resp, err := client.KPIEntriesApi.GetAllKPIEntries(
		context.Background(),
		"2020-01-01",
		"2020-02-01",
		params,
	)
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode > 299 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(kpientries)

	if 1 == 0 {
		test := simplekpi.KpiEntry{
			KpiId:     111,
			UserId:    222,
			EntryDate: "2020-01-01",
			Actual:    12345.0,
		}
		kpi, resp, err := client.KPIEntriesApi.AddKPIEntry(
			context.Background(),
			test,
		)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode > 299 {
			log.Fatal(resp.StatusCode)
		}
		fmt.Printf("SUCCESS [%v]\n", resp.StatusCode)
		fmtutil.PrintJSON(kpi)
	}

	if 1 == 0 {
		kpientryid := int64(333)
		resp, err := client.KPIEntriesApi.DeleteKPIEntry(
			context.Background(),
			kpientryid,
		)
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode > 299 {
			log.Fatal(resp.StatusCode)
		}
		fmt.Printf("DELETE SUCCESS [%v]\n", resp.StatusCode)
	}

	fmt.Println("DONE")
}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/simplego/config"
	"github.com/grokify/simplego/fmt/fmtutil"
	"github.com/jessevdk/go-flags"
)

type Options struct {
	EnvPath  string `short:"e" long:"envpath" description:".env Filepath" required:"false"`
	Site     string `short:"s" long:"site" description:"Your site" required:"false"`
	Username string `short:"u" long:"username" description:"Your username" required:"false"`
	Password string `short:"p" long:"password" description:"Your password" required:"false"`
}

func main() {
	opts := Options{}
	_, err := flags.Parse(&opts)
	if err != nil {
		log.Fatal(err)
	}

	if len(opts.Site) == 0 {
		err := config.LoadDotEnvSkipEmpty(opts.EnvPath, ".env", os.Getenv("ENV_PATH"))
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

	info, resp, err := client.UsersApi.GetAllUsers(context.Background())
	if err != nil {
		log.Fatal(err)
	} else if resp.StatusCode > 299 {
		log.Fatal(resp.StatusCode)
	}
	fmtutil.PrintJSON(info)
	if 1 == 0 {
		kpis, resp, err := client.KPIsApi.GetAllKPIs(context.Background())
		if err != nil {
			log.Fatal(err)
		} else if resp.StatusCode > 299 {
			log.Fatal(resp.StatusCode)
		}
		fmtutil.PrintJSON(kpis)

		for _, kpi := range kpis {
			if strings.Index(kpi.Name, "MAU") > -1 {
				fmtutil.PrintJSON(kpi)
			}
		}
	}
	fmt.Println("DONE")
}

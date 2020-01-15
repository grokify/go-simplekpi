package simplekpiutil

import (
	"fmt"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2more"
)

const baseUrlFormat string = `https://%s.simplekpi.com/api`

func NewClient(site, username, password string) (*simplekpi.APIClient, error) {
	headerVal, err := oauth2more.BasicAuthHeader(username, password)
	if err != nil {
		return nil, err
	}

	cfg := &simplekpi.Configuration{
		BasePath: fmt.Sprintf(baseUrlFormat, site),
		DefaultHeader: map[string]string{
			httputilmore.HeaderAuthorization: headerVal},
		UserAgent: "openapi-generator/1.0.0/go; simplekpi-client-go/1.0.0",
	}

	client := simplekpi.NewAPIClient(cfg)
	return client, nil
}

type Options struct {
	Site     string `short:"s" long:"site" description:"Your site" required:"true"`
	Username string `short:"u" long:"username" description:"Your username" required:"true"`
	Password string `short:"p" long:"password" description:"Your password" required:"true"`
}

func NewClientOptions(opts Options) (*simplekpi.APIClient, error) {
	return NewClient(opts.Site, opts.Username, opts.Password)
}

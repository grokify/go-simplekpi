package simplekpiutil

import (
	"fmt"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gotilla/net/httputilmore"
	"github.com/grokify/oauth2more"
)

const baseUrlFormat string = `https://%s.simplekpi.com/api`

func NewApiClient(site, username, token string) (*simplekpi.APIClient, error) {
	headerVal, err := oauth2more.BasicAuthHeader(username, token)
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

type Config struct {
	Site     string `short:"s" long:"site" description:"Your site" required:"true"`
	Username string `short:"u" long:"username" description:"Your username" required:"true"`
	Token    string `short:"t" long:"token" description:"Your token" required:"true"`
}

func NewApiClientConfig(opts Config) (*simplekpi.APIClient, error) {
	return NewApiClient(opts.Site, opts.Username, opts.Token)
}

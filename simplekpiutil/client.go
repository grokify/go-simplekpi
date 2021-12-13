package simplekpiutil

import (
	"fmt"
	"os"
	"strconv"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/goauth"
	"github.com/grokify/mogo/net/httputilmore"
	"github.com/grokify/mogo/time/timeutil"
)

const (
	baseUrlFormat = `https://%s.simplekpi.com/api`
	ApiDateFormat = timeutil.RFC3339FullDate
	ApiTimeFormat = timeutil.ISO8601NoTZ

	EnvSimplekpiSite     = "SIMPLEKPI_SITE"
	EnvSimplekpiToken    = "SIMPLEKPI_TOKEN"
	EnvSimplekpiUsername = "SIMPLEKPI_USERNAME"
	EnvSimplekpiUserID   = "SIMPLEKPI_USERID"
)

func NewApiClient(site, username, token string) (*simplekpi.APIClient, error) {
	headerVal, err := goauth.BasicAuthHeader(username, token)
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

func NewApiClientEnv() (*simplekpi.APIClient, error) {
	return NewApiClient(
		os.Getenv(EnvSimplekpiSite),
		os.Getenv(EnvSimplekpiUsername),
		os.Getenv(EnvSimplekpiToken))
}

func GetUserIDEnv() (uint, error) {
	userIDString := os.Getenv(EnvSimplekpiUserID)
	if len(userIDString) == 0 {
		return 0, fmt.Errorf("E_NO_USER_ID_ENV_VAR")
	}
	num, err := strconv.Atoi(userIDString)
	if err != nil {
		return 0, err
	}
	return uint(num), nil
}

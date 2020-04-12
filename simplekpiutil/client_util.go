package simplekpiutil

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gotilla/time/timeutil"
)

type ClientUtil struct {
	Client *simplekpi.APIClient
}

func (sku *ClientUtil) GetKPI(kpiId uint64) (simplekpi.Kpi, error) {
	kpi, resp, err := sku.Client.KPIsApi.GetKPI(
		context.Background(), int64(kpiId))
	if err != nil {
		return kpi, err
	} else if resp.StatusCode > 299 {
		return kpi, errors.New(
			fmt.Sprintf("E_SIMPLEKPI_STATUS_CODE [%v]", resp.StatusCode))
	}
	return kpi, nil
}

func (sku *ClientUtil) GetAllKPIEntries(kpiId uint64, startDate, endDate time.Time) ([]simplekpi.KpiEntry, error) {
	params := &simplekpi.GetAllKPIEntriesOpts{}
	if kpiId > 0 {
		params.Kpiid = optional.NewInt32(int32(kpiId))
	}

	kpientries, resp, err := sku.Client.KPIEntriesApi.GetAllKPIEntries(
		context.Background(),
		startDate.Format(timeutil.RFC3339FullDate),
		endDate.Format(timeutil.RFC3339FullDate),
		params)
	if err != nil {
		return []simplekpi.KpiEntry{}, err
	} else if resp.StatusCode > 299 {
		return []simplekpi.KpiEntry{}, errors.New(
			fmt.Sprintf("E_SIMPLEKPI_STATUS_CODE [%v]", resp.StatusCode))
	}
	return kpientries, nil
}

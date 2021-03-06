package simplekpiutil

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/simplego/time/timeutil"
)

type ClientUtil struct {
	APIClient *simplekpi.APIClient
}

func (sku *ClientUtil) GetKPI(kpiId uint64) (simplekpi.Kpi, error) {
	kpi, resp, err := sku.APIClient.KPIsApi.GetKPI(
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

	kpientries, resp, err := sku.APIClient.KPIEntriesApi.GetAllKPIEntries(
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

func (sku *ClientUtil) GetKPIEntriesAsDataSeries(kpiId uint64, startDate, endDate time.Time) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	kentries, err := sku.GetAllKPIEntries(kpiId, startDate, endDate)
	if err != nil {
		return ds, err
	}
	kpi, err := sku.GetKPI(kpiId)
	if err != nil {
		return ds, err
	}
	ds.SeriesName = kpi.Name
	ds.Interval = FrequencyIDToInterval(kpi.FrequencyId)
	ds, err = DataSeriesAddKPIEntries(ds, kentries...)
	if err != nil {
		return ds, err
	}
	return ds, nil
}

func DataSeriesAddKPIEntries(ds statictimeseries.DataSeries, kentries ...simplekpi.KpiEntry) (statictimeseries.DataSeries, error) {
	for _, kentry := range kentries {
		dt, err := time.Parse(ApiTimeFormat, kentry.EntryDate)
		if err != nil {
			return ds, err
		}
		ds.AddItem(statictimeseries.DataItem{
			Time:  dt,
			Value: int64(kentry.Actual)})
	}
	return ds, nil
}

func KPIEntriesToDataSeries(kentries []simplekpi.KpiEntry) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	return DataSeriesAddKPIEntries(ds, kentries...)
}

func FrequencyIDToInterval(frequencyId string) timeutil.Interval {
	frequencyId = strings.ToUpper(strings.TrimSpace(frequencyId))
	if frequencyId == "Q" {
		return timeutil.Quarter
	} else if frequencyId == "M" {
		return timeutil.Month
	}
	return timeutil.Day
}

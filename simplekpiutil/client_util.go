package simplekpiutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/v2/data/timeseries"
	"github.com/grokify/mogo/time/timeutil"
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
		return kpi, fmt.Errorf("E_SIMPLEKPI_STATUS_CODE [%v]", resp.StatusCode)
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
		return []simplekpi.KpiEntry{}, fmt.Errorf("E_SIMPLEKPI_STATUS_CODE [%v]", resp.StatusCode)
	}
	return kpientries, nil
}

func (sku *ClientUtil) GetKPIEntriesAsDataSeries(kpiId uint64, startDate, endDate time.Time) (timeseries.TimeSeries, error) {
	ts := timeseries.NewTimeSeries("KPI Entries")
	kentries, err := sku.GetAllKPIEntries(kpiId, startDate, endDate)
	if err != nil {
		return ts, err
	}
	kpi, err := sku.GetKPI(kpiId)
	if err != nil {
		return ts, err
	}
	ts.SeriesName = kpi.Name
	ts.Interval = FrequencyIDToInterval(kpi.FrequencyId)
	ts, err = DataSeriesAddKPIEntries(ts, kentries...)
	if err != nil {
		return ts, err
	}
	return ts, nil
}

func DataSeriesAddKPIEntries(ts timeseries.TimeSeries, kentries ...simplekpi.KpiEntry) (timeseries.TimeSeries, error) {
	for _, kentry := range kentries {
		dt, err := time.Parse(ApiTimeFormat, kentry.EntryDate)
		if err != nil {
			return ts, err
		}
		ts.AddItems(timeseries.TimeItem{
			Time:  dt,
			Value: int64(kentry.Actual)})
	}
	return ts, nil
}

func KPIEntriesToDataSeries(kentries []simplekpi.KpiEntry) (timeseries.TimeSeries, error) {
	ds := timeseries.NewTimeSeries("KPI Enries")
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

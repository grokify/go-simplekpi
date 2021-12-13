package charts

import (
	"encoding/json"
	"strconv"
	"strings"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/timeseries"
	"github.com/grokify/mogo/time/timeutil"
	"github.com/pkg/errors"
)

func GetKpiAsDataSeries(skApiClient *simplekpi.APIClient, kpiId uint64, startDate, endDate time.Time) (timeseries.TimeSeries, error) {
	ds := timeseries.NewTimeSeries("KPI " + strconv.Itoa(int(kpiId)))
	sku := simplekpiutil.ClientUtil{APIClient: skApiClient}
	kpi, err := sku.GetKPI(kpiId)
	if err != nil {
		return ds, err
	}
	entries, err := sku.GetAllKPIEntries(kpiId, startDate, endDate)
	if err != nil {
		return ds, err
	}
	return KpiAndEntriesToDataSeries(kpi, entries)
}

func KpiAndEntriesToDataSeries(kpi simplekpi.Kpi, entries []simplekpi.KpiEntry) (timeseries.TimeSeries, error) {
	interval := timeutil.Month
	if strings.ToUpper(strings.TrimSpace(kpi.FrequencyId)) == "Q" {
		interval = timeutil.Quarter
	} else if strings.ToUpper(strings.TrimSpace(kpi.FrequencyId)) == "M" {
		interval = timeutil.Month
	}
	return KpiEntriesToDataSeries(kpi.Name, entries, interval)
}

// KpiEntriesToDataSeries converets a slice of KpiEntry to
// `statictimeseris.DataSeries`
func KpiEntriesToDataSeries(seriesName string, kpiEntries []simplekpi.KpiEntry, interval timeutil.Interval) (timeseries.TimeSeries, error) {
	ts := timeseries.NewTimeSeries(strings.TrimSpace(seriesName))
	ts.Interval = interval
	for _, kpie := range kpiEntries {
		dataItem, err := KpiEntryToDataItem(ts.SeriesName, kpie)
		if err != nil {
			return ts, err
		}
		ts.AddItems(dataItem)
	}
	return ts, nil
}

// KpiEntryToDataItem converts a simplekpi.KpiEentry to
// a timeseries.TimeItem.
func KpiEntryToDataItem(seriesName string, entry simplekpi.KpiEntry) (timeseries.TimeItem, error) {
	entryDate := strings.TrimSpace(entry.EntryDate)
	if len(entryDate) == 0 {
		bytes, err := json.Marshal(entry)
		errMsg := "Entry_No_Time"
		if err == nil {
			errMsg += " " + string(bytes)
		}
		return timeseries.TimeItem{}, errors.New(errMsg)
	}
	dt, err := time.Parse(simplekpiutil.ApiTimeFormat, entryDate)
	if err != nil {
		return timeseries.TimeItem{}, errors.Wrap(err, "KpiEntryToDataItem")
	}
	return timeseries.TimeItem{
		SeriesName: strings.TrimSpace(seriesName),
		Time:       dt,
		Value:      int64(entry.Actual),
	}, nil
}

package charts

import (
	"encoding/json"
	"strings"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/simplego/time/timeutil"
	"github.com/pkg/errors"
)

func GetKpiAsDataSeries(skApiClient *simplekpi.APIClient, kpiId uint64, startDate, endDate time.Time) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
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

func KpiAndEntriesToDataSeries(kpi simplekpi.Kpi, entries []simplekpi.KpiEntry) (statictimeseries.DataSeries, error) {
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
func KpiEntriesToDataSeries(seriesName string, kpiEntries []simplekpi.KpiEntry, interval timeutil.Interval) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	ds.SeriesName = strings.TrimSpace(seriesName)
	ds.Interval = interval
	for _, kpie := range kpiEntries {
		dataItem, err := KpiEntryToDataItem(ds.SeriesName, kpie)
		if err != nil {
			return ds, err
		}
		ds.AddItem(dataItem)
	}
	return ds, nil
}

// KpiEntryToDataItem converts a simplekpi.KpiEentry to
// a statictimeseries.DataItem.
func KpiEntryToDataItem(seriesName string, entry simplekpi.KpiEntry) (statictimeseries.DataItem, error) {
	entryDate := strings.TrimSpace(entry.EntryDate)
	if len(entryDate) == 0 {
		bytes, err := json.Marshal(entry)
		errMsg := "Entry_No_Time"
		if err == nil {
			errMsg += " " + string(bytes)
		}
		return statictimeseries.DataItem{}, errors.New(errMsg)
	}
	dt, err := time.Parse(simplekpiutil.ApiTimeFormat, entryDate)
	if err != nil {
		return statictimeseries.DataItem{}, errors.Wrap(err, "KpiEntryToDataItem")
	}
	return statictimeseries.DataItem{
		SeriesName: strings.TrimSpace(seriesName),
		Time:       dt,
		Value:      int64(entry.Actual),
	}, nil
}

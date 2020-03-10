package simplekpiutil

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/data/statictimeseries"
	"github.com/grokify/gotilla/time/timeutil"
)

//func GetKpiIdAsSTS(client *simplekpi.APIClient, dsi metabase2simplekpi.DatasetInfo) error {
func GetKpiIdAsSTS(client *simplekpi.APIClient, kpiId int32, startDate, endDate time.Time) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	kpi, resp, err := client.KPIsApi.GetKPI(
		context.Background(),
		int64(kpiId))
	if err != nil {
		return ds, err
	} else if resp.StatusCode >= 300 {
		return ds, fmt.Errorf("E_SIMPLEKPI_STATUS [%v]", resp.StatusCode)
	}
	kpiName := strings.TrimSpace(kpi.Name)

	kpiEntries, resp, err := client.KPIEntriesApi.GetAllKPIEntries(
		context.Background(),
		startDate.Format(timeutil.RFC3339FullDate),
		endDate.Format(timeutil.RFC3339FullDate),
		&simplekpi.GetAllKPIEntriesOpts{
			Kpiid: optional.NewInt32(kpiId),
		},
	)

	if err != nil {
		return ds, err
	} else if resp.StatusCode >= 300 {
		return ds, fmt.Errorf("E_SIMPLEKPI_STATUS [%v]", resp.StatusCode)
	}
	return KpiEntriesToDataSeries(kpiName, kpiEntries)
}

// KpiEntriesToDataSeries converets a slice of KpiEntry to
// `statictimeseris.DataSeries`
func KpiEntriesToDataSeries(seriesName string, kpiEntries []simplekpi.KpiEntry) (statictimeseries.DataSeries, error) {
	ds := statictimeseries.NewDataSeries()
	ds.SeriesName = seriesName
	for _, kpie := range kpiEntries {
		entryDate, err := time.Parse(ApiTimeFormat, kpie.EntryDate)
		if err != nil {
			return ds, err
		}
		ds.AddItem(statictimeseries.DataItem{
			SeriesName: seriesName,
			Time:       entryDate,
			Value:      int64(kpie.Actual)})
	}
	return ds, nil
}

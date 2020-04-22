package simplekpiutil

import (
	"fmt"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/data/statictimeseries"
)

type KpiEntriesClient struct {
	ApiClient *simplekpi.APIClient
	UserID    int64
}

func NewKpiEntriesClientEnv() (KpiEntriesClient, error) {
	kec := KpiEntriesClient{}
	apiClient, err := NewApiClientEnv()
	if err != nil {
		return kec, err
	}
	kec.ApiClient = apiClient
	userID, err := GetUserIDEnv()
	if err != nil {
		return kec, err
	}
	kec.UserID = int64(userID)
	return kec, nil
}

func (kec *KpiEntriesClient) UpsertKpiEntriesDataSeries(kpiID int64,
	ds statictimeseries.DataSeries) (KpiEntryQueries, []KpiEntryResponse, error) {
	return UpsertKpiEntriesStaticTimeSeries(kec.ApiClient, kec.UserID, kpiID, ds)
}

func (kec *KpiEntriesClient) UpsertKpiEntriesDataSeriesSetSimple(name2KpiID map[string]int64,
	ds3 statictimeseries.DataSeriesSetSimple) ([]KpiEntryQueries, [][]KpiEntryResponse, error) {

	queries := []KpiEntryQueries{}
	responses := [][]KpiEntryResponse{}
	for seriesName, kpiID := range name2KpiID {
		ds, ok := ds3.Series[seriesName]
		if !ok {
			return queries, responses, fmt.Errorf("E_SERIES_NOT_FOUND SeriesName [%s] UpsertKpiEntriesDataSeriesSetSimple", seriesName)
		}
		qry, res, err := kec.UpsertKpiEntriesDataSeries(kpiID, ds)
		if err != nil {
			return queries, responses, fmt.Errorf("E_SERIES_UPSERT_FAILED SeriesName [%v,%s] UpsertKpiEntriesDataSeriesSetSimple", kpiID, seriesName)
		}
		queries = append(queries, qry)
		responses = append(responses, res)
	}

	return queries, responses, nil
}

package simplekpiutil

import (
	"fmt"
	"log"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/v2/data/timeseries"
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
	ds timeseries.TimeSeries) (KpiEntryQueries, []KpiEntryResponse, error) {
	return UpsertKpiEntriesStaticTimeSeries(kec.ApiClient, kec.UserID, kpiID, ds)
}

func (kec *KpiEntriesClient) UpsertKpiEntriesDataSeriesSetSimple(name2KpiID map[string]int64,
	ds3 timeseries.TimeSeriesSet) ([]KpiEntryQueries, [][]KpiEntryResponse, error) {

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

func WriteKpisXlsx(apiClient *simplekpi.APIClient, filename string, kpiIDs []uint64, dateStart, dateEnd time.Time) error {
	cu := ClientUtil{APIClient: apiClient}
	data := []timeseries.TimeSeries{}
	for _, kpiID := range kpiIDs {
		ds, err := cu.GetKPIEntriesAsDataSeries(kpiID, dateStart, dateEnd)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, ds)
	}
	return timeseries.TimeSeriesSliceWriteXLSX(filename, data)
}

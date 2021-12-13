package simplekpiutil

import (
	"context"
	"fmt"

	"github.com/antihax/optional"
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/data/timeseries"
	"github.com/grokify/mogo/time/timeutil"
)

// KpiEntryQueries represents a set of API KPI Entry queries
// to be performed.
type KpiEntryQueries struct {
	Create []simplekpi.KpiEntry
	Update []simplekpi.KpiEntry
}

// RunQueries is a low level function that executes a set of queries.
func RunQueries(client *simplekpi.APIClient, qrys KpiEntryQueries) []KpiEntryResponse {
	resps := []KpiEntryResponse{}
	if len(qrys.Create) > 0 {
		resCreate := CreateKpiEntries(client, qrys.Create)
		resps = append(resps, resCreate...)
	}
	if len(qrys.Update) > 0 {
		resUpdate := UpdateKpiEntries(client, qrys.Update)
		resps = append(resps, resUpdate...)
	}
	return resps
}

// UpsertKpiEntriesStaticTimeSeries is a high level function that
// takse requests and executes them.
func UpsertKpiEntriesStaticTimeSeries(
	client *simplekpi.APIClient,
	userID, kpiID int64,
	ds timeseries.TimeSeries) (KpiEntryQueries, []KpiEntryResponse, error) {

	qrys := KpiEntryQueries{}
	resps := []KpiEntryResponse{}

	if len(ds.ItemMap) == 0 {
		return qrys, resps, nil
	}
	minTime, maxTime := ds.MinMaxTimes()
	return UpsertKpiEntriesStaticTimeSeriesTimes(
		client, userID, kpiID,
		minTime.Format(timeutil.RFC3339FullDate),
		maxTime.Format(timeutil.RFC3339FullDate),
		ds)
}

// UpsertKpiEntriesStaticTimeSeriesTimes is a high level function that
// takse requests and executes them.
func UpsertKpiEntriesStaticTimeSeriesTimes(
	client *simplekpi.APIClient,
	userID, kpiID int64,
	oldDateFrom, oldDateTo string,
	ds timeseries.TimeSeries) (KpiEntryQueries, []KpiEntryResponse, error) {

	qrys := KpiEntryQueries{}
	resps := []KpiEntryResponse{}

	opts := &simplekpi.GetAllKPIEntriesOpts{
		Kpiid: optional.NewInt32(int32(kpiID)),
		Rows:  optional.NewInt32(500)}
	entries, resp, err := client.KPIEntriesApi.GetAllKPIEntries(
		context.Background(), oldDateFrom, oldDateTo, opts,
	)
	if err != nil {
		return qrys, resps, err
	} else if resp.StatusCode >= 300 {
		return qrys, resps, fmt.Errorf("E_SIMPLEKPI_API_STATUS_CODE [%v]", resp.StatusCode)
	}
	qrys = GenerateKpiEntryQueriesYMD(userID, kpiID, entries, ds)
	resps = RunQueries(client, qrys)
	return qrys, resps, nil
}

// GenerateKpiEntryQueriesYMD returns a set of queries given a set
// of current KPI Entries and a timeseries.TimeSeries containing
// new data.
func GenerateKpiEntryQueriesYMD(userid, kpiid int64, existing []simplekpi.KpiEntry, new timeseries.TimeSeries) KpiEntryQueries {
	qrys := KpiEntryQueries{
		Create: []simplekpi.KpiEntry{},
		Update: []simplekpi.KpiEntry{}}

	existingMap := map[string]simplekpi.KpiEntry{}
	for _, kentry := range existing {
		if kentry.KpiId != kpiid {
			continue
		}
		existingMap[kentry.EntryDate] = kentry
	}

	for _, newEntrySTS := range new.ItemMap {
		ymd := newEntrySTS.Time.Format(timeutil.RFC3339FullDate)
		if kentry, ok := existingMap[ymd]; ok {
			oldVal := int64(kentry.Actual)
			if oldVal == newEntrySTS.Value {
				continue
			}
			newEntry := simplekpi.KpiEntry{
				Id:        kentry.Id,
				UserId:    kentry.UserId,
				KpiId:     kentry.KpiId,
				EntryDate: kentry.EntryDate,
				Notes:     kentry.Notes,
				Actual:    float64(newEntrySTS.Value)}
			qrys.Update = append(qrys.Update, newEntry)
		} else {
			newEntry := simplekpi.KpiEntry{
				UserId:    userid,
				KpiId:     kpiid,
				EntryDate: ymd,
				Actual:    float64(newEntrySTS.Value)}
			qrys.Create = append(qrys.Create, newEntry)
		}
	}
	return qrys
}

// KpiEntryResponseErrors returns a collapsed error slice.
func KpiEntryResponseErrors(resps []KpiEntryResponse) []error {
	errs := []error{}
	for _, res := range resps {
		if res.Error != nil {
			errs = append(errs, res.Error)
		}
	}
	return errs
}

// KpiEntryResponse is a wrapper for a batch API response set.
type KpiEntryResponse struct {
	KpiEntry   simplekpi.KpiEntry
	StatusCode int
	Error      error
}

// CreateKpiEntries handles multiple creates.
func CreateKpiEntries(client *simplekpi.APIClient, entries []simplekpi.KpiEntry) []KpiEntryResponse {
	resEntries := []KpiEntryResponse{}
	for _, entry := range entries {
		resEntryAPI, resp, err := client.KPIEntriesApi.AddKPIEntry(
			context.Background(), entry)
		resEntry := KpiEntryResponse{}
		if err != nil {
			resEntry.Error = err
		}
		resEntry.StatusCode = resp.StatusCode
		resEntry.KpiEntry = resEntryAPI
		resEntries = append(resEntries, resEntry)
	}
	return resEntries
}

// UpdateKpiEntries handles multiple updates.
func UpdateKpiEntries(client *simplekpi.APIClient, entries []simplekpi.KpiEntry) []KpiEntryResponse {
	resEntries := []KpiEntryResponse{}
	for _, entry := range entries {
		resEntryAPI, resp, err := client.KPIEntriesApi.UpdateKPIEntry(
			context.Background(), entry.KpiId, entry)
		resEntry := KpiEntryResponse{}
		if err != nil {
			resEntry.Error = err
		}
		resEntry.StatusCode = resp.StatusCode
		resEntry.KpiEntry = resEntryAPI
		resEntries = append(resEntries, resEntry)
	}
	return resEntries
}

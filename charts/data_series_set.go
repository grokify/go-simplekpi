package charts

import (
	"context"
	"fmt"
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gocharts/data/timeseries/interval"
	"github.com/grokify/mogo/errors/errorsutil"
	"github.com/grokify/mogo/time/timeutil"
)

type SimplekpiDataSeriesSet struct {
	StartTime     time.Time
	EndTime       time.Time
	KpiIDs        []int32
	KpiInfos      map[int32]simplekpi.Kpi
	DataSeriesSet interval.TimeSeriesSet
}

func NewSimplekpiDataSeriesSet(iterval timeutil.Interval, weekStart time.Weekday) SimplekpiDataSeriesSet {
	return SimplekpiDataSeriesSet{
		KpiIDs:        []int32{},
		KpiInfos:      map[int32]simplekpi.Kpi{},
		DataSeriesSet: interval.NewTimeSeriesSet(iterval, weekStart)}
}

func (dss *SimplekpiDataSeriesSet) LoadData(client *simplekpi.APIClient) error {
	funcName := "SimplekpiDataSeriesSet.LoadData()"
	for _, kpiID := range dss.KpiIDs {
		kpi, resp, err := client.KPIsApi.GetKPI(context.Background(),
			int64(kpiID))
		if err != nil {
			return errorsutil.Wrap(err, funcName)
		} else if resp.StatusCode >= 300 {
			return fmt.Errorf("E_SIMPLEKPI_API_RESP [%v]: %s", resp.StatusCode, funcName)
		}
		if dss.KpiInfos == nil {
			dss.KpiInfos = map[int32]simplekpi.Kpi{}
		}
		dss.KpiInfos[kpiID] = kpi
		ds, err := GetKpiAsDataSeries(client, uint64(kpiID), dss.StartTime, dss.EndTime)
		if err != nil {
			return errorsutil.Wrap(err, funcName)
		}
		dss.DataSeriesSet.SourceSeriesMap[ds.SeriesName] = ds
	}
	return nil
}

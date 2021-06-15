package charts

import (
	"time"

	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/go-simplekpi/simplekpiutil"
	"github.com/grokify/gocharts/data/timeseries"
	"github.com/grokify/simplego/errors/errorsutil"
	"github.com/grokify/simplego/time/timeutil"
)

func PercentTwoKPIs(skClient *simplekpi.APIClient, numerKpiId1, denomKpiId2 uint64, t0, t1 time.Time) (timeseries.TimeSeries, timeseries.TimeSeries, timeseries.TimeSeries, error) {
	cu := simplekpiutil.ClientUtil{
		APIClient: skClient}
	if t0.Equal(t1) {
		t0 = timeutil.TimeZeroRFC3339()
		t1 = time.Now()
	}
	ds1, err1 := cu.GetKPIEntriesAsDataSeries(numerKpiId1, t0, t1)
	ds2, err2 := cu.GetKPIEntriesAsDataSeries(denomKpiId2, t0, t1)
	err := errorsutil.Join(false, err1, err2)
	if err != nil {
		return ds1, ds2, ds2, err
	}
	ds3, err := timeseries.TimeSeriesDivide(ds1, ds2)
	return ds1, ds2, ds3, err
}

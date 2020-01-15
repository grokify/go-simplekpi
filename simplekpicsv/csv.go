package simplekpicsv

import "time"

type Entry struct {
	KpiId  int32     // use `KPI ID` as title
	Date   time.Time // format as RFC-3339 `date` in CSV
	Actual float64
}

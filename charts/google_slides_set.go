package charts

import (
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/gogoogle/slidesutil/v1"
)

type SlidesInfoSet struct {
	ImageBaseURL     string
	Verbose          bool
	KpiSlideOptsList []KpiSlideOpts
}

func NewSlidesInfoSet() SlidesInfoSet {
	return SlidesInfoSet{KpiSlideOptsList: []KpiSlideOpts{}}
}

func (set *SlidesInfoSet) Inflate() {
	for i, opts := range set.KpiSlideOptsList {
		opts.ImageBaseURL = set.ImageBaseURL
		set.KpiSlideOptsList[i] = opts
	}
}

func CreateKPISlides(skClient *simplekpi.APIClient, pc *slidesutil.PresentationCreator, set SlidesInfoSet) error {
	set.Inflate()
	for _, opts := range set.KpiSlideOptsList {
		_, err := CreateKPISlide(skClient, pc, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

func (set *SlidesInfoSet) Count() int {
	return len(set.KpiSlideOptsList)
}

func (set *SlidesInfoSet) Filter(kpiIDs []uint64) SlidesInfoSet {
	if len(kpiIDs) == 0 {
		return *set
	}
	newSlidesInfoSet := SlidesInfoSet{
		ImageBaseURL:     set.ImageBaseURL,
		Verbose:          set.Verbose,
		KpiSlideOptsList: []KpiSlideOpts{}}
	mapKpis := map[uint64]int{}
	for _, kpiID := range kpiIDs {
		mapKpis[kpiID] = 1
	}
	for _, opts := range set.KpiSlideOptsList {
		if _, ok := mapKpis[opts.KpiID]; ok {
			newSlidesInfoSet.KpiSlideOptsList = append(
				newSlidesInfoSet.KpiSlideOptsList, opts)
		}
	}
	return newSlidesInfoSet
}

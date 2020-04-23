package charts

import (
	"github.com/grokify/go-simplekpi/simplekpi"
	"github.com/grokify/googleutil/slidesutil/v1"
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
		err := CreateKPISlide(skClient, pc, opts)
		if err != nil {
			return err
		}
	}
	return nil
}

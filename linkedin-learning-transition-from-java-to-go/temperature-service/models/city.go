package models

var (
	beachVacationThreshold float64 = 22
	skiVacationThreshold   float64 = -2
)

type city struct {
	name        string
	tempC       float64
	hasBeach    bool
	hasMountain bool
}

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

func NewCity(
	n string,
	t float64,
	hasBeach bool,
	hasMountain bool) CityTemp {
	return &city{
		name:        n,
		tempC:       t,
		hasBeach:    hasBeach,
		hasMountain: hasMountain,
	}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return c.tempC*1.8 + 32
}

func (c city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC >= beachVacationThreshold
}

func (c city) SkiVacationReady() bool {
	return c.hasMountain && c.tempC <= skiVacationThreshold
}

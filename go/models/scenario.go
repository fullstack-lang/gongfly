package models

import "time"

// Scenario holds all infos about a scenario
// swagger:model Scenario
type Scenario struct {
	Name string // mandatory

	start time.Time
	end   time.Time

	Lat, Lng  float64 // map center
	ZoomLevel float64 // zoom level at the initialisation

	// Message Visual Speed, in km/h
	// this field can be tweaked by the end user
	MessageVisualSpeed float64
}

func (scenario *Scenario) GetStart() time.Time {
	return scenario.start
}

func (scenario *Scenario) SetStart(time time.Time) {
	scenario.start = time
}

func (scenario *Scenario) GetEnd() time.Time {
	return scenario.end
}

func (scenario *Scenario) SetEnd(time time.Time) {
	scenario.end = time
}

func (scenario *Scenario) GetLat() float64 { return scenario.Lat }
func (scenario *Scenario) GetLng() float64 { return scenario.Lng }

func (scenario *Scenario) GetZoomLevel() float64 { return scenario.ZoomLevel }

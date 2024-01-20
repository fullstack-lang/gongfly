// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type CivilianAirport_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Name string
}

type Liner_WOP struct {
	// insertion point
	Name string
	Lat float64
	Lng float64
	Heading float64
	Level float64
	Speed float64
	State LinerStateEnum
	TargetHeading float64
	TargetLocationLat float64
	TargetLocationLng float64
	DistanceToTarget float64
	MaxRotationalSpeed float64
	VerticalSpeed float64
	Timestampstring string
}

type Message_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Heading float64
	Level float64
	Speed float64
	State MessageStateEnum
	Name string
	TargetLocationLat float64
	TargetLocationLng float64
	DistanceToTarget float64
	Timestampstring string
	DurationSinceSimulationStart time.Duration
	Timestampstartstring string
	Source string
	Destination string
	Content string
	About_string string
	Display bool
}

type OpsLine_WOP struct {
	// insertion point
	IsTransmitting bool
	TransmissionMessage string
	IsTransmittingBackward bool
	TransmissionMessageBackward string
	State OperationalLineStateEnum
	Name string
}

type Radar_WOP struct {
	// insertion point
	State RadarStateEnum
	Name string
	Lat float64
	Lng float64
	Range float64
}

type Satellite_WOP struct {
	// insertion point
	Name string
	Line1 string
	Line2 string
	Lat float64
	Lng float64
	Heading float64
	Level float64
	Speed float64
	VerticalSpeed float64
	Timestampstring string
}

type Scenario_WOP struct {
	// insertion point
	Name string
	Lat float64
	Lng float64
	ZoomLevel float64
	MessageVisualSpeed float64
}


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

func (from *CivilianAirport) CopyBasicFields(to *CivilianAirport) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Name = from.Name
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

func (from *Liner) CopyBasicFields(to *Liner) {
	// insertion point
	to.Name = from.Name
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Heading = from.Heading
	to.Level = from.Level
	to.Speed = from.Speed
	to.State = from.State
	to.TargetHeading = from.TargetHeading
	to.TargetLocationLat = from.TargetLocationLat
	to.TargetLocationLng = from.TargetLocationLng
	to.DistanceToTarget = from.DistanceToTarget
	to.MaxRotationalSpeed = from.MaxRotationalSpeed
	to.VerticalSpeed = from.VerticalSpeed
	to.Timestampstring = from.Timestampstring
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

func (from *Message) CopyBasicFields(to *Message) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Heading = from.Heading
	to.Level = from.Level
	to.Speed = from.Speed
	to.State = from.State
	to.Name = from.Name
	to.TargetLocationLat = from.TargetLocationLat
	to.TargetLocationLng = from.TargetLocationLng
	to.DistanceToTarget = from.DistanceToTarget
	to.Timestampstring = from.Timestampstring
	to.DurationSinceSimulationStart = from.DurationSinceSimulationStart
	to.Timestampstartstring = from.Timestampstartstring
	to.Source = from.Source
	to.Destination = from.Destination
	to.Content = from.Content
	to.About_string = from.About_string
	to.Display = from.Display
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

func (from *OpsLine) CopyBasicFields(to *OpsLine) {
	// insertion point
	to.IsTransmitting = from.IsTransmitting
	to.TransmissionMessage = from.TransmissionMessage
	to.IsTransmittingBackward = from.IsTransmittingBackward
	to.TransmissionMessageBackward = from.TransmissionMessageBackward
	to.State = from.State
	to.Name = from.Name
}

type Radar_WOP struct {
	// insertion point
	State RadarStateEnum
	Name string
	Lat float64
	Lng float64
	Range float64
}

func (from *Radar) CopyBasicFields(to *Radar) {
	// insertion point
	to.State = from.State
	to.Name = from.Name
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Range = from.Range
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

func (from *Satellite) CopyBasicFields(to *Satellite) {
	// insertion point
	to.Name = from.Name
	to.Line1 = from.Line1
	to.Line2 = from.Line2
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Heading = from.Heading
	to.Level = from.Level
	to.Speed = from.Speed
	to.VerticalSpeed = from.VerticalSpeed
	to.Timestampstring = from.Timestampstring
}

type Scenario_WOP struct {
	// insertion point
	Name string
	Lat float64
	Lng float64
	ZoomLevel float64
	MessageVisualSpeed float64
}

func (from *Scenario) CopyBasicFields(to *Scenario) {
	// insertion point
	to.Name = from.Name
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.ZoomLevel = from.ZoomLevel
	to.MessageVisualSpeed = from.MessageVisualSpeed
}


// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

// insertion point
type Circle_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Name string
	Radius float64
	ColorEnum ColorEnum
	DashStyleEnum DashStyleEnum
}

type DivIcon_WOP struct {
	// insertion point
	Name string
	SVG string
}

type LayerGroup_WOP struct {
	// insertion point
	Name string
	DisplayName string
}

type LayerGroupUse_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

type MapOptions_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Name string
	ZoomLevel float64
	UrlTemplate string
	Attribution string
	MaxZoom int
	ZoomControl bool
	AttributionControl bool
	ZoomSnap int
}

type Marker_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Name string
	ColorEnum ColorEnum
}

type UserClick_WOP struct {
	// insertion point
	Name string
	Lat float64
	Lng float64
	TimeOfClick time.Time
}

type VLine_WOP struct {
	// insertion point
	StartLat float64
	StartLng float64
	EndLat float64
	EndLng float64
	Name string
	ColorEnum ColorEnum
	DashStyleEnum DashStyleEnum
	IsTransmitting TransmittingEnum
	Message string
	IsTransmittingBackward TransmittingEnum
	MessageBackward string
}

type VisualTrack_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Heading float64
	Level float64
	Speed float64
	VerticalSpeed float64
	Name string
	ColorEnum ColorEnum
	DisplayTrackHistory bool
	DisplayLevelAndSpeed bool
}


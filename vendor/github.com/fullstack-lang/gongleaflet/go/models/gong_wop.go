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

func (from *Circle) CopyBasicFields(to *Circle) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Name = from.Name
	to.Radius = from.Radius
	to.ColorEnum = from.ColorEnum
	to.DashStyleEnum = from.DashStyleEnum
}

type DivIcon_WOP struct {
	// insertion point
	Name string
	SVG string
}

func (from *DivIcon) CopyBasicFields(to *DivIcon) {
	// insertion point
	to.Name = from.Name
	to.SVG = from.SVG
}

type LayerGroup_WOP struct {
	// insertion point
	Name string
	DisplayName string
}

func (from *LayerGroup) CopyBasicFields(to *LayerGroup) {
	// insertion point
	to.Name = from.Name
	to.DisplayName = from.DisplayName
}

type LayerGroupUse_WOP struct {
	// insertion point
	Name string
	IsDisplayed bool
}

func (from *LayerGroupUse) CopyBasicFields(to *LayerGroupUse) {
	// insertion point
	to.Name = from.Name
	to.IsDisplayed = from.IsDisplayed
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

func (from *MapOptions) CopyBasicFields(to *MapOptions) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Name = from.Name
	to.ZoomLevel = from.ZoomLevel
	to.UrlTemplate = from.UrlTemplate
	to.Attribution = from.Attribution
	to.MaxZoom = from.MaxZoom
	to.ZoomControl = from.ZoomControl
	to.AttributionControl = from.AttributionControl
	to.ZoomSnap = from.ZoomSnap
}

type Marker_WOP struct {
	// insertion point
	Lat float64
	Lng float64
	Name string
	ColorEnum ColorEnum
}

func (from *Marker) CopyBasicFields(to *Marker) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Name = from.Name
	to.ColorEnum = from.ColorEnum
}

type UserClick_WOP struct {
	// insertion point
	Name string
	Lat float64
	Lng float64
	TimeOfClick time.Time
}

func (from *UserClick) CopyBasicFields(to *UserClick) {
	// insertion point
	to.Name = from.Name
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.TimeOfClick = from.TimeOfClick
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

func (from *VLine) CopyBasicFields(to *VLine) {
	// insertion point
	to.StartLat = from.StartLat
	to.StartLng = from.StartLng
	to.EndLat = from.EndLat
	to.EndLng = from.EndLng
	to.Name = from.Name
	to.ColorEnum = from.ColorEnum
	to.DashStyleEnum = from.DashStyleEnum
	to.IsTransmitting = from.IsTransmitting
	to.Message = from.Message
	to.IsTransmittingBackward = from.IsTransmittingBackward
	to.MessageBackward = from.MessageBackward
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

func (from *VisualTrack) CopyBasicFields(to *VisualTrack) {
	// insertion point
	to.Lat = from.Lat
	to.Lng = from.Lng
	to.Heading = from.Heading
	to.Level = from.Level
	to.Speed = from.Speed
	to.VerticalSpeed = from.VerticalSpeed
	to.Name = from.Name
	to.ColorEnum = from.ColorEnum
	to.DisplayTrackHistory = from.DisplayTrackHistory
	to.DisplayLevelAndSpeed = from.DisplayLevelAndSpeed
}


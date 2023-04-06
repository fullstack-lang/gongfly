package models

import "log"

// Marker provides all necessary elements to the front to display a track
//
// L.Marker is used to display clickable/draggable icons on the map. Extends Layer.
//
// swagger:model Marker
type Marker struct {
	Lat, Lng float64
	Name     string

	ColorEnum ColorEnum

	// LayerGroup the object belongs to
	LayerGroup *LayerGroup

	// DivIcon
	DivIcon *DivIcon

	// swagger:ignore
	// access to the models instance that contains the original information
	MarkerInteface MarkerInterface `gorm:"-"`

	// the Stage_ of the visual track
	Stage_ *StageStruct
}

type MarkerInterface interface {
	GetLat() (lat float64)
	GetLng() (lng float64)
	GetName() (name string)
	GetLayerGroupName() string
}

func (marker *Marker) UpdateMarker() {
	if marker.MarkerInteface != nil {
		marker.Name = marker.MarkerInteface.GetName()

		marker.Lat = marker.MarkerInteface.GetLat()
		marker.Lng = marker.MarkerInteface.GetLng()

		marker.LayerGroup =
			ComputeLayerGroupFromLayerGroupName(
				marker.Stage_,
				marker.MarkerInteface.GetLayerGroupName())
	}
}

// attach visual center to center
func AttachMarker(
	gongleafletStage *StageStruct,
	markerInterface MarkerInterface,
	colorEnum ColorEnum,
	divIcon *DivIcon) (marker *Marker) {

	// check icon is present
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}
	marker = new(Marker).Stage(gongleafletStage)
	marker.MarkerInteface = markerInterface
	marker.ColorEnum = colorEnum
	marker.DivIcon = divIcon
	marker.Stage_ = gongleafletStage
	marker.UpdateMarker()

	return
}

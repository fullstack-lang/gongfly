package models

import "log"

// VisualTrack provides all necessary elements to the front to display a track
//
// # In leaflet, it is translated into a MovingMarker
//
// swagger:model VisualTrack
type VisualTrack struct {
	Lat, Lng, Heading, Level, Speed, VerticalSpeed float64
	Name                                           string

	ColorEnum ColorEnum

	// LayerGroup the object belongs to
	LayerGroup *LayerGroup

	// access to the models instance that contains the original information
	// swagger:ignore
	// a moving object has to implement this interface to
	// be displayed as a visual interface
	//
	VisualTrackInterface VisualTrackInterface `gorm:"-"`

	DivIcon *DivIcon

	// if true display dots from the trajectory
	DisplayTrackHistory bool

	// if true, display level and speed below the icon
	DisplayLevelAndSpeed bool

	// the stage of the visual track
	stage *StageStruct
}

func (visualTrack *VisualTrack) UpdateTrack() {
	if visualTrack.VisualTrackInterface != nil {
		visualTrack.Name = visualTrack.VisualTrackInterface.GetName()

		visualTrack.Lat = visualTrack.VisualTrackInterface.GetLat()
		visualTrack.Lng = visualTrack.VisualTrackInterface.GetLng()
		visualTrack.Heading = visualTrack.VisualTrackInterface.GetHeading()
		visualTrack.Level = visualTrack.VisualTrackInterface.GetLevel()
		visualTrack.Speed = visualTrack.VisualTrackInterface.GetSpeed()
		visualTrack.VerticalSpeed = visualTrack.VisualTrackInterface.GetVerticalSpeed()

		visualTrack.LayerGroup = ComputeLayerGroupFromLayerGroupName(visualTrack.stage, visualTrack.VisualTrackInterface.GetLayerGroupName())
	}
}

// AttachVisualTrack attaches an object that match the visualTrackInterface to a visual track
// the visual track will request information from the object in order to display the track
func AttachVisualTrack(
	gongleafletStage *StageStruct,
	track VisualTrackInterface,
	divIcon *DivIcon,
	colorEnum ColorEnum,
	displayTrackHistory bool,
	displayLevelAndSpeed bool) (visualTrack *VisualTrack) {

	// sometimes, the visual icon is nil (not reproductible bug)
	if divIcon == nil {
		log.Fatal("nil visual icon")
	}

	visualTrack = new(VisualTrack).Stage(gongleafletStage)
	visualTrack.VisualTrackInterface = track
	visualTrack.DivIcon = divIcon
	visualTrack.DisplayTrackHistory = displayTrackHistory
	visualTrack.DisplayLevelAndSpeed = displayLevelAndSpeed
	visualTrack.ColorEnum = colorEnum
	visualTrack.stage = gongleafletStage
	visualTrack.UpdateTrack()

	return
}

package models

// VisualTrack provides all necessary elements to the front to display a track
//
// In leaflet, it is translated into a MovingMarker
//
// swagger:model visualtrack
type VisualTrack struct {
	Lat, Lng, Heading, Level, Speed, VerticalSpeed float64
	Name                                           string

	VisualColorEnum VisualColorEnum

	// LayerGroup the object belongs to
	LayerGroup *LayerGroup

	// access to the models instance that contains the original information
	// swagger:ignore
	VisualTrackInterface VisualTrackInterface `gorm:"-"`

	DivIcon *DivIcon

	Display bool

	// if true display dots from the trajectory
	DisplayTrackHistory bool

	// if true, display level and speed below the icon
	DisplayLevelAndSpeed bool
}

type VisualTrackInterface interface {

	// position
	GetLat() (lat float64)
	GetLng() (lng float64)

	// cinemetic
	GetHeading() (heading float64)
	GetSpeed() (speed float64)
	GetVerticalSpeed() (verticalSpeed float64)
	GetLevel() (level float64)

	GetName() (name string)

	// the models name knows wether it has to be displayed
	GetDisplay() bool
	GetLayerGroupName() string
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
		visualTrack.Display = visualTrack.VisualTrackInterface.GetDisplay()

		visualTrack.LayerGroup = computeLayerGroupFromLayerGroupName(visualTrack.VisualTrackInterface.GetLayerGroupName())
	}
}
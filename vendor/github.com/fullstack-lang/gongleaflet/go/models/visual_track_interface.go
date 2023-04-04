package models

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
	GetLayerGroupName() string
}

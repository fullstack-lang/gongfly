package models

// VisualMap provides all necessary elements to the front to display a track
//
// swagger:model visualmap
type VisualMap struct {
	Lat, Lng  float64 // map center
	Name      string
	ZoomLevel float64 // zoom level at the initialisation

	UrlTemplate        string
	Attribution        string
	MaxZoom            int
	ZoomControl        bool
	AttributionControl bool
	ZoomSnap           bool

	// swagger:ignore
	// access to the models instance that contains the original information
	VisualMapInterface VisualMapInterface `gorm:"-"`
}

type VisualMapInterface interface {
	GetLat() (lat float64)
	GetLng() (lng float64)
	GetName() (name string)

	GetZoomLevel() (lng float64)
}

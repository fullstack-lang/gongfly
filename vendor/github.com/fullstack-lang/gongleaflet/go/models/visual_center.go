package models

// VisualCenter provides all necessary elements to the front to display a track
//
// swagger:model visualcenter
type VisualCenter struct {
	Lat, Lng float64
	Name     string

	VisualColorEnum VisualColorEnum

	// VisualLayer the object belongs to
	VisualLayer *VisualLayer
	// VisualIcon
	VisualIcon *VisualIcon

	// swagger:ignore
	// access to the models instance that contains the original information
	VisualCenterInteface VisualCenterInterface `gorm:"-"`
}

type VisualCenterInterface interface {
	GetLat() (lat float64)
	GetLng() (lng float64)
	GetName() (name string)
	GetVisualLayerName() string
}

func (visualCenter *VisualCenter) UpdateCenter() {
	if visualCenter.VisualCenterInteface != nil {
		visualCenter.Name = visualCenter.VisualCenterInteface.GetName()

		visualCenter.Lat = visualCenter.VisualCenterInteface.GetLat()
		visualCenter.Lng = visualCenter.VisualCenterInteface.GetLng()

		visualCenter.VisualLayer =
			computeVisualLayerFromVisualLayerName(visualCenter.VisualCenterInteface.GetVisualLayerName())
	}
}

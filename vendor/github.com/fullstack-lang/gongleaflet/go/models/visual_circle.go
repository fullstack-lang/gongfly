package models

// VisualCircle provides all necessary elements to the front to display a track
//
// swagger:model visualcenter
type VisualCircle struct {
	Lat, Lng float64
	Name     string

	Radius float64

	VisualColorEnum VisualColorEnum
	DashStyleEnum   DashStyleEnum

	// VisualLayer the object belongs to
	VisualLayer *VisualLayer

	// swagger:ignore
	// access to the models instance that contains the original information
	Circle VisualCircleInterface `gorm:"-"`
}

type VisualCircleInterface interface {
	GetLat() (lat float64)
	GetLng() (lng float64)
	GetRadius() (lng float64)

	GetName() (name string)
}

func (VisualCircle *VisualCircle) UpdateCircle() {
	if VisualCircle.Circle != nil {
		VisualCircle.Name = VisualCircle.Circle.GetName()

		VisualCircle.Lat = VisualCircle.Circle.GetLat()
		VisualCircle.Lng = VisualCircle.Circle.GetLng()

		VisualCircle.Radius = VisualCircle.Circle.GetRadius()

		VisualCircle.VisualLayer = DefaultVisualLayer
	}
}

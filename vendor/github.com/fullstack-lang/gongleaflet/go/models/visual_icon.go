package models

// VisualIcon provides all necessary elements to the front to display a track
//
// swagger:model visualicon
type VisualIcon struct {
	Name string

	SVG string // the SVG description

	// swagger:ignore
	// access to the models instance that contains the original information
	VisualIconInterface VisualIconInterface `gorm:"-"`
}

type VisualIconInterface interface {
	GetIconName() (name string)
	GetSVG() (name string)
}

func (VisualIcon *VisualIcon) UpdateVisualIcon() {
	if VisualIcon.VisualIconInterface != nil {
		VisualIcon.Name = VisualIcon.VisualIconInterface.GetIconName()
	}
}

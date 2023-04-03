package models

// LayerGroup is the gong version of the leaflet LayerGroup
//
// swagger:model LayerGroup
type LayerGroup struct {
	Name string

	DisplayName string

	// swagger:ignore
	// access to the models instance that contains the original information
	LayerGroupInterface LayerGroupInterface `gorm:"-"`
}

type LayerGroupInterface interface {
	GetLayerName() (name string)
}

func (LayerGroup *LayerGroup) UpdateLayerGroup() {
	if LayerGroup.LayerGroupInterface != nil {
		LayerGroup.Name = LayerGroup.LayerGroupInterface.GetLayerName()
		LayerGroup.DisplayName = LayerGroup.LayerGroupInterface.GetLayerName()
	}
}

var DefaultLayerGroup *LayerGroup

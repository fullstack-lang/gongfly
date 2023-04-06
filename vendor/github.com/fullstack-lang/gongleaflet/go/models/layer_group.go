package models

import "log"

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

// little simple algo for the visual layer computation
func ComputeLayerGroupFromLayerGroupName(gongleafletStage *StageStruct, layerGroupName string) (layerGroup *LayerGroup) {

	for _layerGroup := range gongleafletStage.LayerGroups {
		if _layerGroup.Name == layerGroupName {
			layerGroup = _layerGroup
			continue
		}
	}
	if layerGroup == nil {
		log.Printf("Unknown layer %s ", layerGroupName)
	}
	return
}

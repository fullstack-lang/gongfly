package models

import "log"

// VisualLayer provides all necessary elements to the front to display a track
//
// swagger:model visuallayer
type VisualLayer struct {
	Name string

	DisplayName string

	// swagger:ignore
	// access to the models instance that contains the original information
	VisualLayerInterface VisualLayerInterface `gorm:"-"`
}

type VisualLayerInterface interface {
	GetLayerName() (name string)
}

func (VisualLayer *VisualLayer) UpdateVisualLayer() {
	if VisualLayer.VisualLayerInterface != nil {
		VisualLayer.Name = VisualLayer.VisualLayerInterface.GetLayerName()
		VisualLayer.DisplayName = VisualLayer.VisualLayerInterface.GetLayerName()
	}
}

// little simple algo for the visual layer computation
func computeVisualLayerFromVisualLayerName(visualLayerName string) (visualLayer *VisualLayer) {

	for _visualLayer := range Stage.VisualLayers {
		if _visualLayer.Name == visualLayerName {
			visualLayer = _visualLayer
			continue
		}
	}
	if visualLayer == nil {
		log.Printf("Unknown layer %s ", visualLayerName)
	}
	return
}

var DefaultVisualLayer = (&VisualLayer{
	Name: "Default",
}).StageCopy()

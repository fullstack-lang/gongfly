package models

import gongsim_models "github.com/fullstack-lang/gongsim/go/models"

// swagger:model MovingObject
type MovingObject struct {

	// current position
	Lat float64
	Lng float64

	// heading in degrees
	Heading float64

	// level
	// for instance 210 means 21 000 feets
	Level float64

	// Horizonatl Speed, in km/h
	Speed float64
}

func (*MovingObject) GetConcept() ConceptEnum           { return Aircraft_ }
func (*MovingObject) GetVisualLayerName() (name string) { return string(Aircraft_) }

type MovingObjectInterface interface {
	GetName() string
	GetLat() float64
	GetLng() float64
	QueueEvent(gongsim_models.EventInterface)
	TransfertAndQueueEvent(gongsim_models.EventInterface)
}

package models

import gongsim_models "github.com/fullstack-lang/gongsim/go/models"

// LatLngAgentInterface is an interface to another agent
// that have a name, a position and who can queue/fire events
// usefull for handling message from one agent to another
type LatLngAgentInterface interface {
	GetName() string
	GetLat() float64
	GetLng() float64
	QueueEvent(gongsim_models.EventInterface)
	TransfertAndQueueEvent(gongsim_models.EventInterface)
}

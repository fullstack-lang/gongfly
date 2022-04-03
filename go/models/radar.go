package models

import (
	"fmt"
	"log"

	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// Radar
// swagger:model radar
type Radar struct {
	System // concept

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	State RadarStateEnum

	Name string

	// current position
	Lat float64
	Lng float64

	// range in km
	Range float64
}

func (Radar *Radar) Register() (res *Radar) {
	gongsim_models.AppendToSingloton(Radar)
	res = Radar
	return
}

// state
const (
	WORKING RadarStateEnum = "WORKING"
)

// RadarStateEnum ..
// swagger:enum RadarStateEnum
type RadarStateEnum string

// FireNextEvent fire next Event
func (radar *Radar) FireNextEvent() {

	event, _ := radar.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		radar.QueueEvent(checkStateEvent)

		// update state vector
		switch radar.State {
		}
	default:
		err := fmt.Sprintf("unkown event type %T", event)
		log.Panic(err)
	}
}

// functions to satisty the visual interface for center
func (radar *Radar) GetLat() float64 { return radar.Lat }
func (radar *Radar) GetLng() float64 { return radar.Lng }

// functions to satisty the visual interface for circle
func (radar *Radar) GetRadius() float64 { return radar.Range }

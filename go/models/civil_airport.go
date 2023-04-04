package models

import gongsim_models "github.com/fullstack-lang/gongsim/go/models"

// CivilianAirport
type CivilianAirport struct {
	CenterConcept // allow for selecting visual objects

	// Agent
	// swagger:ignore
	gongsim_models.Agent

	Name string
}

func (civilianAirport *CivilianAirport) Register(engine *gongsim_models.Engine) (res *CivilianAirport) {
	gongsim_models.AppendToSingloton(engine, civilianAirport)
	res = civilianAirport
	return
}

// FireNextEvent fire next Event
// basicaly, it does nothing (yet to be implemented)
func (civilianAirport *CivilianAirport) FireNextEvent() {

	event, _ := civilianAirport.GetNextEventAndRemoveIt()

	switch event.(type) {
	case *gongsim_models.UpdateState:
		checkStateEvent := event.(*gongsim_models.UpdateState)

		// post next event
		checkStateEvent.SetFireTime(checkStateEvent.GetFireTime().Add(checkStateEvent.Period))
		civilianAirport.QueueEvent(checkStateEvent)
	}
}

// functions to satisty the visual interface for center
func (civilianAirport *CivilianAirport) GetLat() float64 { return civilianAirport.Lat }
func (civilianAirport *CivilianAirport) GetLng() float64 { return civilianAirport.Lng }

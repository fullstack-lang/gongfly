package simulation

import (
	"log"
	"time"

	"github.com/fullstack-lang/gongfly/go/icons"
	"github.com/fullstack-lang/gongfly/go/models"
	"github.com/fullstack-lang/gongfly/go/reference"

	target_models "github.com/fullstack-lang/gongfly/go/models"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// Simulation meets the gongsim_models.SimulationInterface
//
// it is in a separate package from the models package because it handles the relation
// to the visual representation of the simulaton object with leaflet
type Simulation struct {
	scenario *target_models.Scenario
}

func (engineSpecific *Simulation) setInitialStateVectorOfAgentsAndSimulation() {

	engine := gongsim_models.EngineSingloton

	// start and end date
	engineSpecific.scenario = reference.Scenario1
	engine.SetStartTime(engineSpecific.scenario.GetStart())
	engine.SetCurrentTime(engine.GetStartTime())
	engine.State = gongsim_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", engine.GetStartTime())

	engine.SetEndTime(engineSpecific.scenario.GetEnd())
	log.Printf("Sim end  \t\t\t%s\n", engine.EndTime)

	engine.Speed = 60

	// liner MDM

	*reference.Sc1_AF_3577_MDM = reference.Sc1_AF_CDG_HYE_ref
	reference.Sc1_AF_3577_MDM.QueueUpdateEvent(1 * time.Second)

}

// NewSimulation ...
func NewSimulation() (simulation *Simulation) {

	//
	// simulation generic initialisation steps
	//
	engine := gongsim_models.EngineSingloton
	simulation = &Simulation{}

	//
	// simulation initialisation
	//
	simulation.Reset(engine)

	return
}

// HasAnyStateChanged is used when the simulation user ask to advance up to the next
// state change. It is not implemented in this simulation
func (engineSpecific *Simulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {

	return false
}

// map to store dynamicaly association between message and visual tracks
var MapMessageVisualTrack = make(map[*target_models.Message]*gongleaflet_models.VisualTrack)

// CommitAgent is called by the generic engine anytime a visual refresh is needed
// Commit will perform a write of all agent state to the back repo and the back repo
// will be checked out by the front repo (for display sake)
func (engineSpecific *Simulation) CommitAgents(engine *gongsim_models.Engine) {

	// if there are new messages, attach them to visual tracks
	for message := range target_models.Stage.Messages {
		if message.Display == true && MapMessageVisualTrack[message] == nil {
			visualTrack := new(gongleaflet_models.VisualTrack).Stage()
			visualTrack.Name = message.Content
			visualTrack.VisualTrackInterface = message
			display := true
			visualTrack.Display = display
			visualTrack.DivIcon = icons.Arrow
			_false := false
			visualTrack.DisplayTrackHistory = _false
			visualTrack.DisplayLevelAndSpeed = _false

			visualTrack.UpdateTrack()

			// append the visual track to the map between message & visual tracks
			MapMessageVisualTrack[message] = visualTrack
		}
	}

	nbMessages := len(target_models.Stage.Messages)
	nbMessagesInMap := len(MapMessageVisualTrack)

	if nbMessages < nbMessagesInMap {
		log.Fatal("problem with messages visualisations")
	}

	// if a message has a display value at false, delete the visual track
	for message := range target_models.Stage.Messages {
		if message.Display == false {
			if visualTrack := MapMessageVisualTrack[message]; visualTrack != nil {
				visualTrack.Unstage()
				delete(MapMessageVisualTrack, message)
			}
		}
	}

	// update all visual tracks
	for visualTrack := range gongleaflet_models.Stage.VisualTracks {
		visualTrack.UpdateTrack()
	}

	// update all visual centers
	for visualCenter := range gongleaflet_models.Stage.Markers {
		visualCenter.UpdateMarker()
	}

	// update all visual lines
	for visualLine := range gongleaflet_models.Stage.VisualLines {
		visualLine.UpdateLine()
	}

	// update all visual circles
	for visualCircle := range gongleaflet_models.Stage.VisualCircles {
		visualCircle.UpdateCircle()
	}

	// put all to database
	target_models.Stage.Commit()
	gongsim_models.Stage.Commit()
	gongleaflet_models.Stage.Commit()
}

// Reset simulation
func (simulation *Simulation) Reset(engine *gongsim_models.Engine) {

	// remove all events from agents
	for _, agent := range engine.Agents() {
		agent.Reset()
	}

	// remove all message and their VisualTracks
	for message := range target_models.Stage.Messages {
		message.Unstage()
		if visualTrack := MapMessageVisualTrack[message]; visualTrack != nil {
			visualTrack.Unstage()
			delete(MapMessageVisualTrack, message)
		}
	}

	// set initial conditions
	simulation.setInitialStateVectorOfAgentsAndSimulation()

	// set the engine of all agents
	for _, agent := range engine.Agents() {
		agent.SetEngine(engine)
	}

	log.Printf("Simulation reset")
}

// GetLastCommitNb is used by the front repo to computer wether a checkout is necessary
// It does so by comparing successive commit nb
func (simulation *Simulation) GetLastCommitNb() (commitNb uint) {

	if models.Stage.BackRepo != nil {
		commitNb = models.Stage.BackRepo.GetLastCommitNb()
	}

	return
}

func (simulation *Simulation) GetLastCommitNbFromFront() (commitNb uint) {

	if models.Stage.BackRepo != nil {
		commitNb = models.Stage.BackRepo.GetLastPushFromFrontNb()
	}
	return
}

// CheckoutAgents checkout all staged agents to the back repo
func (simulation *Simulation) CheckoutAgents(engine *gongsim_models.Engine) {

	// checkout the simulation agent states
	models.Stage.Checkout()
}

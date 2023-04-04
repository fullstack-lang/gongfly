package simulation

import (
	"log"

	"github.com/fullstack-lang/gongfly/go/icons"
	"github.com/fullstack-lang/gongfly/go/reference"

	gongfly_models "github.com/fullstack-lang/gongfly/go/models"
	gongleaflet_models "github.com/fullstack-lang/gongleaflet/go/models"
	gongsim_models "github.com/fullstack-lang/gongsim/go/models"
)

// Simulation meets the gongsim_models.SimulationInterface
//
// it is in a separate package from the models package because it handles the relation
// to the visual representation of the simulaton object with leaflet
type Simulation struct {
	scenario         *gongfly_models.Scenario
	gongflyStage     *gongfly_models.StageStruct
	gongleafletStage *gongleaflet_models.StageStruct
	gongsimStage     *gongsim_models.StageStruct
	engine           *gongsim_models.Engine
}

func (simulation *Simulation) GetEngine() (engine *gongsim_models.Engine) {
	engine = simulation.engine
	return
}

func (simulation *Simulation) setInitialStateVectorOfAgentsAndSimulation() {

	// start and end date
	simulation.scenario = reference.Scenario1
	simulation.engine.SetStartTime(simulation.scenario.GetStart())
	simulation.engine.SetCurrentTime(simulation.engine.GetStartTime())
	simulation.engine.State = gongsim_models.PAUSED
	log.Printf("Sim start \t\t\t%s\n", simulation.engine.GetStartTime())

	simulation.engine.SetEndTime(simulation.scenario.GetEnd())
	log.Printf("Sim end  \t\t\t%s\n", simulation.engine.EndTime)

	simulation.engine.Speed = 120
}

// NewSimulation ...
func NewSimulation(
	gongflyStage *gongfly_models.StageStruct,
	gongsimStage *gongsim_models.StageStruct,
	gongleafletStage *gongleaflet_models.StageStruct) (simulation *Simulation) {

	//
	// simulation generic initialisation steps
	//
	engine := new(gongsim_models.Engine).Stage(gongsimStage)
	simulation = &Simulation{
		gongflyStage:     gongflyStage,
		gongleafletStage: gongleafletStage,
		gongsimStage:     gongsimStage,
		engine:           engine,
	}

	//
	// simulation initialisation
	//
	simulation.Reset()

	return
}

// HasAnyStateChanged is used when the simulation user ask to advance up to the next
// state change. It is not implemented in this simulation
func (simulation *Simulation) HasAnyStateChanged(engine *gongsim_models.Engine) bool {

	return false
}

// map to store dynamicaly association between message and visual tracks
var MapMessageVisualTrack = make(map[*gongfly_models.Message]*gongleaflet_models.VisualTrack)

// CommitAgent is called by the generic engine anytime a visual refresh is needed
// Commit will perform a write of all agent state to the back repo and the back repo
// will be checked out by the front repo (for display sake)
func (simulation *Simulation) CommitAgents(engine *gongsim_models.Engine) {

	// if there are new messages, attach them to visual tracks
	for message := range simulation.gongflyStage.Messages {
		if message.Display == true && MapMessageVisualTrack[message] == nil {
			visualTrack := new(gongleaflet_models.VisualTrack).Stage(simulation.gongleafletStage)
			visualTrack.Name = message.Content
			visualTrack.VisualTrackInterface = message
			visualTrack.DivIcon = icons.Arrow
			_false := false
			visualTrack.DisplayTrackHistory = _false
			visualTrack.DisplayLevelAndSpeed = _false

			visualTrack.UpdateTrack()

			// append the visual track to the map between message & visual tracks
			MapMessageVisualTrack[message] = visualTrack
		}
	}

	nbMessages := len(simulation.gongflyStage.Messages)
	nbMessagesInMap := len(MapMessageVisualTrack)

	if nbMessages < nbMessagesInMap {
		log.Fatal("problem with messages visualisations")
	}

	// if a message has a display value at false, delete the visual track
	for message := range simulation.gongflyStage.Messages {
		if message.Display == false {
			if visualTrack := MapMessageVisualTrack[message]; visualTrack != nil {
				visualTrack.Unstage(simulation.gongleafletStage)
				delete(MapMessageVisualTrack, message)
			}
		}
	}

	// update all visual tracks
	for visualTrack := range simulation.gongleafletStage.VisualTracks {
		visualTrack.UpdateTrack()
	}

	// update all visual centers
	for visualCenter := range simulation.gongleafletStage.Markers {
		visualCenter.UpdateMarker()
	}

	// update all visual lines
	for visualLine := range simulation.gongleafletStage.VLines {
		visualLine.UpdateLine()
	}

	// update all visual circles
	for visualCircle := range simulation.gongleafletStage.Circles {
		visualCircle.UpdateCircle()
	}

	// put all to database
	simulation.gongflyStage.Commit()
	simulation.gongsimStage.Commit()
	simulation.gongleafletStage.Commit()
}

// Reset simulation
func (simulation *Simulation) Reset() {

	// remove all events from agents
	for _, agent := range simulation.engine.Agents() {
		agent.Reset()
	}

	// remove all message and their VisualTracks
	for message := range simulation.gongflyStage.Messages {
		message.Unstage(simulation.gongflyStage)
		if visualTrack := MapMessageVisualTrack[message]; visualTrack != nil {
			visualTrack.Unstage(simulation.gongleafletStage)
			delete(MapMessageVisualTrack, message)
		}
	}

	// set initial conditions
	simulation.setInitialStateVectorOfAgentsAndSimulation()

	// set the engine of all agents
	for _, agent := range simulation.engine.Agents() {
		agent.SetEngine(simulation.engine)
	}

	log.Printf("Simulation reset")
}

// GetLastCommitNb is used by the front repo to computer wether a checkout is necessary
// It does so by comparing successive commit nb
func (simulation *Simulation) GetLastCommitNb() (commitNb uint) {

	if simulation.gongflyStage.BackRepo != nil {
		commitNb = simulation.gongflyStage.BackRepo.GetLastCommitFromBackNb()
	}

	return
}

func (simulation *Simulation) GetLastCommitNbFromFront() (commitNb uint) {

	if simulation.gongflyStage.BackRepo != nil {
		commitNb = simulation.gongflyStage.BackRepo.GetLastPushFromFrontNb()
	}
	return
}

// CheckoutAgents checkout all staged agents to the back repo
func (simulation *Simulation) CheckoutAgents(engine *gongsim_models.Engine) {

	// checkout the simulation agent states
	simulation.gongflyStage.Checkout()
}

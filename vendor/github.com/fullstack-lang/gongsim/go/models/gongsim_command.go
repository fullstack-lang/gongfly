package models

import (
	"log"
	"time"
)

// GongsimCommand is the struct of the instance that is updated by the front for issuing commands
// swagger:model GongsimCommand
type GongsimCommand struct {
	Name             string
	Command          GongsimCommandType
	CommandDate      string
	SpeedCommandType SpeedCommandType
	DateSpeedCommand string
	Engine           *Engine

	stage *StageStruct
}

func NewGongSimCommand(stage *StageStruct, engine *Engine) (gongsimCommand *GongsimCommand) {
	gongsimCommand = &(GongsimCommand{
		Name:        "Gongsim Command Singloton",
		Command:     COMMAND_PAUSE,
		CommandDate: "",
		Engine:      engine,
		stage:       stage,
	})

	gongsimCommand.Stage(stage).SetupGongsimThreads()
	return
}

//
// SetupGongsimThreads enables GongsimCommand to periodicaly watch the GongsimCommand in the Repo
//
// It is set up with three threads:
// - The "command pooler"
// - The "commit scheduler"
// - The "engine driver"

var Quit chan bool

// SetupGongsimThreads schedules gongsim threads
func (gongsimCommand *GongsimCommand) SetupGongsimThreads() *GongsimCommand {

	Quit = make(chan bool)

	go gongsimCommand.commandPooler()
	go gongsimCommand.commitScheduler()
	go gongsimCommand.checkoutScheduler()
	go gongsimCommand.watcher()

	go func() {

		var nextState EngineDriverState
		var nextMode EngineRunMode
		var engineStopMode EngineStopMode

		// when the engine receive a ADVANCE_10_MIN or FIRE_TILL_STATE_CHANGE
		// it can take into account this command only once
		// commandCompletionDate records at which date the command was completed
		var commandCompletionDate string
		for {

			//
			// compute engine driver next state
			//
			// next state is either
			/*

				RESET_SIMULATION
				COMMIT_AGENT_STATES
				CHECKOUT_AGENT_STATES
				FIRE_ONE_EVENT
				SLEEP_100_MS


			*/
			//
			//
			select {
			case <-Quit:
				return
			default:
				nextState = UNKOWN
				_, nextSimTime, _ := gongsimCommand.Engine.GetNextEvent()
				var currentTimePlus10Minute time.Time

				if gongsimCommand.Engine.nextCommitDate.After(gongsimCommand.Engine.lastCommitDate) {
					// log.Printf("Commit agent states scheduled at event # %d and commit nb # %d",
					// 	gongsimCommand.Engine.Fired,
					// 	gongsimCommand.Engine.GetLastCommitNb())
					nextState = COMMIT_AGENT_STATES
				} else {
					if gongsimCommand.Engine.nextCheckoutDate.After(gongsimCommand.Engine.lastCheckoutDate) {
						log.Printf("Checkout agent states scheduled at event # %d and commit nb # %d",
							gongsimCommand.Engine.Fired,
							gongsimCommand.Engine.GetLastCommitNb())
						nextState = CHECKOUT_AGENT_STATES
					} else if nextSimTime.Before(gongsimCommand.Engine.GetEndTime()) {
						switch gongsimCommand.Command {
						case COMMAND_PLAY:
							nextState = FIRE_ONE_EVENT
							nextMode = RELATIVE_SPEED
						case COMMAND_ADVANCE_10_MIN:
							if commandCompletionDate != gongsimCommand.CommandDate {
								nextState = FIRE_ONE_EVENT
								nextMode = FULL_SPEED
								engineStopMode = TEN_MINUTES
								currentTimePlus10Minute = nextSimTime.Add(time.Minute * 10)
								commandCompletionDate = gongsimCommand.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						case COMMAND_FIRE_EVENT_TILL_STATES_CHANGE:
							if commandCompletionDate != gongsimCommand.CommandDate {
								nextState = FIRE_ONE_EVENT
								nextMode = FULL_SPEED
								engineStopMode = STATE_CHANGED
								commandCompletionDate = gongsimCommand.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						case COMMAND_RESET:
							if commandCompletionDate != gongsimCommand.CommandDate {
								nextState = RESET_SIMULATION
								commandCompletionDate = gongsimCommand.CommandDate
							} else {
								nextState = SLEEP_100_MS
							}
						default:
							nextState = SLEEP_100_MS
						}
					} else {
						nextState = SLEEP_100_MS
					}
				}
				if nextState == UNKOWN {
					log.Panicln("Unkown end state for engine driver")
				}

				//
				// perform what is expected in the state
				//
				switch nextState {
				case RESET_SIMULATION:
					// reset all agents
					for _, agent := range gongsimCommand.Engine.Agents() {
						agent.Reset()
					}

					// call specific Reset function
					if gongsimCommand.Engine.Simulation != nil {
						gongsimCommand.Engine.Simulation.Reset(gongsimCommand.Engine)
					}
					// commit the engine state
					gongsimCommand.Engine.Commit(gongsimCommand.stage)
				case COMMIT_AGENT_STATES:
					if gongsimCommand.Engine.Simulation != nil {
						gongsimCommand.Engine.Simulation.CommitAgents(gongsimCommand.Engine)
						gongsimCommand.Engine.lastCommitDate = gongsimCommand.Engine.nextCommitDate
						simulationEventForLastEngineCommit = gongsimCommand.Engine.Fired
						commitFromFrontNbAfterLastEngineCommitOrCheckout = gongsimCommand.Engine.GetLastCommitNbFromFront()
					}
				case CHECKOUT_AGENT_STATES:
					if gongsimCommand.Engine.Simulation != nil {
						gongsimCommand.Engine.Simulation.CheckoutAgents(gongsimCommand.Engine)
						gongsimCommand.Engine.lastCheckoutDate = gongsimCommand.Engine.nextCheckoutDate
						commitFromFrontNbAfterLastEngineCommitOrCheckout = gongsimCommand.Engine.GetLastCommitNbFromFront()
					}
				case FIRE_ONE_EVENT:
					if gongsimCommand.Engine.State != RUNNING {
						gongsimCommand.Engine.State = RUNNING
						gongsimCommand.Engine.Commit(gongsimCommand.stage)
					}

					if nextMode == RELATIVE_SPEED {

						realtimeDurationBetweenHorizons := 500 * time.Millisecond

						gongsimCommand.Engine.nextRealtimeHorizon = time.Now().Add(realtimeDurationBetweenHorizons)
						gongsimCommand.Engine.nextSimulatedTimeHorizon =
							gongsimCommand.Engine.currentTime.Add(time.Duration(int64(gongsimCommand.Engine.Speed * float64(realtimeDurationBetweenHorizons))))

						for nextSimTime.Before(gongsimCommand.Engine.nextSimulatedTimeHorizon) {
							var agent AgentInterface
							agent, nextSimTime, _ = gongsimCommand.Engine.FireNextEvent()

							if agent == nil {
								return
							}
						}
						gongsimCommand.Engine.currentTime = gongsimCommand.Engine.nextSimulatedTimeHorizon

						sleepTime := gongsimCommand.Engine.nextRealtimeHorizon.Sub(time.Now())
						time.Sleep(sleepTime)

						gongsimCommand.Engine.Commit(gongsimCommand.stage)

						// // log.Printf(lastSimTime.String() + " " + nextSimTime.String())
						// if nextSimTime.Sub(gongsimCommand.Engine.GetCurrentTime()) > 0 {
						// 	simTimeAdvance := nextSimTime.Sub(gongsimCommand.Engine.GetCurrentTime())

						// 	sleepDuration := time.Duration(float64(simTimeAdvance) / gongsimCommand.Engine.Speed)
						// 	// log.Printf("total sleep duration " + sleepDuration.String())

						// 	// in order for the end user to see progress in the simulation time
						// 	// the egine time is updated periodicaly
						// 	maxSleepAtATime := time.Duration(500 * time.Millisecond)

						// 	cumulatedSleepTime := time.Duration(0 * time.Millisecond)
						// 	for cumulatedSleepTime < sleepDuration && gongsimCommand.Command == COMMAND_PLAY {
						// 		sleepTime := Min(sleepDuration-cumulatedSleepTime, maxSleepAtATime)
						// 		// log.Printf("Stepped sleep time " + sleepTime.String() + " cumulated " + cumulatedSleepTime.String())
						// 		time.Sleep(sleepTime)
						// 		cumulatedSleepTime += sleepTime

						// 		// update engine current time
						// 		progressInSimulatedTimeInMiliseconds := gongsimCommand.Engine.Speed *
						// 			float64(sleepTime.Milliseconds())
						// 		gongsimCommand.Engine.SetCurrentTime(gongsimCommand.Engine.GetCurrentTime().Add(
						// 			time.Duration(progressInSimulatedTimeInMiliseconds) * time.Millisecond))
						// 		// log.Printf("Engine current time " + gongsimCommand.Engine.CurrentTime.String())
						// 		gongsimCommand.Engine.Commit(gongsimCommand.stage)
						// 	}
						// }
						// _, nextSimTime, _ = gongsimCommand.Engine.FireNextEvent()
					} else { // FULL SPEED
						if engineStopMode == TEN_MINUTES {
							for nextSimTime.Before(currentTimePlus10Minute) {
								var agent AgentInterface
								agent, nextSimTime, _ = gongsimCommand.Engine.FireNextEvent()

								if agent == nil {
									return
								}
							}
						} else if engineStopMode == STATE_CHANGED {
							hasAnyStateHasChanged := false
							for nextSimTime.Before(gongsimCommand.Engine.GetEndTime()) && !hasAnyStateHasChanged {

								var agent AgentInterface
								agent, nextSimTime, _ = gongsimCommand.Engine.FireNextEvent()

								if agent == nil {
									return
								}

								if gongsimCommand.Engine.Simulation != nil {
									hasAnyStateHasChanged = gongsimCommand.Engine.Simulation.HasAnyStateChanged(gongsimCommand.Engine)
								}
							}
						} else {
							log.Panicf("should not reach this piont")
						}

						// time has progressed, therefore an update is necessary
						gongsimCommand.Engine.Commit(gongsimCommand.stage)
					}

				case SLEEP_100_MS:
					if gongsimCommand.Engine.State != PAUSED {
						gongsimCommand.Engine.State = PAUSED
						gongsimCommand.Engine.Commit(gongsimCommand.stage)
					}
					time.Sleep(time.Duration(100 * time.Millisecond))
				default:
				}
			}
		}
	}()

	return gongsimCommand
}

// Min returns the larger of x or y.
func Min(x, y time.Duration) time.Duration {
	if x > y {
		return y
	}
	return x
}

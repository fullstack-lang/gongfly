package models

import (
	"log"
	"time"
)

// commandPooler" is a simple scheduled thread that manages the checkout
// of the current command from the back repo to the stage. If the front client writes a "PLAY" or "PAUSE"
// command, it is written to the back repo and the command pooler retrieves this command to the stage where it
// can be read by the "engine driver". The "command pooler" is a scheduled task
func (gongsimCommand *GongsimCommand) commandPooler() {

	// commandPoolerPeriod is the period of the "command pooler"
	//
	// The period shall not too short for performance sake but not too long because the end user needs a responsive application
	var CommandPoolerPeriod = time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case t := <-CommandPoolerPeriod.C:

			gongsimCommand.Checkout(gongsimCommand.stage)
			if GongsimStatusSingloton.CompletionDate != gongsimCommand.CommandDate {
				log.Println("commandPooler reads new command ", gongsimCommand.Command, "  timestamp ", gongsimCommand.CommandDate, " at ", t)

				GongsimStatusSingloton.CurrentCommand = gongsimCommand.Command
				GongsimStatusSingloton.CompletionDate = gongsimCommand.CommandDate
				GongsimStatusSingloton.Commit(gongsimCommand.stage)
			}
			if GongsimStatusSingloton.SpeedCommandCompletionDate != gongsimCommand.DateSpeedCommand {
				log.Println("commandPooler reads new speed command ", gongsimCommand.SpeedCommandType, "  timestamp ", gongsimCommand.CommandDate, " at ", t)

				switch gongsimCommand.SpeedCommandType {
				case COMMAND_DECREASE_SPEED_50_PERCENTS:
					gongsimCommand.Engine.Speed *= 0.5
					gongsimCommand.Engine.Commit(gongsimCommand.stage)
				case COMMAND_INCREASE_SPEED_100_PERCENTS:
					gongsimCommand.Engine.Speed *= 2.0
					gongsimCommand.Engine.Commit(gongsimCommand.stage)
				}

				GongsimStatusSingloton.CurrentSpeedCommand = gongsimCommand.SpeedCommandType
				GongsimStatusSingloton.SpeedCommandCompletionDate = gongsimCommand.DateSpeedCommand
				GongsimStatusSingloton.Commit(gongsimCommand.stage)
			}
		}
	}
}

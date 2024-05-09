// generated code - do not edit

import { GongsimCommandAPI } from './gongsimcommand-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { Engine } from './engine'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongsimCommand {

	static GONGSTRUCT_NAME = "GongsimCommand"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Command: string = ""
	CommandDate: string = ""
	SpeedCommandType: string = ""
	DateSpeedCommand: string = ""

	// insertion point for pointers and slices of pointers declarations
	Engine?: Engine

}

export function CopyGongsimCommandToGongsimCommandAPI(gongsimcommand: GongsimCommand, gongsimcommandAPI: GongsimCommandAPI) {

	gongsimcommandAPI.CreatedAt = gongsimcommand.CreatedAt
	gongsimcommandAPI.DeletedAt = gongsimcommand.DeletedAt
	gongsimcommandAPI.ID = gongsimcommand.ID

	// insertion point for basic fields copy operations
	gongsimcommandAPI.Name = gongsimcommand.Name
	gongsimcommandAPI.Command = gongsimcommand.Command
	gongsimcommandAPI.CommandDate = gongsimcommand.CommandDate
	gongsimcommandAPI.SpeedCommandType = gongsimcommand.SpeedCommandType
	gongsimcommandAPI.DateSpeedCommand = gongsimcommand.DateSpeedCommand

	// insertion point for pointer fields encoding
	gongsimcommandAPI.GongsimCommandPointersEncoding.EngineID.Valid = true
	if (gongsimcommand.Engine != undefined) {
		gongsimcommandAPI.GongsimCommandPointersEncoding.EngineID.Int64 = gongsimcommand.Engine.ID  
	} else {
		gongsimcommandAPI.GongsimCommandPointersEncoding.EngineID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyGongsimCommandAPIToGongsimCommand update basic, pointers and slice of pointers fields of gongsimcommand
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongsimcommandAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongsimCommandAPIToGongsimCommand(gongsimcommandAPI: GongsimCommandAPI, gongsimcommand: GongsimCommand, frontRepo: FrontRepo) {

	gongsimcommand.CreatedAt = gongsimcommandAPI.CreatedAt
	gongsimcommand.DeletedAt = gongsimcommandAPI.DeletedAt
	gongsimcommand.ID = gongsimcommandAPI.ID

	// insertion point for basic fields copy operations
	gongsimcommand.Name = gongsimcommandAPI.Name
	gongsimcommand.Command = gongsimcommandAPI.Command
	gongsimcommand.CommandDate = gongsimcommandAPI.CommandDate
	gongsimcommand.SpeedCommandType = gongsimcommandAPI.SpeedCommandType
	gongsimcommand.DateSpeedCommand = gongsimcommandAPI.DateSpeedCommand

	// insertion point for pointer fields encoding
	gongsimcommand.Engine = frontRepo.map_ID_Engine.get(gongsimcommandAPI.GongsimCommandPointersEncoding.EngineID.Int64)

	// insertion point for slice of pointers fields encoding
}

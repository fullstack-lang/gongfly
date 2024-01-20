// generated code - do not edit

import { GongsimCommandDB } from './gongsimcommand-db'
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

export function CopyGongsimCommandToGongsimCommandDB(gongsimcommand: GongsimCommand, gongsimcommandDB: GongsimCommandDB) {

	gongsimcommandDB.CreatedAt = gongsimcommand.CreatedAt
	gongsimcommandDB.DeletedAt = gongsimcommand.DeletedAt
	gongsimcommandDB.ID = gongsimcommand.ID
	
	// insertion point for basic fields copy operations
	gongsimcommandDB.Name = gongsimcommand.Name
	gongsimcommandDB.Command = gongsimcommand.Command
	gongsimcommandDB.CommandDate = gongsimcommand.CommandDate
	gongsimcommandDB.SpeedCommandType = gongsimcommand.SpeedCommandType
	gongsimcommandDB.DateSpeedCommand = gongsimcommand.DateSpeedCommand

	// insertion point for pointer fields encoding
    gongsimcommandDB.GongsimCommandPointersEncoding.EngineID.Valid = true
	if (gongsimcommand.Engine != undefined) {
		gongsimcommandDB.GongsimCommandPointersEncoding.EngineID.Int64 = gongsimcommand.Engine.ID  
    } else {
		gongsimcommandDB.GongsimCommandPointersEncoding.EngineID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyGongsimCommandDBToGongsimCommand(gongsimcommandDB: GongsimCommandDB, gongsimcommand: GongsimCommand, frontRepo: FrontRepo) {

	gongsimcommand.CreatedAt = gongsimcommandDB.CreatedAt
	gongsimcommand.DeletedAt = gongsimcommandDB.DeletedAt
	gongsimcommand.ID = gongsimcommandDB.ID
	
	// insertion point for basic fields copy operations
	gongsimcommand.Name = gongsimcommandDB.Name
	gongsimcommand.Command = gongsimcommandDB.Command
	gongsimcommand.CommandDate = gongsimcommandDB.CommandDate
	gongsimcommand.SpeedCommandType = gongsimcommandDB.SpeedCommandType
	gongsimcommand.DateSpeedCommand = gongsimcommandDB.DateSpeedCommand

	// insertion point for pointer fields encoding
	gongsimcommand.Engine = frontRepo.Engines.get(gongsimcommandDB.GongsimCommandPointersEncoding.EngineID.Int64)

	// insertion point for slice of pointers fields encoding
}

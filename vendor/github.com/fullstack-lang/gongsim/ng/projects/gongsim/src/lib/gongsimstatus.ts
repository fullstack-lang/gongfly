// generated code - do not edit

import { GongsimStatusDB } from './gongsimstatus-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongsimStatus {

	static GONGSTRUCT_NAME = "GongsimStatus"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	CurrentCommand: string = ""
	CompletionDate: string = ""
	CurrentSpeedCommand: string = ""
	SpeedCommandCompletionDate: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyGongsimStatusToGongsimStatusDB(gongsimstatus: GongsimStatus, gongsimstatusDB: GongsimStatusDB) {

	gongsimstatusDB.CreatedAt = gongsimstatus.CreatedAt
	gongsimstatusDB.DeletedAt = gongsimstatus.DeletedAt
	gongsimstatusDB.ID = gongsimstatus.ID
	
	// insertion point for basic fields copy operations
	gongsimstatusDB.Name = gongsimstatus.Name
	gongsimstatusDB.CurrentCommand = gongsimstatus.CurrentCommand
	gongsimstatusDB.CompletionDate = gongsimstatus.CompletionDate
	gongsimstatusDB.CurrentSpeedCommand = gongsimstatus.CurrentSpeedCommand
	gongsimstatusDB.SpeedCommandCompletionDate = gongsimstatus.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyGongsimStatusDBToGongsimStatus(gongsimstatusDB: GongsimStatusDB, gongsimstatus: GongsimStatus, frontRepo: FrontRepo) {

	gongsimstatus.CreatedAt = gongsimstatusDB.CreatedAt
	gongsimstatus.DeletedAt = gongsimstatusDB.DeletedAt
	gongsimstatus.ID = gongsimstatusDB.ID
	
	// insertion point for basic fields copy operations
	gongsimstatus.Name = gongsimstatusDB.Name
	gongsimstatus.CurrentCommand = gongsimstatusDB.CurrentCommand
	gongsimstatus.CompletionDate = gongsimstatusDB.CompletionDate
	gongsimstatus.CurrentSpeedCommand = gongsimstatusDB.CurrentSpeedCommand
	gongsimstatus.SpeedCommandCompletionDate = gongsimstatusDB.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

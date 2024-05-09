// generated code - do not edit

import { GongsimStatusAPI } from './gongsimstatus-api'
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

export function CopyGongsimStatusToGongsimStatusAPI(gongsimstatus: GongsimStatus, gongsimstatusAPI: GongsimStatusAPI) {

	gongsimstatusAPI.CreatedAt = gongsimstatus.CreatedAt
	gongsimstatusAPI.DeletedAt = gongsimstatus.DeletedAt
	gongsimstatusAPI.ID = gongsimstatus.ID

	// insertion point for basic fields copy operations
	gongsimstatusAPI.Name = gongsimstatus.Name
	gongsimstatusAPI.CurrentCommand = gongsimstatus.CurrentCommand
	gongsimstatusAPI.CompletionDate = gongsimstatus.CompletionDate
	gongsimstatusAPI.CurrentSpeedCommand = gongsimstatus.CurrentSpeedCommand
	gongsimstatusAPI.SpeedCommandCompletionDate = gongsimstatus.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyGongsimStatusAPIToGongsimStatus update basic, pointers and slice of pointers fields of gongsimstatus
// from respectively the basic fields and encoded fields of pointers and slices of pointers of gongsimstatusAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyGongsimStatusAPIToGongsimStatus(gongsimstatusAPI: GongsimStatusAPI, gongsimstatus: GongsimStatus, frontRepo: FrontRepo) {

	gongsimstatus.CreatedAt = gongsimstatusAPI.CreatedAt
	gongsimstatus.DeletedAt = gongsimstatusAPI.DeletedAt
	gongsimstatus.ID = gongsimstatusAPI.ID

	// insertion point for basic fields copy operations
	gongsimstatus.Name = gongsimstatusAPI.Name
	gongsimstatus.CurrentCommand = gongsimstatusAPI.CurrentCommand
	gongsimstatus.CompletionDate = gongsimstatusAPI.CompletionDate
	gongsimstatus.CurrentSpeedCommand = gongsimstatusAPI.CurrentSpeedCommand
	gongsimstatus.SpeedCommandCompletionDate = gongsimstatusAPI.SpeedCommandCompletionDate

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

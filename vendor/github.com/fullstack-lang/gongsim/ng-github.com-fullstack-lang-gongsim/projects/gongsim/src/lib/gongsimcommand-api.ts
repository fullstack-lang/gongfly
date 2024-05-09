// insertion point for imports
import { EngineAPI } from './engine-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongsimCommandAPI {

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

	// insertion point for other decls

	GongsimCommandPointersEncoding: GongsimCommandPointersEncoding = new GongsimCommandPointersEncoding
}

export class GongsimCommandPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	EngineID: NullInt64 = new NullInt64 // if pointer is null, Engine.ID = 0

}

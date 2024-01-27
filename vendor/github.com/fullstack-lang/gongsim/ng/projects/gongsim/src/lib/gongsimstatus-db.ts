// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class GongsimStatusDB {

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

	// insertion point for other decls

	GongsimStatusPointersEncoding: GongsimStatusPointersEncoding = new GongsimStatusPointersEncoding
}

export class GongsimStatusPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}

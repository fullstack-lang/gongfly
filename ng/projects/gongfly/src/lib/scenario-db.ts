// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ScenarioDB {

	static GONGSTRUCT_NAME = "Scenario"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Lat: number = 0
	Lng: number = 0
	ZoomLevel: number = 0
	MessageVisualSpeed: number = 0

	// insertion point for pointers and slices of pointers declarations

	ScenarioPointersEncoding: ScenarioPointersEncoding = new ScenarioPointersEncoding
}

export class ScenarioPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RadarDB {

	static GONGSTRUCT_NAME = "Radar"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	State: string = ""
	Name: string = ""
	Lat: number = 0
	Lng: number = 0
	Range: number = 0

	// insertion point for pointers and slices of pointers declarations

	RadarPointersEncoding: RadarPointersEncoding = new RadarPointersEncoding
}

export class RadarPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}

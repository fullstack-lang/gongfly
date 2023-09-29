// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SatelliteDB {

	static GONGSTRUCT_NAME = "Satellite"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Line1: string = ""
	Line2: string = ""
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	VerticalSpeed: number = 0
	Timestampstring: string = ""

	// insertion point for other declarations
}

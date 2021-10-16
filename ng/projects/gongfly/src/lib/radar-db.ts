// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RadarDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""
	State: string = ""
	Name: string = ""
	Lat: number = 0
	Lng: number = 0
	Range: number = 0

	// insertion point for other declarations
}

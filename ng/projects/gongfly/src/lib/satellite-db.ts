// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class SatelliteDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	Line1: string = ""
	Line2: string = ""
	TechName: string = ""
	Name: string = ""
	TargetHeading: number = 0
	TargetLocationLat: number = 0
	TargetLocationLng: number = 0
	DistanceToTarget: number = 0
	MaxRotationalSpeed: number = 0
	VerticalSpeed: number = 0
	Timestampstring: string = ""

	// insertion point for other declarations
}

// insertion point for imports
import { OpsLineDB } from './opsline-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LinerDB {

	static GONGSTRUCT_NAME = "Liner"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	State: string = ""
	TargetHeading: number = 0
	TargetLocationLat: number = 0
	TargetLocationLng: number = 0
	DistanceToTarget: number = 0
	MaxRotationalSpeed: number = 0
	VerticalSpeed: number = 0
	Timestampstring: string = ""

	// insertion point for pointers and slices of pointers declarations
	ReporingLine?: OpsLineDB


	LinerPointersEncoding: LinerPointersEncoding = new LinerPointersEncoding
}

export class LinerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	ReporingLineID: NullInt64 = new NullInt64 // if pointer is null, ReporingLine.ID = 0

}

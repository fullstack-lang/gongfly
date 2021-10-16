// insertion point for imports
import { LinerDB } from './liner-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class OrderDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0
	OrderMessage: string = ""
	Number: number = 0
	Type: string = ""
	TargetLat: number = 0
	TargetLng: number = 0

	// insertion point for other declarations
	Duration_string?: string
	Target?: LinerDB
	TargetID: NullInt64 = new NullInt64 // if pointer is null, Target.ID = 0

}

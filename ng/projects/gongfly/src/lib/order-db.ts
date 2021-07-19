// insertion point for imports
import { LinerDB } from './liner-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class OrderDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Duration?: number
	OrderMessage?: string
	Number?: number
	Type?: string
	TargetLat?: number
	TargetLng?: number

	// insertion point for other declarations
	Duration_string?: string
	Target?: LinerDB
	TargetID?: NullInt64

}

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class MovingObjectDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Lat?: number
	Lng?: number
	Heading?: number
	Level?: number
	Speed?: number

	// insertion point for other declarations
}

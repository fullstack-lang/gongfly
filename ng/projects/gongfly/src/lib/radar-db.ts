// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class RadarDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	TechName?: string
	State?: string
	Name?: string
	Lat?: number
	Lng?: number
	Range?: number

	// insertion point for other declarations
}

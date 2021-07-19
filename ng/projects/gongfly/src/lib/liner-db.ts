// insertion point for imports
import { OpsLineDB } from './opsline-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class LinerDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Lat?: number
	Lng?: number
	Heading?: number
	Level?: number
	Speed?: number
	TechName?: string
	State?: string
	Name?: string
	TargetHeading?: number
	TargetLocationLat?: number
	TargetLocationLng?: number
	DistanceToTarget?: number
	MaxRotationalSpeed?: number
	VerticalSpeed?: number
	Timestampstring?: string

	// insertion point for other declarations
	ReporingLine?: OpsLineDB
	ReporingLineID?: NullInt64

}

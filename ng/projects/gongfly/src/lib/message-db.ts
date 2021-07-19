// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class MessageDB {
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
	TargetLocationLat?: number
	TargetLocationLng?: number
	DistanceToTarget?: number
	Timestampstring?: string
	DurationSinceSimulationStart?: number
	Timestampstartstring?: string
	Source?: string
	Destination?: string
	Content?: string
	About_string?: string
	Display?: string

	// insertion point for other declarations
	DurationSinceSimulationStart_string?: string
}

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MessageDB {

	static GONGSTRUCT_NAME = "Message"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	State: string = ""
	Name: string = ""
	TargetLocationLat: number = 0
	TargetLocationLng: number = 0
	DistanceToTarget: number = 0
	Timestampstring: string = ""
	DurationSinceSimulationStart: number = 0
	Timestampstartstring: string = ""
	Source: string = ""
	Destination: string = ""
	Content: string = ""
	About_string: string = ""
	Display: boolean = false

	// insertion point for other declarations
	DurationSinceSimulationStart_string?: string
}

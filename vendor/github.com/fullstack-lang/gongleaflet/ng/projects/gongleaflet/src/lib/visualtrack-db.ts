// insertion point for imports
import { VisualLayerDB } from './visuallayer-db'
import { VisualIconDB } from './visualicon-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VisualTrackDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	VerticalSpeed: number = 0
	Name: string = ""
	VisualColorEnum: string = ""
	Display: boolean = false
	DisplayTrackHistory: boolean = false
	DisplayLevelAndSpeed: boolean = false

	// insertion point for other declarations
	VisualLayer?: VisualLayerDB
	VisualLayerID: NullInt64 = new NullInt64 // if pointer is null, VisualLayer.ID = 0

	VisualIcon?: VisualIconDB
	VisualIconID: NullInt64 = new NullInt64 // if pointer is null, VisualIcon.ID = 0

}

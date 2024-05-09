// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'
import { DivIconAPI } from './divicon-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VisualTrackAPI {

	static GONGSTRUCT_NAME = "VisualTrack"

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
	ColorEnum: string = ""
	DisplayTrackHistory: boolean = false
	DisplayLevelAndSpeed: boolean = false

	// insertion point for other decls

	VisualTrackPointersEncoding: VisualTrackPointersEncoding = new VisualTrackPointersEncoding
}

export class VisualTrackPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupID: NullInt64 = new NullInt64 // if pointer is null, LayerGroup.ID = 0

	DivIconID: NullInt64 = new NullInt64 // if pointer is null, DivIcon.ID = 0

}

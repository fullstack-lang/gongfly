// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'
import { DivIconAPI } from './divicon-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkerAPI {

	static GONGSTRUCT_NAME = "Marker"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	ColorEnum: string = ""

	// insertion point for other decls

	MarkerPointersEncoding: MarkerPointersEncoding = new MarkerPointersEncoding
}

export class MarkerPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupID: NullInt64 = new NullInt64 // if pointer is null, LayerGroup.ID = 0

	DivIconID: NullInt64 = new NullInt64 // if pointer is null, DivIcon.ID = 0

}

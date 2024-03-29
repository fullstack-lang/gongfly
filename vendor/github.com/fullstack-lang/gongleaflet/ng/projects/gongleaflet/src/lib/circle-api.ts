// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CircleAPI {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	Radius: number = 0
	ColorEnum: string = ""
	DashStyleEnum: string = ""

	// insertion point for other decls

	CirclePointersEncoding: CirclePointersEncoding = new CirclePointersEncoding
}

export class CirclePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupID: NullInt64 = new NullInt64 // if pointer is null, LayerGroup.ID = 0

}

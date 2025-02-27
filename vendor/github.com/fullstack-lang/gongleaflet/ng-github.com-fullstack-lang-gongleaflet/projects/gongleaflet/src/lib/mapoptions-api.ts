// insertion point for imports
import { LayerGroupUseAPI } from './layergroupuse-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MapOptionsAPI {

	static GONGSTRUCT_NAME = "MapOptions"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	ZoomLevel: number = 0
	UrlTemplate: string = ""
	Attribution: string = ""
	MaxZoom: number = 0
	ZoomControl: boolean = false
	AttributionControl: boolean = false
	ZoomSnap: number = 0

	// insertion point for other decls

	MapOptionsPointersEncoding: MapOptionsPointersEncoding = new MapOptionsPointersEncoding
}

export class MapOptionsPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupUses: number[] = []
}

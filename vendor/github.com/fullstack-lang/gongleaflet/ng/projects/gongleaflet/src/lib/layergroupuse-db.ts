// insertion point for imports
import { LayerGroupDB } from './layergroup-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerGroupUseDB {

	static GONGSTRUCT_NAME = "LayerGroupUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroupDB


	LayerGroupUsePointersEncoding: LayerGroupUsePointersEncoding = new LayerGroupUsePointersEncoding
}

export class LayerGroupUsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupID: NullInt64 = new NullInt64 // if pointer is null, LayerGroup.ID = 0

}

// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerGroupUseAPI {

	static GONGSTRUCT_NAME = "LayerGroupUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for other decls

	LayerGroupUsePointersEncoding: LayerGroupUsePointersEncoding = new LayerGroupUsePointersEncoding
}

export class LayerGroupUsePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	LayerGroupID: NullInt64 = new NullInt64 // if pointer is null, LayerGroup.ID = 0

}

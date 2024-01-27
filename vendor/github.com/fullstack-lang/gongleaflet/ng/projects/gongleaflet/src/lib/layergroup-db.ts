// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerGroupDB {

	static GONGSTRUCT_NAME = "LayerGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DisplayName: string = ""

	// insertion point for other decls

	LayerGroupPointersEncoding: LayerGroupPointersEncoding = new LayerGroupPointersEncoding
}

export class LayerGroupPointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
}

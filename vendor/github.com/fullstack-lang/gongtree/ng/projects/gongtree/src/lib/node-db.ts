// insertion point for imports
import { ButtonDB } from './button-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NodeDB {

	static GONGSTRUCT_NAME = "Node"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	BackgroundColor: string = ""
	IsExpanded: boolean = false
	HasCheckboxButton: boolean = false
	IsChecked: boolean = false
	IsCheckboxDisabled: boolean = false
	IsInEditMode: boolean = false
	IsNodeClickable: boolean = false

	// insertion point for pointers and slices of pointers declarations
	Children: Array<NodeDB> = []
	Buttons: Array<ButtonDB> = []

	NodePointersEncoding: NodePointersEncoding = new NodePointersEncoding
}

export class NodePointersEncoding {
	// insertion point for pointers and slices of pointers encoding fields
	Children: number[] = []
	Buttons: number[] = []
}
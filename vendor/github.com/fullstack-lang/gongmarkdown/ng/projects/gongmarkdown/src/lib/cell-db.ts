// insertion point for imports
import { RowDB } from './row-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class CellDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Row_CellsDBID: NullInt64 = new NullInt64
	Row_CellsDBID_Index: NullInt64  = new NullInt64 // store the index of the cell instance in Row.Cells
	Row_Cells_reverse?: RowDB 

}

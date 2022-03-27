// insertion point for imports
import { CellDB } from './cell-db'
import { ElementDB } from './element-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class RowDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""

	// insertion point for other declarations
	Cells?: Array<CellDB>
	Element_RowsDBID: NullInt64 = new NullInt64
	Element_RowsDBID_Index: NullInt64  = new NullInt64 // store the index of the row instance in Element.Rows
	Element_Rows_reverse?: ElementDB 

}

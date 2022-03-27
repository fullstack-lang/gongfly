// insertion point for imports
import { ElementDB } from './element-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MarkdownContentDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Content: string = ""

	// insertion point for other declarations
	Root?: ElementDB
	RootID: NullInt64 = new NullInt64 // if pointer is null, Root.ID = 0

}

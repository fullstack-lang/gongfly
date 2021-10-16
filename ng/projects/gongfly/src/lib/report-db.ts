// insertion point for imports
import { LinerDB } from './liner-db'
import { OpsLineDB } from './opsline-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class ReportDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0
	ReportMessage: string = ""
	Number: number = 0
	Type: string = ""

	// insertion point for other declarations
	Duration_string?: string
	About?: LinerDB
	AboutID: NullInt64 = new NullInt64 // if pointer is null, About.ID = 0

	OpsLine?: OpsLineDB
	OpsLineID: NullInt64 = new NullInt64 // if pointer is null, OpsLine.ID = 0

}

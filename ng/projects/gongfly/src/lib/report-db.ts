// insertion point for imports
import { LinerDB } from './liner-db'
import { OpsLineDB } from './opsline-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class ReportDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	Name?: string
	Duration?: number
	ReportMessage?: string
	Number?: number
	Type?: string

	// insertion point for other declarations
	Duration_string?: string
	About?: LinerDB
	AboutID?: NullInt64

	OpsLine?: OpsLineDB
	OpsLineID?: NullInt64

}

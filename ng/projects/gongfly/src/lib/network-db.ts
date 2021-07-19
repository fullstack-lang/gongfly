// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class NetworkDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	IsTransmitting?: string
	TransmissionMessage?: string
	IsTransmittingBackward?: string
	TransmissionMessageBackward?: string

	// insertion point for other declarations
}

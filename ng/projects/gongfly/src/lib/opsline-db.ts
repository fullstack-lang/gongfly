// insertion point for imports
import { ScenarioDB } from './scenario-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './front-repo.service'

export class OpsLineDB {
	CreatedAt?: string;
	DeletedAt?: string;
	ID?: number;

	// insertion point for basic fields declarations
	IsTransmitting?: string
	TransmissionMessage?: string
	IsTransmittingBackward?: string
	TransmissionMessageBackward?: string
	TechName?: string
	State?: string
	Name?: string

	// insertion point for other declarations
	Scenario?: ScenarioDB
	ScenarioID?: NullInt64

}

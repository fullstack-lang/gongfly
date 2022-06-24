// insertion point for imports
import { ScenarioDB } from './scenario-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class OpsLineDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	IsTransmitting: boolean = false
	TransmissionMessage: string = ""
	IsTransmittingBackward: boolean = false
	TransmissionMessageBackward: string = ""
	State: string = ""
	Name: string = ""

	// insertion point for other declarations
	Scenario?: ScenarioDB
	ScenarioID: NullInt64 = new NullInt64 // if pointer is null, Scenario.ID = 0

}

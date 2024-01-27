// generated code - do not edit

import { DummyAgentDB } from './dummyagent-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DummyAgent {

	static GONGSTRUCT_NAME = "DummyAgent"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	TechName: string = ""
	Name: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDummyAgentToDummyAgentDB(dummyagent: DummyAgent, dummyagentDB: DummyAgentDB) {

	dummyagentDB.CreatedAt = dummyagent.CreatedAt
	dummyagentDB.DeletedAt = dummyagent.DeletedAt
	dummyagentDB.ID = dummyagent.ID

	// insertion point for basic fields copy operations
	dummyagentDB.TechName = dummyagent.TechName
	dummyagentDB.Name = dummyagent.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDummyAgentDBToDummyAgent update basic, pointers and slice of pointers fields of dummyagent
// from respectively the basic fields and encoded fields of pointers and slices of pointers of dummyagentDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDummyAgentDBToDummyAgent(dummyagentDB: DummyAgentDB, dummyagent: DummyAgent, frontRepo: FrontRepo) {

	dummyagent.CreatedAt = dummyagentDB.CreatedAt
	dummyagent.DeletedAt = dummyagentDB.DeletedAt
	dummyagent.ID = dummyagentDB.ID

	// insertion point for basic fields copy operations
	dummyagent.TechName = dummyagentDB.TechName
	dummyagent.Name = dummyagentDB.Name

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

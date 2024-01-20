// generated code - do not edit

import { EngineDB } from './engine-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Engine {

	static GONGSTRUCT_NAME = "Engine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	EndTime: string = ""
	CurrentTime: string = ""
	SecondsSinceStart: number = 0
	Fired: number = 0
	ControlMode: string = ""
	State: string = ""
	Speed: number = 0

	// insertion point for pointers and slices of pointers declarations
}

export function CopyEngineToEngineDB(engine: Engine, engineDB: EngineDB) {

	engineDB.CreatedAt = engine.CreatedAt
	engineDB.DeletedAt = engine.DeletedAt
	engineDB.ID = engine.ID

	// insertion point for basic fields copy operations
	engineDB.Name = engine.Name
	engineDB.EndTime = engine.EndTime
	engineDB.CurrentTime = engine.CurrentTime
	engineDB.SecondsSinceStart = engine.SecondsSinceStart
	engineDB.Fired = engine.Fired
	engineDB.ControlMode = engine.ControlMode
	engineDB.State = engine.State
	engineDB.Speed = engine.Speed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

export function CopyEngineDBToEngine(engineDB: EngineDB, engine: Engine, frontRepo: FrontRepo) {

	engine.CreatedAt = engineDB.CreatedAt
	engine.DeletedAt = engineDB.DeletedAt
	engine.ID = engineDB.ID

	// insertion point for basic fields copy operations
	engine.Name = engineDB.Name
	engine.EndTime = engineDB.EndTime
	engine.CurrentTime = engineDB.CurrentTime
	engine.SecondsSinceStart = engineDB.SecondsSinceStart
	engine.Fired = engineDB.Fired
	engine.ControlMode = engineDB.ControlMode
	engine.State = engineDB.State
	engine.Speed = engineDB.Speed

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

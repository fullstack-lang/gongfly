// generated code - do not edit

import { EventDB } from './event-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Event {

	static GONGSTRUCT_NAME = "Event"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Duration: number = 0

	// insertion point for pointers and slices of pointers declarations
	Duration_string?: string
}

export function CopyEventToEventDB(event: Event, eventDB: EventDB) {

	eventDB.CreatedAt = event.CreatedAt
	eventDB.DeletedAt = event.DeletedAt
	eventDB.ID = event.ID

	// insertion point for basic fields copy operations
	eventDB.Name = event.Name
	eventDB.Duration = event.Duration

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyEventDBToEvent update basic, pointers and slice of pointers fields of event
// from respectively the basic fields and encoded fields of pointers and slices of pointers of eventDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyEventDBToEvent(eventDB: EventDB, event: Event, frontRepo: FrontRepo) {

	event.CreatedAt = eventDB.CreatedAt
	event.DeletedAt = eventDB.DeletedAt
	event.ID = eventDB.ID

	// insertion point for basic fields copy operations
	event.Name = eventDB.Name
	event.Duration = eventDB.Duration

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

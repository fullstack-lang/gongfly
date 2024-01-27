// generated code - do not edit

import { UserClickDB } from './userclick-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class UserClick {

	static GONGSTRUCT_NAME = "UserClick"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	Lat: number = 0
	Lng: number = 0
	TimeOfClick: Date = new Date

	// insertion point for pointers and slices of pointers declarations
}

export function CopyUserClickToUserClickDB(userclick: UserClick, userclickDB: UserClickDB) {

	userclickDB.CreatedAt = userclick.CreatedAt
	userclickDB.DeletedAt = userclick.DeletedAt
	userclickDB.ID = userclick.ID

	// insertion point for basic fields copy operations
	userclickDB.Name = userclick.Name
	userclickDB.Lat = userclick.Lat
	userclickDB.Lng = userclick.Lng

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyUserClickDBToUserClick update basic, pointers and slice of pointers fields of userclick
// from respectively the basic fields and encoded fields of pointers and slices of pointers of userclickDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyUserClickDBToUserClick(userclickDB: UserClickDB, userclick: UserClick, frontRepo: FrontRepo) {

	userclick.CreatedAt = userclickDB.CreatedAt
	userclick.DeletedAt = userclickDB.DeletedAt
	userclick.ID = userclickDB.ID

	// insertion point for basic fields copy operations
	userclick.Name = userclickDB.Name
	userclick.Lat = userclickDB.Lat
	userclick.Lng = userclickDB.Lng

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

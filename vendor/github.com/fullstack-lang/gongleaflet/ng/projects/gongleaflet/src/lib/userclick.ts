// generated code - do not edit

import { UserClickAPI } from './userclick-api'
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

export function CopyUserClickToUserClickAPI(userclick: UserClick, userclickAPI: UserClickAPI) {

	userclickAPI.CreatedAt = userclick.CreatedAt
	userclickAPI.DeletedAt = userclick.DeletedAt
	userclickAPI.ID = userclick.ID

	// insertion point for basic fields copy operations
	userclickAPI.Name = userclick.Name
	userclickAPI.Lat = userclick.Lat
	userclickAPI.Lng = userclick.Lng

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyUserClickAPIToUserClick update basic, pointers and slice of pointers fields of userclick
// from respectively the basic fields and encoded fields of pointers and slices of pointers of userclickAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyUserClickAPIToUserClick(userclickAPI: UserClickAPI, userclick: UserClick, frontRepo: FrontRepo) {

	userclick.CreatedAt = userclickAPI.CreatedAt
	userclick.DeletedAt = userclickAPI.DeletedAt
	userclick.ID = userclickAPI.ID

	// insertion point for basic fields copy operations
	userclick.Name = userclickAPI.Name
	userclick.Lat = userclickAPI.Lat
	userclick.Lng = userclickAPI.Lng

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

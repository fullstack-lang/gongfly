// generated code - do not edit

import { DivIconAPI } from './divicon-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class DivIcon {

	static GONGSTRUCT_NAME = "DivIcon"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	SVG: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyDivIconToDivIconAPI(divicon: DivIcon, diviconAPI: DivIconAPI) {

	diviconAPI.CreatedAt = divicon.CreatedAt
	diviconAPI.DeletedAt = divicon.DeletedAt
	diviconAPI.ID = divicon.ID

	// insertion point for basic fields copy operations
	diviconAPI.Name = divicon.Name
	diviconAPI.SVG = divicon.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDivIconAPIToDivIcon update basic, pointers and slice of pointers fields of divicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of diviconAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDivIconAPIToDivIcon(diviconAPI: DivIconAPI, divicon: DivIcon, frontRepo: FrontRepo) {

	divicon.CreatedAt = diviconAPI.CreatedAt
	divicon.DeletedAt = diviconAPI.DeletedAt
	divicon.ID = diviconAPI.ID

	// insertion point for basic fields copy operations
	divicon.Name = diviconAPI.Name
	divicon.SVG = diviconAPI.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

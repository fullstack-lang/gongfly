// generated code - do not edit

import { DivIconDB } from './divicon-db'
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

export function CopyDivIconToDivIconDB(divicon: DivIcon, diviconDB: DivIconDB) {

	diviconDB.CreatedAt = divicon.CreatedAt
	diviconDB.DeletedAt = divicon.DeletedAt
	diviconDB.ID = divicon.ID

	// insertion point for basic fields copy operations
	diviconDB.Name = divicon.Name
	diviconDB.SVG = divicon.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyDivIconDBToDivIcon update basic, pointers and slice of pointers fields of divicon
// from respectively the basic fields and encoded fields of pointers and slices of pointers of diviconDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyDivIconDBToDivIcon(diviconDB: DivIconDB, divicon: DivIcon, frontRepo: FrontRepo) {

	divicon.CreatedAt = diviconDB.CreatedAt
	divicon.DeletedAt = diviconDB.DeletedAt
	divicon.ID = diviconDB.ID

	// insertion point for basic fields copy operations
	divicon.Name = diviconDB.Name
	divicon.SVG = diviconDB.SVG

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}
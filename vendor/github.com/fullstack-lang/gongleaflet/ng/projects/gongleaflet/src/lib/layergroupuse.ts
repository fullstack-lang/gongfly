// generated code - do not edit

import { LayerGroupUseDB } from './layergroupuse-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroup } from './layergroup'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerGroupUse {

	static GONGSTRUCT_NAME = "LayerGroupUse"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsDisplayed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroup

}

export function CopyLayerGroupUseToLayerGroupUseDB(layergroupuse: LayerGroupUse, layergroupuseDB: LayerGroupUseDB) {

	layergroupuseDB.CreatedAt = layergroupuse.CreatedAt
	layergroupuseDB.DeletedAt = layergroupuse.DeletedAt
	layergroupuseDB.ID = layergroupuse.ID

	// insertion point for basic fields copy operations
	layergroupuseDB.Name = layergroupuse.Name
	layergroupuseDB.IsDisplayed = layergroupuse.IsDisplayed

	// insertion point for pointer fields encoding
	layergroupuseDB.LayerGroupUsePointersEncoding.LayerGroupID.Valid = true
	if (layergroupuse.LayerGroup != undefined) {
		layergroupuseDB.LayerGroupUsePointersEncoding.LayerGroupID.Int64 = layergroupuse.LayerGroup.ID  
	} else {
		layergroupuseDB.LayerGroupUsePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyLayerGroupUseDBToLayerGroupUse update basic, pointers and slice of pointers fields of layergroupuse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layergroupuseDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerGroupUseDBToLayerGroupUse(layergroupuseDB: LayerGroupUseDB, layergroupuse: LayerGroupUse, frontRepo: FrontRepo) {

	layergroupuse.CreatedAt = layergroupuseDB.CreatedAt
	layergroupuse.DeletedAt = layergroupuseDB.DeletedAt
	layergroupuse.ID = layergroupuseDB.ID

	// insertion point for basic fields copy operations
	layergroupuse.Name = layergroupuseDB.Name
	layergroupuse.IsDisplayed = layergroupuseDB.IsDisplayed

	// insertion point for pointer fields encoding
	layergroupuse.LayerGroup = frontRepo.map_ID_LayerGroup.get(layergroupuseDB.LayerGroupUsePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

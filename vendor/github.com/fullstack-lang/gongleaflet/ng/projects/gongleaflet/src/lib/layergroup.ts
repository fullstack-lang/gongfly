// generated code - do not edit

import { LayerGroupDB } from './layergroup-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class LayerGroup {

	static GONGSTRUCT_NAME = "LayerGroup"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	DisplayName: string = ""

	// insertion point for pointers and slices of pointers declarations
}

export function CopyLayerGroupToLayerGroupDB(layergroup: LayerGroup, layergroupDB: LayerGroupDB) {

	layergroupDB.CreatedAt = layergroup.CreatedAt
	layergroupDB.DeletedAt = layergroup.DeletedAt
	layergroupDB.ID = layergroup.ID

	// insertion point for basic fields copy operations
	layergroupDB.Name = layergroup.Name
	layergroupDB.DisplayName = layergroup.DisplayName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLayerGroupDBToLayerGroup update basic, pointers and slice of pointers fields of layergroup
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layergroupDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerGroupDBToLayerGroup(layergroupDB: LayerGroupDB, layergroup: LayerGroup, frontRepo: FrontRepo) {

	layergroup.CreatedAt = layergroupDB.CreatedAt
	layergroup.DeletedAt = layergroupDB.DeletedAt
	layergroup.ID = layergroupDB.ID

	// insertion point for basic fields copy operations
	layergroup.Name = layergroupDB.Name
	layergroup.DisplayName = layergroupDB.DisplayName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

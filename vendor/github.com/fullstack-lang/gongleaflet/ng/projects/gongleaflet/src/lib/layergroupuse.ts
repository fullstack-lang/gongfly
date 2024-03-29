// generated code - do not edit

import { LayerGroupUseAPI } from './layergroupuse-api'
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

export function CopyLayerGroupUseToLayerGroupUseAPI(layergroupuse: LayerGroupUse, layergroupuseAPI: LayerGroupUseAPI) {

	layergroupuseAPI.CreatedAt = layergroupuse.CreatedAt
	layergroupuseAPI.DeletedAt = layergroupuse.DeletedAt
	layergroupuseAPI.ID = layergroupuse.ID

	// insertion point for basic fields copy operations
	layergroupuseAPI.Name = layergroupuse.Name
	layergroupuseAPI.IsDisplayed = layergroupuse.IsDisplayed

	// insertion point for pointer fields encoding
	layergroupuseAPI.LayerGroupUsePointersEncoding.LayerGroupID.Valid = true
	if (layergroupuse.LayerGroup != undefined) {
		layergroupuseAPI.LayerGroupUsePointersEncoding.LayerGroupID.Int64 = layergroupuse.LayerGroup.ID  
	} else {
		layergroupuseAPI.LayerGroupUsePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyLayerGroupUseAPIToLayerGroupUse update basic, pointers and slice of pointers fields of layergroupuse
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layergroupuseAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerGroupUseAPIToLayerGroupUse(layergroupuseAPI: LayerGroupUseAPI, layergroupuse: LayerGroupUse, frontRepo: FrontRepo) {

	layergroupuse.CreatedAt = layergroupuseAPI.CreatedAt
	layergroupuse.DeletedAt = layergroupuseAPI.DeletedAt
	layergroupuse.ID = layergroupuseAPI.ID

	// insertion point for basic fields copy operations
	layergroupuse.Name = layergroupuseAPI.Name
	layergroupuse.IsDisplayed = layergroupuseAPI.IsDisplayed

	// insertion point for pointer fields encoding
	layergroupuse.LayerGroup = frontRepo.map_ID_LayerGroup.get(layergroupuseAPI.LayerGroupUsePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

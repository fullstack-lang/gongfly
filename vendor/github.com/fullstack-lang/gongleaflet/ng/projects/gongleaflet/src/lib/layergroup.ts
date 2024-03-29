// generated code - do not edit

import { LayerGroupAPI } from './layergroup-api'
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

export function CopyLayerGroupToLayerGroupAPI(layergroup: LayerGroup, layergroupAPI: LayerGroupAPI) {

	layergroupAPI.CreatedAt = layergroup.CreatedAt
	layergroupAPI.DeletedAt = layergroup.DeletedAt
	layergroupAPI.ID = layergroup.ID

	// insertion point for basic fields copy operations
	layergroupAPI.Name = layergroup.Name
	layergroupAPI.DisplayName = layergroup.DisplayName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

// CopyLayerGroupAPIToLayerGroup update basic, pointers and slice of pointers fields of layergroup
// from respectively the basic fields and encoded fields of pointers and slices of pointers of layergroupAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyLayerGroupAPIToLayerGroup(layergroupAPI: LayerGroupAPI, layergroup: LayerGroup, frontRepo: FrontRepo) {

	layergroup.CreatedAt = layergroupAPI.CreatedAt
	layergroup.DeletedAt = layergroupAPI.DeletedAt
	layergroup.ID = layergroupAPI.ID

	// insertion point for basic fields copy operations
	layergroup.Name = layergroupAPI.Name
	layergroup.DisplayName = layergroupAPI.DisplayName

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
}

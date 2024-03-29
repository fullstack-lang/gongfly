// generated code - do not edit

import { VLineAPI } from './vline-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroup } from './layergroup'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VLine {

	static GONGSTRUCT_NAME = "VLine"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	StartLat: number = 0
	StartLng: number = 0
	EndLat: number = 0
	EndLng: number = 0
	Name: string = ""
	ColorEnum: string = ""
	DashStyleEnum: string = ""
	IsTransmitting: string = ""
	Message: string = ""
	IsTransmittingBackward: string = ""
	MessageBackward: string = ""

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroup

}

export function CopyVLineToVLineAPI(vline: VLine, vlineAPI: VLineAPI) {

	vlineAPI.CreatedAt = vline.CreatedAt
	vlineAPI.DeletedAt = vline.DeletedAt
	vlineAPI.ID = vline.ID

	// insertion point for basic fields copy operations
	vlineAPI.StartLat = vline.StartLat
	vlineAPI.StartLng = vline.StartLng
	vlineAPI.EndLat = vline.EndLat
	vlineAPI.EndLng = vline.EndLng
	vlineAPI.Name = vline.Name
	vlineAPI.ColorEnum = vline.ColorEnum
	vlineAPI.DashStyleEnum = vline.DashStyleEnum
	vlineAPI.IsTransmitting = vline.IsTransmitting
	vlineAPI.Message = vline.Message
	vlineAPI.IsTransmittingBackward = vline.IsTransmittingBackward
	vlineAPI.MessageBackward = vline.MessageBackward

	// insertion point for pointer fields encoding
	vlineAPI.VLinePointersEncoding.LayerGroupID.Valid = true
	if (vline.LayerGroup != undefined) {
		vlineAPI.VLinePointersEncoding.LayerGroupID.Int64 = vline.LayerGroup.ID  
	} else {
		vlineAPI.VLinePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyVLineAPIToVLine update basic, pointers and slice of pointers fields of vline
// from respectively the basic fields and encoded fields of pointers and slices of pointers of vlineAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVLineAPIToVLine(vlineAPI: VLineAPI, vline: VLine, frontRepo: FrontRepo) {

	vline.CreatedAt = vlineAPI.CreatedAt
	vline.DeletedAt = vlineAPI.DeletedAt
	vline.ID = vlineAPI.ID

	// insertion point for basic fields copy operations
	vline.StartLat = vlineAPI.StartLat
	vline.StartLng = vlineAPI.StartLng
	vline.EndLat = vlineAPI.EndLat
	vline.EndLng = vlineAPI.EndLng
	vline.Name = vlineAPI.Name
	vline.ColorEnum = vlineAPI.ColorEnum
	vline.DashStyleEnum = vlineAPI.DashStyleEnum
	vline.IsTransmitting = vlineAPI.IsTransmitting
	vline.Message = vlineAPI.Message
	vline.IsTransmittingBackward = vlineAPI.IsTransmittingBackward
	vline.MessageBackward = vlineAPI.MessageBackward

	// insertion point for pointer fields encoding
	vline.LayerGroup = frontRepo.map_ID_LayerGroup.get(vlineAPI.VLinePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

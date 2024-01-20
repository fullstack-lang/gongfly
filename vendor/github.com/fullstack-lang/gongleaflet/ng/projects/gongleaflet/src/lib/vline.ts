// generated code - do not edit

import { VLineDB } from './vline-db'
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

export function CopyVLineToVLineDB(vline: VLine, vlineDB: VLineDB) {

	vlineDB.CreatedAt = vline.CreatedAt
	vlineDB.DeletedAt = vline.DeletedAt
	vlineDB.ID = vline.ID
	
	// insertion point for basic fields copy operations
	vlineDB.StartLat = vline.StartLat
	vlineDB.StartLng = vline.StartLng
	vlineDB.EndLat = vline.EndLat
	vlineDB.EndLng = vline.EndLng
	vlineDB.Name = vline.Name
	vlineDB.ColorEnum = vline.ColorEnum
	vlineDB.DashStyleEnum = vline.DashStyleEnum
	vlineDB.IsTransmitting = vline.IsTransmitting
	vlineDB.Message = vline.Message
	vlineDB.IsTransmittingBackward = vline.IsTransmittingBackward
	vlineDB.MessageBackward = vline.MessageBackward

	// insertion point for pointer fields encoding
    vlineDB.VLinePointersEncoding.LayerGroupID.Valid = true
	if (vline.LayerGroup != undefined) {
		vlineDB.VLinePointersEncoding.LayerGroupID.Int64 = vline.LayerGroup.ID  
    } else {
		vlineDB.VLinePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

export function CopyVLineDBToVLine(vlineDB: VLineDB, vline: VLine, frontRepo: FrontRepo) {

	vline.CreatedAt = vlineDB.CreatedAt
	vline.DeletedAt = vlineDB.DeletedAt
	vline.ID = vlineDB.ID
	
	// insertion point for basic fields copy operations
	vline.StartLat = vlineDB.StartLat
	vline.StartLng = vlineDB.StartLng
	vline.EndLat = vlineDB.EndLat
	vline.EndLng = vlineDB.EndLng
	vline.Name = vlineDB.Name
	vline.ColorEnum = vlineDB.ColorEnum
	vline.DashStyleEnum = vlineDB.DashStyleEnum
	vline.IsTransmitting = vlineDB.IsTransmitting
	vline.Message = vlineDB.Message
	vline.IsTransmittingBackward = vlineDB.IsTransmittingBackward
	vline.MessageBackward = vlineDB.MessageBackward

	// insertion point for pointer fields encoding
	vline.LayerGroup = frontRepo.LayerGroups.get(vlineDB.VLinePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

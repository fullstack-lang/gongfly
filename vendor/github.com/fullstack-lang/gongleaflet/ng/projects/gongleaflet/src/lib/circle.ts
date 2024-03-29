// generated code - do not edit

import { CircleAPI } from './circle-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroup } from './layergroup'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Circle {

	static GONGSTRUCT_NAME = "Circle"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	Radius: number = 0
	ColorEnum: string = ""
	DashStyleEnum: string = ""

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroup

}

export function CopyCircleToCircleAPI(circle: Circle, circleAPI: CircleAPI) {

	circleAPI.CreatedAt = circle.CreatedAt
	circleAPI.DeletedAt = circle.DeletedAt
	circleAPI.ID = circle.ID

	// insertion point for basic fields copy operations
	circleAPI.Lat = circle.Lat
	circleAPI.Lng = circle.Lng
	circleAPI.Name = circle.Name
	circleAPI.Radius = circle.Radius
	circleAPI.ColorEnum = circle.ColorEnum
	circleAPI.DashStyleEnum = circle.DashStyleEnum

	// insertion point for pointer fields encoding
	circleAPI.CirclePointersEncoding.LayerGroupID.Valid = true
	if (circle.LayerGroup != undefined) {
		circleAPI.CirclePointersEncoding.LayerGroupID.Int64 = circle.LayerGroup.ID  
	} else {
		circleAPI.CirclePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyCircleAPIToCircle update basic, pointers and slice of pointers fields of circle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circleAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleAPIToCircle(circleAPI: CircleAPI, circle: Circle, frontRepo: FrontRepo) {

	circle.CreatedAt = circleAPI.CreatedAt
	circle.DeletedAt = circleAPI.DeletedAt
	circle.ID = circleAPI.ID

	// insertion point for basic fields copy operations
	circle.Lat = circleAPI.Lat
	circle.Lng = circleAPI.Lng
	circle.Name = circleAPI.Name
	circle.Radius = circleAPI.Radius
	circle.ColorEnum = circleAPI.ColorEnum
	circle.DashStyleEnum = circleAPI.DashStyleEnum

	// insertion point for pointer fields encoding
	circle.LayerGroup = frontRepo.map_ID_LayerGroup.get(circleAPI.CirclePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

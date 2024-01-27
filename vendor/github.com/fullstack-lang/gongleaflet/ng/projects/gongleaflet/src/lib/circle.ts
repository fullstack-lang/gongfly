// generated code - do not edit

import { CircleDB } from './circle-db'
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

export function CopyCircleToCircleDB(circle: Circle, circleDB: CircleDB) {

	circleDB.CreatedAt = circle.CreatedAt
	circleDB.DeletedAt = circle.DeletedAt
	circleDB.ID = circle.ID

	// insertion point for basic fields copy operations
	circleDB.Lat = circle.Lat
	circleDB.Lng = circle.Lng
	circleDB.Name = circle.Name
	circleDB.Radius = circle.Radius
	circleDB.ColorEnum = circle.ColorEnum
	circleDB.DashStyleEnum = circle.DashStyleEnum

	// insertion point for pointer fields encoding
	circleDB.CirclePointersEncoding.LayerGroupID.Valid = true
	if (circle.LayerGroup != undefined) {
		circleDB.CirclePointersEncoding.LayerGroupID.Int64 = circle.LayerGroup.ID  
	} else {
		circleDB.CirclePointersEncoding.LayerGroupID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyCircleDBToCircle update basic, pointers and slice of pointers fields of circle
// from respectively the basic fields and encoded fields of pointers and slices of pointers of circleDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyCircleDBToCircle(circleDB: CircleDB, circle: Circle, frontRepo: FrontRepo) {

	circle.CreatedAt = circleDB.CreatedAt
	circle.DeletedAt = circleDB.DeletedAt
	circle.ID = circleDB.ID

	// insertion point for basic fields copy operations
	circle.Lat = circleDB.Lat
	circle.Lng = circleDB.Lng
	circle.Name = circleDB.Name
	circle.Radius = circleDB.Radius
	circle.ColorEnum = circleDB.ColorEnum
	circle.DashStyleEnum = circleDB.DashStyleEnum

	// insertion point for pointer fields encoding
	circle.LayerGroup = frontRepo.map_ID_LayerGroup.get(circleDB.CirclePointersEncoding.LayerGroupID.Int64)

	// insertion point for slice of pointers fields encoding
}

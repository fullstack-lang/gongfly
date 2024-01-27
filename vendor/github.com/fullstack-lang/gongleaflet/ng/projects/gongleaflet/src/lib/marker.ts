// generated code - do not edit

import { MarkerDB } from './marker-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroup } from './layergroup'
import { DivIcon } from './divicon'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class Marker {

	static GONGSTRUCT_NAME = "Marker"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	ColorEnum: string = ""

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroup

	DivIcon?: DivIcon

}

export function CopyMarkerToMarkerDB(marker: Marker, markerDB: MarkerDB) {

	markerDB.CreatedAt = marker.CreatedAt
	markerDB.DeletedAt = marker.DeletedAt
	markerDB.ID = marker.ID

	// insertion point for basic fields copy operations
	markerDB.Lat = marker.Lat
	markerDB.Lng = marker.Lng
	markerDB.Name = marker.Name
	markerDB.ColorEnum = marker.ColorEnum

	// insertion point for pointer fields encoding
	markerDB.MarkerPointersEncoding.LayerGroupID.Valid = true
	if (marker.LayerGroup != undefined) {
		markerDB.MarkerPointersEncoding.LayerGroupID.Int64 = marker.LayerGroup.ID  
	} else {
		markerDB.MarkerPointersEncoding.LayerGroupID.Int64 = 0 		
	}

	markerDB.MarkerPointersEncoding.DivIconID.Valid = true
	if (marker.DivIcon != undefined) {
		markerDB.MarkerPointersEncoding.DivIconID.Int64 = marker.DivIcon.ID  
	} else {
		markerDB.MarkerPointersEncoding.DivIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyMarkerDBToMarker update basic, pointers and slice of pointers fields of marker
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markerDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMarkerDBToMarker(markerDB: MarkerDB, marker: Marker, frontRepo: FrontRepo) {

	marker.CreatedAt = markerDB.CreatedAt
	marker.DeletedAt = markerDB.DeletedAt
	marker.ID = markerDB.ID

	// insertion point for basic fields copy operations
	marker.Lat = markerDB.Lat
	marker.Lng = markerDB.Lng
	marker.Name = markerDB.Name
	marker.ColorEnum = markerDB.ColorEnum

	// insertion point for pointer fields encoding
	marker.LayerGroup = frontRepo.map_ID_LayerGroup.get(markerDB.MarkerPointersEncoding.LayerGroupID.Int64)
	marker.DivIcon = frontRepo.map_ID_DivIcon.get(markerDB.MarkerPointersEncoding.DivIconID.Int64)

	// insertion point for slice of pointers fields encoding
}

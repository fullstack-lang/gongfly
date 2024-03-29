// generated code - do not edit

import { MarkerAPI } from './marker-api'
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

export function CopyMarkerToMarkerAPI(marker: Marker, markerAPI: MarkerAPI) {

	markerAPI.CreatedAt = marker.CreatedAt
	markerAPI.DeletedAt = marker.DeletedAt
	markerAPI.ID = marker.ID

	// insertion point for basic fields copy operations
	markerAPI.Lat = marker.Lat
	markerAPI.Lng = marker.Lng
	markerAPI.Name = marker.Name
	markerAPI.ColorEnum = marker.ColorEnum

	// insertion point for pointer fields encoding
	markerAPI.MarkerPointersEncoding.LayerGroupID.Valid = true
	if (marker.LayerGroup != undefined) {
		markerAPI.MarkerPointersEncoding.LayerGroupID.Int64 = marker.LayerGroup.ID  
	} else {
		markerAPI.MarkerPointersEncoding.LayerGroupID.Int64 = 0 		
	}

	markerAPI.MarkerPointersEncoding.DivIconID.Valid = true
	if (marker.DivIcon != undefined) {
		markerAPI.MarkerPointersEncoding.DivIconID.Int64 = marker.DivIcon.ID  
	} else {
		markerAPI.MarkerPointersEncoding.DivIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyMarkerAPIToMarker update basic, pointers and slice of pointers fields of marker
// from respectively the basic fields and encoded fields of pointers and slices of pointers of markerAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMarkerAPIToMarker(markerAPI: MarkerAPI, marker: Marker, frontRepo: FrontRepo) {

	marker.CreatedAt = markerAPI.CreatedAt
	marker.DeletedAt = markerAPI.DeletedAt
	marker.ID = markerAPI.ID

	// insertion point for basic fields copy operations
	marker.Lat = markerAPI.Lat
	marker.Lng = markerAPI.Lng
	marker.Name = markerAPI.Name
	marker.ColorEnum = markerAPI.ColorEnum

	// insertion point for pointer fields encoding
	marker.LayerGroup = frontRepo.map_ID_LayerGroup.get(markerAPI.MarkerPointersEncoding.LayerGroupID.Int64)
	marker.DivIcon = frontRepo.map_ID_DivIcon.get(markerAPI.MarkerPointersEncoding.DivIconID.Int64)

	// insertion point for slice of pointers fields encoding
}

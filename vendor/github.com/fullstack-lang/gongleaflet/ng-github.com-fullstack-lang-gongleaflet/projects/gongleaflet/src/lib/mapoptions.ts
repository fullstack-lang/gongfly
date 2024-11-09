// generated code - do not edit

import { MapOptionsAPI } from './mapoptions-api'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroupUse } from './layergroupuse'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class MapOptions {

	static GONGSTRUCT_NAME = "MapOptions"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Name: string = ""
	ZoomLevel: number = 0
	UrlTemplate: string = ""
	Attribution: string = ""
	MaxZoom: number = 0
	ZoomControl: boolean = false
	AttributionControl: boolean = false
	ZoomSnap: number = 0

	// insertion point for pointers and slices of pointers declarations
	LayerGroupUses: Array<LayerGroupUse> = []
}

export function CopyMapOptionsToMapOptionsAPI(mapoptions: MapOptions, mapoptionsAPI: MapOptionsAPI) {

	mapoptionsAPI.CreatedAt = mapoptions.CreatedAt
	mapoptionsAPI.DeletedAt = mapoptions.DeletedAt
	mapoptionsAPI.ID = mapoptions.ID

	// insertion point for basic fields copy operations
	mapoptionsAPI.Lat = mapoptions.Lat
	mapoptionsAPI.Lng = mapoptions.Lng
	mapoptionsAPI.Name = mapoptions.Name
	mapoptionsAPI.ZoomLevel = mapoptions.ZoomLevel
	mapoptionsAPI.UrlTemplate = mapoptions.UrlTemplate
	mapoptionsAPI.Attribution = mapoptions.Attribution
	mapoptionsAPI.MaxZoom = mapoptions.MaxZoom
	mapoptionsAPI.ZoomControl = mapoptions.ZoomControl
	mapoptionsAPI.AttributionControl = mapoptions.AttributionControl
	mapoptionsAPI.ZoomSnap = mapoptions.ZoomSnap

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	mapoptionsAPI.MapOptionsPointersEncoding.LayerGroupUses = []
	for (let _layergroupuse of mapoptions.LayerGroupUses) {
		mapoptionsAPI.MapOptionsPointersEncoding.LayerGroupUses.push(_layergroupuse.ID)
	}

}

// CopyMapOptionsAPIToMapOptions update basic, pointers and slice of pointers fields of mapoptions
// from respectively the basic fields and encoded fields of pointers and slices of pointers of mapoptionsAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyMapOptionsAPIToMapOptions(mapoptionsAPI: MapOptionsAPI, mapoptions: MapOptions, frontRepo: FrontRepo) {

	mapoptions.CreatedAt = mapoptionsAPI.CreatedAt
	mapoptions.DeletedAt = mapoptionsAPI.DeletedAt
	mapoptions.ID = mapoptionsAPI.ID

	// insertion point for basic fields copy operations
	mapoptions.Lat = mapoptionsAPI.Lat
	mapoptions.Lng = mapoptionsAPI.Lng
	mapoptions.Name = mapoptionsAPI.Name
	mapoptions.ZoomLevel = mapoptionsAPI.ZoomLevel
	mapoptions.UrlTemplate = mapoptionsAPI.UrlTemplate
	mapoptions.Attribution = mapoptionsAPI.Attribution
	mapoptions.MaxZoom = mapoptionsAPI.MaxZoom
	mapoptions.ZoomControl = mapoptionsAPI.ZoomControl
	mapoptions.AttributionControl = mapoptionsAPI.AttributionControl
	mapoptions.ZoomSnap = mapoptionsAPI.ZoomSnap

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	if (!Array.isArray(mapoptionsAPI.MapOptionsPointersEncoding.LayerGroupUses)) {
		console.error('Rects is not an array:', mapoptionsAPI.MapOptionsPointersEncoding.LayerGroupUses);
		return;
	}

	mapoptions.LayerGroupUses = new Array<LayerGroupUse>()
	for (let _id of mapoptionsAPI.MapOptionsPointersEncoding.LayerGroupUses) {
		let _layergroupuse = frontRepo.map_ID_LayerGroupUse.get(_id)
		if (_layergroupuse != undefined) {
			mapoptions.LayerGroupUses.push(_layergroupuse!)
		}
	}
}

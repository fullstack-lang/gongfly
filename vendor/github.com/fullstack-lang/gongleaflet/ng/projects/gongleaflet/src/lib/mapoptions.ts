// generated code - do not edit

import { MapOptionsDB } from './mapoptions-db'
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

export function CopyMapOptionsToMapOptionsDB(mapoptions: MapOptions, mapoptionsDB: MapOptionsDB) {

	mapoptionsDB.CreatedAt = mapoptions.CreatedAt
	mapoptionsDB.DeletedAt = mapoptions.DeletedAt
	mapoptionsDB.ID = mapoptions.ID
	
	// insertion point for basic fields copy operations
	mapoptionsDB.Lat = mapoptions.Lat
	mapoptionsDB.Lng = mapoptions.Lng
	mapoptionsDB.Name = mapoptions.Name
	mapoptionsDB.ZoomLevel = mapoptions.ZoomLevel
	mapoptionsDB.UrlTemplate = mapoptions.UrlTemplate
	mapoptionsDB.Attribution = mapoptions.Attribution
	mapoptionsDB.MaxZoom = mapoptions.MaxZoom
	mapoptionsDB.ZoomControl = mapoptions.ZoomControl
	mapoptionsDB.AttributionControl = mapoptions.AttributionControl
	mapoptionsDB.ZoomSnap = mapoptions.ZoomSnap

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	mapoptionsDB.MapOptionsPointersEncoding.LayerGroupUses = []
    for (let _layergroupuse of mapoptions.LayerGroupUses) {
		mapoptionsDB.MapOptionsPointersEncoding.LayerGroupUses.push(_layergroupuse.ID)
    }
	
}

export function CopyMapOptionsDBToMapOptions(mapoptionsDB: MapOptionsDB, mapoptions: MapOptions, frontRepo: FrontRepo) {

	mapoptions.CreatedAt = mapoptionsDB.CreatedAt
	mapoptions.DeletedAt = mapoptionsDB.DeletedAt
	mapoptions.ID = mapoptionsDB.ID
	
	// insertion point for basic fields copy operations
	mapoptions.Lat = mapoptionsDB.Lat
	mapoptions.Lng = mapoptionsDB.Lng
	mapoptions.Name = mapoptionsDB.Name
	mapoptions.ZoomLevel = mapoptionsDB.ZoomLevel
	mapoptions.UrlTemplate = mapoptionsDB.UrlTemplate
	mapoptions.Attribution = mapoptionsDB.Attribution
	mapoptions.MaxZoom = mapoptionsDB.MaxZoom
	mapoptions.ZoomControl = mapoptionsDB.ZoomControl
	mapoptions.AttributionControl = mapoptionsDB.AttributionControl
	mapoptions.ZoomSnap = mapoptionsDB.ZoomSnap

	// insertion point for pointer fields encoding

	// insertion point for slice of pointers fields encoding
	mapoptions.LayerGroupUses = new Array<LayerGroupUse>()
	for (let _id of mapoptionsDB.MapOptionsPointersEncoding.LayerGroupUses) {
	  let _layergroupuse = frontRepo.LayerGroupUses.get(_id)
	  if (_layergroupuse != undefined) {
		mapoptions.LayerGroupUses.push(_layergroupuse!)
	  }
	}
}

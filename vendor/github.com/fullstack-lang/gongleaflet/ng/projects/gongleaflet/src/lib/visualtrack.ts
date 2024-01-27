// generated code - do not edit

import { VisualTrackDB } from './visualtrack-db'
import { FrontRepo } from './front-repo.service';

// insertion point for imports
import { LayerGroup } from './layergroup'
import { DivIcon } from './divicon'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class VisualTrack {

	static GONGSTRUCT_NAME = "VisualTrack"

	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Lat: number = 0
	Lng: number = 0
	Heading: number = 0
	Level: number = 0
	Speed: number = 0
	VerticalSpeed: number = 0
	Name: string = ""
	ColorEnum: string = ""
	DisplayTrackHistory: boolean = false
	DisplayLevelAndSpeed: boolean = false

	// insertion point for pointers and slices of pointers declarations
	LayerGroup?: LayerGroup

	DivIcon?: DivIcon

}

export function CopyVisualTrackToVisualTrackDB(visualtrack: VisualTrack, visualtrackDB: VisualTrackDB) {

	visualtrackDB.CreatedAt = visualtrack.CreatedAt
	visualtrackDB.DeletedAt = visualtrack.DeletedAt
	visualtrackDB.ID = visualtrack.ID

	// insertion point for basic fields copy operations
	visualtrackDB.Lat = visualtrack.Lat
	visualtrackDB.Lng = visualtrack.Lng
	visualtrackDB.Heading = visualtrack.Heading
	visualtrackDB.Level = visualtrack.Level
	visualtrackDB.Speed = visualtrack.Speed
	visualtrackDB.VerticalSpeed = visualtrack.VerticalSpeed
	visualtrackDB.Name = visualtrack.Name
	visualtrackDB.ColorEnum = visualtrack.ColorEnum
	visualtrackDB.DisplayTrackHistory = visualtrack.DisplayTrackHistory
	visualtrackDB.DisplayLevelAndSpeed = visualtrack.DisplayLevelAndSpeed

	// insertion point for pointer fields encoding
	visualtrackDB.VisualTrackPointersEncoding.LayerGroupID.Valid = true
	if (visualtrack.LayerGroup != undefined) {
		visualtrackDB.VisualTrackPointersEncoding.LayerGroupID.Int64 = visualtrack.LayerGroup.ID  
	} else {
		visualtrackDB.VisualTrackPointersEncoding.LayerGroupID.Int64 = 0 		
	}

	visualtrackDB.VisualTrackPointersEncoding.DivIconID.Valid = true
	if (visualtrack.DivIcon != undefined) {
		visualtrackDB.VisualTrackPointersEncoding.DivIconID.Int64 = visualtrack.DivIcon.ID  
	} else {
		visualtrackDB.VisualTrackPointersEncoding.DivIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyVisualTrackDBToVisualTrack update basic, pointers and slice of pointers fields of visualtrack
// from respectively the basic fields and encoded fields of pointers and slices of pointers of visualtrackDB
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVisualTrackDBToVisualTrack(visualtrackDB: VisualTrackDB, visualtrack: VisualTrack, frontRepo: FrontRepo) {

	visualtrack.CreatedAt = visualtrackDB.CreatedAt
	visualtrack.DeletedAt = visualtrackDB.DeletedAt
	visualtrack.ID = visualtrackDB.ID

	// insertion point for basic fields copy operations
	visualtrack.Lat = visualtrackDB.Lat
	visualtrack.Lng = visualtrackDB.Lng
	visualtrack.Heading = visualtrackDB.Heading
	visualtrack.Level = visualtrackDB.Level
	visualtrack.Speed = visualtrackDB.Speed
	visualtrack.VerticalSpeed = visualtrackDB.VerticalSpeed
	visualtrack.Name = visualtrackDB.Name
	visualtrack.ColorEnum = visualtrackDB.ColorEnum
	visualtrack.DisplayTrackHistory = visualtrackDB.DisplayTrackHistory
	visualtrack.DisplayLevelAndSpeed = visualtrackDB.DisplayLevelAndSpeed

	// insertion point for pointer fields encoding
	visualtrack.LayerGroup = frontRepo.map_ID_LayerGroup.get(visualtrackDB.VisualTrackPointersEncoding.LayerGroupID.Int64)
	visualtrack.DivIcon = frontRepo.map_ID_DivIcon.get(visualtrackDB.VisualTrackPointersEncoding.DivIconID.Int64)

	// insertion point for slice of pointers fields encoding
}

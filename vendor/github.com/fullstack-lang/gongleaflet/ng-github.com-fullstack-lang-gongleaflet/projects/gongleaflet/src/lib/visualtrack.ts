// generated code - do not edit

import { VisualTrackAPI } from './visualtrack-api'
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

export function CopyVisualTrackToVisualTrackAPI(visualtrack: VisualTrack, visualtrackAPI: VisualTrackAPI) {

	visualtrackAPI.CreatedAt = visualtrack.CreatedAt
	visualtrackAPI.DeletedAt = visualtrack.DeletedAt
	visualtrackAPI.ID = visualtrack.ID

	// insertion point for basic fields copy operations
	visualtrackAPI.Lat = visualtrack.Lat
	visualtrackAPI.Lng = visualtrack.Lng
	visualtrackAPI.Heading = visualtrack.Heading
	visualtrackAPI.Level = visualtrack.Level
	visualtrackAPI.Speed = visualtrack.Speed
	visualtrackAPI.VerticalSpeed = visualtrack.VerticalSpeed
	visualtrackAPI.Name = visualtrack.Name
	visualtrackAPI.ColorEnum = visualtrack.ColorEnum
	visualtrackAPI.DisplayTrackHistory = visualtrack.DisplayTrackHistory
	visualtrackAPI.DisplayLevelAndSpeed = visualtrack.DisplayLevelAndSpeed

	// insertion point for pointer fields encoding
	visualtrackAPI.VisualTrackPointersEncoding.LayerGroupID.Valid = true
	if (visualtrack.LayerGroup != undefined) {
		visualtrackAPI.VisualTrackPointersEncoding.LayerGroupID.Int64 = visualtrack.LayerGroup.ID  
	} else {
		visualtrackAPI.VisualTrackPointersEncoding.LayerGroupID.Int64 = 0 		
	}

	visualtrackAPI.VisualTrackPointersEncoding.DivIconID.Valid = true
	if (visualtrack.DivIcon != undefined) {
		visualtrackAPI.VisualTrackPointersEncoding.DivIconID.Int64 = visualtrack.DivIcon.ID  
	} else {
		visualtrackAPI.VisualTrackPointersEncoding.DivIconID.Int64 = 0 		
	}


	// insertion point for slice of pointers fields encoding
}

// CopyVisualTrackAPIToVisualTrack update basic, pointers and slice of pointers fields of visualtrack
// from respectively the basic fields and encoded fields of pointers and slices of pointers of visualtrackAPI
// this function uses frontRepo.map_ID_<structname> to decode the encoded fields
// a condition is that those maps has to be initialized before
export function CopyVisualTrackAPIToVisualTrack(visualtrackAPI: VisualTrackAPI, visualtrack: VisualTrack, frontRepo: FrontRepo) {

	visualtrack.CreatedAt = visualtrackAPI.CreatedAt
	visualtrack.DeletedAt = visualtrackAPI.DeletedAt
	visualtrack.ID = visualtrackAPI.ID

	// insertion point for basic fields copy operations
	visualtrack.Lat = visualtrackAPI.Lat
	visualtrack.Lng = visualtrackAPI.Lng
	visualtrack.Heading = visualtrackAPI.Heading
	visualtrack.Level = visualtrackAPI.Level
	visualtrack.Speed = visualtrackAPI.Speed
	visualtrack.VerticalSpeed = visualtrackAPI.VerticalSpeed
	visualtrack.Name = visualtrackAPI.Name
	visualtrack.ColorEnum = visualtrackAPI.ColorEnum
	visualtrack.DisplayTrackHistory = visualtrackAPI.DisplayTrackHistory
	visualtrack.DisplayLevelAndSpeed = visualtrackAPI.DisplayLevelAndSpeed

	// insertion point for pointer fields encoding
	visualtrack.LayerGroup = frontRepo.map_ID_LayerGroup.get(visualtrackAPI.VisualTrackPointersEncoding.LayerGroupID.Int64)
	visualtrack.DivIcon = frontRepo.map_ID_DivIcon.get(visualtrackAPI.VisualTrackPointersEncoding.DivIconID.Int64)

	// insertion point for slice of pointers fields encoding
}

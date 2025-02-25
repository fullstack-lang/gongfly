// generated code - do not edit

//insertion point for imports
import { CircleAPI } from './circle-api'

import { DivIconAPI } from './divicon-api'

import { LayerGroupAPI } from './layergroup-api'

import { LayerGroupUseAPI } from './layergroupuse-api'

import { MapOptionsAPI } from './mapoptions-api'

import { MarkerAPI } from './marker-api'

import { UserClickAPI } from './userclick-api'

import { VLineAPI } from './vline-api'

import { VisualTrackAPI } from './visualtrack-api'


export class BackRepoData {
	// insertion point for declarations
	CircleAPIs = new Array<CircleAPI>()

	DivIconAPIs = new Array<DivIconAPI>()

	LayerGroupAPIs = new Array<LayerGroupAPI>()

	LayerGroupUseAPIs = new Array<LayerGroupUseAPI>()

	MapOptionsAPIs = new Array<MapOptionsAPI>()

	MarkerAPIs = new Array<MarkerAPI>()

	UserClickAPIs = new Array<UserClickAPI>()

	VLineAPIs = new Array<VLineAPI>()

	VisualTrackAPIs = new Array<VisualTrackAPI>()


	// index of the web socket for this stack type (unique among all stack instances)
	GONG__Index : number

	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CircleAPIs = data?.CircleAPIs || [];

		this.DivIconAPIs = data?.DivIconAPIs || [];

		this.LayerGroupAPIs = data?.LayerGroupAPIs || [];

		this.LayerGroupUseAPIs = data?.LayerGroupUseAPIs || [];

		this.MapOptionsAPIs = data?.MapOptionsAPIs || [];

		this.MarkerAPIs = data?.MarkerAPIs || [];

		this.UserClickAPIs = data?.UserClickAPIs || [];

		this.VLineAPIs = data?.VLineAPIs || [];

		this.VisualTrackAPIs = data?.VisualTrackAPIs || [];

		this.GONG__Index = data?.GONG__Index ?? -1;   // Assign Index here
	}

}
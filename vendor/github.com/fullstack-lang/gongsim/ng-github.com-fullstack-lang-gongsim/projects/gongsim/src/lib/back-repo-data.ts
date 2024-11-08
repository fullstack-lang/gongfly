// generated code - do not edit

//insertion point for imports
import { DummyAgentAPI } from './dummyagent-api'

import { EngineAPI } from './engine-api'

import { EventAPI } from './event-api'

import { GongsimCommandAPI } from './gongsimcommand-api'

import { GongsimStatusAPI } from './gongsimstatus-api'

import { UpdateStateAPI } from './updatestate-api'


export class BackRepoData {
	// insertion point for declarations
	DummyAgentAPIs = new Array<DummyAgentAPI>()

	EngineAPIs = new Array<EngineAPI>()

	EventAPIs = new Array<EventAPI>()

	GongsimCommandAPIs = new Array<GongsimCommandAPI>()

	GongsimStatusAPIs = new Array<GongsimStatusAPI>()

	UpdateStateAPIs = new Array<UpdateStateAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.DummyAgentAPIs = data?.DummyAgentAPIs || [];

		this.EngineAPIs = data?.EngineAPIs || [];

		this.EventAPIs = data?.EventAPIs || [];

		this.GongsimCommandAPIs = data?.GongsimCommandAPIs || [];

		this.GongsimStatusAPIs = data?.GongsimStatusAPIs || [];

		this.UpdateStateAPIs = data?.UpdateStateAPIs || [];

	}

}
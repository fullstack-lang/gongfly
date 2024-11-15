// generated code - do not edit

//insertion point for imports
import { CommandAPI } from './command-api'

import { DummyAgentAPI } from './dummyagent-api'

import { EngineAPI } from './engine-api'

import { EventAPI } from './event-api'

import { StatusAPI } from './status-api'

import { UpdateStateAPI } from './updatestate-api'


export class BackRepoData {
	// insertion point for declarations
	CommandAPIs = new Array<CommandAPI>()

	DummyAgentAPIs = new Array<DummyAgentAPI>()

	EngineAPIs = new Array<EngineAPI>()

	EventAPIs = new Array<EventAPI>()

	StatusAPIs = new Array<StatusAPI>()

	UpdateStateAPIs = new Array<UpdateStateAPI>()



	constructor(data?: Partial<BackRepoData>) {
		// insertion point for copies
		this.CommandAPIs = data?.CommandAPIs || [];

		this.DummyAgentAPIs = data?.DummyAgentAPIs || [];

		this.EngineAPIs = data?.EngineAPIs || [];

		this.EventAPIs = data?.EventAPIs || [];

		this.StatusAPIs = data?.StatusAPIs || [];

		this.UpdateStateAPIs = data?.UpdateStateAPIs || [];

	}

}
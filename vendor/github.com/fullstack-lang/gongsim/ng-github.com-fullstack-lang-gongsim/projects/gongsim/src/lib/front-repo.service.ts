import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders, HttpParams } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { DummyAgentAPI } from './dummyagent-api'
import { DummyAgent, CopyDummyAgentAPIToDummyAgent } from './dummyagent'
import { DummyAgentService } from './dummyagent.service'

import { EngineAPI } from './engine-api'
import { Engine, CopyEngineAPIToEngine } from './engine'
import { EngineService } from './engine.service'

import { EventAPI } from './event-api'
import { Event, CopyEventAPIToEvent } from './event'
import { EventService } from './event.service'

import { GongsimCommandAPI } from './gongsimcommand-api'
import { GongsimCommand, CopyGongsimCommandAPIToGongsimCommand } from './gongsimcommand'
import { GongsimCommandService } from './gongsimcommand.service'

import { GongsimStatusAPI } from './gongsimstatus-api'
import { GongsimStatus, CopyGongsimStatusAPIToGongsimStatus } from './gongsimstatus'
import { GongsimStatusService } from './gongsimstatus.service'

import { UpdateStateAPI } from './updatestate-api'
import { UpdateState, CopyUpdateStateAPIToUpdateState } from './updatestate'
import { UpdateStateService } from './updatestate.service'


import { BackRepoData } from './back-repo-data'

export const StackType = "github.com/fullstack-lang/gongsim/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_DummyAgents = new Array<DummyAgent>() // array of front instances
	map_ID_DummyAgent = new Map<number, DummyAgent>() // map of front instances

	array_Engines = new Array<Engine>() // array of front instances
	map_ID_Engine = new Map<number, Engine>() // map of front instances

	array_Events = new Array<Event>() // array of front instances
	map_ID_Event = new Map<number, Event>() // map of front instances

	array_GongsimCommands = new Array<GongsimCommand>() // array of front instances
	map_ID_GongsimCommand = new Map<number, GongsimCommand>() // map of front instances

	array_GongsimStatuss = new Array<GongsimStatus>() // array of front instances
	map_ID_GongsimStatus = new Map<number, GongsimStatus>() // map of front instances

	array_UpdateStates = new Array<UpdateState>() // array of front instances
	map_ID_UpdateState = new Map<number, UpdateState>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'DummyAgent':
				return this.array_DummyAgents as unknown as Array<Type>
			case 'Engine':
				return this.array_Engines as unknown as Array<Type>
			case 'Event':
				return this.array_Events as unknown as Array<Type>
			case 'GongsimCommand':
				return this.array_GongsimCommands as unknown as Array<Type>
			case 'GongsimStatus':
				return this.array_GongsimStatuss as unknown as Array<Type>
			case 'UpdateState':
				return this.array_UpdateStates as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}

	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'DummyAgent':
				return this.map_ID_DummyAgent as unknown as Map<number, Type>
			case 'Engine':
				return this.map_ID_Engine as unknown as Map<number, Type>
			case 'Event':
				return this.map_ID_Event as unknown as Map<number, Type>
			case 'GongsimCommand':
				return this.map_ID_GongsimCommand as unknown as Map<number, Type>
			case 'GongsimStatus':
				return this.map_ID_GongsimStatus as unknown as Map<number, Type>
			case 'UpdateState':
				return this.map_ID_UpdateState as unknown as Map<number, Type>
			default:
				throw new Error("Type not recognized");
		}
	}
}

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
	ID: number = 0 // ID of the calling instance

	// the reverse pointer is the name of the generated field on the destination
	// struct of the ONE-MANY association
	ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
	OrderingMode: boolean = false // if true, this is for ordering items

	// there are different selection mode : ONE_MANY or MANY_MANY
	SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

	// used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
	//
	// In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
	// 
	// in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
	// at the end of the ONE-MANY association
	SourceStruct: string = ""	// The "Aclass"
	SourceField: string = "" // the "AnarrayofbUse"
	IntermediateStruct: string = "" // the "AclassBclassUse" 
	IntermediateStructField: string = "" // the "Bclass" as field
	NextAssociationStruct: string = "" // the "Bclass"

	GONG__StackPath: string = ""
}

export enum SelectionMode {
	ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
	MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
	providedIn: 'root'
})
export class FrontRepoService {

	GONG__StackPath: string = ""
	private socket: WebSocket | undefined

	httpOptions = {
		headers: new HttpHeaders({ 'Content-Type': 'application/json' })
	};

	//
	// Store of all instances of the stack
	//
	frontRepo = new (FrontRepo)

	constructor(
		private http: HttpClient, // insertion point sub template 
		private dummyagentService: DummyAgentService,
		private engineService: EngineService,
		private eventService: EventService,
		private gongsimcommandService: GongsimCommandService,
		private gongsimstatusService: GongsimStatusService,
		private updatestateService: UpdateStateService,
	) { }

	// postService provides a post function for each struct name
	postService(structName: string, instanceToBePosted: any) {
		let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
		let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

		servicePostFunction(instanceToBePosted).subscribe(
			instance => {
				let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("post")
			}
		);
	}

	// deleteService provides a delete function for each struct name
	deleteService(structName: string, instanceToBeDeleted: any) {
		let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
		let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

		serviceDeleteFunction(instanceToBeDeleted).subscribe(
			instance => {
				let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
				behaviorSubject.next("delete")
			}
		);
	}

	// typing of observable can be messy in typescript. Therefore, one force the type
	observableFrontRepo: [
		Observable<null>, // see below for the of(null) observable
		// insertion point sub template 
		Observable<DummyAgentAPI[]>,
		Observable<EngineAPI[]>,
		Observable<EventAPI[]>,
		Observable<GongsimCommandAPI[]>,
		Observable<GongsimStatusAPI[]>,
		Observable<UpdateStateAPI[]>,
	] = [
			// Using "combineLatest" with a placeholder observable.
			//
			// This allows the typescript compiler to pass when no GongStruct is present in the front API
			//
			// The "of(null)" is a "meaningless" observable that emits a single value (null) and completes.
			// This is used as a workaround to satisfy TypeScript requirements and the "combineLatest" 
			// expectation for a non-empty array of observables.
			of(null), // 
			// insertion point sub template
			this.dummyagentService.getDummyAgents(this.GONG__StackPath, this.frontRepo),
			this.engineService.getEngines(this.GONG__StackPath, this.frontRepo),
			this.eventService.getEvents(this.GONG__StackPath, this.frontRepo),
			this.gongsimcommandService.getGongsimCommands(this.GONG__StackPath, this.frontRepo),
			this.gongsimstatusService.getGongsimStatuss(this.GONG__StackPath, this.frontRepo),
			this.updatestateService.getUpdateStates(this.GONG__StackPath, this.frontRepo),
		];

	//
	// pull performs a GET on all struct of the stack and redeem association pointers 
	//
	// This is an observable. Therefore, the control flow forks with
	// - pull() return immediatly the observable
	// - the observable observer, if it subscribe, is called when all GET calls are performs
	pull(GONG__StackPath: string = ""): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath

		this.observableFrontRepo = [
			of(null), // see above for justification
			// insertion point sub template
			this.dummyagentService.getDummyAgents(this.GONG__StackPath, this.frontRepo),
			this.engineService.getEngines(this.GONG__StackPath, this.frontRepo),
			this.eventService.getEvents(this.GONG__StackPath, this.frontRepo),
			this.gongsimcommandService.getGongsimCommands(this.GONG__StackPath, this.frontRepo),
			this.gongsimstatusService.getGongsimStatuss(this.GONG__StackPath, this.frontRepo),
			this.updatestateService.getUpdateStates(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						dummyagents_,
						engines_,
						events_,
						gongsimcommands_,
						gongsimstatuss_,
						updatestates_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var dummyagents: DummyAgentAPI[]
						dummyagents = dummyagents_ as DummyAgentAPI[]
						var engines: EngineAPI[]
						engines = engines_ as EngineAPI[]
						var events: EventAPI[]
						events = events_ as EventAPI[]
						var gongsimcommands: GongsimCommandAPI[]
						gongsimcommands = gongsimcommands_ as GongsimCommandAPI[]
						var gongsimstatuss: GongsimStatusAPI[]
						gongsimstatuss = gongsimstatuss_ as GongsimStatusAPI[]
						var updatestates: UpdateStateAPI[]
						updatestates = updatestates_ as UpdateStateAPI[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_DummyAgents = []
						this.frontRepo.map_ID_DummyAgent.clear()

						dummyagents.forEach(
							dummyagentAPI => {
								let dummyagent = new DummyAgent
								this.frontRepo.array_DummyAgents.push(dummyagent)
								this.frontRepo.map_ID_DummyAgent.set(dummyagentAPI.ID, dummyagent)
							}
						)

						// init the arrays
						this.frontRepo.array_Engines = []
						this.frontRepo.map_ID_Engine.clear()

						engines.forEach(
							engineAPI => {
								let engine = new Engine
								this.frontRepo.array_Engines.push(engine)
								this.frontRepo.map_ID_Engine.set(engineAPI.ID, engine)
							}
						)

						// init the arrays
						this.frontRepo.array_Events = []
						this.frontRepo.map_ID_Event.clear()

						events.forEach(
							eventAPI => {
								let event = new Event
								this.frontRepo.array_Events.push(event)
								this.frontRepo.map_ID_Event.set(eventAPI.ID, event)
							}
						)

						// init the arrays
						this.frontRepo.array_GongsimCommands = []
						this.frontRepo.map_ID_GongsimCommand.clear()

						gongsimcommands.forEach(
							gongsimcommandAPI => {
								let gongsimcommand = new GongsimCommand
								this.frontRepo.array_GongsimCommands.push(gongsimcommand)
								this.frontRepo.map_ID_GongsimCommand.set(gongsimcommandAPI.ID, gongsimcommand)
							}
						)

						// init the arrays
						this.frontRepo.array_GongsimStatuss = []
						this.frontRepo.map_ID_GongsimStatus.clear()

						gongsimstatuss.forEach(
							gongsimstatusAPI => {
								let gongsimstatus = new GongsimStatus
								this.frontRepo.array_GongsimStatuss.push(gongsimstatus)
								this.frontRepo.map_ID_GongsimStatus.set(gongsimstatusAPI.ID, gongsimstatus)
							}
						)

						// init the arrays
						this.frontRepo.array_UpdateStates = []
						this.frontRepo.map_ID_UpdateState.clear()

						updatestates.forEach(
							updatestateAPI => {
								let updatestate = new UpdateState
								this.frontRepo.array_UpdateStates.push(updatestate)
								this.frontRepo.map_ID_UpdateState.set(updatestateAPI.ID, updatestate)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						dummyagents.forEach(
							dummyagentAPI => {
								let dummyagent = this.frontRepo.map_ID_DummyAgent.get(dummyagentAPI.ID)
								CopyDummyAgentAPIToDummyAgent(dummyagentAPI, dummyagent!, this.frontRepo)
							}
						)

						// fill up front objects
						engines.forEach(
							engineAPI => {
								let engine = this.frontRepo.map_ID_Engine.get(engineAPI.ID)
								CopyEngineAPIToEngine(engineAPI, engine!, this.frontRepo)
							}
						)

						// fill up front objects
						events.forEach(
							eventAPI => {
								let event = this.frontRepo.map_ID_Event.get(eventAPI.ID)
								CopyEventAPIToEvent(eventAPI, event!, this.frontRepo)
							}
						)

						// fill up front objects
						gongsimcommands.forEach(
							gongsimcommandAPI => {
								let gongsimcommand = this.frontRepo.map_ID_GongsimCommand.get(gongsimcommandAPI.ID)
								CopyGongsimCommandAPIToGongsimCommand(gongsimcommandAPI, gongsimcommand!, this.frontRepo)
							}
						)

						// fill up front objects
						gongsimstatuss.forEach(
							gongsimstatusAPI => {
								let gongsimstatus = this.frontRepo.map_ID_GongsimStatus.get(gongsimstatusAPI.ID)
								CopyGongsimStatusAPIToGongsimStatus(gongsimstatusAPI, gongsimstatus!, this.frontRepo)
							}
						)

						// fill up front objects
						updatestates.forEach(
							updatestateAPI => {
								let updatestate = this.frontRepo.map_ID_UpdateState.get(updatestateAPI.ID)
								CopyUpdateStateAPIToUpdateState(updatestateAPI, updatestate!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
	}

	public connectToWebSocket(GONG__StackPath: string): Observable<FrontRepo> {

		this.GONG__StackPath = GONG__StackPath


		let params = new HttpParams().set("GONG__StackPath", this.GONG__StackPath)
		let basePath = 'ws://localhost:8080/api/github.com/fullstack-lang/gongsim/go/v1/ws/stage'
		let paramString = params.toString()
		let url = `${basePath}?${paramString}`
		this.socket = new WebSocket(url)

		return new Observable(observer => {
			this.socket!.onmessage = event => {
				let _this = this

				const backRepoData = new BackRepoData(JSON.parse(event.data))

				// 
				// First Step: init map of instances
				// insertion point sub template for init 
				// init the arrays
				// insertion point sub template for init 
				// init the arrays
				this.frontRepo.array_DummyAgents = []
				this.frontRepo.map_ID_DummyAgent.clear()

				backRepoData.DummyAgentAPIs.forEach(
					dummyagentAPI => {
						let dummyagent = new DummyAgent
						this.frontRepo.array_DummyAgents.push(dummyagent)
						this.frontRepo.map_ID_DummyAgent.set(dummyagentAPI.ID, dummyagent)
					}
				)

				// init the arrays
				this.frontRepo.array_Engines = []
				this.frontRepo.map_ID_Engine.clear()

				backRepoData.EngineAPIs.forEach(
					engineAPI => {
						let engine = new Engine
						this.frontRepo.array_Engines.push(engine)
						this.frontRepo.map_ID_Engine.set(engineAPI.ID, engine)
					}
				)

				// init the arrays
				this.frontRepo.array_Events = []
				this.frontRepo.map_ID_Event.clear()

				backRepoData.EventAPIs.forEach(
					eventAPI => {
						let event = new Event
						this.frontRepo.array_Events.push(event)
						this.frontRepo.map_ID_Event.set(eventAPI.ID, event)
					}
				)

				// init the arrays
				this.frontRepo.array_GongsimCommands = []
				this.frontRepo.map_ID_GongsimCommand.clear()

				backRepoData.GongsimCommandAPIs.forEach(
					gongsimcommandAPI => {
						let gongsimcommand = new GongsimCommand
						this.frontRepo.array_GongsimCommands.push(gongsimcommand)
						this.frontRepo.map_ID_GongsimCommand.set(gongsimcommandAPI.ID, gongsimcommand)
					}
				)

				// init the arrays
				this.frontRepo.array_GongsimStatuss = []
				this.frontRepo.map_ID_GongsimStatus.clear()

				backRepoData.GongsimStatusAPIs.forEach(
					gongsimstatusAPI => {
						let gongsimstatus = new GongsimStatus
						this.frontRepo.array_GongsimStatuss.push(gongsimstatus)
						this.frontRepo.map_ID_GongsimStatus.set(gongsimstatusAPI.ID, gongsimstatus)
					}
				)

				// init the arrays
				this.frontRepo.array_UpdateStates = []
				this.frontRepo.map_ID_UpdateState.clear()

				backRepoData.UpdateStateAPIs.forEach(
					updatestateAPI => {
						let updatestate = new UpdateState
						this.frontRepo.array_UpdateStates.push(updatestate)
						this.frontRepo.map_ID_UpdateState.set(updatestateAPI.ID, updatestate)
					}
				)


				// 
				// Second Step: reddeem front objects
				// insertion point sub template for redeem 
				// fill up front objects
				// insertion point sub template for redeem 
				// fill up front objects
				backRepoData.DummyAgentAPIs.forEach(
					dummyagentAPI => {
						let dummyagent = this.frontRepo.map_ID_DummyAgent.get(dummyagentAPI.ID)
						CopyDummyAgentAPIToDummyAgent(dummyagentAPI, dummyagent!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.EngineAPIs.forEach(
					engineAPI => {
						let engine = this.frontRepo.map_ID_Engine.get(engineAPI.ID)
						CopyEngineAPIToEngine(engineAPI, engine!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.EventAPIs.forEach(
					eventAPI => {
						let event = this.frontRepo.map_ID_Event.get(eventAPI.ID)
						CopyEventAPIToEvent(eventAPI, event!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongsimCommandAPIs.forEach(
					gongsimcommandAPI => {
						let gongsimcommand = this.frontRepo.map_ID_GongsimCommand.get(gongsimcommandAPI.ID)
						CopyGongsimCommandAPIToGongsimCommand(gongsimcommandAPI, gongsimcommand!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.GongsimStatusAPIs.forEach(
					gongsimstatusAPI => {
						let gongsimstatus = this.frontRepo.map_ID_GongsimStatus.get(gongsimstatusAPI.ID)
						CopyGongsimStatusAPIToGongsimStatus(gongsimstatusAPI, gongsimstatus!, this.frontRepo)
					}
				)

				// fill up front objects
				backRepoData.UpdateStateAPIs.forEach(
					updatestateAPI => {
						let updatestate = this.frontRepo.map_ID_UpdateState.get(updatestateAPI.ID)
						CopyUpdateStateAPIToUpdateState(updatestateAPI, updatestate!, this.frontRepo)
					}
				)



				observer.next(this.frontRepo)
			}
			this.socket!.onerror = event => {
				observer.error(event)
			}
			this.socket!.onclose = event => {
				observer.complete()
			}

			return () => {
				this.socket!.close()
			}
		})
	}
}

// insertion point for get unique ID per struct 
export function getDummyAgentUniqueID(id: number): number {
	return 31 * id
}
export function getEngineUniqueID(id: number): number {
	return 37 * id
}
export function getEventUniqueID(id: number): number {
	return 41 * id
}
export function getGongsimCommandUniqueID(id: number): number {
	return 43 * id
}
export function getGongsimStatusUniqueID(id: number): number {
	return 47 * id
}
export function getUpdateStateUniqueID(id: number): number {
	return 53 * id
}

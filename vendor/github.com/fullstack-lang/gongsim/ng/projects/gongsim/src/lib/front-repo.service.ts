import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { DummyAgentDB } from './dummyagent-db'
import { DummyAgent, CopyDummyAgentDBToDummyAgent } from './dummyagent'
import { DummyAgentService } from './dummyagent.service'

import { EngineDB } from './engine-db'
import { Engine, CopyEngineDBToEngine } from './engine'
import { EngineService } from './engine.service'

import { EventDB } from './event-db'
import { Event, CopyEventDBToEvent } from './event'
import { EventService } from './event.service'

import { GongsimCommandDB } from './gongsimcommand-db'
import { GongsimCommand, CopyGongsimCommandDBToGongsimCommand } from './gongsimcommand'
import { GongsimCommandService } from './gongsimcommand.service'

import { GongsimStatusDB } from './gongsimstatus-db'
import { GongsimStatus, CopyGongsimStatusDBToGongsimStatus } from './gongsimstatus'
import { GongsimStatusService } from './gongsimstatus.service'

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
		Observable<DummyAgentDB[]>,
		Observable<EngineDB[]>,
		Observable<EventDB[]>,
		Observable<GongsimCommandDB[]>,
		Observable<GongsimStatusDB[]>,
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
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var dummyagents: DummyAgentDB[]
						dummyagents = dummyagents_ as DummyAgentDB[]
						var engines: EngineDB[]
						engines = engines_ as EngineDB[]
						var events: EventDB[]
						events = events_ as EventDB[]
						var gongsimcommands: GongsimCommandDB[]
						gongsimcommands = gongsimcommands_ as GongsimCommandDB[]
						var gongsimstatuss: GongsimStatusDB[]
						gongsimstatuss = gongsimstatuss_ as GongsimStatusDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_DummyAgents = []
						this.frontRepo.map_ID_DummyAgent.clear()

						dummyagents.forEach(
							dummyagentDB => {
								let dummyagent = new DummyAgent
								this.frontRepo.array_DummyAgents.push(dummyagent)
								this.frontRepo.map_ID_DummyAgent.set(dummyagentDB.ID, dummyagent)
							}
						)

						// init the arrays
						this.frontRepo.array_Engines = []
						this.frontRepo.map_ID_Engine.clear()

						engines.forEach(
							engineDB => {
								let engine = new Engine
								this.frontRepo.array_Engines.push(engine)
								this.frontRepo.map_ID_Engine.set(engineDB.ID, engine)
							}
						)

						// init the arrays
						this.frontRepo.array_Events = []
						this.frontRepo.map_ID_Event.clear()

						events.forEach(
							eventDB => {
								let event = new Event
								this.frontRepo.array_Events.push(event)
								this.frontRepo.map_ID_Event.set(eventDB.ID, event)
							}
						)

						// init the arrays
						this.frontRepo.array_GongsimCommands = []
						this.frontRepo.map_ID_GongsimCommand.clear()

						gongsimcommands.forEach(
							gongsimcommandDB => {
								let gongsimcommand = new GongsimCommand
								this.frontRepo.array_GongsimCommands.push(gongsimcommand)
								this.frontRepo.map_ID_GongsimCommand.set(gongsimcommandDB.ID, gongsimcommand)
							}
						)

						// init the arrays
						this.frontRepo.array_GongsimStatuss = []
						this.frontRepo.map_ID_GongsimStatus.clear()

						gongsimstatuss.forEach(
							gongsimstatusDB => {
								let gongsimstatus = new GongsimStatus
								this.frontRepo.array_GongsimStatuss.push(gongsimstatus)
								this.frontRepo.map_ID_GongsimStatus.set(gongsimstatusDB.ID, gongsimstatus)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						dummyagents.forEach(
							dummyagentDB => {
								let dummyagent = this.frontRepo.map_ID_DummyAgent.get(dummyagentDB.ID)
								CopyDummyAgentDBToDummyAgent(dummyagentDB, dummyagent!, this.frontRepo)
							}
						)

						// fill up front objects
						engines.forEach(
							engineDB => {
								let engine = this.frontRepo.map_ID_Engine.get(engineDB.ID)
								CopyEngineDBToEngine(engineDB, engine!, this.frontRepo)
							}
						)

						// fill up front objects
						events.forEach(
							eventDB => {
								let event = this.frontRepo.map_ID_Event.get(eventDB.ID)
								CopyEventDBToEvent(eventDB, event!, this.frontRepo)
							}
						)

						// fill up front objects
						gongsimcommands.forEach(
							gongsimcommandDB => {
								let gongsimcommand = this.frontRepo.map_ID_GongsimCommand.get(gongsimcommandDB.ID)
								CopyGongsimCommandDBToGongsimCommand(gongsimcommandDB, gongsimcommand!, this.frontRepo)
							}
						)

						// fill up front objects
						gongsimstatuss.forEach(
							gongsimstatusDB => {
								let gongsimstatus = this.frontRepo.map_ID_GongsimStatus.get(gongsimstatusDB.ID)
								CopyGongsimStatusDBToGongsimStatus(gongsimstatusDB, gongsimstatus!, this.frontRepo)
							}
						)


						// hand over control flow to observer
						observer.next(this.frontRepo)
					}
				)
			}
		)
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

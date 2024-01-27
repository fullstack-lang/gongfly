import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CircleDB } from './circle-db'
import { Circle, CopyCircleDBToCircle } from './circle'
import { CircleService } from './circle.service'

import { DivIconDB } from './divicon-db'
import { DivIcon, CopyDivIconDBToDivIcon } from './divicon'
import { DivIconService } from './divicon.service'

import { LayerGroupDB } from './layergroup-db'
import { LayerGroup, CopyLayerGroupDBToLayerGroup } from './layergroup'
import { LayerGroupService } from './layergroup.service'

import { LayerGroupUseDB } from './layergroupuse-db'
import { LayerGroupUse, CopyLayerGroupUseDBToLayerGroupUse } from './layergroupuse'
import { LayerGroupUseService } from './layergroupuse.service'

import { MapOptionsDB } from './mapoptions-db'
import { MapOptions, CopyMapOptionsDBToMapOptions } from './mapoptions'
import { MapOptionsService } from './mapoptions.service'

import { MarkerDB } from './marker-db'
import { Marker, CopyMarkerDBToMarker } from './marker'
import { MarkerService } from './marker.service'

import { UserClickDB } from './userclick-db'
import { UserClick, CopyUserClickDBToUserClick } from './userclick'
import { UserClickService } from './userclick.service'

import { VLineDB } from './vline-db'
import { VLine, CopyVLineDBToVLine } from './vline'
import { VLineService } from './vline.service'

import { VisualTrackDB } from './visualtrack-db'
import { VisualTrack, CopyVisualTrackDBToVisualTrack } from './visualtrack'
import { VisualTrackService } from './visualtrack.service'

export const StackType = "github.com/fullstack-lang/gongleaflet/go/models"

// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
	array_Circles = new Array<Circle>() // array of front instances
	map_ID_Circle = new Map<number, Circle>() // map of front instances

	array_DivIcons = new Array<DivIcon>() // array of front instances
	map_ID_DivIcon = new Map<number, DivIcon>() // map of front instances

	array_LayerGroups = new Array<LayerGroup>() // array of front instances
	map_ID_LayerGroup = new Map<number, LayerGroup>() // map of front instances

	array_LayerGroupUses = new Array<LayerGroupUse>() // array of front instances
	map_ID_LayerGroupUse = new Map<number, LayerGroupUse>() // map of front instances

	array_MapOptionss = new Array<MapOptions>() // array of front instances
	map_ID_MapOptions = new Map<number, MapOptions>() // map of front instances

	array_Markers = new Array<Marker>() // array of front instances
	map_ID_Marker = new Map<number, Marker>() // map of front instances

	array_UserClicks = new Array<UserClick>() // array of front instances
	map_ID_UserClick = new Map<number, UserClick>() // map of front instances

	array_VLines = new Array<VLine>() // array of front instances
	map_ID_VLine = new Map<number, VLine>() // map of front instances

	array_VisualTracks = new Array<VisualTrack>() // array of front instances
	map_ID_VisualTrack = new Map<number, VisualTrack>() // map of front instances


	// getFrontArray allows for a get function that is robust to refactoring of the named struct name
	// for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
	// contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
	getFrontArray<Type>(gongStructName: string): Array<Type> {
		switch (gongStructName) {
			// insertion point
			case 'Circle':
				return this.array_Circles as unknown as Array<Type>
			case 'DivIcon':
				return this.array_DivIcons as unknown as Array<Type>
			case 'LayerGroup':
				return this.array_LayerGroups as unknown as Array<Type>
			case 'LayerGroupUse':
				return this.array_LayerGroupUses as unknown as Array<Type>
			case 'MapOptions':
				return this.array_MapOptionss as unknown as Array<Type>
			case 'Marker':
				return this.array_Markers as unknown as Array<Type>
			case 'UserClick':
				return this.array_UserClicks as unknown as Array<Type>
			case 'VLine':
				return this.array_VLines as unknown as Array<Type>
			case 'VisualTrack':
				return this.array_VisualTracks as unknown as Array<Type>
			default:
				throw new Error("Type not recognized");
		}
	}
	
	getFrontMap<Type>(gongStructName: string): Map<number, Type> {
		switch (gongStructName) {
			// insertion point
			case 'Circle':
				return this.map_ID_Circle as unknown as Map<number, Type>
			case 'DivIcon':
				return this.map_ID_DivIcon as unknown as Map<number, Type>
			case 'LayerGroup':
				return this.map_ID_LayerGroup as unknown as Map<number, Type>
			case 'LayerGroupUse':
				return this.map_ID_LayerGroupUse as unknown as Map<number, Type>
			case 'MapOptions':
				return this.map_ID_MapOptions as unknown as Map<number, Type>
			case 'Marker':
				return this.map_ID_Marker as unknown as Map<number, Type>
			case 'UserClick':
				return this.map_ID_UserClick as unknown as Map<number, Type>
			case 'VLine':
				return this.map_ID_VLine as unknown as Map<number, Type>
			case 'VisualTrack':
				return this.map_ID_VisualTrack as unknown as Map<number, Type>
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
		private circleService: CircleService,
		private diviconService: DivIconService,
		private layergroupService: LayerGroupService,
		private layergroupuseService: LayerGroupUseService,
		private mapoptionsService: MapOptionsService,
		private markerService: MarkerService,
		private userclickService: UserClickService,
		private vlineService: VLineService,
		private visualtrackService: VisualTrackService,
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
		Observable<CircleDB[]>,
		Observable<DivIconDB[]>,
		Observable<LayerGroupDB[]>,
		Observable<LayerGroupUseDB[]>,
		Observable<MapOptionsDB[]>,
		Observable<MarkerDB[]>,
		Observable<UserClickDB[]>,
		Observable<VLineDB[]>,
		Observable<VisualTrackDB[]>,
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
			this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
			this.diviconService.getDivIcons(this.GONG__StackPath, this.frontRepo),
			this.layergroupService.getLayerGroups(this.GONG__StackPath, this.frontRepo),
			this.layergroupuseService.getLayerGroupUses(this.GONG__StackPath, this.frontRepo),
			this.mapoptionsService.getMapOptionss(this.GONG__StackPath, this.frontRepo),
			this.markerService.getMarkers(this.GONG__StackPath, this.frontRepo),
			this.userclickService.getUserClicks(this.GONG__StackPath, this.frontRepo),
			this.vlineService.getVLines(this.GONG__StackPath, this.frontRepo),
			this.visualtrackService.getVisualTracks(this.GONG__StackPath, this.frontRepo),
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
			this.circleService.getCircles(this.GONG__StackPath, this.frontRepo),
			this.diviconService.getDivIcons(this.GONG__StackPath, this.frontRepo),
			this.layergroupService.getLayerGroups(this.GONG__StackPath, this.frontRepo),
			this.layergroupuseService.getLayerGroupUses(this.GONG__StackPath, this.frontRepo),
			this.mapoptionsService.getMapOptionss(this.GONG__StackPath, this.frontRepo),
			this.markerService.getMarkers(this.GONG__StackPath, this.frontRepo),
			this.userclickService.getUserClicks(this.GONG__StackPath, this.frontRepo),
			this.vlineService.getVLines(this.GONG__StackPath, this.frontRepo),
			this.visualtrackService.getVisualTracks(this.GONG__StackPath, this.frontRepo),
		]

		return new Observable<FrontRepo>(
			(observer) => {
				combineLatest(
					this.observableFrontRepo
				).subscribe(
					([
						___of_null, // see above for the explanation about of
						// insertion point sub template for declarations 
						circles_,
						divicons_,
						layergroups_,
						layergroupuses_,
						mapoptionss_,
						markers_,
						userclicks_,
						vlines_,
						visualtracks_,
					]) => {
						let _this = this
						// Typing can be messy with many items. Therefore, type casting is necessary here
						// insertion point sub template for type casting 
						var circles: CircleDB[]
						circles = circles_ as CircleDB[]
						var divicons: DivIconDB[]
						divicons = divicons_ as DivIconDB[]
						var layergroups: LayerGroupDB[]
						layergroups = layergroups_ as LayerGroupDB[]
						var layergroupuses: LayerGroupUseDB[]
						layergroupuses = layergroupuses_ as LayerGroupUseDB[]
						var mapoptionss: MapOptionsDB[]
						mapoptionss = mapoptionss_ as MapOptionsDB[]
						var markers: MarkerDB[]
						markers = markers_ as MarkerDB[]
						var userclicks: UserClickDB[]
						userclicks = userclicks_ as UserClickDB[]
						var vlines: VLineDB[]
						vlines = vlines_ as VLineDB[]
						var visualtracks: VisualTrackDB[]
						visualtracks = visualtracks_ as VisualTrackDB[]

						// 
						// First Step: init map of instances
						// insertion point sub template for init 
						// init the arrays
						this.frontRepo.array_Circles = []
						this.frontRepo.map_ID_Circle.clear()

						circles.forEach(
							circleDB => {
								let circle = new Circle
								this.frontRepo.array_Circles.push(circle)
								this.frontRepo.map_ID_Circle.set(circleDB.ID, circle)
							}
						)

						// init the arrays
						this.frontRepo.array_DivIcons = []
						this.frontRepo.map_ID_DivIcon.clear()

						divicons.forEach(
							diviconDB => {
								let divicon = new DivIcon
								this.frontRepo.array_DivIcons.push(divicon)
								this.frontRepo.map_ID_DivIcon.set(diviconDB.ID, divicon)
							}
						)

						// init the arrays
						this.frontRepo.array_LayerGroups = []
						this.frontRepo.map_ID_LayerGroup.clear()

						layergroups.forEach(
							layergroupDB => {
								let layergroup = new LayerGroup
								this.frontRepo.array_LayerGroups.push(layergroup)
								this.frontRepo.map_ID_LayerGroup.set(layergroupDB.ID, layergroup)
							}
						)

						// init the arrays
						this.frontRepo.array_LayerGroupUses = []
						this.frontRepo.map_ID_LayerGroupUse.clear()

						layergroupuses.forEach(
							layergroupuseDB => {
								let layergroupuse = new LayerGroupUse
								this.frontRepo.array_LayerGroupUses.push(layergroupuse)
								this.frontRepo.map_ID_LayerGroupUse.set(layergroupuseDB.ID, layergroupuse)
							}
						)

						// init the arrays
						this.frontRepo.array_MapOptionss = []
						this.frontRepo.map_ID_MapOptions.clear()

						mapoptionss.forEach(
							mapoptionsDB => {
								let mapoptions = new MapOptions
								this.frontRepo.array_MapOptionss.push(mapoptions)
								this.frontRepo.map_ID_MapOptions.set(mapoptionsDB.ID, mapoptions)
							}
						)

						// init the arrays
						this.frontRepo.array_Markers = []
						this.frontRepo.map_ID_Marker.clear()

						markers.forEach(
							markerDB => {
								let marker = new Marker
								this.frontRepo.array_Markers.push(marker)
								this.frontRepo.map_ID_Marker.set(markerDB.ID, marker)
							}
						)

						// init the arrays
						this.frontRepo.array_UserClicks = []
						this.frontRepo.map_ID_UserClick.clear()

						userclicks.forEach(
							userclickDB => {
								let userclick = new UserClick
								this.frontRepo.array_UserClicks.push(userclick)
								this.frontRepo.map_ID_UserClick.set(userclickDB.ID, userclick)
							}
						)

						// init the arrays
						this.frontRepo.array_VLines = []
						this.frontRepo.map_ID_VLine.clear()

						vlines.forEach(
							vlineDB => {
								let vline = new VLine
								this.frontRepo.array_VLines.push(vline)
								this.frontRepo.map_ID_VLine.set(vlineDB.ID, vline)
							}
						)

						// init the arrays
						this.frontRepo.array_VisualTracks = []
						this.frontRepo.map_ID_VisualTrack.clear()

						visualtracks.forEach(
							visualtrackDB => {
								let visualtrack = new VisualTrack
								this.frontRepo.array_VisualTracks.push(visualtrack)
								this.frontRepo.map_ID_VisualTrack.set(visualtrackDB.ID, visualtrack)
							}
						)


						// 
						// Second Step: reddeem front objects
						// insertion point sub template for redeem 
						// fill up front objects
						circles.forEach(
							circleDB => {
								let circle = this.frontRepo.map_ID_Circle.get(circleDB.ID)
								CopyCircleDBToCircle(circleDB, circle!, this.frontRepo)
							}
						)

						// fill up front objects
						divicons.forEach(
							diviconDB => {
								let divicon = this.frontRepo.map_ID_DivIcon.get(diviconDB.ID)
								CopyDivIconDBToDivIcon(diviconDB, divicon!, this.frontRepo)
							}
						)

						// fill up front objects
						layergroups.forEach(
							layergroupDB => {
								let layergroup = this.frontRepo.map_ID_LayerGroup.get(layergroupDB.ID)
								CopyLayerGroupDBToLayerGroup(layergroupDB, layergroup!, this.frontRepo)
							}
						)

						// fill up front objects
						layergroupuses.forEach(
							layergroupuseDB => {
								let layergroupuse = this.frontRepo.map_ID_LayerGroupUse.get(layergroupuseDB.ID)
								CopyLayerGroupUseDBToLayerGroupUse(layergroupuseDB, layergroupuse!, this.frontRepo)
							}
						)

						// fill up front objects
						mapoptionss.forEach(
							mapoptionsDB => {
								let mapoptions = this.frontRepo.map_ID_MapOptions.get(mapoptionsDB.ID)
								CopyMapOptionsDBToMapOptions(mapoptionsDB, mapoptions!, this.frontRepo)
							}
						)

						// fill up front objects
						markers.forEach(
							markerDB => {
								let marker = this.frontRepo.map_ID_Marker.get(markerDB.ID)
								CopyMarkerDBToMarker(markerDB, marker!, this.frontRepo)
							}
						)

						// fill up front objects
						userclicks.forEach(
							userclickDB => {
								let userclick = this.frontRepo.map_ID_UserClick.get(userclickDB.ID)
								CopyUserClickDBToUserClick(userclickDB, userclick!, this.frontRepo)
							}
						)

						// fill up front objects
						vlines.forEach(
							vlineDB => {
								let vline = this.frontRepo.map_ID_VLine.get(vlineDB.ID)
								CopyVLineDBToVLine(vlineDB, vline!, this.frontRepo)
							}
						)

						// fill up front objects
						visualtracks.forEach(
							visualtrackDB => {
								let visualtrack = this.frontRepo.map_ID_VisualTrack.get(visualtrackDB.ID)
								CopyVisualTrackDBToVisualTrack(visualtrackDB, visualtrack!, this.frontRepo)
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
export function getCircleUniqueID(id: number): number {
	return 31 * id
}
export function getDivIconUniqueID(id: number): number {
	return 37 * id
}
export function getLayerGroupUniqueID(id: number): number {
	return 41 * id
}
export function getLayerGroupUseUniqueID(id: number): number {
	return 43 * id
}
export function getMapOptionsUniqueID(id: number): number {
	return 47 * id
}
export function getMarkerUniqueID(id: number): number {
	return 53 * id
}
export function getUserClickUniqueID(id: number): number {
	return 59 * id
}
export function getVLineUniqueID(id: number): number {
	return 61 * id
}
export function getVisualTrackUniqueID(id: number): number {
	return 67 * id
}

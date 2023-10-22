import { Injectable } from '@angular/core'
import { HttpClient, HttpHeaders } from '@angular/common/http'

import { Observable, combineLatest, BehaviorSubject, of } from 'rxjs'

// insertion point sub template for services imports
import { CivilianAirportDB } from './civilianairport-db'
import { CivilianAirportService } from './civilianairport.service'

import { LinerDB } from './liner-db'
import { LinerService } from './liner.service'

import { MessageDB } from './message-db'
import { MessageService } from './message.service'

import { OpsLineDB } from './opsline-db'
import { OpsLineService } from './opsline.service'

import { RadarDB } from './radar-db'
import { RadarService } from './radar.service'

import { SatelliteDB } from './satellite-db'
import { SatelliteService } from './satellite.service'

import { ScenarioDB } from './scenario-db'
import { ScenarioService } from './scenario.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template
  CivilianAirports_array = new Array<CivilianAirportDB>() // array of repo instances
  CivilianAirports = new Map<number, CivilianAirportDB>() // map of repo instances
  CivilianAirports_batch = new Map<number, CivilianAirportDB>() // same but only in last GET (for finding repo instances to delete)

  Liners_array = new Array<LinerDB>() // array of repo instances
  Liners = new Map<number, LinerDB>() // map of repo instances
  Liners_batch = new Map<number, LinerDB>() // same but only in last GET (for finding repo instances to delete)

  Messages_array = new Array<MessageDB>() // array of repo instances
  Messages = new Map<number, MessageDB>() // map of repo instances
  Messages_batch = new Map<number, MessageDB>() // same but only in last GET (for finding repo instances to delete)

  OpsLines_array = new Array<OpsLineDB>() // array of repo instances
  OpsLines = new Map<number, OpsLineDB>() // map of repo instances
  OpsLines_batch = new Map<number, OpsLineDB>() // same but only in last GET (for finding repo instances to delete)

  Radars_array = new Array<RadarDB>() // array of repo instances
  Radars = new Map<number, RadarDB>() // map of repo instances
  Radars_batch = new Map<number, RadarDB>() // same but only in last GET (for finding repo instances to delete)

  Satellites_array = new Array<SatelliteDB>() // array of repo instances
  Satellites = new Map<number, SatelliteDB>() // map of repo instances
  Satellites_batch = new Map<number, SatelliteDB>() // same but only in last GET (for finding repo instances to delete)

  Scenarios_array = new Array<ScenarioDB>() // array of repo instances
  Scenarios = new Map<number, ScenarioDB>() // map of repo instances
  Scenarios_batch = new Map<number, ScenarioDB>() // same but only in last GET (for finding repo instances to delete)


  // getArray allows for a get function that is robust to refactoring of the named struct name
  // for instance frontRepo.getArray<Astruct>( Astruct.GONGSTRUCT_NAME), is robust to a refactoring of Astruct identifier
  // contrary to frontRepo.Astructs_array which is not refactored when Astruct identifier is modified
  getArray<Type>(gongStructName: string): Array<Type> {
    switch (gongStructName) {
      // insertion point
      case 'CivilianAirport':
        return this.CivilianAirports_array as unknown as Array<Type>
      case 'Liner':
        return this.Liners_array as unknown as Array<Type>
      case 'Message':
        return this.Messages_array as unknown as Array<Type>
      case 'OpsLine':
        return this.OpsLines_array as unknown as Array<Type>
      case 'Radar':
        return this.Radars_array as unknown as Array<Type>
      case 'Satellite':
        return this.Satellites_array as unknown as Array<Type>
      case 'Scenario':
        return this.Scenarios_array as unknown as Array<Type>
      default:
        throw new Error("Type not recognized");
    }
  }

  // getMap allows for a get function that is robust to refactoring of the named struct name
  getMap<Type>(gongStructName: string): Map<number, Type> {
    switch (gongStructName) {
      // insertion point
      case 'CivilianAirport':
        return this.CivilianAirports_array as unknown as Map<number, Type>
      case 'Liner':
        return this.Liners_array as unknown as Map<number, Type>
      case 'Message':
        return this.Messages_array as unknown as Map<number, Type>
      case 'OpsLine':
        return this.OpsLines_array as unknown as Map<number, Type>
      case 'Radar':
        return this.Radars_array as unknown as Map<number, Type>
      case 'Satellite':
        return this.Satellites_array as unknown as Map<number, Type>
      case 'Scenario':
        return this.Scenarios_array as unknown as Map<number, Type>
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
  SourceStruct: string = ""  // The "Aclass"
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
    private civilianairportService: CivilianAirportService,
    private linerService: LinerService,
    private messageService: MessageService,
    private opslineService: OpsLineService,
    private radarService: RadarService,
    private satelliteService: SatelliteService,
    private scenarioService: ScenarioService,
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
    Observable<CivilianAirportDB[]>,
    Observable<LinerDB[]>,
    Observable<MessageDB[]>,
    Observable<OpsLineDB[]>,
    Observable<RadarDB[]>,
    Observable<SatelliteDB[]>,
    Observable<ScenarioDB[]>,
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
      this.civilianairportService.getCivilianAirports(this.GONG__StackPath, this.frontRepo),
      this.linerService.getLiners(this.GONG__StackPath, this.frontRepo),
      this.messageService.getMessages(this.GONG__StackPath, this.frontRepo),
      this.opslineService.getOpsLines(this.GONG__StackPath, this.frontRepo),
      this.radarService.getRadars(this.GONG__StackPath, this.frontRepo),
      this.satelliteService.getSatellites(this.GONG__StackPath, this.frontRepo),
      this.scenarioService.getScenarios(this.GONG__StackPath, this.frontRepo),
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
      this.civilianairportService.getCivilianAirports(this.GONG__StackPath, this.frontRepo),
      this.linerService.getLiners(this.GONG__StackPath, this.frontRepo),
      this.messageService.getMessages(this.GONG__StackPath, this.frontRepo),
      this.opslineService.getOpsLines(this.GONG__StackPath, this.frontRepo),
      this.radarService.getRadars(this.GONG__StackPath, this.frontRepo),
      this.satelliteService.getSatellites(this.GONG__StackPath, this.frontRepo),
      this.scenarioService.getScenarios(this.GONG__StackPath, this.frontRepo),
    ]

    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([
            ___of_null, // see above for the explanation about of
            // insertion point sub template for declarations 
            civilianairports_,
            liners_,
            messages_,
            opslines_,
            radars_,
            satellites_,
            scenarios_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var civilianairports: CivilianAirportDB[]
            civilianairports = civilianairports_ as CivilianAirportDB[]
            var liners: LinerDB[]
            liners = liners_ as LinerDB[]
            var messages: MessageDB[]
            messages = messages_ as MessageDB[]
            var opslines: OpsLineDB[]
            opslines = opslines_ as OpsLineDB[]
            var radars: RadarDB[]
            radars = radars_ as RadarDB[]
            var satellites: SatelliteDB[]
            satellites = satellites_ as SatelliteDB[]
            var scenarios: ScenarioDB[]
            scenarios = scenarios_ as ScenarioDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            this.frontRepo.CivilianAirports_array = civilianairports

            // clear the map that counts CivilianAirport in the GET
            this.frontRepo.CivilianAirports_batch.clear()

            civilianairports.forEach(
              civilianairport => {
                this.frontRepo.CivilianAirports.set(civilianairport.ID, civilianairport)
                this.frontRepo.CivilianAirports_batch.set(civilianairport.ID, civilianairport)
              }
            )

            // clear civilianairports that are absent from the batch
            this.frontRepo.CivilianAirports.forEach(
              civilianairport => {
                if (this.frontRepo.CivilianAirports_batch.get(civilianairport.ID) == undefined) {
                  this.frontRepo.CivilianAirports.delete(civilianairport.ID)
                }
              }
            )

            // sort CivilianAirports_array array
            this.frontRepo.CivilianAirports_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Liners_array = liners

            // clear the map that counts Liner in the GET
            this.frontRepo.Liners_batch.clear()

            liners.forEach(
              liner => {
                this.frontRepo.Liners.set(liner.ID, liner)
                this.frontRepo.Liners_batch.set(liner.ID, liner)
              }
            )

            // clear liners that are absent from the batch
            this.frontRepo.Liners.forEach(
              liner => {
                if (this.frontRepo.Liners_batch.get(liner.ID) == undefined) {
                  this.frontRepo.Liners.delete(liner.ID)
                }
              }
            )

            // sort Liners_array array
            this.frontRepo.Liners_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Messages_array = messages

            // clear the map that counts Message in the GET
            this.frontRepo.Messages_batch.clear()

            messages.forEach(
              message => {
                this.frontRepo.Messages.set(message.ID, message)
                this.frontRepo.Messages_batch.set(message.ID, message)
              }
            )

            // clear messages that are absent from the batch
            this.frontRepo.Messages.forEach(
              message => {
                if (this.frontRepo.Messages_batch.get(message.ID) == undefined) {
                  this.frontRepo.Messages.delete(message.ID)
                }
              }
            )

            // sort Messages_array array
            this.frontRepo.Messages_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.OpsLines_array = opslines

            // clear the map that counts OpsLine in the GET
            this.frontRepo.OpsLines_batch.clear()

            opslines.forEach(
              opsline => {
                this.frontRepo.OpsLines.set(opsline.ID, opsline)
                this.frontRepo.OpsLines_batch.set(opsline.ID, opsline)
              }
            )

            // clear opslines that are absent from the batch
            this.frontRepo.OpsLines.forEach(
              opsline => {
                if (this.frontRepo.OpsLines_batch.get(opsline.ID) == undefined) {
                  this.frontRepo.OpsLines.delete(opsline.ID)
                }
              }
            )

            // sort OpsLines_array array
            this.frontRepo.OpsLines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Radars_array = radars

            // clear the map that counts Radar in the GET
            this.frontRepo.Radars_batch.clear()

            radars.forEach(
              radar => {
                this.frontRepo.Radars.set(radar.ID, radar)
                this.frontRepo.Radars_batch.set(radar.ID, radar)
              }
            )

            // clear radars that are absent from the batch
            this.frontRepo.Radars.forEach(
              radar => {
                if (this.frontRepo.Radars_batch.get(radar.ID) == undefined) {
                  this.frontRepo.Radars.delete(radar.ID)
                }
              }
            )

            // sort Radars_array array
            this.frontRepo.Radars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Satellites_array = satellites

            // clear the map that counts Satellite in the GET
            this.frontRepo.Satellites_batch.clear()

            satellites.forEach(
              satellite => {
                this.frontRepo.Satellites.set(satellite.ID, satellite)
                this.frontRepo.Satellites_batch.set(satellite.ID, satellite)
              }
            )

            // clear satellites that are absent from the batch
            this.frontRepo.Satellites.forEach(
              satellite => {
                if (this.frontRepo.Satellites_batch.get(satellite.ID) == undefined) {
                  this.frontRepo.Satellites.delete(satellite.ID)
                }
              }
            )

            // sort Satellites_array array
            this.frontRepo.Satellites_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            this.frontRepo.Scenarios_array = scenarios

            // clear the map that counts Scenario in the GET
            this.frontRepo.Scenarios_batch.clear()

            scenarios.forEach(
              scenario => {
                this.frontRepo.Scenarios.set(scenario.ID, scenario)
                this.frontRepo.Scenarios_batch.set(scenario.ID, scenario)
              }
            )

            // clear scenarios that are absent from the batch
            this.frontRepo.Scenarios.forEach(
              scenario => {
                if (this.frontRepo.Scenarios_batch.get(scenario.ID) == undefined) {
                  this.frontRepo.Scenarios.delete(scenario.ID)
                }
              }
            )

            // sort Scenarios_array array
            this.frontRepo.Scenarios_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: reddeem slice of pointers fields
            // insertion point sub template for redeem 
            civilianairports.forEach(
              civilianairport => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            liners.forEach(
              liner => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field ReporingLine redeeming
                {
                  let _opsline = this.frontRepo.OpsLines.get(liner.LinerPointersEncoding.ReporingLineID.Int64)
                  if (_opsline) {
                    liner.ReporingLine = _opsline
                  }
                }
                // insertion point for pointers decoding
              }
            )
            messages.forEach(
              message => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            opslines.forEach(
              opsline => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Scenario redeeming
                {
                  let _scenario = this.frontRepo.Scenarios.get(opsline.OpsLinePointersEncoding.ScenarioID.Int64)
                  if (_scenario) {
                    opsline.Scenario = _scenario
                  }
                }
                // insertion point for pointers decoding
              }
            )
            radars.forEach(
              radar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            satellites.forEach(
              satellite => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )
            scenarios.forEach(
              scenario => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointers decoding
              }
            )

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // CivilianAirportPull performs a GET on CivilianAirport of the stack and redeem association pointers 
  CivilianAirportPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.civilianairportService.getCivilianAirports(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            civilianairports,
          ]) => {
            // init the array
            this.frontRepo.CivilianAirports_array = civilianairports

            // clear the map that counts CivilianAirport in the GET
            this.frontRepo.CivilianAirports_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            civilianairports.forEach(
              civilianairport => {
                this.frontRepo.CivilianAirports.set(civilianairport.ID, civilianairport)
                this.frontRepo.CivilianAirports_batch.set(civilianairport.ID, civilianairport)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear civilianairports that are absent from the GET
            this.frontRepo.CivilianAirports.forEach(
              civilianairport => {
                if (this.frontRepo.CivilianAirports_batch.get(civilianairport.ID) == undefined) {
                  this.frontRepo.CivilianAirports.delete(civilianairport.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // LinerPull performs a GET on Liner of the stack and redeem association pointers 
  LinerPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.linerService.getLiners(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            liners,
          ]) => {
            // init the array
            this.frontRepo.Liners_array = liners

            // clear the map that counts Liner in the GET
            this.frontRepo.Liners_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            liners.forEach(
              liner => {
                this.frontRepo.Liners.set(liner.ID, liner)
                this.frontRepo.Liners_batch.set(liner.ID, liner)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field ReporingLine redeeming
                {
                  let _opsline = this.frontRepo.OpsLines.get(liner.LinerPointersEncoding.ReporingLineID.Int64)
                  if (_opsline) {
                    liner.ReporingLine = _opsline
                  }
                }
              }
            )

            // clear liners that are absent from the GET
            this.frontRepo.Liners.forEach(
              liner => {
                if (this.frontRepo.Liners_batch.get(liner.ID) == undefined) {
                  this.frontRepo.Liners.delete(liner.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // MessagePull performs a GET on Message of the stack and redeem association pointers 
  MessagePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.messageService.getMessages(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            messages,
          ]) => {
            // init the array
            this.frontRepo.Messages_array = messages

            // clear the map that counts Message in the GET
            this.frontRepo.Messages_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            messages.forEach(
              message => {
                this.frontRepo.Messages.set(message.ID, message)
                this.frontRepo.Messages_batch.set(message.ID, message)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear messages that are absent from the GET
            this.frontRepo.Messages.forEach(
              message => {
                if (this.frontRepo.Messages_batch.get(message.ID) == undefined) {
                  this.frontRepo.Messages.delete(message.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // OpsLinePull performs a GET on OpsLine of the stack and redeem association pointers 
  OpsLinePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.opslineService.getOpsLines(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            opslines,
          ]) => {
            // init the array
            this.frontRepo.OpsLines_array = opslines

            // clear the map that counts OpsLine in the GET
            this.frontRepo.OpsLines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            opslines.forEach(
              opsline => {
                this.frontRepo.OpsLines.set(opsline.ID, opsline)
                this.frontRepo.OpsLines_batch.set(opsline.ID, opsline)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Scenario redeeming
                {
                  let _scenario = this.frontRepo.Scenarios.get(opsline.OpsLinePointersEncoding.ScenarioID.Int64)
                  if (_scenario) {
                    opsline.Scenario = _scenario
                  }
                }
              }
            )

            // clear opslines that are absent from the GET
            this.frontRepo.OpsLines.forEach(
              opsline => {
                if (this.frontRepo.OpsLines_batch.get(opsline.ID) == undefined) {
                  this.frontRepo.OpsLines.delete(opsline.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // RadarPull performs a GET on Radar of the stack and redeem association pointers 
  RadarPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.radarService.getRadars(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            radars,
          ]) => {
            // init the array
            this.frontRepo.Radars_array = radars

            // clear the map that counts Radar in the GET
            this.frontRepo.Radars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            radars.forEach(
              radar => {
                this.frontRepo.Radars.set(radar.ID, radar)
                this.frontRepo.Radars_batch.set(radar.ID, radar)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear radars that are absent from the GET
            this.frontRepo.Radars.forEach(
              radar => {
                if (this.frontRepo.Radars_batch.get(radar.ID) == undefined) {
                  this.frontRepo.Radars.delete(radar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // SatellitePull performs a GET on Satellite of the stack and redeem association pointers 
  SatellitePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.satelliteService.getSatellites(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            satellites,
          ]) => {
            // init the array
            this.frontRepo.Satellites_array = satellites

            // clear the map that counts Satellite in the GET
            this.frontRepo.Satellites_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            satellites.forEach(
              satellite => {
                this.frontRepo.Satellites.set(satellite.ID, satellite)
                this.frontRepo.Satellites_batch.set(satellite.ID, satellite)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear satellites that are absent from the GET
            this.frontRepo.Satellites.forEach(
              satellite => {
                if (this.frontRepo.Satellites_batch.get(satellite.ID) == undefined) {
                  this.frontRepo.Satellites.delete(satellite.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }

  // ScenarioPull performs a GET on Scenario of the stack and redeem association pointers 
  ScenarioPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.scenarioService.getScenarios(this.GONG__StackPath, this.frontRepo)
        ]).subscribe(
          ([ // insertion point sub template 
            scenarios,
          ]) => {
            // init the array
            this.frontRepo.Scenarios_array = scenarios

            // clear the map that counts Scenario in the GET
            this.frontRepo.Scenarios_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            scenarios.forEach(
              scenario => {
                this.frontRepo.Scenarios.set(scenario.ID, scenario)
                this.frontRepo.Scenarios_batch.set(scenario.ID, scenario)

                // insertion point for redeeming ONE/ZERO-ONE associations
              }
            )

            // clear scenarios that are absent from the GET
            this.frontRepo.Scenarios.forEach(
              scenario => {
                if (this.frontRepo.Scenarios_batch.get(scenario.ID) == undefined) {
                  this.frontRepo.Scenarios.delete(scenario.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(this.frontRepo)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getCivilianAirportUniqueID(id: number): number {
  return 31 * id
}
export function getLinerUniqueID(id: number): number {
  return 37 * id
}
export function getMessageUniqueID(id: number): number {
  return 41 * id
}
export function getOpsLineUniqueID(id: number): number {
  return 43 * id
}
export function getRadarUniqueID(id: number): number {
  return 47 * id
}
export function getSatelliteUniqueID(id: number): number {
  return 53 * id
}
export function getScenarioUniqueID(id: number): number {
  return 59 * id
}

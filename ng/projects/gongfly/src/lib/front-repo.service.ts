import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { CivilianAirportDB } from './civilianairport-db'
import { CivilianAirportService } from './civilianairport.service'

import { LinerDB } from './liner-db'
import { LinerService } from './liner.service'

import { MessageDB } from './message-db'
import { MessageService } from './message.service'

import { OpsLineDB } from './opsline-db'
import { OpsLineService } from './opsline.service'

import { OrderDB } from './order-db'
import { OrderService } from './order.service'

import { RadarDB } from './radar-db'
import { RadarService } from './radar.service'

import { ReportDB } from './report-db'
import { ReportService } from './report.service'

import { SatelliteDB } from './satellite-db'
import { SatelliteService } from './satellite.service'

import { ScenarioDB } from './scenario-db'
import { ScenarioService } from './scenario.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  CivilianAirports_array = new Array<CivilianAirportDB>(); // array of repo instances
  CivilianAirports = new Map<number, CivilianAirportDB>(); // map of repo instances
  CivilianAirports_batch = new Map<number, CivilianAirportDB>(); // same but only in last GET (for finding repo instances to delete)
  Liners_array = new Array<LinerDB>(); // array of repo instances
  Liners = new Map<number, LinerDB>(); // map of repo instances
  Liners_batch = new Map<number, LinerDB>(); // same but only in last GET (for finding repo instances to delete)
  Messages_array = new Array<MessageDB>(); // array of repo instances
  Messages = new Map<number, MessageDB>(); // map of repo instances
  Messages_batch = new Map<number, MessageDB>(); // same but only in last GET (for finding repo instances to delete)
  OpsLines_array = new Array<OpsLineDB>(); // array of repo instances
  OpsLines = new Map<number, OpsLineDB>(); // map of repo instances
  OpsLines_batch = new Map<number, OpsLineDB>(); // same but only in last GET (for finding repo instances to delete)
  Orders_array = new Array<OrderDB>(); // array of repo instances
  Orders = new Map<number, OrderDB>(); // map of repo instances
  Orders_batch = new Map<number, OrderDB>(); // same but only in last GET (for finding repo instances to delete)
  Radars_array = new Array<RadarDB>(); // array of repo instances
  Radars = new Map<number, RadarDB>(); // map of repo instances
  Radars_batch = new Map<number, RadarDB>(); // same but only in last GET (for finding repo instances to delete)
  Reports_array = new Array<ReportDB>(); // array of repo instances
  Reports = new Map<number, ReportDB>(); // map of repo instances
  Reports_batch = new Map<number, ReportDB>(); // same but only in last GET (for finding repo instances to delete)
  Satellites_array = new Array<SatelliteDB>(); // array of repo instances
  Satellites = new Map<number, SatelliteDB>(); // map of repo instances
  Satellites_batch = new Map<number, SatelliteDB>(); // same but only in last GET (for finding repo instances to delete)
  Scenarios_array = new Array<ScenarioDB>(); // array of repo instances
  Scenarios = new Map<number, ScenarioDB>(); // map of repo instances
  Scenarios_batch = new Map<number, ScenarioDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

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

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private civilianairportService: CivilianAirportService,
    private linerService: LinerService,
    private messageService: MessageService,
    private opslineService: OpsLineService,
    private orderService: OrderService,
    private radarService: RadarService,
    private reportService: ReportService,
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
  observableFrontRepo: [ // insertion point sub template 
    Observable<CivilianAirportDB[]>,
    Observable<LinerDB[]>,
    Observable<MessageDB[]>,
    Observable<OpsLineDB[]>,
    Observable<OrderDB[]>,
    Observable<RadarDB[]>,
    Observable<ReportDB[]>,
    Observable<SatelliteDB[]>,
    Observable<ScenarioDB[]>,
  ] = [ // insertion point sub template 
      this.civilianairportService.getCivilianAirports(),
      this.linerService.getLiners(),
      this.messageService.getMessages(),
      this.opslineService.getOpsLines(),
      this.orderService.getOrders(),
      this.radarService.getRadars(),
      this.reportService.getReports(),
      this.satelliteService.getSatellites(),
      this.scenarioService.getScenarios(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            civilianairports_,
            liners_,
            messages_,
            opslines_,
            orders_,
            radars_,
            reports_,
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
            var orders: OrderDB[]
            orders = orders_ as OrderDB[]
            var radars: RadarDB[]
            radars = radars_ as RadarDB[]
            var reports: ReportDB[]
            reports = reports_ as ReportDB[]
            var satellites: SatelliteDB[]
            satellites = satellites_ as SatelliteDB[]
            var scenarios: ScenarioDB[]
            scenarios = scenarios_ as ScenarioDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.CivilianAirports_array = civilianairports

            // clear the map that counts CivilianAirport in the GET
            FrontRepoSingloton.CivilianAirports_batch.clear()

            civilianairports.forEach(
              civilianairport => {
                FrontRepoSingloton.CivilianAirports.set(civilianairport.ID, civilianairport)
                FrontRepoSingloton.CivilianAirports_batch.set(civilianairport.ID, civilianairport)
              }
            )

            // clear civilianairports that are absent from the batch
            FrontRepoSingloton.CivilianAirports.forEach(
              civilianairport => {
                if (FrontRepoSingloton.CivilianAirports_batch.get(civilianairport.ID) == undefined) {
                  FrontRepoSingloton.CivilianAirports.delete(civilianairport.ID)
                }
              }
            )

            // sort CivilianAirports_array array
            FrontRepoSingloton.CivilianAirports_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Liners_array = liners

            // clear the map that counts Liner in the GET
            FrontRepoSingloton.Liners_batch.clear()

            liners.forEach(
              liner => {
                FrontRepoSingloton.Liners.set(liner.ID, liner)
                FrontRepoSingloton.Liners_batch.set(liner.ID, liner)
              }
            )

            // clear liners that are absent from the batch
            FrontRepoSingloton.Liners.forEach(
              liner => {
                if (FrontRepoSingloton.Liners_batch.get(liner.ID) == undefined) {
                  FrontRepoSingloton.Liners.delete(liner.ID)
                }
              }
            )

            // sort Liners_array array
            FrontRepoSingloton.Liners_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Messages_array = messages

            // clear the map that counts Message in the GET
            FrontRepoSingloton.Messages_batch.clear()

            messages.forEach(
              message => {
                FrontRepoSingloton.Messages.set(message.ID, message)
                FrontRepoSingloton.Messages_batch.set(message.ID, message)
              }
            )

            // clear messages that are absent from the batch
            FrontRepoSingloton.Messages.forEach(
              message => {
                if (FrontRepoSingloton.Messages_batch.get(message.ID) == undefined) {
                  FrontRepoSingloton.Messages.delete(message.ID)
                }
              }
            )

            // sort Messages_array array
            FrontRepoSingloton.Messages_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.OpsLines_array = opslines

            // clear the map that counts OpsLine in the GET
            FrontRepoSingloton.OpsLines_batch.clear()

            opslines.forEach(
              opsline => {
                FrontRepoSingloton.OpsLines.set(opsline.ID, opsline)
                FrontRepoSingloton.OpsLines_batch.set(opsline.ID, opsline)
              }
            )

            // clear opslines that are absent from the batch
            FrontRepoSingloton.OpsLines.forEach(
              opsline => {
                if (FrontRepoSingloton.OpsLines_batch.get(opsline.ID) == undefined) {
                  FrontRepoSingloton.OpsLines.delete(opsline.ID)
                }
              }
            )

            // sort OpsLines_array array
            FrontRepoSingloton.OpsLines_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Orders_array = orders

            // clear the map that counts Order in the GET
            FrontRepoSingloton.Orders_batch.clear()

            orders.forEach(
              order => {
                FrontRepoSingloton.Orders.set(order.ID, order)
                FrontRepoSingloton.Orders_batch.set(order.ID, order)
              }
            )

            // clear orders that are absent from the batch
            FrontRepoSingloton.Orders.forEach(
              order => {
                if (FrontRepoSingloton.Orders_batch.get(order.ID) == undefined) {
                  FrontRepoSingloton.Orders.delete(order.ID)
                }
              }
            )

            // sort Orders_array array
            FrontRepoSingloton.Orders_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Radars_array = radars

            // clear the map that counts Radar in the GET
            FrontRepoSingloton.Radars_batch.clear()

            radars.forEach(
              radar => {
                FrontRepoSingloton.Radars.set(radar.ID, radar)
                FrontRepoSingloton.Radars_batch.set(radar.ID, radar)
              }
            )

            // clear radars that are absent from the batch
            FrontRepoSingloton.Radars.forEach(
              radar => {
                if (FrontRepoSingloton.Radars_batch.get(radar.ID) == undefined) {
                  FrontRepoSingloton.Radars.delete(radar.ID)
                }
              }
            )

            // sort Radars_array array
            FrontRepoSingloton.Radars_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Reports_array = reports

            // clear the map that counts Report in the GET
            FrontRepoSingloton.Reports_batch.clear()

            reports.forEach(
              report => {
                FrontRepoSingloton.Reports.set(report.ID, report)
                FrontRepoSingloton.Reports_batch.set(report.ID, report)
              }
            )

            // clear reports that are absent from the batch
            FrontRepoSingloton.Reports.forEach(
              report => {
                if (FrontRepoSingloton.Reports_batch.get(report.ID) == undefined) {
                  FrontRepoSingloton.Reports.delete(report.ID)
                }
              }
            )

            // sort Reports_array array
            FrontRepoSingloton.Reports_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Satellites_array = satellites

            // clear the map that counts Satellite in the GET
            FrontRepoSingloton.Satellites_batch.clear()

            satellites.forEach(
              satellite => {
                FrontRepoSingloton.Satellites.set(satellite.ID, satellite)
                FrontRepoSingloton.Satellites_batch.set(satellite.ID, satellite)
              }
            )

            // clear satellites that are absent from the batch
            FrontRepoSingloton.Satellites.forEach(
              satellite => {
                if (FrontRepoSingloton.Satellites_batch.get(satellite.ID) == undefined) {
                  FrontRepoSingloton.Satellites.delete(satellite.ID)
                }
              }
            )

            // sort Satellites_array array
            FrontRepoSingloton.Satellites_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Scenarios_array = scenarios

            // clear the map that counts Scenario in the GET
            FrontRepoSingloton.Scenarios_batch.clear()

            scenarios.forEach(
              scenario => {
                FrontRepoSingloton.Scenarios.set(scenario.ID, scenario)
                FrontRepoSingloton.Scenarios_batch.set(scenario.ID, scenario)
              }
            )

            // clear scenarios that are absent from the batch
            FrontRepoSingloton.Scenarios.forEach(
              scenario => {
                if (FrontRepoSingloton.Scenarios_batch.get(scenario.ID) == undefined) {
                  FrontRepoSingloton.Scenarios.delete(scenario.ID)
                }
              }
            )

            // sort Scenarios_array array
            FrontRepoSingloton.Scenarios_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            civilianairports.forEach(
              civilianairport => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            liners.forEach(
              liner => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field ReporingLine redeeming
                {
                  let _opsline = FrontRepoSingloton.OpsLines.get(liner.ReporingLineID.Int64)
                  if (_opsline) {
                    liner.ReporingLine = _opsline
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            messages.forEach(
              message => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            opslines.forEach(
              opsline => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Scenario redeeming
                {
                  let _scenario = FrontRepoSingloton.Scenarios.get(opsline.ScenarioID.Int64)
                  if (_scenario) {
                    opsline.Scenario = _scenario
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            orders.forEach(
              order => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Target redeeming
                {
                  let _liner = FrontRepoSingloton.Liners.get(order.TargetID.Int64)
                  if (_liner) {
                    order.Target = _liner
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            radars.forEach(
              radar => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            reports.forEach(
              report => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field About redeeming
                {
                  let _liner = FrontRepoSingloton.Liners.get(report.AboutID.Int64)
                  if (_liner) {
                    report.About = _liner
                  }
                }
                // insertion point for pointer field OpsLine redeeming
                {
                  let _opsline = FrontRepoSingloton.OpsLines.get(report.OpsLineID.Int64)
                  if (_opsline) {
                    report.OpsLine = _opsline
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            satellites.forEach(
              satellite => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            scenarios.forEach(
              scenario => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.civilianairportService.getCivilianAirports()
        ]).subscribe(
          ([ // insertion point sub template 
            civilianairports,
          ]) => {
            // init the array
            FrontRepoSingloton.CivilianAirports_array = civilianairports

            // clear the map that counts CivilianAirport in the GET
            FrontRepoSingloton.CivilianAirports_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            civilianairports.forEach(
              civilianairport => {
                FrontRepoSingloton.CivilianAirports.set(civilianairport.ID, civilianairport)
                FrontRepoSingloton.CivilianAirports_batch.set(civilianairport.ID, civilianairport)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear civilianairports that are absent from the GET
            FrontRepoSingloton.CivilianAirports.forEach(
              civilianairport => {
                if (FrontRepoSingloton.CivilianAirports_batch.get(civilianairport.ID) == undefined) {
                  FrontRepoSingloton.CivilianAirports.delete(civilianairport.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.linerService.getLiners()
        ]).subscribe(
          ([ // insertion point sub template 
            liners,
          ]) => {
            // init the array
            FrontRepoSingloton.Liners_array = liners

            // clear the map that counts Liner in the GET
            FrontRepoSingloton.Liners_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            liners.forEach(
              liner => {
                FrontRepoSingloton.Liners.set(liner.ID, liner)
                FrontRepoSingloton.Liners_batch.set(liner.ID, liner)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field ReporingLine redeeming
                {
                  let _opsline = FrontRepoSingloton.OpsLines.get(liner.ReporingLineID.Int64)
                  if (_opsline) {
                    liner.ReporingLine = _opsline
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear liners that are absent from the GET
            FrontRepoSingloton.Liners.forEach(
              liner => {
                if (FrontRepoSingloton.Liners_batch.get(liner.ID) == undefined) {
                  FrontRepoSingloton.Liners.delete(liner.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.messageService.getMessages()
        ]).subscribe(
          ([ // insertion point sub template 
            messages,
          ]) => {
            // init the array
            FrontRepoSingloton.Messages_array = messages

            // clear the map that counts Message in the GET
            FrontRepoSingloton.Messages_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            messages.forEach(
              message => {
                FrontRepoSingloton.Messages.set(message.ID, message)
                FrontRepoSingloton.Messages_batch.set(message.ID, message)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear messages that are absent from the GET
            FrontRepoSingloton.Messages.forEach(
              message => {
                if (FrontRepoSingloton.Messages_batch.get(message.ID) == undefined) {
                  FrontRepoSingloton.Messages.delete(message.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.opslineService.getOpsLines()
        ]).subscribe(
          ([ // insertion point sub template 
            opslines,
          ]) => {
            // init the array
            FrontRepoSingloton.OpsLines_array = opslines

            // clear the map that counts OpsLine in the GET
            FrontRepoSingloton.OpsLines_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            opslines.forEach(
              opsline => {
                FrontRepoSingloton.OpsLines.set(opsline.ID, opsline)
                FrontRepoSingloton.OpsLines_batch.set(opsline.ID, opsline)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Scenario redeeming
                {
                  let _scenario = FrontRepoSingloton.Scenarios.get(opsline.ScenarioID.Int64)
                  if (_scenario) {
                    opsline.Scenario = _scenario
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear opslines that are absent from the GET
            FrontRepoSingloton.OpsLines.forEach(
              opsline => {
                if (FrontRepoSingloton.OpsLines_batch.get(opsline.ID) == undefined) {
                  FrontRepoSingloton.OpsLines.delete(opsline.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // OrderPull performs a GET on Order of the stack and redeem association pointers 
  OrderPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.orderService.getOrders()
        ]).subscribe(
          ([ // insertion point sub template 
            orders,
          ]) => {
            // init the array
            FrontRepoSingloton.Orders_array = orders

            // clear the map that counts Order in the GET
            FrontRepoSingloton.Orders_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            orders.forEach(
              order => {
                FrontRepoSingloton.Orders.set(order.ID, order)
                FrontRepoSingloton.Orders_batch.set(order.ID, order)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Target redeeming
                {
                  let _liner = FrontRepoSingloton.Liners.get(order.TargetID.Int64)
                  if (_liner) {
                    order.Target = _liner
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear orders that are absent from the GET
            FrontRepoSingloton.Orders.forEach(
              order => {
                if (FrontRepoSingloton.Orders_batch.get(order.ID) == undefined) {
                  FrontRepoSingloton.Orders.delete(order.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.radarService.getRadars()
        ]).subscribe(
          ([ // insertion point sub template 
            radars,
          ]) => {
            // init the array
            FrontRepoSingloton.Radars_array = radars

            // clear the map that counts Radar in the GET
            FrontRepoSingloton.Radars_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            radars.forEach(
              radar => {
                FrontRepoSingloton.Radars.set(radar.ID, radar)
                FrontRepoSingloton.Radars_batch.set(radar.ID, radar)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear radars that are absent from the GET
            FrontRepoSingloton.Radars.forEach(
              radar => {
                if (FrontRepoSingloton.Radars_batch.get(radar.ID) == undefined) {
                  FrontRepoSingloton.Radars.delete(radar.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // ReportPull performs a GET on Report of the stack and redeem association pointers 
  ReportPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.reportService.getReports()
        ]).subscribe(
          ([ // insertion point sub template 
            reports,
          ]) => {
            // init the array
            FrontRepoSingloton.Reports_array = reports

            // clear the map that counts Report in the GET
            FrontRepoSingloton.Reports_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            reports.forEach(
              report => {
                FrontRepoSingloton.Reports.set(report.ID, report)
                FrontRepoSingloton.Reports_batch.set(report.ID, report)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field About redeeming
                {
                  let _liner = FrontRepoSingloton.Liners.get(report.AboutID.Int64)
                  if (_liner) {
                    report.About = _liner
                  }
                }
                // insertion point for pointer field OpsLine redeeming
                {
                  let _opsline = FrontRepoSingloton.OpsLines.get(report.OpsLineID.Int64)
                  if (_opsline) {
                    report.OpsLine = _opsline
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear reports that are absent from the GET
            FrontRepoSingloton.Reports.forEach(
              report => {
                if (FrontRepoSingloton.Reports_batch.get(report.ID) == undefined) {
                  FrontRepoSingloton.Reports.delete(report.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.satelliteService.getSatellites()
        ]).subscribe(
          ([ // insertion point sub template 
            satellites,
          ]) => {
            // init the array
            FrontRepoSingloton.Satellites_array = satellites

            // clear the map that counts Satellite in the GET
            FrontRepoSingloton.Satellites_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            satellites.forEach(
              satellite => {
                FrontRepoSingloton.Satellites.set(satellite.ID, satellite)
                FrontRepoSingloton.Satellites_batch.set(satellite.ID, satellite)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear satellites that are absent from the GET
            FrontRepoSingloton.Satellites.forEach(
              satellite => {
                if (FrontRepoSingloton.Satellites_batch.get(satellite.ID) == undefined) {
                  FrontRepoSingloton.Satellites.delete(satellite.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
          this.scenarioService.getScenarios()
        ]).subscribe(
          ([ // insertion point sub template 
            scenarios,
          ]) => {
            // init the array
            FrontRepoSingloton.Scenarios_array = scenarios

            // clear the map that counts Scenario in the GET
            FrontRepoSingloton.Scenarios_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            scenarios.forEach(
              scenario => {
                FrontRepoSingloton.Scenarios.set(scenario.ID, scenario)
                FrontRepoSingloton.Scenarios_batch.set(scenario.ID, scenario)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear scenarios that are absent from the GET
            FrontRepoSingloton.Scenarios.forEach(
              scenario => {
                if (FrontRepoSingloton.Scenarios_batch.get(scenario.ID) == undefined) {
                  FrontRepoSingloton.Scenarios.delete(scenario.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
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
export function getOrderUniqueID(id: number): number {
  return 47 * id
}
export function getRadarUniqueID(id: number): number {
  return 53 * id
}
export function getReportUniqueID(id: number): number {
  return 59 * id
}
export function getSatelliteUniqueID(id: number): number {
  return 61 * id
}
export function getScenarioUniqueID(id: number): number {
  return 67 * id
}

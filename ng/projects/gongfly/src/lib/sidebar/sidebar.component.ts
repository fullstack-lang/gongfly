import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbFromBackService } from '../commitnbfromback.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { CivilianAirportService } from '../civilianairport.service'
import { getCivilianAirportUniqueID } from '../front-repo.service'
import { LinerService } from '../liner.service'
import { getLinerUniqueID } from '../front-repo.service'
import { MessageService } from '../message.service'
import { getMessageUniqueID } from '../front-repo.service'
import { OpsLineService } from '../opsline.service'
import { getOpsLineUniqueID } from '../front-repo.service'
import { RadarService } from '../radar.service'
import { getRadarUniqueID } from '../front-repo.service'
import { SatelliteService } from '../satellite.service'
import { getSatelliteUniqueID } from '../front-repo.service'
import { ScenarioService } from '../scenario.service'
import { getScenarioUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongfly-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNbFromBack: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbFromBackService: CommitNbFromBackService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private civilianairportService: CivilianAirportService,
    private linerService: LinerService,
    private messageService: MessageService,
    private opslineService: OpsLineService,
    private radarService: RadarService,
    private satelliteService: SatelliteService,
    private scenarioService: ScenarioService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.civilianairportService.CivilianAirportServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.linerService.LinerServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.messageService.MessageServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.opslineService.OpsLineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.radarService.RadarServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.satelliteService.SatelliteServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.scenarioService.ScenarioServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the CivilianAirport part of the mat tree
      */
      let civilianairportGongNodeStruct: GongNode = {
        name: "CivilianAirport",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "CivilianAirport",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(civilianairportGongNodeStruct)

      this.frontRepo.CivilianAirports_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.CivilianAirports_array.forEach(
        civilianairportDB => {
          let civilianairportGongNodeInstance: GongNode = {
            name: civilianairportDB.Name,
            type: GongNodeType.INSTANCE,
            id: civilianairportDB.ID,
            uniqueIdPerStack: getCivilianAirportUniqueID(civilianairportDB.ID),
            structName: "CivilianAirport",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          civilianairportGongNodeStruct.children!.push(civilianairportGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Liner part of the mat tree
      */
      let linerGongNodeStruct: GongNode = {
        name: "Liner",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Liner",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(linerGongNodeStruct)

      this.frontRepo.Liners_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Liners_array.forEach(
        linerDB => {
          let linerGongNodeInstance: GongNode = {
            name: linerDB.Name,
            type: GongNodeType.INSTANCE,
            id: linerDB.ID,
            uniqueIdPerStack: getLinerUniqueID(linerDB.ID),
            structName: "Liner",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          linerGongNodeStruct.children!.push(linerGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association ReporingLine
          */
          let ReporingLineGongNodeAssociation: GongNode = {
            name: "(OpsLine) ReporingLine",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: linerDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Liner",
            associationField: "ReporingLine",
            associatedStructName: "OpsLine",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          linerGongNodeInstance.children!.push(ReporingLineGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation ReporingLine
            */
          if (linerDB.ReporingLine != undefined) {
            let linerGongNodeInstance_ReporingLine: GongNode = {
              name: linerDB.ReporingLine.Name,
              type: GongNodeType.INSTANCE,
              id: linerDB.ReporingLine.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getLinerUniqueID(linerDB.ID)
                + 5 * getOpsLineUniqueID(linerDB.ReporingLine.ID),
              structName: "OpsLine",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ReporingLineGongNodeAssociation.children.push(linerGongNodeInstance_ReporingLine)
          }

        }
      )

      /**
      * fill up the Message part of the mat tree
      */
      let messageGongNodeStruct: GongNode = {
        name: "Message",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Message",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(messageGongNodeStruct)

      this.frontRepo.Messages_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Messages_array.forEach(
        messageDB => {
          let messageGongNodeInstance: GongNode = {
            name: messageDB.Name,
            type: GongNodeType.INSTANCE,
            id: messageDB.ID,
            uniqueIdPerStack: getMessageUniqueID(messageDB.ID),
            structName: "Message",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          messageGongNodeStruct.children!.push(messageGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the OpsLine part of the mat tree
      */
      let opslineGongNodeStruct: GongNode = {
        name: "OpsLine",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "OpsLine",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(opslineGongNodeStruct)

      this.frontRepo.OpsLines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.OpsLines_array.forEach(
        opslineDB => {
          let opslineGongNodeInstance: GongNode = {
            name: opslineDB.Name,
            type: GongNodeType.INSTANCE,
            id: opslineDB.ID,
            uniqueIdPerStack: getOpsLineUniqueID(opslineDB.ID),
            structName: "OpsLine",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          opslineGongNodeStruct.children!.push(opslineGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Scenario
          */
          let ScenarioGongNodeAssociation: GongNode = {
            name: "(Scenario) Scenario",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: opslineDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "OpsLine",
            associationField: "Scenario",
            associatedStructName: "Scenario",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          opslineGongNodeInstance.children!.push(ScenarioGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Scenario
            */
          if (opslineDB.Scenario != undefined) {
            let opslineGongNodeInstance_Scenario: GongNode = {
              name: opslineDB.Scenario.Name,
              type: GongNodeType.INSTANCE,
              id: opslineDB.Scenario.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getOpsLineUniqueID(opslineDB.ID)
                + 5 * getScenarioUniqueID(opslineDB.Scenario.ID),
              structName: "Scenario",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ScenarioGongNodeAssociation.children.push(opslineGongNodeInstance_Scenario)
          }

        }
      )

      /**
      * fill up the Radar part of the mat tree
      */
      let radarGongNodeStruct: GongNode = {
        name: "Radar",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Radar",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(radarGongNodeStruct)

      this.frontRepo.Radars_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Radars_array.forEach(
        radarDB => {
          let radarGongNodeInstance: GongNode = {
            name: radarDB.Name,
            type: GongNodeType.INSTANCE,
            id: radarDB.ID,
            uniqueIdPerStack: getRadarUniqueID(radarDB.ID),
            structName: "Radar",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          radarGongNodeStruct.children!.push(radarGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Satellite part of the mat tree
      */
      let satelliteGongNodeStruct: GongNode = {
        name: "Satellite",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Satellite",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(satelliteGongNodeStruct)

      this.frontRepo.Satellites_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Satellites_array.forEach(
        satelliteDB => {
          let satelliteGongNodeInstance: GongNode = {
            name: satelliteDB.Name,
            type: GongNodeType.INSTANCE,
            id: satelliteDB.ID,
            uniqueIdPerStack: getSatelliteUniqueID(satelliteDB.ID),
            structName: "Satellite",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          satelliteGongNodeStruct.children!.push(satelliteGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Scenario part of the mat tree
      */
      let scenarioGongNodeStruct: GongNode = {
        name: "Scenario",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Scenario",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(scenarioGongNodeStruct)

      this.frontRepo.Scenarios_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Scenarios_array.forEach(
        scenarioDB => {
          let scenarioGongNodeInstance: GongNode = {
            name: scenarioDB.Name,
            type: GongNodeType.INSTANCE,
            id: scenarioDB.ID,
            uniqueIdPerStack: getScenarioUniqueID(scenarioDB.ID),
            structName: "Scenario",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          scenarioGongNodeStruct.children!.push(scenarioGongNodeInstance)

          // insertion point for per field code
        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbFromBackService.getCommitNbFromBack().subscribe(
      commitNbFromBack => {
        this.commitNbFromBack = commitNbFromBack
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_table: ["github_com_fullstack_lang_gongfly_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongfly_go_table: ["github_com_fullstack_lang_gongfly_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + structName.toLowerCase() + "-detail", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}

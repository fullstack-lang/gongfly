import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { VisualCenterService } from '../visualcenter.service'
import { getVisualCenterUniqueID } from '../front-repo.service'
import { VisualCircleService } from '../visualcircle.service'
import { getVisualCircleUniqueID } from '../front-repo.service'
import { VisualIconService } from '../visualicon.service'
import { getVisualIconUniqueID } from '../front-repo.service'
import { VisualLayerService } from '../visuallayer.service'
import { getVisualLayerUniqueID } from '../front-repo.service'
import { VisualLineService } from '../visualline.service'
import { getVisualLineUniqueID } from '../front-repo.service'
import { VisualMapService } from '../visualmap.service'
import { getVisualMapUniqueID } from '../front-repo.service'
import { VisualTrackService } from '../visualtrack.service'
import { getVisualTrackUniqueID } from '../front-repo.service'

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
  selector: 'app-gongleaflet-sidebar',
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
  commitNb: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,

    // insertion point for per struct service declaration
    private visualcenterService: VisualCenterService,
    private visualcircleService: VisualCircleService,
    private visualiconService: VisualIconService,
    private visuallayerService: VisualLayerService,
    private visuallineService: VisualLineService,
    private visualmapService: VisualMapService,
    private visualtrackService: VisualTrackService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.visualcenterService.VisualCenterServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visualcircleService.VisualCircleServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visualiconService.VisualIconServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visuallayerService.VisualLayerServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visuallineService.VisualLineServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visualmapService.VisualMapServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.visualtrackService.VisualTrackServiceChanged.subscribe(
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
      * fill up the VisualCenter part of the mat tree
      */
      let visualcenterGongNodeStruct: GongNode = {
        name: "VisualCenter",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualCenter",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visualcenterGongNodeStruct)

      this.frontRepo.VisualCenters_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualCenters_array.forEach(
        visualcenterDB => {
          let visualcenterGongNodeInstance: GongNode = {
            name: visualcenterDB.Name,
            type: GongNodeType.INSTANCE,
            id: visualcenterDB.ID,
            uniqueIdPerStack: getVisualCenterUniqueID(visualcenterDB.ID),
            structName: "VisualCenter",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visualcenterGongNodeStruct.children!.push(visualcenterGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association VisualLayer
          */
          let VisualLayerGongNodeAssociation: GongNode = {
            name: "(VisualLayer) VisualLayer",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visualcenterDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualCenter",
            associationField: "VisualLayer",
            associatedStructName: "VisualLayer",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visualcenterGongNodeInstance.children!.push(VisualLayerGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualLayer
            */
          if (visualcenterDB.VisualLayer != undefined) {
            let visualcenterGongNodeInstance_VisualLayer: GongNode = {
              name: visualcenterDB.VisualLayer.Name,
              type: GongNodeType.INSTANCE,
              id: visualcenterDB.VisualLayer.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualCenterUniqueID(visualcenterDB.ID)
                + 5 * getVisualLayerUniqueID(visualcenterDB.VisualLayer.ID),
              structName: "VisualLayer",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualLayerGongNodeAssociation.children.push(visualcenterGongNodeInstance_VisualLayer)
          }

          /**
          * let append a node for the association VisualIcon
          */
          let VisualIconGongNodeAssociation: GongNode = {
            name: "(VisualIcon) VisualIcon",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visualcenterDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualCenter",
            associationField: "VisualIcon",
            associatedStructName: "VisualIcon",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visualcenterGongNodeInstance.children!.push(VisualIconGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualIcon
            */
          if (visualcenterDB.VisualIcon != undefined) {
            let visualcenterGongNodeInstance_VisualIcon: GongNode = {
              name: visualcenterDB.VisualIcon.Name,
              type: GongNodeType.INSTANCE,
              id: visualcenterDB.VisualIcon.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualCenterUniqueID(visualcenterDB.ID)
                + 5 * getVisualIconUniqueID(visualcenterDB.VisualIcon.ID),
              structName: "VisualIcon",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualIconGongNodeAssociation.children.push(visualcenterGongNodeInstance_VisualIcon)
          }

        }
      )

      /**
      * fill up the VisualCircle part of the mat tree
      */
      let visualcircleGongNodeStruct: GongNode = {
        name: "VisualCircle",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualCircle",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visualcircleGongNodeStruct)

      this.frontRepo.VisualCircles_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualCircles_array.forEach(
        visualcircleDB => {
          let visualcircleGongNodeInstance: GongNode = {
            name: visualcircleDB.Name,
            type: GongNodeType.INSTANCE,
            id: visualcircleDB.ID,
            uniqueIdPerStack: getVisualCircleUniqueID(visualcircleDB.ID),
            structName: "VisualCircle",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visualcircleGongNodeStruct.children!.push(visualcircleGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association VisualLayer
          */
          let VisualLayerGongNodeAssociation: GongNode = {
            name: "(VisualLayer) VisualLayer",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visualcircleDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualCircle",
            associationField: "VisualLayer",
            associatedStructName: "VisualLayer",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visualcircleGongNodeInstance.children!.push(VisualLayerGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualLayer
            */
          if (visualcircleDB.VisualLayer != undefined) {
            let visualcircleGongNodeInstance_VisualLayer: GongNode = {
              name: visualcircleDB.VisualLayer.Name,
              type: GongNodeType.INSTANCE,
              id: visualcircleDB.VisualLayer.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualCircleUniqueID(visualcircleDB.ID)
                + 5 * getVisualLayerUniqueID(visualcircleDB.VisualLayer.ID),
              structName: "VisualLayer",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualLayerGongNodeAssociation.children.push(visualcircleGongNodeInstance_VisualLayer)
          }

        }
      )

      /**
      * fill up the VisualIcon part of the mat tree
      */
      let visualiconGongNodeStruct: GongNode = {
        name: "VisualIcon",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualIcon",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visualiconGongNodeStruct)

      this.frontRepo.VisualIcons_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualIcons_array.forEach(
        visualiconDB => {
          let visualiconGongNodeInstance: GongNode = {
            name: visualiconDB.Name,
            type: GongNodeType.INSTANCE,
            id: visualiconDB.ID,
            uniqueIdPerStack: getVisualIconUniqueID(visualiconDB.ID),
            structName: "VisualIcon",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visualiconGongNodeStruct.children!.push(visualiconGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the VisualLayer part of the mat tree
      */
      let visuallayerGongNodeStruct: GongNode = {
        name: "VisualLayer",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualLayer",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visuallayerGongNodeStruct)

      this.frontRepo.VisualLayers_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualLayers_array.forEach(
        visuallayerDB => {
          let visuallayerGongNodeInstance: GongNode = {
            name: visuallayerDB.Name,
            type: GongNodeType.INSTANCE,
            id: visuallayerDB.ID,
            uniqueIdPerStack: getVisualLayerUniqueID(visuallayerDB.ID),
            structName: "VisualLayer",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visuallayerGongNodeStruct.children!.push(visuallayerGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the VisualLine part of the mat tree
      */
      let visuallineGongNodeStruct: GongNode = {
        name: "VisualLine",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualLine",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visuallineGongNodeStruct)

      this.frontRepo.VisualLines_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualLines_array.forEach(
        visuallineDB => {
          let visuallineGongNodeInstance: GongNode = {
            name: visuallineDB.Name,
            type: GongNodeType.INSTANCE,
            id: visuallineDB.ID,
            uniqueIdPerStack: getVisualLineUniqueID(visuallineDB.ID),
            structName: "VisualLine",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visuallineGongNodeStruct.children!.push(visuallineGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association VisualLayer
          */
          let VisualLayerGongNodeAssociation: GongNode = {
            name: "(VisualLayer) VisualLayer",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visuallineDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualLine",
            associationField: "VisualLayer",
            associatedStructName: "VisualLayer",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visuallineGongNodeInstance.children!.push(VisualLayerGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualLayer
            */
          if (visuallineDB.VisualLayer != undefined) {
            let visuallineGongNodeInstance_VisualLayer: GongNode = {
              name: visuallineDB.VisualLayer.Name,
              type: GongNodeType.INSTANCE,
              id: visuallineDB.VisualLayer.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualLineUniqueID(visuallineDB.ID)
                + 5 * getVisualLayerUniqueID(visuallineDB.VisualLayer.ID),
              structName: "VisualLayer",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualLayerGongNodeAssociation.children.push(visuallineGongNodeInstance_VisualLayer)
          }

        }
      )

      /**
      * fill up the VisualMap part of the mat tree
      */
      let visualmapGongNodeStruct: GongNode = {
        name: "VisualMap",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualMap",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visualmapGongNodeStruct)

      this.frontRepo.VisualMaps_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualMaps_array.forEach(
        visualmapDB => {
          let visualmapGongNodeInstance: GongNode = {
            name: visualmapDB.Name,
            type: GongNodeType.INSTANCE,
            id: visualmapDB.ID,
            uniqueIdPerStack: getVisualMapUniqueID(visualmapDB.ID),
            structName: "VisualMap",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visualmapGongNodeStruct.children!.push(visualmapGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the VisualTrack part of the mat tree
      */
      let visualtrackGongNodeStruct: GongNode = {
        name: "VisualTrack",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "VisualTrack",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(visualtrackGongNodeStruct)

      this.frontRepo.VisualTracks_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.VisualTracks_array.forEach(
        visualtrackDB => {
          let visualtrackGongNodeInstance: GongNode = {
            name: visualtrackDB.Name,
            type: GongNodeType.INSTANCE,
            id: visualtrackDB.ID,
            uniqueIdPerStack: getVisualTrackUniqueID(visualtrackDB.ID),
            structName: "VisualTrack",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          visualtrackGongNodeStruct.children!.push(visualtrackGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association VisualLayer
          */
          let VisualLayerGongNodeAssociation: GongNode = {
            name: "(VisualLayer) VisualLayer",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visualtrackDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualTrack",
            associationField: "VisualLayer",
            associatedStructName: "VisualLayer",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visualtrackGongNodeInstance.children!.push(VisualLayerGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualLayer
            */
          if (visualtrackDB.VisualLayer != undefined) {
            let visualtrackGongNodeInstance_VisualLayer: GongNode = {
              name: visualtrackDB.VisualLayer.Name,
              type: GongNodeType.INSTANCE,
              id: visualtrackDB.VisualLayer.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualTrackUniqueID(visualtrackDB.ID)
                + 5 * getVisualLayerUniqueID(visualtrackDB.VisualLayer.ID),
              structName: "VisualLayer",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualLayerGongNodeAssociation.children.push(visualtrackGongNodeInstance_VisualLayer)
          }

          /**
          * let append a node for the association VisualIcon
          */
          let VisualIconGongNodeAssociation: GongNode = {
            name: "(VisualIcon) VisualIcon",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: visualtrackDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "VisualTrack",
            associationField: "VisualIcon",
            associatedStructName: "VisualIcon",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          visualtrackGongNodeInstance.children!.push(VisualIconGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation VisualIcon
            */
          if (visualtrackDB.VisualIcon != undefined) {
            let visualtrackGongNodeInstance_VisualIcon: GongNode = {
              name: visualtrackDB.VisualIcon.Name,
              type: GongNodeType.INSTANCE,
              id: visualtrackDB.VisualIcon.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getVisualTrackUniqueID(visualtrackDB.ID)
                + 5 * getVisualIconUniqueID(visualtrackDB.VisualIcon.ID),
              structName: "VisualIcon",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            VisualIconGongNodeAssociation.children.push(visualtrackGongNodeInstance_VisualIcon)
          }

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
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
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
        github_com_fullstack_lang_gongleaflet_go_table: ["github_com_fullstack_lang_gongleaflet_go-" + path]
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
          github_com_fullstack_lang_gongleaflet_go_table: ["github_com_fullstack_lang_gongleaflet_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongleaflet_go_presentation: ["github_com_fullstack_lang_gongleaflet_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}

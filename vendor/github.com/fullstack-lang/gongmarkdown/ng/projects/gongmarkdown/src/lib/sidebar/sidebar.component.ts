import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'

// insertion point for per struct import code
import { CellService } from '../cell.service'
import { getCellUniqueID } from '../front-repo.service'
import { ElementService } from '../element.service'
import { getElementUniqueID } from '../front-repo.service'
import { MarkdownContentService } from '../markdowncontent.service'
import { getMarkdownContentUniqueID } from '../front-repo.service'
import { RowService } from '../row.service'
import { getRowUniqueID } from '../front-repo.service'

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
  selector: 'app-gongmarkdown-sidebar',
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
    private cellService: CellService,
    private elementService: ElementService,
    private markdowncontentService: MarkdownContentService,
    private rowService: RowService,
  ) { }

  ngOnInit(): void {
    this.refresh()

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.cellService.CellServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.elementService.ElementServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.markdowncontentService.MarkdownContentServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.rowService.RowServiceChanged.subscribe(
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
      * fill up the Cell part of the mat tree
      */
      let cellGongNodeStruct: GongNode = {
        name: "Cell",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Cell",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(cellGongNodeStruct)

      this.frontRepo.Cells_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Cells_array.forEach(
        cellDB => {
          let cellGongNodeInstance: GongNode = {
            name: cellDB.Name,
            type: GongNodeType.INSTANCE,
            id: cellDB.ID,
            uniqueIdPerStack: getCellUniqueID(cellDB.ID),
            structName: "Cell",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          cellGongNodeStruct.children!.push(cellGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Element part of the mat tree
      */
      let elementGongNodeStruct: GongNode = {
        name: "Element",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Element",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(elementGongNodeStruct)

      this.frontRepo.Elements_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Elements_array.forEach(
        elementDB => {
          let elementGongNodeInstance: GongNode = {
            name: elementDB.Name,
            type: GongNodeType.INSTANCE,
            id: elementDB.ID,
            uniqueIdPerStack: getElementUniqueID(elementDB.ID),
            structName: "Element",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          elementGongNodeStruct.children!.push(elementGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer SubElements
          */
          let SubElementsGongNodeAssociation: GongNode = {
            name: "(Element) SubElements",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: elementDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Element",
            associationField: "SubElements",
            associatedStructName: "Element",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          elementGongNodeInstance.children.push(SubElementsGongNodeAssociation)

          elementDB.SubElements?.forEach(elementDB => {
            let elementNode: GongNode = {
              name: elementDB.Name,
              type: GongNodeType.INSTANCE,
              id: elementDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getElementUniqueID(elementDB.ID)
                + 11 * getElementUniqueID(elementDB.ID),
              structName: "Element",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            SubElementsGongNodeAssociation.children.push(elementNode)
          })

          /**
          * let append a node for the slide of pointer Rows
          */
          let RowsGongNodeAssociation: GongNode = {
            name: "(Row) Rows",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: elementDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Element",
            associationField: "Rows",
            associatedStructName: "Row",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          elementGongNodeInstance.children.push(RowsGongNodeAssociation)

          elementDB.Rows?.forEach(rowDB => {
            let rowNode: GongNode = {
              name: rowDB.Name,
              type: GongNodeType.INSTANCE,
              id: rowDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getElementUniqueID(elementDB.ID)
                + 11 * getRowUniqueID(rowDB.ID),
              structName: "Row",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RowsGongNodeAssociation.children.push(rowNode)
          })

        }
      )

      /**
      * fill up the MarkdownContent part of the mat tree
      */
      let markdowncontentGongNodeStruct: GongNode = {
        name: "MarkdownContent",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "MarkdownContent",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(markdowncontentGongNodeStruct)

      this.frontRepo.MarkdownContents_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.MarkdownContents_array.forEach(
        markdowncontentDB => {
          let markdowncontentGongNodeInstance: GongNode = {
            name: markdowncontentDB.Name,
            type: GongNodeType.INSTANCE,
            id: markdowncontentDB.ID,
            uniqueIdPerStack: getMarkdownContentUniqueID(markdowncontentDB.ID),
            structName: "MarkdownContent",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          markdowncontentGongNodeStruct.children!.push(markdowncontentGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association Root
          */
          let RootGongNodeAssociation: GongNode = {
            name: "(Element) Root",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: markdowncontentDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "MarkdownContent",
            associationField: "Root",
            associatedStructName: "Element",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          markdowncontentGongNodeInstance.children!.push(RootGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Root
            */
          if (markdowncontentDB.Root != undefined) {
            let markdowncontentGongNodeInstance_Root: GongNode = {
              name: markdowncontentDB.Root.Name,
              type: GongNodeType.INSTANCE,
              id: markdowncontentDB.Root.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getMarkdownContentUniqueID(markdowncontentDB.ID)
                + 5 * getElementUniqueID(markdowncontentDB.Root.ID),
              structName: "Element",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RootGongNodeAssociation.children.push(markdowncontentGongNodeInstance_Root)
          }

        }
      )

      /**
      * fill up the Row part of the mat tree
      */
      let rowGongNodeStruct: GongNode = {
        name: "Row",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Row",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(rowGongNodeStruct)

      this.frontRepo.Rows_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Rows_array.forEach(
        rowDB => {
          let rowGongNodeInstance: GongNode = {
            name: rowDB.Name,
            type: GongNodeType.INSTANCE,
            id: rowDB.ID,
            uniqueIdPerStack: getRowUniqueID(rowDB.ID),
            structName: "Row",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          rowGongNodeStruct.children!.push(rowGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Cells
          */
          let CellsGongNodeAssociation: GongNode = {
            name: "(Cell) Cells",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: rowDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Row",
            associationField: "Cells",
            associatedStructName: "Cell",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          rowGongNodeInstance.children.push(CellsGongNodeAssociation)

          rowDB.Cells?.forEach(cellDB => {
            let cellNode: GongNode = {
              name: cellDB.Name,
              type: GongNodeType.INSTANCE,
              id: cellDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getRowUniqueID(rowDB.ID)
                + 11 * getCellUniqueID(cellDB.ID),
              structName: "Cell",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            CellsGongNodeAssociation.children.push(cellNode)
          })

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
        github_com_fullstack_lang_gongmarkdown_go_table: ["github_com_fullstack_lang_gongmarkdown_go-" + path]
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
          github_com_fullstack_lang_gongmarkdown_go_table: ["github_com_fullstack_lang_gongmarkdown_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongmarkdown_go_presentation: ["github_com_fullstack_lang_gongmarkdown_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}

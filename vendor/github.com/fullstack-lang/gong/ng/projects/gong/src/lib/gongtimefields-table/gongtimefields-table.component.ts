// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { GongTimeFieldDB } from '../gongtimefield-db'
import { GongTimeFieldService } from '../gongtimefield.service'

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-gongtimefieldstable',
  templateUrl: './gongtimefields-table.component.html',
  styleUrls: ['./gongtimefields-table.component.css'],
})
export class GongTimeFieldsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of GongTimeField instances
  selection: SelectionModel<GongTimeFieldDB> = new (SelectionModel)
  initialSelection = new Array<GongTimeFieldDB>()

  // the data source for the table
  gongtimefields: GongTimeFieldDB[] = []
  matTableDataSource: MatTableDataSource<GongTimeFieldDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.gongtimefields
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (gongtimefieldDB: GongTimeFieldDB, property: string) => {
      switch (property) {
        case 'ID':
          return gongtimefieldDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return gongtimefieldDB.Name;

        case 'Index':
          return gongtimefieldDB.Index;

        case 'GongStruct_GongTimeFields':
          return this.frontRepo.GongStructs.get(gongtimefieldDB.GongStruct_GongTimeFieldsDBID.Int64)!.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (gongtimefieldDB: GongTimeFieldDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the gongtimefieldDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += gongtimefieldDB.Name.toLowerCase()
      mergedContent += gongtimefieldDB.Index.toString()
      if (gongtimefieldDB.GongStruct_GongTimeFieldsDBID.Int64 != 0) {
        mergedContent += this.frontRepo.GongStructs.get(gongtimefieldDB.GongStruct_GongTimeFieldsDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private gongtimefieldService: GongTimeFieldService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of gongtimefield instances
    public dialogRef: MatDialogRef<GongTimeFieldsTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.gongtimefieldService.GongTimeFieldServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getGongTimeFields()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Index",
        "GongStruct_GongTimeFields",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Index",
        "GongStruct_GongTimeFields",
      ]
      this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getGongTimeFields()
    this.matTableDataSource = new MatTableDataSource(this.gongtimefields)
  }

  getGongTimeFields(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.gongtimefields = this.frontRepo.GongTimeFields_array;

        // insertion point for variables Recoveries

        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let gongtimefield of this.gongtimefields) {
            let ID = this.dialogData.ID
            let revPointer = gongtimefield[this.dialogData.ReversePointer as keyof GongTimeFieldDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(gongtimefield)
            }
            this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongTimeFieldDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as GongTimeFieldDB[]
          for (let associationInstance of sourceField) {
            let gongtimefield = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongTimeFieldDB
            this.initialSelection.push(gongtimefield)
          }

          this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.gongtimefields
      }
    )
  }

  // newGongTimeField initiate a new gongtimefield
  // create a new GongTimeField objet
  newGongTimeField() {
  }

  deleteGongTimeField(gongtimefieldID: number, gongtimefield: GongTimeFieldDB) {
    // list of gongtimefields is truncated of gongtimefield before the delete
    this.gongtimefields = this.gongtimefields.filter(h => h !== gongtimefield);

    this.gongtimefieldService.deleteGongTimeField(gongtimefieldID).subscribe(
      gongtimefield => {
        this.gongtimefieldService.GongTimeFieldServiceChanged.next("delete")
      }
    );
  }

  editGongTimeField(gongtimefieldID: number, gongtimefield: GongTimeFieldDB) {

  }

  // display gongtimefield in router
  displayGongTimeFieldInRouter(gongtimefieldID: number) {
    this.router.navigate(["github_com_fullstack_lang_gong_go-" + "gongtimefield-display", gongtimefieldID])
  }

  // set editor outlet
  setEditorRouterOutlet(gongtimefieldID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_editor: ["github_com_fullstack_lang_gong_go-" + "gongtimefield-detail", gongtimefieldID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(gongtimefieldID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gong_go_presentation: ["github_com_fullstack_lang_gong_go-" + "gongtimefield-presentation", gongtimefieldID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.gongtimefields.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.gongtimefields.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<GongTimeFieldDB>()

      // reset all initial selection of gongtimefield that belong to gongtimefield
      for (let gongtimefield of this.initialSelection) {
        let index = gongtimefield[this.dialogData.ReversePointer as keyof GongTimeFieldDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(gongtimefield)

      }

      // from selection, set gongtimefield that belong to gongtimefield
      for (let gongtimefield of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = gongtimefield[this.dialogData.ReversePointer as keyof GongTimeFieldDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(gongtimefield)
      }


      // update all gongtimefield (only update selection & initial selection)
      for (let gongtimefield of toUpdate) {
        this.gongtimefieldService.updateGongTimeField(gongtimefield)
          .subscribe(gongtimefield => {
            this.gongtimefieldService.GongTimeFieldServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, GongTimeFieldDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedGongTimeField = new Set<number>()
      for (let gongtimefield of this.initialSelection) {
        if (this.selection.selected.includes(gongtimefield)) {
          // console.log("gongtimefield " + gongtimefield.Name + " is still selected")
        } else {
          console.log("gongtimefield " + gongtimefield.Name + " has been unselected")
          unselectedGongTimeField.add(gongtimefield.ID)
          console.log("is unselected " + unselectedGongTimeField.has(gongtimefield.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let gongtimefield = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as GongTimeFieldDB
      if (unselectedGongTimeField.has(gongtimefield.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<GongTimeFieldDB>) = new Array<GongTimeFieldDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          gongtimefield => {
            if (!this.initialSelection.includes(gongtimefield)) {
              // console.log("gongtimefield " + gongtimefield.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + gongtimefield.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = gongtimefield.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = gongtimefield.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("gongtimefield " + gongtimefield.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<GongTimeFieldDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
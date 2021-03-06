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
import { LinerDB } from '../liner-db'
import { LinerService } from '../liner.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-linerstable',
  templateUrl: './liners-table.component.html',
  styleUrls: ['./liners-table.component.css'],
})
export class LinersTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Liner instances
  selection: SelectionModel<LinerDB> = new (SelectionModel)
  initialSelection = new Array<LinerDB>()

  // the data source for the table
  liners: LinerDB[] = []
  matTableDataSource: MatTableDataSource<LinerDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.liners
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
    this.matTableDataSource.sortingDataAccessor = (linerDB: LinerDB, property: string) => {
      switch (property) {
        case 'ID':
          return linerDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return linerDB.Name;

        case 'Lat':
          return linerDB.Lat;

        case 'Lng':
          return linerDB.Lng;

        case 'Heading':
          return linerDB.Heading;

        case 'Level':
          return linerDB.Level;

        case 'Speed':
          return linerDB.Speed;

        case 'State':
          return linerDB.State;

        case 'TargetHeading':
          return linerDB.TargetHeading;

        case 'TargetLocationLat':
          return linerDB.TargetLocationLat;

        case 'TargetLocationLng':
          return linerDB.TargetLocationLng;

        case 'DistanceToTarget':
          return linerDB.DistanceToTarget;

        case 'MaxRotationalSpeed':
          return linerDB.MaxRotationalSpeed;

        case 'VerticalSpeed':
          return linerDB.VerticalSpeed;

        case 'Timestampstring':
          return linerDB.Timestampstring;

        case 'ReporingLine':
          return (linerDB.ReporingLine ? linerDB.ReporingLine.Name : '');

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (linerDB: LinerDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the linerDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += linerDB.Name.toLowerCase()
      mergedContent += linerDB.Lat.toString()
      mergedContent += linerDB.Lng.toString()
      mergedContent += linerDB.Heading.toString()
      mergedContent += linerDB.Level.toString()
      mergedContent += linerDB.Speed.toString()
      mergedContent += linerDB.State.toLowerCase()
      mergedContent += linerDB.TargetHeading.toString()
      mergedContent += linerDB.TargetLocationLat.toString()
      mergedContent += linerDB.TargetLocationLng.toString()
      mergedContent += linerDB.DistanceToTarget.toString()
      mergedContent += linerDB.MaxRotationalSpeed.toString()
      mergedContent += linerDB.VerticalSpeed.toString()
      mergedContent += linerDB.Timestampstring.toLowerCase()
      if (linerDB.ReporingLine) {
        mergedContent += linerDB.ReporingLine.Name.toLowerCase()
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
    private linerService: LinerService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of liner instances
    public dialogRef: MatDialogRef<LinersTableComponent>,
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
    this.linerService.LinerServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getLiners()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "Lat",
        "Lng",
        "Heading",
        "Level",
        "Speed",
        "State",
        "TargetHeading",
        "TargetLocationLat",
        "TargetLocationLng",
        "DistanceToTarget",
        "MaxRotationalSpeed",
        "VerticalSpeed",
        "Timestampstring",
        "ReporingLine",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "Lat",
        "Lng",
        "Heading",
        "Level",
        "Speed",
        "State",
        "TargetHeading",
        "TargetLocationLat",
        "TargetLocationLng",
        "DistanceToTarget",
        "MaxRotationalSpeed",
        "VerticalSpeed",
        "Timestampstring",
        "ReporingLine",
      ]
      this.selection = new SelectionModel<LinerDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getLiners()
    this.matTableDataSource = new MatTableDataSource(this.liners)
  }

  getLiners(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.liners = this.frontRepo.Liners_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let liner of this.liners) {
            let ID = this.dialogData.ID
            let revPointer = liner[this.dialogData.ReversePointer as keyof LinerDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(liner)
            }
            this.selection = new SelectionModel<LinerDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinerDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as LinerDB[]
          for (let associationInstance of sourceField) {
            let liner = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinerDB
            this.initialSelection.push(liner)
          }

          this.selection = new SelectionModel<LinerDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.liners
      }
    )
  }

  // newLiner initiate a new liner
  // create a new Liner objet
  newLiner() {
  }

  deleteLiner(linerID: number, liner: LinerDB) {
    // list of liners is truncated of liner before the delete
    this.liners = this.liners.filter(h => h !== liner);

    this.linerService.deleteLiner(linerID).subscribe(
      liner => {
        this.linerService.LinerServiceChanged.next("delete")
      }
    );
  }

  editLiner(linerID: number, liner: LinerDB) {

  }

  // display liner in router
  displayLinerInRouter(linerID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongfly_go-" + "liner-display", linerID])
  }

  // set editor outlet
  setEditorRouterOutlet(linerID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "liner-detail", linerID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(linerID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_presentation: ["github_com_fullstack_lang_gongfly_go-" + "liner-presentation", linerID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.liners.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.liners.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<LinerDB>()

      // reset all initial selection of liner that belong to liner
      for (let liner of this.initialSelection) {
        let index = liner[this.dialogData.ReversePointer as keyof LinerDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(liner)

      }

      // from selection, set liner that belong to liner
      for (let liner of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = liner[this.dialogData.ReversePointer as keyof LinerDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(liner)
      }


      // update all liner (only update selection & initial selection)
      for (let liner of toUpdate) {
        this.linerService.updateLiner(liner)
          .subscribe(liner => {
            this.linerService.LinerServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, LinerDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedLiner = new Set<number>()
      for (let liner of this.initialSelection) {
        if (this.selection.selected.includes(liner)) {
          // console.log("liner " + liner.Name + " is still selected")
        } else {
          console.log("liner " + liner.Name + " has been unselected")
          unselectedLiner.add(liner.ID)
          console.log("is unselected " + unselectedLiner.has(liner.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let liner = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as LinerDB
      if (unselectedLiner.has(liner.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<LinerDB>) = new Array<LinerDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          liner => {
            if (!this.initialSelection.includes(liner)) {
              // console.log("liner " + liner.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + liner.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = liner.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = liner.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("liner " + liner.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<LinerDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}

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
import { RadarDB } from '../radar-db'
import { RadarService } from '../radar.service'

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
  selector: 'app-radarstable',
  templateUrl: './radars-table.component.html',
  styleUrls: ['./radars-table.component.css'],
})
export class RadarsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Radar instances
  selection: SelectionModel<RadarDB> = new (SelectionModel)
  initialSelection = new Array<RadarDB>()

  // the data source for the table
  radars: RadarDB[] = []
  matTableDataSource: MatTableDataSource<RadarDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.radars
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
    this.matTableDataSource.sortingDataAccessor = (radarDB: RadarDB, property: string) => {
      switch (property) {
        case 'ID':
          return radarDB.ID

        // insertion point for specific sorting accessor
        case 'State':
          return radarDB.State;

        case 'Name':
          return radarDB.Name;

        case 'Lat':
          return radarDB.Lat;

        case 'Lng':
          return radarDB.Lng;

        case 'Range':
          return radarDB.Range;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (radarDB: RadarDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the radarDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += radarDB.State.toLowerCase()
      mergedContent += radarDB.Name.toLowerCase()
      mergedContent += radarDB.Lat.toString()
      mergedContent += radarDB.Lng.toString()
      mergedContent += radarDB.Range.toString()

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
    private radarService: RadarService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of radar instances
    public dialogRef: MatDialogRef<RadarsTableComponent>,
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
    this.radarService.RadarServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getRadars()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "State",
        "Name",
        "Lat",
        "Lng",
        "Range",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "State",
        "Name",
        "Lat",
        "Lng",
        "Range",
      ]
      this.selection = new SelectionModel<RadarDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getRadars()
    this.matTableDataSource = new MatTableDataSource(this.radars)
  }

  getRadars(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.radars = this.frontRepo.Radars_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let radar of this.radars) {
            let ID = this.dialogData.ID
            let revPointer = radar[this.dialogData.ReversePointer as keyof RadarDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(radar)
            }
            this.selection = new SelectionModel<RadarDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, RadarDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as RadarDB[]
          for (let associationInstance of sourceField) {
            let radar = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as RadarDB
            this.initialSelection.push(radar)
          }

          this.selection = new SelectionModel<RadarDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.radars
      }
    )
  }

  // newRadar initiate a new radar
  // create a new Radar objet
  newRadar() {
  }

  deleteRadar(radarID: number, radar: RadarDB) {
    // list of radars is truncated of radar before the delete
    this.radars = this.radars.filter(h => h !== radar);

    this.radarService.deleteRadar(radarID).subscribe(
      radar => {
        this.radarService.RadarServiceChanged.next("delete")
      }
    );
  }

  editRadar(radarID: number, radar: RadarDB) {

  }

  // display radar in router
  displayRadarInRouter(radarID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongfly_go-" + "radar-display", radarID])
  }

  // set editor outlet
  setEditorRouterOutlet(radarID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "radar-detail", radarID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(radarID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongfly_go_presentation: ["github_com_fullstack_lang_gongfly_go-" + "radar-presentation", radarID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.radars.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.radars.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<RadarDB>()

      // reset all initial selection of radar that belong to radar
      for (let radar of this.initialSelection) {
        let index = radar[this.dialogData.ReversePointer as keyof RadarDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(radar)

      }

      // from selection, set radar that belong to radar
      for (let radar of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = radar[this.dialogData.ReversePointer as keyof RadarDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(radar)
      }


      // update all radar (only update selection & initial selection)
      for (let radar of toUpdate) {
        this.radarService.updateRadar(radar)
          .subscribe(radar => {
            this.radarService.RadarServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, RadarDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedRadar = new Set<number>()
      for (let radar of this.initialSelection) {
        if (this.selection.selected.includes(radar)) {
          // console.log("radar " + radar.Name + " is still selected")
        } else {
          console.log("radar " + radar.Name + " has been unselected")
          unselectedRadar.add(radar.ID)
          console.log("is unselected " + unselectedRadar.has(radar.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let radar = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as RadarDB
      if (unselectedRadar.has(radar.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<RadarDB>) = new Array<RadarDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          radar => {
            if (!this.initialSelection.includes(radar)) {
              // console.log("radar " + radar.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + radar.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = radar.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = radar.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("radar " + radar.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<RadarDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}

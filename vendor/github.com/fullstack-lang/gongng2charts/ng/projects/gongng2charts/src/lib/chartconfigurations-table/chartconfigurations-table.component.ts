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
import { ChartConfigurationDB } from '../chartconfiguration-db'
import { ChartConfigurationService } from '../chartconfiguration.service'

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
  selector: 'app-chartconfigurationstable',
  templateUrl: './chartconfigurations-table.component.html',
  styleUrls: ['./chartconfigurations-table.component.css'],
})
export class ChartConfigurationsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of ChartConfiguration instances
  selection: SelectionModel<ChartConfigurationDB> = new (SelectionModel)
  initialSelection = new Array<ChartConfigurationDB>()

  // the data source for the table
  chartconfigurations: ChartConfigurationDB[] = []
  matTableDataSource: MatTableDataSource<ChartConfigurationDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.chartconfigurations
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
    this.matTableDataSource.sortingDataAccessor = (chartconfigurationDB: ChartConfigurationDB, property: string) => {
      switch (property) {
        case 'ID':
          return chartconfigurationDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return chartconfigurationDB.Name;

        case 'ChartType':
          return chartconfigurationDB.ChartType;

        case 'Width':
          return chartconfigurationDB.Width;

        case 'Heigth':
          return chartconfigurationDB.Heigth;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (chartconfigurationDB: ChartConfigurationDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the chartconfigurationDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += chartconfigurationDB.Name.toLowerCase()
      mergedContent += chartconfigurationDB.ChartType.toLowerCase()
      mergedContent += chartconfigurationDB.Width.toString()
      mergedContent += chartconfigurationDB.Heigth.toString()

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
    private chartconfigurationService: ChartConfigurationService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of chartconfiguration instances
    public dialogRef: MatDialogRef<ChartConfigurationsTableComponent>,
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
    this.chartconfigurationService.ChartConfigurationServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getChartConfigurations()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "ChartType",
        "Width",
        "Heigth",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "ChartType",
        "Width",
        "Heigth",
      ]
      this.selection = new SelectionModel<ChartConfigurationDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getChartConfigurations()
    this.matTableDataSource = new MatTableDataSource(this.chartconfigurations)
  }

  getChartConfigurations(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.chartconfigurations = this.frontRepo.ChartConfigurations_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let chartconfiguration of this.chartconfigurations) {
            let ID = this.dialogData.ID
            let revPointer = chartconfiguration[this.dialogData.ReversePointer as keyof ChartConfigurationDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(chartconfiguration)
            }
            this.selection = new SelectionModel<ChartConfigurationDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ChartConfigurationDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as ChartConfigurationDB[]
          for (let associationInstance of sourceField) {
            let chartconfiguration = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ChartConfigurationDB
            this.initialSelection.push(chartconfiguration)
          }

          this.selection = new SelectionModel<ChartConfigurationDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.chartconfigurations
      }
    )
  }

  // newChartConfiguration initiate a new chartconfiguration
  // create a new ChartConfiguration objet
  newChartConfiguration() {
  }

  deleteChartConfiguration(chartconfigurationID: number, chartconfiguration: ChartConfigurationDB) {
    // list of chartconfigurations is truncated of chartconfiguration before the delete
    this.chartconfigurations = this.chartconfigurations.filter(h => h !== chartconfiguration);

    this.chartconfigurationService.deleteChartConfiguration(chartconfigurationID).subscribe(
      chartconfiguration => {
        this.chartconfigurationService.ChartConfigurationServiceChanged.next("delete")
      }
    );
  }

  editChartConfiguration(chartconfigurationID: number, chartconfiguration: ChartConfigurationDB) {

  }

  // display chartconfiguration in router
  displayChartConfigurationInRouter(chartconfigurationID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongng2charts_go-" + "chartconfiguration-display", chartconfigurationID])
  }

  // set editor outlet
  setEditorRouterOutlet(chartconfigurationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_editor: ["github_com_fullstack_lang_gongng2charts_go-" + "chartconfiguration-detail", chartconfigurationID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(chartconfigurationID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongng2charts_go_presentation: ["github_com_fullstack_lang_gongng2charts_go-" + "chartconfiguration-presentation", chartconfigurationID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.chartconfigurations.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.chartconfigurations.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<ChartConfigurationDB>()

      // reset all initial selection of chartconfiguration that belong to chartconfiguration
      for (let chartconfiguration of this.initialSelection) {
        let index = chartconfiguration[this.dialogData.ReversePointer as keyof ChartConfigurationDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(chartconfiguration)

      }

      // from selection, set chartconfiguration that belong to chartconfiguration
      for (let chartconfiguration of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = chartconfiguration[this.dialogData.ReversePointer as keyof ChartConfigurationDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(chartconfiguration)
      }


      // update all chartconfiguration (only update selection & initial selection)
      for (let chartconfiguration of toUpdate) {
        this.chartconfigurationService.updateChartConfiguration(chartconfiguration)
          .subscribe(chartconfiguration => {
            this.chartconfigurationService.ChartConfigurationServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ChartConfigurationDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedChartConfiguration = new Set<number>()
      for (let chartconfiguration of this.initialSelection) {
        if (this.selection.selected.includes(chartconfiguration)) {
          // console.log("chartconfiguration " + chartconfiguration.Name + " is still selected")
        } else {
          console.log("chartconfiguration " + chartconfiguration.Name + " has been unselected")
          unselectedChartConfiguration.add(chartconfiguration.ID)
          console.log("is unselected " + unselectedChartConfiguration.has(chartconfiguration.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let chartconfiguration = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ChartConfigurationDB
      if (unselectedChartConfiguration.has(chartconfiguration.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<ChartConfigurationDB>) = new Array<ChartConfigurationDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          chartconfiguration => {
            if (!this.initialSelection.includes(chartconfiguration)) {
              // console.log("chartconfiguration " + chartconfiguration.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + chartconfiguration.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = chartconfiguration.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = chartconfiguration.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("chartconfiguration " + chartconfiguration.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<ChartConfigurationDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
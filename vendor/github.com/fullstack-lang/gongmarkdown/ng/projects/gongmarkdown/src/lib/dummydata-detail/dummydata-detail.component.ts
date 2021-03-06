// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { DummyDataDB } from '../dummydata-db'
import { DummyDataService } from '../dummydata.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ElementTypeSelect, ElementTypeList } from '../ElementType'
import { DummnyTypeIntSelect, DummnyTypeIntList } from '../DummnyTypeInt'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// DummyDataDetailComponent is initilizaed from different routes
// DummyDataDetailComponentState detail different cases 
enum DummyDataDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-dummydata-detail',
	templateUrl: './dummydata-detail.component.html',
	styleUrls: ['./dummydata-detail.component.css'],
})
export class DummyDataDetailComponent implements OnInit {

	// insertion point for declarations
	DummyBoolFormControl = new FormControl(false);
	ElementTypeList: ElementTypeSelect[] = []
	DummnyTypeIntList: DummnyTypeIntSelect[] = []
	DummyDuration_Hours: number = 0
	DummyDuration_Minutes: number = 0
	DummyDuration_Seconds: number = 0

	// the DummyDataDB of interest
	dummydata: DummyDataDB = new DummyDataDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: DummyDataDetailComponentState = DummyDataDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private dummydataService: DummyDataService,
		private frontRepoService: FrontRepoService,
		public dialog: MatDialog,
		private route: ActivatedRoute,
		private router: Router,
	) {
	}

	ngOnInit(): void {

		// compute state
		this.id = +this.route.snapshot.paramMap.get('id')!;
		this.originStruct = this.route.snapshot.paramMap.get('originStruct')!;
		this.originStructFieldName = this.route.snapshot.paramMap.get('originStructFieldName')!;

		const association = this.route.snapshot.paramMap.get('association');
		if (this.id == 0) {
			this.state = DummyDataDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = DummyDataDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getDummyData()

		// observable for changes in structs
		this.dummydataService.DummyDataServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getDummyData()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ElementTypeList = ElementTypeList
		this.DummnyTypeIntList = DummnyTypeIntList
	}

	getDummyData(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case DummyDataDetailComponentState.CREATE_INSTANCE:
						this.dummydata = new (DummyDataDB)
						break;
					case DummyDataDetailComponentState.UPDATE_INSTANCE:
						let dummydata = frontRepo.DummyDatas.get(this.id)
						console.assert(dummydata != undefined, "missing dummydata with id:" + this.id)
						this.dummydata = dummydata!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				this.DummyBoolFormControl.setValue(this.dummydata.DummyBool)
				// computation of Hours, Minutes, Seconds for DummyDuration
				this.DummyDuration_Hours = Math.floor(this.dummydata.DummyDuration / (3600 * 1000 * 1000 * 1000))
				this.DummyDuration_Minutes = Math.floor(this.dummydata.DummyDuration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.DummyDuration_Seconds = this.dummydata.DummyDuration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.dummydata.DummyBool = this.DummyBoolFormControl.value
		this.dummydata.DummyDuration =
			this.DummyDuration_Hours * (3600 * 1000 * 1000 * 1000) +
			this.DummyDuration_Minutes * (60 * 1000 * 1000 * 1000) +
			this.DummyDuration_Seconds * (1000 * 1000 * 1000)
		if (this.dummydata.DummyPointerToGongStructID == undefined) {
			this.dummydata.DummyPointerToGongStructID = new NullInt64
		}
		if (this.dummydata.DummyPointerToGongStruct != undefined) {
			this.dummydata.DummyPointerToGongStructID.Int64 = this.dummydata.DummyPointerToGongStruct.ID
			this.dummydata.DummyPointerToGongStructID.Valid = true
		} else {
			this.dummydata.DummyPointerToGongStructID.Int64 = 0
			this.dummydata.DummyPointerToGongStructID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case DummyDataDetailComponentState.UPDATE_INSTANCE:
				this.dummydataService.updateDummyData(this.dummydata)
					.subscribe(dummydata => {
						this.dummydataService.DummyDataServiceChanged.next("update")
					});
				break;
			default:
				this.dummydataService.postDummyData(this.dummydata).subscribe(dummydata => {
					this.dummydataService.DummyDataServiceChanged.next("post")
					this.dummydata = new (DummyDataDB) // reset fields
				});
		}
	}

	// openReverseSelection is a generic function that calls dialog for the edition of 
	// ONE-MANY association
	// It uses the MapOfComponent provided by the front repo
	openReverseSelection(AssociatedStruct: string, reverseField: string, selectionMode: string,
		sourceField: string, intermediateStructField: string, nextAssociatedStruct: string) {

		console.log("mode " + selectionMode)

		const dialogConfig = new MatDialogConfig();

		let dialogData = new DialogData();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.width = "50%"
		dialogConfig.height = "50%"
		if (selectionMode == SelectionMode.ONE_MANY_ASSOCIATION_MODE) {

			dialogData.ID = this.dummydata.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(AssociatedStruct).get(
					AssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}
		if (selectionMode == SelectionMode.MANY_MANY_ASSOCIATION_MODE) {
			dialogData.ID = this.dummydata.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "DummyData"
			dialogData.SourceField = sourceField

			// set up the intermediate struct
			dialogData.IntermediateStruct = AssociatedStruct
			dialogData.IntermediateStructField = intermediateStructField

			// set up the end struct
			dialogData.NextAssociationStruct = nextAssociatedStruct

			dialogConfig.data = dialogData
			const dialogRef: MatDialogRef<string, any> = this.dialog.open(
				MapOfComponents.get(nextAssociatedStruct).get(
					nextAssociatedStruct + 'sTableComponent'
				),
				dialogConfig
			);
			dialogRef.afterClosed().subscribe(result => {
			});
		}

	}

	openDragAndDropOrdering(AssociatedStruct: string, reverseField: string) {

		const dialogConfig = new MatDialogConfig();

		// dialogConfig.disableClose = true;
		dialogConfig.autoFocus = true;
		dialogConfig.data = {
			ID: this.dummydata.ID,
			ReversePointer: reverseField,
			OrderingMode: true,
		};
		const dialogRef: MatDialogRef<string, any> = this.dialog.open(
			MapOfSortingComponents.get(AssociatedStruct).get(
				AssociatedStruct + 'SortingComponent'
			),
			dialogConfig
		);

		dialogRef.afterClosed().subscribe(result => {
		});
	}

	fillUpNameIfEmpty(event: { value: { Name: string; }; }) {
		if (this.dummydata.Name == "") {
			this.dummydata.Name = event.value.Name
		}
	}

	toggleTextArea(fieldName: string) {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			let displayAsTextArea = this.mapFields_displayAsTextArea.get(fieldName)
			this.mapFields_displayAsTextArea.set(fieldName, !displayAsTextArea)
		} else {
			this.mapFields_displayAsTextArea.set(fieldName, true)
		}
	}

	isATextArea(fieldName: string): boolean {
		if (this.mapFields_displayAsTextArea.has(fieldName)) {
			return this.mapFields_displayAsTextArea.get(fieldName)!
		} else {
			return false
		}
	}

	compareObjects(o1: any, o2: any) {
		if (o1?.ID == o2?.ID) {
			return true;
		}
		else {
			return false
		}
	}
}

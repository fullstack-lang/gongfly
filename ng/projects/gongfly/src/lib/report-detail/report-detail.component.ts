// generated from NgDetailTemplateTS
import { Component, OnInit } from '@angular/core';
import { FormControl } from '@angular/forms';

import { ReportDB } from '../report-db'
import { ReportService } from '../report.service'

import { FrontRepoService, FrontRepo, SelectionMode, DialogData } from '../front-repo.service'
import { MapOfComponents } from '../map-components'
import { MapOfSortingComponents } from '../map-components'

// insertion point for imports
import { ReportEnumSelect, ReportEnumList } from '../ReportEnum'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

import { MatDialog, MAT_DIALOG_DATA, MatDialogRef, MatDialogConfig } from '@angular/material/dialog';

import { NullInt64 } from '../null-int64'

// ReportDetailComponent is initilizaed from different routes
// ReportDetailComponentState detail different cases 
enum ReportDetailComponentState {
	CREATE_INSTANCE,
	UPDATE_INSTANCE,
	// insertion point for declarations of enum values of state
}

@Component({
	selector: 'app-report-detail',
	templateUrl: './report-detail.component.html',
	styleUrls: ['./report-detail.component.css'],
})
export class ReportDetailComponent implements OnInit {

	// insertion point for declarations
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0
	ReportEnumList: ReportEnumSelect[] = []

	// the ReportDB of interest
	report: ReportDB = new ReportDB

	// front repo
	frontRepo: FrontRepo = new FrontRepo

	// this stores the information related to string fields
	// if false, the field is inputed with an <input ...> form 
	// if true, it is inputed with a <textarea ...> </textarea>
	mapFields_displayAsTextArea = new Map<string, boolean>()

	// the state at initialization (CREATION, UPDATE or CREATE with one association set)
	state: ReportDetailComponentState = ReportDetailComponentState.CREATE_INSTANCE

	// in UDPATE state, if is the id of the instance to update
	// in CREATE state with one association set, this is the id of the associated instance
	id: number = 0

	// in CREATE state with one association set, this is the id of the associated instance
	originStruct: string = ""
	originStructFieldName: string = ""

	constructor(
		private reportService: ReportService,
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
			this.state = ReportDetailComponentState.CREATE_INSTANCE
		} else {
			if (this.originStruct == undefined) {
				this.state = ReportDetailComponentState.UPDATE_INSTANCE
			} else {
				switch (this.originStructFieldName) {
					// insertion point for state computation
					default:
						console.log(this.originStructFieldName + " is unkown association")
				}
			}
		}

		this.getReport()

		// observable for changes in structs
		this.reportService.ReportServiceChanged.subscribe(
			message => {
				if (message == "post" || message == "update" || message == "delete") {
					this.getReport()
				}
			}
		)

		// insertion point for initialisation of enums list
		this.ReportEnumList = ReportEnumList
	}

	getReport(): void {

		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				switch (this.state) {
					case ReportDetailComponentState.CREATE_INSTANCE:
						this.report = new (ReportDB)
						break;
					case ReportDetailComponentState.UPDATE_INSTANCE:
						let report = frontRepo.Reports.get(this.id)
						console.assert(report != undefined, "missing report with id:" + this.id)
						this.report = report!
						break;
					// insertion point for init of association field
					default:
						console.log(this.state + " is unkown state")
				}

				// insertion point for recovery of form controls value for bool fields
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.report.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.report.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.report.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
			}
		)


	}

	save(): void {

		// some fields needs to be translated into serializable forms
		// pointers fields, after the translation, are nulled in order to perform serialization

		// insertion point for translation/nullation of each field
		this.report.Duration =
			this.Duration_Hours * (3600 * 1000 * 1000 * 1000) +
			this.Duration_Minutes * (60 * 1000 * 1000 * 1000) +
			this.Duration_Seconds * (1000 * 1000 * 1000)
		if (this.report.AboutID == undefined) {
			this.report.AboutID = new NullInt64
		}
		if (this.report.About != undefined) {
			this.report.AboutID.Int64 = this.report.About.ID
			this.report.AboutID.Valid = true
		} else {
			this.report.AboutID.Int64 = 0
			this.report.AboutID.Valid = true
		}
		if (this.report.OpsLineID == undefined) {
			this.report.OpsLineID = new NullInt64
		}
		if (this.report.OpsLine != undefined) {
			this.report.OpsLineID.Int64 = this.report.OpsLine.ID
			this.report.OpsLineID.Valid = true
		} else {
			this.report.OpsLineID.Int64 = 0
			this.report.OpsLineID.Valid = true
		}

		// save from the front pointer space to the non pointer space for serialization

		// insertion point for translation/nullation of each pointers

		switch (this.state) {
			case ReportDetailComponentState.UPDATE_INSTANCE:
				this.reportService.updateReport(this.report)
					.subscribe(report => {
						this.reportService.ReportServiceChanged.next("update")
					});
				break;
			default:
				this.reportService.postReport(this.report).subscribe(report => {
					this.reportService.ReportServiceChanged.next("post")
					this.report = new (ReportDB) // reset fields
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

			dialogData.ID = this.report.ID!
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
			dialogData.ID = this.report.ID!
			dialogData.ReversePointer = reverseField
			dialogData.OrderingMode = false
			dialogData.SelectionMode = selectionMode

			// set up the source
			dialogData.SourceStruct = "Report"
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
			ID: this.report.ID,
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
		if (this.report.Name == "") {
			this.report.Name = event.value.Name
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

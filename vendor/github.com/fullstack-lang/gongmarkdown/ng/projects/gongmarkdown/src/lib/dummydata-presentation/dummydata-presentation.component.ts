import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { DummyDataDB } from '../dummydata-db'
import { DummyDataService } from '../dummydata.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports
import { DummnyTypeIntList } from '../DummnyTypeInt'

export interface dummydataDummyElement {
}

const ELEMENT_DATA: dummydataDummyElement[] = [
];

@Component({
	selector: 'app-dummydata-presentation',
	templateUrl: './dummydata-presentation.component.html',
	styleUrls: ['./dummydata-presentation.component.css'],
})
export class DummyDataPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// fields from DummyDuration
	DummyDuration_Hours: number = 0
	DummyDuration_Minutes: number = 0
	DummyDuration_Seconds: number = 0
	// insertion point for additionnal enum int field declarations
	DummyEnumInt_Value : string = ""

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	dummydata: DummyDataDB = new (DummyDataDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private dummydataService: DummyDataService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getDummyData();

		// observable for changes in 
		this.dummydataService.DummyDataServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getDummyData()
				}
			}
		)
	}

	getDummyData(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.dummydata = this.frontRepo.DummyDatas.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for DummyDuration
				this.DummyDuration_Hours = Math.floor(this.dummydata.DummyDuration / (3600 * 1000 * 1000 * 1000))
				this.DummyDuration_Minutes = Math.floor(this.dummydata.DummyDuration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.DummyDuration_Seconds = this.dummydata.DummyDuration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
				// insertion point for recovery of enum tint
				this.DummyEnumInt_Value = DummnyTypeIntList[this.dummydata.DummyEnumInt].viewValue
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongmarkdown_go_presentation: ["github_com_fullstack_lang_gongmarkdown_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "dummydata-detail", ID]
			}
		}]);
	}
}

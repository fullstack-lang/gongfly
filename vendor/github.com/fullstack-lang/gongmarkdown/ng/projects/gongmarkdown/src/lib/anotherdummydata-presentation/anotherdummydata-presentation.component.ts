import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { AnotherDummyDataDB } from '../anotherdummydata-db'
import { AnotherDummyDataService } from '../anotherdummydata.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface anotherdummydataDummyElement {
}

const ELEMENT_DATA: anotherdummydataDummyElement[] = [
];

@Component({
	selector: 'app-anotherdummydata-presentation',
	templateUrl: './anotherdummydata-presentation.component.html',
	styleUrls: ['./anotherdummydata-presentation.component.css'],
})
export class AnotherDummyDataPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	anotherdummydata: AnotherDummyDataDB = new (AnotherDummyDataDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private anotherdummydataService: AnotherDummyDataService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getAnotherDummyData();

		// observable for changes in 
		this.anotherdummydataService.AnotherDummyDataServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getAnotherDummyData()
				}
			}
		)
	}

	getAnotherDummyData(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.anotherdummydata = this.frontRepo.AnotherDummyDatas.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
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
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "anotherdummydata-detail", ID]
			}
		}]);
	}
}

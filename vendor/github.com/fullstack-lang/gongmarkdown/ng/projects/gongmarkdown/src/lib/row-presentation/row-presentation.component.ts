import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { RowDB } from '../row-db'
import { RowService } from '../row.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface rowDummyElement {
}

const ELEMENT_DATA: rowDummyElement[] = [
];

@Component({
	selector: 'app-row-presentation',
	templateUrl: './row-presentation.component.html',
	styleUrls: ['./row-presentation.component.css'],
})
export class RowPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	row: RowDB = new (RowDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private rowService: RowService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getRow();

		// observable for changes in 
		this.rowService.RowServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getRow()
				}
			}
		)
	}

	getRow(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.row = this.frontRepo.Rows.get(id)!

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
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "row-detail", ID]
			}
		}]);
	}
}

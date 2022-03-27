import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { CellDB } from '../cell-db'
import { CellService } from '../cell.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface cellDummyElement {
}

const ELEMENT_DATA: cellDummyElement[] = [
];

@Component({
	selector: 'app-cell-presentation',
	templateUrl: './cell-presentation.component.html',
	styleUrls: ['./cell-presentation.component.css'],
})
export class CellPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	cell: CellDB = new (CellDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private cellService: CellService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getCell();

		// observable for changes in 
		this.cellService.CellServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getCell()
				}
			}
		)
	}

	getCell(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.cell = this.frontRepo.Cells.get(id)!

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
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "cell-detail", ID]
			}
		}]);
	}
}

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { LinerDB } from '../liner-db'
import { LinerService } from '../liner.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface linerDummyElement {
}

const ELEMENT_DATA: linerDummyElement[] = [
];

@Component({
	selector: 'app-liner-presentation',
	templateUrl: './liner-presentation.component.html',
	styleUrls: ['./liner-presentation.component.css'],
})
export class LinerPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	liner: LinerDB = new (LinerDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private linerService: LinerService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getLiner();

		// observable for changes in 
		this.linerService.LinerServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getLiner()
				}
			}
		)
	}

	getLiner(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.liner = this.frontRepo.Liners.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongfly_go_presentation: ["github_com_fullstack_lang_gongfly_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "liner-detail", ID]
			}
		}]);
	}
}

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { OpsLineDB } from '../opsline-db'
import { OpsLineService } from '../opsline.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface opslineDummyElement {
}

const ELEMENT_DATA: opslineDummyElement[] = [
];

@Component({
	selector: 'app-opsline-presentation',
	templateUrl: './opsline-presentation.component.html',
	styleUrls: ['./opsline-presentation.component.css'],
})
export class OpsLinePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	opsline: OpsLineDB = new (OpsLineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private opslineService: OpsLineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getOpsLine();

		// observable for changes in 
		this.opslineService.OpsLineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getOpsLine()
				}
			}
		)
	}

	getOpsLine(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.opsline = this.frontRepo.OpsLines.get(id)!

				// insertion point for recovery of durations
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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "opsline-detail", ID]
			}
		}]);
	}
}

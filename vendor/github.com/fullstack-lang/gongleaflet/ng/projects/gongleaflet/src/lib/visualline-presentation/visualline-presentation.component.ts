import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { VisualLineDB } from '../visualline-db'
import { VisualLineService } from '../visualline.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface visuallineDummyElement {
}

const ELEMENT_DATA: visuallineDummyElement[] = [
];

@Component({
	selector: 'app-visualline-presentation',
	templateUrl: './visualline-presentation.component.html',
	styleUrls: ['./visualline-presentation.component.css'],
})
export class VisualLinePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	visualline: VisualLineDB = new (VisualLineDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private visuallineService: VisualLineService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVisualLine();

		// observable for changes in 
		this.visuallineService.VisualLineServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVisualLine()
				}
			}
		)
	}

	getVisualLine(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.visualline = this.frontRepo.VisualLines.get(id)!

				// insertion point for recovery of durations
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongleaflet_go_presentation: ["github_com_fullstack_lang_gongleaflet_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + "visualline-detail", ID]
			}
		}]);
	}
}

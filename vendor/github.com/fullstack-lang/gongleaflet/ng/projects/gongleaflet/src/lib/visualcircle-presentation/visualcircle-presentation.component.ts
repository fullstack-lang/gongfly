import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { VisualCircleDB } from '../visualcircle-db'
import { VisualCircleService } from '../visualcircle.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface visualcircleDummyElement {
}

const ELEMENT_DATA: visualcircleDummyElement[] = [
];

@Component({
	selector: 'app-visualcircle-presentation',
	templateUrl: './visualcircle-presentation.component.html',
	styleUrls: ['./visualcircle-presentation.component.css'],
})
export class VisualCirclePresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	visualcircle: VisualCircleDB = new (VisualCircleDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private visualcircleService: VisualCircleService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVisualCircle();

		// observable for changes in 
		this.visualcircleService.VisualCircleServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVisualCircle()
				}
			}
		)
	}

	getVisualCircle(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.visualcircle = this.frontRepo.VisualCircles.get(id)!

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
				github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + "visualcircle-detail", ID]
			}
		}]);
	}
}

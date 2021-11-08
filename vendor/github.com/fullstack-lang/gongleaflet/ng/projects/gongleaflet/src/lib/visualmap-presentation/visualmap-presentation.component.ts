import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { VisualMapDB } from '../visualmap-db'
import { VisualMapService } from '../visualmap.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface visualmapDummyElement {
}

const ELEMENT_DATA: visualmapDummyElement[] = [
];

@Component({
	selector: 'app-visualmap-presentation',
	templateUrl: './visualmap-presentation.component.html',
	styleUrls: ['./visualmap-presentation.component.css'],
})
export class VisualMapPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	visualmap: VisualMapDB = new (VisualMapDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private visualmapService: VisualMapService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVisualMap();

		// observable for changes in 
		this.visualmapService.VisualMapServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVisualMap()
				}
			}
		)
	}

	getVisualMap(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.visualmap = this.frontRepo.VisualMaps.get(id)!

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
				github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + "visualmap-detail", ID]
			}
		}]);
	}
}

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { RadarDB } from '../radar-db'
import { RadarService } from '../radar.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface radarDummyElement {
}

const ELEMENT_DATA: radarDummyElement[] = [
];

@Component({
	selector: 'app-radar-presentation',
	templateUrl: './radar-presentation.component.html',
	styleUrls: ['./radar-presentation.component.css'],
})
export class RadarPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	radar: RadarDB = new (RadarDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private radarService: RadarService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getRadar();

		// observable for changes in 
		this.radarService.RadarServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getRadar()
				}
			}
		)
	}

	getRadar(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.radar = this.frontRepo.Radars.get(id)!

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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "radar-detail", ID]
			}
		}]);
	}
}

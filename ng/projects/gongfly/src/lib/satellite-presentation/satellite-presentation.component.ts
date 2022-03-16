import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { SatelliteDB } from '../satellite-db'
import { SatelliteService } from '../satellite.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface satelliteDummyElement {
}

const ELEMENT_DATA: satelliteDummyElement[] = [
];

@Component({
	selector: 'app-satellite-presentation',
	templateUrl: './satellite-presentation.component.html',
	styleUrls: ['./satellite-presentation.component.css'],
})
export class SatellitePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	satellite: SatelliteDB = new (SatelliteDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private satelliteService: SatelliteService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getSatellite();

		// observable for changes in 
		this.satelliteService.SatelliteServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getSatellite()
				}
			}
		)
	}

	getSatellite(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.satellite = this.frontRepo.Satellites.get(id)!

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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "satellite-detail", ID]
			}
		}]);
	}
}

import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { CivilianAirportDB } from '../civilianairport-db'
import { CivilianAirportService } from '../civilianairport.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface civilianairportDummyElement {
}

const ELEMENT_DATA: civilianairportDummyElement[] = [
];

@Component({
	selector: 'app-civilianairport-presentation',
	templateUrl: './civilianairport-presentation.component.html',
	styleUrls: ['./civilianairport-presentation.component.css'],
})
export class CivilianAirportPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	civilianairport: CivilianAirportDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private civilianairportService: CivilianAirportService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getCivilianAirport();

		// observable for changes in 
		this.civilianairportService.CivilianAirportServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getCivilianAirport()
				}
			}
		)
	}

	getCivilianAirport(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.civilianairport = this.frontRepo.CivilianAirports.get(id)

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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "civilianairport-detail", ID]
			}
		}]);
	}
}

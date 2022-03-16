import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ScenarioDB } from '../scenario-db'
import { ScenarioService } from '../scenario.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface scenarioDummyElement {
}

const ELEMENT_DATA: scenarioDummyElement[] = [
];

@Component({
	selector: 'app-scenario-presentation',
	templateUrl: './scenario-presentation.component.html',
	styleUrls: ['./scenario-presentation.component.css'],
})
export class ScenarioPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	scenario: ScenarioDB = new (ScenarioDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private scenarioService: ScenarioService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getScenario();

		// observable for changes in 
		this.scenarioService.ScenarioServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getScenario()
				}
			}
		)
	}

	getScenario(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.scenario = this.frontRepo.Scenarios.get(id)!

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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "scenario-detail", ID]
			}
		}]);
	}
}

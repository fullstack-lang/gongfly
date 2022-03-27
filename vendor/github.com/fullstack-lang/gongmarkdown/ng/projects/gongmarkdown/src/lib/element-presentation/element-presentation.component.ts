import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ElementDB } from '../element-db'
import { ElementService } from '../element.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface elementDummyElement {
}

const ELEMENT_DATA: elementDummyElement[] = [
];

@Component({
	selector: 'app-element-presentation',
	templateUrl: './element-presentation.component.html',
	styleUrls: ['./element-presentation.component.css'],
})
export class ElementPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	element: ElementDB = new (ElementDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private elementService: ElementService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getElement();

		// observable for changes in 
		this.elementService.ElementServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getElement()
				}
			}
		)
	}

	getElement(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.element = this.frontRepo.Elements.get(id)!

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
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "element-detail", ID]
			}
		}]);
	}
}

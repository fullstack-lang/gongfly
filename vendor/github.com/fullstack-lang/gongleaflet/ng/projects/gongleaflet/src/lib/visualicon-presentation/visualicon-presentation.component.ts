import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { VisualIconDB } from '../visualicon-db'
import { VisualIconService } from '../visualicon.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface visualiconDummyElement {
}

const ELEMENT_DATA: visualiconDummyElement[] = [
];

@Component({
	selector: 'app-visualicon-presentation',
	templateUrl: './visualicon-presentation.component.html',
	styleUrls: ['./visualicon-presentation.component.css'],
})
export class VisualIconPresentationComponent implements OnInit {

	// insertion point for declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	visualicon: VisualIconDB = new (VisualIconDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private visualiconService: VisualIconService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getVisualIcon();

		// observable for changes in 
		this.visualiconService.VisualIconServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getVisualIcon()
				}
			}
		)
	}

	getVisualIcon(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.visualicon = this.frontRepo.VisualIcons.get(id)!

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
				github_com_fullstack_lang_gongleaflet_go_editor: ["github_com_fullstack_lang_gongleaflet_go-" + "visualicon-detail", ID]
			}
		}]);
	}
}

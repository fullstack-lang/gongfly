import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MarkdownContentDB } from '../markdowncontent-db'
import { MarkdownContentService } from '../markdowncontent.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface markdowncontentDummyElement {
}

const ELEMENT_DATA: markdowncontentDummyElement[] = [
];

@Component({
	selector: 'app-markdowncontent-presentation',
	templateUrl: './markdowncontent-presentation.component.html',
	styleUrls: ['./markdowncontent-presentation.component.css'],
})
export class MarkdownContentPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	markdowncontent: MarkdownContentDB = new (MarkdownContentDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private markdowncontentService: MarkdownContentService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMarkdownContent();

		// observable for changes in 
		this.markdowncontentService.MarkdownContentServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMarkdownContent()
				}
			}
		)
	}

	getMarkdownContent(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.markdowncontent = this.frontRepo.MarkdownContents.get(id)!

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
				github_com_fullstack_lang_gongmarkdown_go_editor: ["github_com_fullstack_lang_gongmarkdown_go-" + "markdowncontent-detail", ID]
			}
		}]);
	}
}

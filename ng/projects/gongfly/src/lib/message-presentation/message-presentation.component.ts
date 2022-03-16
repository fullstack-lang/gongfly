import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { MessageDB } from '../message-db'
import { MessageService } from '../message.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface messageDummyElement {
}

const ELEMENT_DATA: messageDummyElement[] = [
];

@Component({
	selector: 'app-message-presentation',
	templateUrl: './message-presentation.component.html',
	styleUrls: ['./message-presentation.component.css'],
})
export class MessagePresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// fields from DurationSinceSimulationStart
	DurationSinceSimulationStart_Hours: number = 0
	DurationSinceSimulationStart_Minutes: number = 0
	DurationSinceSimulationStart_Seconds: number = 0
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	message: MessageDB = new (MessageDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private messageService: MessageService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getMessage();

		// observable for changes in 
		this.messageService.MessageServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getMessage()
				}
			}
		)
	}

	getMessage(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.message = this.frontRepo.Messages.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for DurationSinceSimulationStart
				this.DurationSinceSimulationStart_Hours = Math.floor(this.message.DurationSinceSimulationStart / (3600 * 1000 * 1000 * 1000))
				this.DurationSinceSimulationStart_Minutes = Math.floor(this.message.DurationSinceSimulationStart % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.DurationSinceSimulationStart_Seconds = this.message.DurationSinceSimulationStart % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "message-detail", ID]
			}
		}]);
	}
}

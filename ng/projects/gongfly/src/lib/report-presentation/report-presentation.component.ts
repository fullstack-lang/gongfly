import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ReportDB } from '../report-db'
import { ReportService } from '../report.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface reportDummyElement {
}

const ELEMENT_DATA: reportDummyElement[] = [
];

@Component({
	selector: 'app-report-presentation',
	templateUrl: './report-presentation.component.html',
	styleUrls: ['./report-presentation.component.css'],
})
export class ReportPresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from Duration
	Duration_Hours: number = 0
	Duration_Minutes: number = 0
	Duration_Seconds: number = 0

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	report: ReportDB = new (ReportDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private reportService: ReportService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getReport();

		// observable for changes in 
		this.reportService.ReportServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getReport()
				}
			}
		)
	}

	getReport(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.report = this.frontRepo.Reports.get(id)!

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.report.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.report.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.report.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "report-detail", ID]
			}
		}]);
	}
}

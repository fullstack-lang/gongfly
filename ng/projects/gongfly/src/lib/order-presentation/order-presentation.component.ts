import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { OrderDB } from '../order-db'
import { OrderService } from '../order.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

export interface orderDummyElement {
}

const ELEMENT_DATA: orderDummyElement[] = [
];

@Component({
	selector: 'app-order-presentation',
	templateUrl: './order-presentation.component.html',
	styleUrls: ['./order-presentation.component.css'],
})
export class OrderPresentationComponent implements OnInit {

	// insertion point for declarations
	// fields from Duration
	Duration_Hours: number
	Duration_Minutes: number
	Duration_Seconds: number

	displayedColumns: string[] = [];
	dataSource = ELEMENT_DATA;

	order: OrderDB;

	// front repo
	frontRepo: FrontRepo
 
	constructor(
		private orderService: OrderService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getOrder();

		// observable for changes in 
		this.orderService.OrderServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getOrder()
				}
			}
		)
	}

	getOrder(): void {
		const id = +this.route.snapshot.paramMap.get('id');
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.order = this.frontRepo.Orders.get(id)

				// insertion point for recovery of durations
				// computation of Hours, Minutes, Seconds for Duration
				this.Duration_Hours = Math.floor(this.order.Duration / (3600 * 1000 * 1000 * 1000))
				this.Duration_Minutes = Math.floor(this.order.Duration % (3600 * 1000 * 1000 * 1000) / (60 * 1000 * 1000 * 1000))
				this.Duration_Seconds = this.order.Duration % (60 * 1000 * 1000 * 1000) / (1000 * 1000 * 1000)
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
				github_com_fullstack_lang_gongfly_go_editor: ["github_com_fullstack_lang_gongfly_go-" + "order-detail", ID]
			}
		}]);
	}
}

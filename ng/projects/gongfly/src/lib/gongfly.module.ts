import { NgModule } from '@angular/core';

import { FormsModule, ReactiveFormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { Routes, RouterModule } from '@angular/router';

// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';

import { AngularSplitModule, SplitComponent } from 'angular-split';

import {
	NgxMatDatetimePickerModule,
	NgxMatNativeDateModule,
	NgxMatTimepickerModule
} from '@angular-material-components/datetime-picker';

import { AppRoutingModule } from './app-routing.module';

import { SplitterComponent } from './splitter/splitter.component'
import { SidebarComponent } from './sidebar/sidebar.component';

// insertion point for imports 
import { CivilianAirportsTableComponent } from './civilianairports-table/civilianairports-table.component'
import { CivilianAirportSortingComponent } from './civilianairport-sorting/civilianairport-sorting.component'
import { CivilianAirportDetailComponent } from './civilianairport-detail/civilianairport-detail.component'
import { CivilianAirportPresentationComponent } from './civilianairport-presentation/civilianairport-presentation.component'

import { LinersTableComponent } from './liners-table/liners-table.component'
import { LinerSortingComponent } from './liner-sorting/liner-sorting.component'
import { LinerDetailComponent } from './liner-detail/liner-detail.component'
import { LinerPresentationComponent } from './liner-presentation/liner-presentation.component'

import { MessagesTableComponent } from './messages-table/messages-table.component'
import { MessageSortingComponent } from './message-sorting/message-sorting.component'
import { MessageDetailComponent } from './message-detail/message-detail.component'
import { MessagePresentationComponent } from './message-presentation/message-presentation.component'

import { OpsLinesTableComponent } from './opslines-table/opslines-table.component'
import { OpsLineSortingComponent } from './opsline-sorting/opsline-sorting.component'
import { OpsLineDetailComponent } from './opsline-detail/opsline-detail.component'
import { OpsLinePresentationComponent } from './opsline-presentation/opsline-presentation.component'

import { OrdersTableComponent } from './orders-table/orders-table.component'
import { OrderSortingComponent } from './order-sorting/order-sorting.component'
import { OrderDetailComponent } from './order-detail/order-detail.component'
import { OrderPresentationComponent } from './order-presentation/order-presentation.component'

import { RadarsTableComponent } from './radars-table/radars-table.component'
import { RadarSortingComponent } from './radar-sorting/radar-sorting.component'
import { RadarDetailComponent } from './radar-detail/radar-detail.component'
import { RadarPresentationComponent } from './radar-presentation/radar-presentation.component'

import { ReportsTableComponent } from './reports-table/reports-table.component'
import { ReportSortingComponent } from './report-sorting/report-sorting.component'
import { ReportDetailComponent } from './report-detail/report-detail.component'
import { ReportPresentationComponent } from './report-presentation/report-presentation.component'

import { ScenariosTableComponent } from './scenarios-table/scenarios-table.component'
import { ScenarioSortingComponent } from './scenario-sorting/scenario-sorting.component'
import { ScenarioDetailComponent } from './scenario-detail/scenario-detail.component'
import { ScenarioPresentationComponent } from './scenario-presentation/scenario-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		CivilianAirportsTableComponent,
		CivilianAirportSortingComponent,
		CivilianAirportDetailComponent,
		CivilianAirportPresentationComponent,

		LinersTableComponent,
		LinerSortingComponent,
		LinerDetailComponent,
		LinerPresentationComponent,

		MessagesTableComponent,
		MessageSortingComponent,
		MessageDetailComponent,
		MessagePresentationComponent,

		OpsLinesTableComponent,
		OpsLineSortingComponent,
		OpsLineDetailComponent,
		OpsLinePresentationComponent,

		OrdersTableComponent,
		OrderSortingComponent,
		OrderDetailComponent,
		OrderPresentationComponent,

		RadarsTableComponent,
		RadarSortingComponent,
		RadarDetailComponent,
		RadarPresentationComponent,

		ReportsTableComponent,
		ReportSortingComponent,
		ReportDetailComponent,
		ReportPresentationComponent,

		ScenariosTableComponent,
		ScenarioSortingComponent,
		ScenarioDetailComponent,
		ScenarioPresentationComponent,


		SplitterComponent,
		SidebarComponent
	],
	imports: [
		FormsModule,
		ReactiveFormsModule,
		CommonModule,
		RouterModule,

		AppRoutingModule,

		MatSliderModule,
		MatSelectModule,
		MatFormFieldModule,
		MatInputModule,
		MatDatepickerModule,
		MatTableModule,
		MatSortModule,
		MatPaginatorModule,
		MatCheckboxModule,
		MatButtonModule,
		MatIconModule,
		MatToolbarModule,
		MatExpansionModule,
		MatListModule,
		MatDialogModule,
		MatGridListModule,
		MatTreeModule,
		DragDropModule,

		NgxMatDatetimePickerModule,
		NgxMatNativeDateModule,
		NgxMatTimepickerModule,

		AngularSplitModule,
	],
	exports: [
		// insertion point for declarations 
		CivilianAirportsTableComponent,
		CivilianAirportSortingComponent,
		CivilianAirportDetailComponent,
		CivilianAirportPresentationComponent,

		LinersTableComponent,
		LinerSortingComponent,
		LinerDetailComponent,
		LinerPresentationComponent,

		MessagesTableComponent,
		MessageSortingComponent,
		MessageDetailComponent,
		MessagePresentationComponent,

		OpsLinesTableComponent,
		OpsLineSortingComponent,
		OpsLineDetailComponent,
		OpsLinePresentationComponent,

		OrdersTableComponent,
		OrderSortingComponent,
		OrderDetailComponent,
		OrderPresentationComponent,

		RadarsTableComponent,
		RadarSortingComponent,
		RadarDetailComponent,
		RadarPresentationComponent,

		ReportsTableComponent,
		ReportSortingComponent,
		ReportDetailComponent,
		ReportPresentationComponent,

		ScenariosTableComponent,
		ScenarioSortingComponent,
		ScenarioDetailComponent,
		ScenarioPresentationComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongflyModule { }

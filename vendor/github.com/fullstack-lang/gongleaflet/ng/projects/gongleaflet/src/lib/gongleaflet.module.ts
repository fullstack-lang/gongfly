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
import { GongstructSelectionService } from './gongstruct-selection.service'

// insertion point for imports 
import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleSortingComponent } from './circle-sorting/circle-sorting.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'

import { DivIconsTableComponent } from './divicons-table/divicons-table.component'
import { DivIconSortingComponent } from './divicon-sorting/divicon-sorting.component'
import { DivIconDetailComponent } from './divicon-detail/divicon-detail.component'

import { LayerGroupsTableComponent } from './layergroups-table/layergroups-table.component'
import { LayerGroupSortingComponent } from './layergroup-sorting/layergroup-sorting.component'
import { LayerGroupDetailComponent } from './layergroup-detail/layergroup-detail.component'

import { LayerGroupUsesTableComponent } from './layergroupuses-table/layergroupuses-table.component'
import { LayerGroupUseSortingComponent } from './layergroupuse-sorting/layergroupuse-sorting.component'
import { LayerGroupUseDetailComponent } from './layergroupuse-detail/layergroupuse-detail.component'

import { MapOptionssTableComponent } from './mapoptionss-table/mapoptionss-table.component'
import { MapOptionsSortingComponent } from './mapoptions-sorting/mapoptions-sorting.component'
import { MapOptionsDetailComponent } from './mapoptions-detail/mapoptions-detail.component'

import { MarkersTableComponent } from './markers-table/markers-table.component'
import { MarkerSortingComponent } from './marker-sorting/marker-sorting.component'
import { MarkerDetailComponent } from './marker-detail/marker-detail.component'

import { UserClicksTableComponent } from './userclicks-table/userclicks-table.component'
import { UserClickSortingComponent } from './userclick-sorting/userclick-sorting.component'
import { UserClickDetailComponent } from './userclick-detail/userclick-detail.component'

import { VLinesTableComponent } from './vlines-table/vlines-table.component'
import { VLineSortingComponent } from './vline-sorting/vline-sorting.component'
import { VLineDetailComponent } from './vline-detail/vline-detail.component'

import { VisualTracksTableComponent } from './visualtracks-table/visualtracks-table.component'
import { VisualTrackSortingComponent } from './visualtrack-sorting/visualtrack-sorting.component'
import { VisualTrackDetailComponent } from './visualtrack-detail/visualtrack-detail.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		CirclesTableComponent,
		CircleSortingComponent,
		CircleDetailComponent,

		DivIconsTableComponent,
		DivIconSortingComponent,
		DivIconDetailComponent,

		LayerGroupsTableComponent,
		LayerGroupSortingComponent,
		LayerGroupDetailComponent,

		LayerGroupUsesTableComponent,
		LayerGroupUseSortingComponent,
		LayerGroupUseDetailComponent,

		MapOptionssTableComponent,
		MapOptionsSortingComponent,
		MapOptionsDetailComponent,

		MarkersTableComponent,
		MarkerSortingComponent,
		MarkerDetailComponent,

		UserClicksTableComponent,
		UserClickSortingComponent,
		UserClickDetailComponent,

		VLinesTableComponent,
		VLineSortingComponent,
		VLineDetailComponent,

		VisualTracksTableComponent,
		VisualTrackSortingComponent,
		VisualTrackDetailComponent,


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
		CirclesTableComponent,
		CircleSortingComponent,
		CircleDetailComponent,

		DivIconsTableComponent,
		DivIconSortingComponent,
		DivIconDetailComponent,

		LayerGroupsTableComponent,
		LayerGroupSortingComponent,
		LayerGroupDetailComponent,

		LayerGroupUsesTableComponent,
		LayerGroupUseSortingComponent,
		LayerGroupUseDetailComponent,

		MapOptionssTableComponent,
		MapOptionsSortingComponent,
		MapOptionsDetailComponent,

		MarkersTableComponent,
		MarkerSortingComponent,
		MarkerDetailComponent,

		UserClicksTableComponent,
		UserClickSortingComponent,
		UserClickDetailComponent,

		VLinesTableComponent,
		VLineSortingComponent,
		VLineDetailComponent,

		VisualTracksTableComponent,
		VisualTrackSortingComponent,
		VisualTrackDetailComponent,


		SplitterComponent,
		SidebarComponent,

	],
	providers: [
		GongstructSelectionService,
		{
			provide: MatDialogRef,
			useValue: {}
		},
	],
})
export class GongleafletModule { }

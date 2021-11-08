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
import { VisualCentersTableComponent } from './visualcenters-table/visualcenters-table.component'
import { VisualCenterSortingComponent } from './visualcenter-sorting/visualcenter-sorting.component'
import { VisualCenterDetailComponent } from './visualcenter-detail/visualcenter-detail.component'
import { VisualCenterPresentationComponent } from './visualcenter-presentation/visualcenter-presentation.component'

import { VisualCirclesTableComponent } from './visualcircles-table/visualcircles-table.component'
import { VisualCircleSortingComponent } from './visualcircle-sorting/visualcircle-sorting.component'
import { VisualCircleDetailComponent } from './visualcircle-detail/visualcircle-detail.component'
import { VisualCirclePresentationComponent } from './visualcircle-presentation/visualcircle-presentation.component'

import { VisualIconsTableComponent } from './visualicons-table/visualicons-table.component'
import { VisualIconSortingComponent } from './visualicon-sorting/visualicon-sorting.component'
import { VisualIconDetailComponent } from './visualicon-detail/visualicon-detail.component'
import { VisualIconPresentationComponent } from './visualicon-presentation/visualicon-presentation.component'

import { VisualLayersTableComponent } from './visuallayers-table/visuallayers-table.component'
import { VisualLayerSortingComponent } from './visuallayer-sorting/visuallayer-sorting.component'
import { VisualLayerDetailComponent } from './visuallayer-detail/visuallayer-detail.component'
import { VisualLayerPresentationComponent } from './visuallayer-presentation/visuallayer-presentation.component'

import { VisualLinesTableComponent } from './visuallines-table/visuallines-table.component'
import { VisualLineSortingComponent } from './visualline-sorting/visualline-sorting.component'
import { VisualLineDetailComponent } from './visualline-detail/visualline-detail.component'
import { VisualLinePresentationComponent } from './visualline-presentation/visualline-presentation.component'

import { VisualMapsTableComponent } from './visualmaps-table/visualmaps-table.component'
import { VisualMapSortingComponent } from './visualmap-sorting/visualmap-sorting.component'
import { VisualMapDetailComponent } from './visualmap-detail/visualmap-detail.component'
import { VisualMapPresentationComponent } from './visualmap-presentation/visualmap-presentation.component'

import { VisualTracksTableComponent } from './visualtracks-table/visualtracks-table.component'
import { VisualTrackSortingComponent } from './visualtrack-sorting/visualtrack-sorting.component'
import { VisualTrackDetailComponent } from './visualtrack-detail/visualtrack-detail.component'
import { VisualTrackPresentationComponent } from './visualtrack-presentation/visualtrack-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		VisualCentersTableComponent,
		VisualCenterSortingComponent,
		VisualCenterDetailComponent,
		VisualCenterPresentationComponent,

		VisualCirclesTableComponent,
		VisualCircleSortingComponent,
		VisualCircleDetailComponent,
		VisualCirclePresentationComponent,

		VisualIconsTableComponent,
		VisualIconSortingComponent,
		VisualIconDetailComponent,
		VisualIconPresentationComponent,

		VisualLayersTableComponent,
		VisualLayerSortingComponent,
		VisualLayerDetailComponent,
		VisualLayerPresentationComponent,

		VisualLinesTableComponent,
		VisualLineSortingComponent,
		VisualLineDetailComponent,
		VisualLinePresentationComponent,

		VisualMapsTableComponent,
		VisualMapSortingComponent,
		VisualMapDetailComponent,
		VisualMapPresentationComponent,

		VisualTracksTableComponent,
		VisualTrackSortingComponent,
		VisualTrackDetailComponent,
		VisualTrackPresentationComponent,


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
		VisualCentersTableComponent,
		VisualCenterSortingComponent,
		VisualCenterDetailComponent,
		VisualCenterPresentationComponent,

		VisualCirclesTableComponent,
		VisualCircleSortingComponent,
		VisualCircleDetailComponent,
		VisualCirclePresentationComponent,

		VisualIconsTableComponent,
		VisualIconSortingComponent,
		VisualIconDetailComponent,
		VisualIconPresentationComponent,

		VisualLayersTableComponent,
		VisualLayerSortingComponent,
		VisualLayerDetailComponent,
		VisualLayerPresentationComponent,

		VisualLinesTableComponent,
		VisualLineSortingComponent,
		VisualLineDetailComponent,
		VisualLinePresentationComponent,

		VisualMapsTableComponent,
		VisualMapSortingComponent,
		VisualMapDetailComponent,
		VisualMapPresentationComponent,

		VisualTracksTableComponent,
		VisualTrackSortingComponent,
		VisualTrackDetailComponent,
		VisualTrackPresentationComponent,


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
export class GongleafletModule { }

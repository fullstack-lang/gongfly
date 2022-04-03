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
import { AnotherDummyDatasTableComponent } from './anotherdummydatas-table/anotherdummydatas-table.component'
import { AnotherDummyDataSortingComponent } from './anotherdummydata-sorting/anotherdummydata-sorting.component'
import { AnotherDummyDataDetailComponent } from './anotherdummydata-detail/anotherdummydata-detail.component'
import { AnotherDummyDataPresentationComponent } from './anotherdummydata-presentation/anotherdummydata-presentation.component'

import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'
import { CellPresentationComponent } from './cell-presentation/cell-presentation.component'

import { DummyDatasTableComponent } from './dummydatas-table/dummydatas-table.component'
import { DummyDataSortingComponent } from './dummydata-sorting/dummydata-sorting.component'
import { DummyDataDetailComponent } from './dummydata-detail/dummydata-detail.component'
import { DummyDataPresentationComponent } from './dummydata-presentation/dummydata-presentation.component'

import { ElementsTableComponent } from './elements-table/elements-table.component'
import { ElementSortingComponent } from './element-sorting/element-sorting.component'
import { ElementDetailComponent } from './element-detail/element-detail.component'
import { ElementPresentationComponent } from './element-presentation/element-presentation.component'

import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
import { MarkdownContentSortingComponent } from './markdowncontent-sorting/markdowncontent-sorting.component'
import { MarkdownContentDetailComponent } from './markdowncontent-detail/markdowncontent-detail.component'
import { MarkdownContentPresentationComponent } from './markdowncontent-presentation/markdowncontent-presentation.component'

import { RowsTableComponent } from './rows-table/rows-table.component'
import { RowSortingComponent } from './row-sorting/row-sorting.component'
import { RowDetailComponent } from './row-detail/row-detail.component'
import { RowPresentationComponent } from './row-presentation/row-presentation.component'


@NgModule({
	declarations: [
		// insertion point for declarations 
		AnotherDummyDatasTableComponent,
		AnotherDummyDataSortingComponent,
		AnotherDummyDataDetailComponent,
		AnotherDummyDataPresentationComponent,

		CellsTableComponent,
		CellSortingComponent,
		CellDetailComponent,
		CellPresentationComponent,

		DummyDatasTableComponent,
		DummyDataSortingComponent,
		DummyDataDetailComponent,
		DummyDataPresentationComponent,

		ElementsTableComponent,
		ElementSortingComponent,
		ElementDetailComponent,
		ElementPresentationComponent,

		MarkdownContentsTableComponent,
		MarkdownContentSortingComponent,
		MarkdownContentDetailComponent,
		MarkdownContentPresentationComponent,

		RowsTableComponent,
		RowSortingComponent,
		RowDetailComponent,
		RowPresentationComponent,


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
		AnotherDummyDatasTableComponent,
		AnotherDummyDataSortingComponent,
		AnotherDummyDataDetailComponent,
		AnotherDummyDataPresentationComponent,

		CellsTableComponent,
		CellSortingComponent,
		CellDetailComponent,
		CellPresentationComponent,

		DummyDatasTableComponent,
		DummyDataSortingComponent,
		DummyDataDetailComponent,
		DummyDataPresentationComponent,

		ElementsTableComponent,
		ElementSortingComponent,
		ElementDetailComponent,
		ElementPresentationComponent,

		MarkdownContentsTableComponent,
		MarkdownContentSortingComponent,
		MarkdownContentDetailComponent,
		MarkdownContentPresentationComponent,

		RowsTableComponent,
		RowSortingComponent,
		RowDetailComponent,
		RowPresentationComponent,


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
export class GongmarkdownModule { }

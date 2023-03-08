import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CheckoutSchedulersTableComponent } from './checkoutschedulers-table/checkoutschedulers-table.component'
import { CheckoutSchedulerDetailComponent } from './checkoutscheduler-detail/checkoutscheduler-detail.component'

import { CirclesTableComponent } from './circles-table/circles-table.component'
import { CircleDetailComponent } from './circle-detail/circle-detail.component'

import { DivIconsTableComponent } from './divicons-table/divicons-table.component'
import { DivIconDetailComponent } from './divicon-detail/divicon-detail.component'

import { LayerGroupsTableComponent } from './layergroups-table/layergroups-table.component'
import { LayerGroupDetailComponent } from './layergroup-detail/layergroup-detail.component'

import { LayerGroupUsesTableComponent } from './layergroupuses-table/layergroupuses-table.component'
import { LayerGroupUseDetailComponent } from './layergroupuse-detail/layergroupuse-detail.component'

import { MapOptionssTableComponent } from './mapoptionss-table/mapoptionss-table.component'
import { MapOptionsDetailComponent } from './mapoptions-detail/mapoptions-detail.component'

import { MarkersTableComponent } from './markers-table/markers-table.component'
import { MarkerDetailComponent } from './marker-detail/marker-detail.component'

import { UserClicksTableComponent } from './userclicks-table/userclicks-table.component'
import { UserClickDetailComponent } from './userclick-detail/userclick-detail.component'

import { VLinesTableComponent } from './vlines-table/vlines-table.component'
import { VLineDetailComponent } from './vline-detail/vline-detail.component'

import { VisualTracksTableComponent } from './visualtracks-table/visualtracks-table.component'
import { VisualTrackDetailComponent } from './visualtrack-detail/visualtrack-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongleaflet_go-checkoutschedulers/:GONG__StackPath', component: CheckoutSchedulersTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-checkoutscheduler-adder/:GONG__StackPath', component: CheckoutSchedulerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-checkoutscheduler-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: CheckoutSchedulerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-checkoutscheduler-detail/:id/:GONG__StackPath', component: CheckoutSchedulerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-circles/:GONG__StackPath', component: CirclesTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-circle-adder/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-circle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-circle-detail/:id/:GONG__StackPath', component: CircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-divicons/:GONG__StackPath', component: DivIconsTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-divicon-adder/:GONG__StackPath', component: DivIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-divicon-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: DivIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-divicon-detail/:id/:GONG__StackPath', component: DivIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroups/:GONG__StackPath', component: LayerGroupsTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroup-adder/:GONG__StackPath', component: LayerGroupDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroup-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LayerGroupDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroup-detail/:id/:GONG__StackPath', component: LayerGroupDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroupuses/:GONG__StackPath', component: LayerGroupUsesTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroupuse-adder/:GONG__StackPath', component: LayerGroupUseDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroupuse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LayerGroupUseDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-layergroupuse-detail/:id/:GONG__StackPath', component: LayerGroupUseDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-mapoptionss/:GONG__StackPath', component: MapOptionssTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-mapoptions-adder/:GONG__StackPath', component: MapOptionsDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-mapoptions-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MapOptionsDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-mapoptions-detail/:id/:GONG__StackPath', component: MapOptionsDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-markers/:GONG__StackPath', component: MarkersTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-marker-adder/:GONG__StackPath', component: MarkerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-marker-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MarkerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-marker-detail/:id/:GONG__StackPath', component: MarkerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-userclicks/:GONG__StackPath', component: UserClicksTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-userclick-adder/:GONG__StackPath', component: UserClickDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-userclick-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: UserClickDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-userclick-detail/:id/:GONG__StackPath', component: UserClickDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-vlines/:GONG__StackPath', component: VLinesTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-vline-adder/:GONG__StackPath', component: VLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-vline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: VLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-vline-detail/:id/:GONG__StackPath', component: VLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtracks/:GONG__StackPath', component: VisualTracksTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-adder/:GONG__StackPath', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-detail/:id/:GONG__StackPath', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

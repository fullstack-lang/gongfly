import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { VisualCentersTableComponent } from './visualcenters-table/visualcenters-table.component'
import { VisualCenterDetailComponent } from './visualcenter-detail/visualcenter-detail.component'
import { VisualCenterPresentationComponent } from './visualcenter-presentation/visualcenter-presentation.component'

import { VisualCirclesTableComponent } from './visualcircles-table/visualcircles-table.component'
import { VisualCircleDetailComponent } from './visualcircle-detail/visualcircle-detail.component'
import { VisualCirclePresentationComponent } from './visualcircle-presentation/visualcircle-presentation.component'

import { VisualIconsTableComponent } from './visualicons-table/visualicons-table.component'
import { VisualIconDetailComponent } from './visualicon-detail/visualicon-detail.component'
import { VisualIconPresentationComponent } from './visualicon-presentation/visualicon-presentation.component'

import { VisualLayersTableComponent } from './visuallayers-table/visuallayers-table.component'
import { VisualLayerDetailComponent } from './visuallayer-detail/visuallayer-detail.component'
import { VisualLayerPresentationComponent } from './visuallayer-presentation/visuallayer-presentation.component'

import { VisualLinesTableComponent } from './visuallines-table/visuallines-table.component'
import { VisualLineDetailComponent } from './visualline-detail/visualline-detail.component'
import { VisualLinePresentationComponent } from './visualline-presentation/visualline-presentation.component'

import { VisualMapsTableComponent } from './visualmaps-table/visualmaps-table.component'
import { VisualMapDetailComponent } from './visualmap-detail/visualmap-detail.component'
import { VisualMapPresentationComponent } from './visualmap-presentation/visualmap-presentation.component'

import { VisualTracksTableComponent } from './visualtracks-table/visualtracks-table.component'
import { VisualTrackDetailComponent } from './visualtrack-detail/visualtrack-detail.component'
import { VisualTrackPresentationComponent } from './visualtrack-presentation/visualtrack-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenters', component: VisualCentersTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenter-adder', component: VisualCenterDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenter-adder/:id/:originStruct/:originStructFieldName', component: VisualCenterDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenter-detail/:id', component: VisualCenterDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenter-presentation/:id', component: VisualCenterPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcenter-presentation-special/:id', component: VisualCenterPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisualcenterpres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircles', component: VisualCirclesTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircle-adder', component: VisualCircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircle-adder/:id/:originStruct/:originStructFieldName', component: VisualCircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircle-detail/:id', component: VisualCircleDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircle-presentation/:id', component: VisualCirclePresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualcircle-presentation-special/:id', component: VisualCirclePresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisualcirclepres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicons', component: VisualIconsTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicon-adder', component: VisualIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicon-adder/:id/:originStruct/:originStructFieldName', component: VisualIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicon-detail/:id', component: VisualIconDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicon-presentation/:id', component: VisualIconPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualicon-presentation-special/:id', component: VisualIconPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisualiconpres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayers', component: VisualLayersTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayer-adder', component: VisualLayerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayer-adder/:id/:originStruct/:originStructFieldName', component: VisualLayerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayer-detail/:id', component: VisualLayerDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayer-presentation/:id', component: VisualLayerPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallayer-presentation-special/:id', component: VisualLayerPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisuallayerpres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visuallines', component: VisualLinesTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualline-adder', component: VisualLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualline-adder/:id/:originStruct/:originStructFieldName', component: VisualLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualline-detail/:id', component: VisualLineDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualline-presentation/:id', component: VisualLinePresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualline-presentation-special/:id', component: VisualLinePresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisuallinepres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmaps', component: VisualMapsTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmap-adder', component: VisualMapDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmap-adder/:id/:originStruct/:originStructFieldName', component: VisualMapDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmap-detail/:id', component: VisualMapDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmap-presentation/:id', component: VisualMapPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualmap-presentation-special/:id', component: VisualMapPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisualmappres' },

	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtracks', component: VisualTracksTableComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_table' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-adder', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-adder/:id/:originStruct/:originStructFieldName', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-detail/:id', component: VisualTrackDetailComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_editor' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-presentation/:id', component: VisualTrackPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongleaflet_go-visualtrack-presentation-special/:id', component: VisualTrackPresentationComponent, outlet: 'github_com_fullstack_lang_gongleaflet_govisualtrackpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

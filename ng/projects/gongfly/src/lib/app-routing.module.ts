import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CivilianAirportsTableComponent } from './civilianairports-table/civilianairports-table.component'
import { CivilianAirportDetailComponent } from './civilianairport-detail/civilianairport-detail.component'

import { LinersTableComponent } from './liners-table/liners-table.component'
import { LinerDetailComponent } from './liner-detail/liner-detail.component'

import { MessagesTableComponent } from './messages-table/messages-table.component'
import { MessageDetailComponent } from './message-detail/message-detail.component'

import { OpsLinesTableComponent } from './opslines-table/opslines-table.component'
import { OpsLineDetailComponent } from './opsline-detail/opsline-detail.component'

import { RadarsTableComponent } from './radars-table/radars-table.component'
import { RadarDetailComponent } from './radar-detail/radar-detail.component'

import { SatellitesTableComponent } from './satellites-table/satellites-table.component'
import { SatelliteDetailComponent } from './satellite-detail/satellite-detail.component'

import { ScenariosTableComponent } from './scenarios-table/scenarios-table.component'
import { ScenarioDetailComponent } from './scenario-detail/scenario-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairports', component: CivilianAirportsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder/:id/:originStruct/:originStructFieldName', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-detail/:id', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-liners', component: LinersTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder/:id/:originStruct/:originStructFieldName', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-detail/:id', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-messages', component: MessagesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder/:id/:originStruct/:originStructFieldName', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-detail/:id', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-opslines', component: OpsLinesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder/:id/:originStruct/:originStructFieldName', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-detail/:id', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-radars', component: RadarsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder/:id/:originStruct/:originStructFieldName', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-detail/:id', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-satellites', component: SatellitesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder/:id/:originStruct/:originStructFieldName', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-detail/:id', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-scenarios', component: ScenariosTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder/:id/:originStruct/:originStructFieldName', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-detail/:id', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

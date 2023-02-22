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
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairports/:GONG__StackPath', component: CivilianAirportsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder/:GONG__StackPath', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-detail/:id/:GONG__StackPath', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-liners/:GONG__StackPath', component: LinersTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder/:GONG__StackPath', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-detail/:id/:GONG__StackPath', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-messages/:GONG__StackPath', component: MessagesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder/:GONG__StackPath', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-detail/:id/:GONG__StackPath', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-opslines/:GONG__StackPath', component: OpsLinesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder/:GONG__StackPath', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-detail/:id/:GONG__StackPath', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-radars/:GONG__StackPath', component: RadarsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder/:GONG__StackPath', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-detail/:id/:GONG__StackPath', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-satellites/:GONG__StackPath', component: SatellitesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder/:GONG__StackPath', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-detail/:id/:GONG__StackPath', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

	{ path: 'github_com_fullstack_lang_gongfly_go-scenarios/:GONG__StackPath', component: ScenariosTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder/:GONG__StackPath', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-detail/:id/:GONG__StackPath', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

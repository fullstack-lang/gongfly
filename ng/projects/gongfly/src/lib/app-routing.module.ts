import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CivilianAirportsTableComponent } from './civilianairports-table/civilianairports-table.component'
import { CivilianAirportDetailComponent } from './civilianairport-detail/civilianairport-detail.component'
import { CivilianAirportPresentationComponent } from './civilianairport-presentation/civilianairport-presentation.component'

import { LinersTableComponent } from './liners-table/liners-table.component'
import { LinerDetailComponent } from './liner-detail/liner-detail.component'
import { LinerPresentationComponent } from './liner-presentation/liner-presentation.component'

import { MessagesTableComponent } from './messages-table/messages-table.component'
import { MessageDetailComponent } from './message-detail/message-detail.component'
import { MessagePresentationComponent } from './message-presentation/message-presentation.component'

import { OpsLinesTableComponent } from './opslines-table/opslines-table.component'
import { OpsLineDetailComponent } from './opsline-detail/opsline-detail.component'
import { OpsLinePresentationComponent } from './opsline-presentation/opsline-presentation.component'

import { RadarsTableComponent } from './radars-table/radars-table.component'
import { RadarDetailComponent } from './radar-detail/radar-detail.component'
import { RadarPresentationComponent } from './radar-presentation/radar-presentation.component'

import { SatellitesTableComponent } from './satellites-table/satellites-table.component'
import { SatelliteDetailComponent } from './satellite-detail/satellite-detail.component'
import { SatellitePresentationComponent } from './satellite-presentation/satellite-presentation.component'

import { ScenariosTableComponent } from './scenarios-table/scenarios-table.component'
import { ScenarioDetailComponent } from './scenario-detail/scenario-detail.component'
import { ScenarioPresentationComponent } from './scenario-presentation/scenario-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairports', component: CivilianAirportsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-adder/:id/:originStruct/:originStructFieldName', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-detail/:id', component: CivilianAirportDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-presentation/:id', component: CivilianAirportPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-civilianairport-presentation-special/:id', component: CivilianAirportPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_gocivilianairportpres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-liners', component: LinersTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-adder/:id/:originStruct/:originStructFieldName', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-detail/:id', component: LinerDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-presentation/:id', component: LinerPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-liner-presentation-special/:id', component: LinerPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_golinerpres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-messages', component: MessagesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-adder/:id/:originStruct/:originStructFieldName', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-detail/:id', component: MessageDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-presentation/:id', component: MessagePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-message-presentation-special/:id', component: MessagePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_gomessagepres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-opslines', component: OpsLinesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-adder/:id/:originStruct/:originStructFieldName', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-detail/:id', component: OpsLineDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-presentation/:id', component: OpsLinePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-opsline-presentation-special/:id', component: OpsLinePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_goopslinepres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-radars', component: RadarsTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-adder/:id/:originStruct/:originStructFieldName', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-detail/:id', component: RadarDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-presentation/:id', component: RadarPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-radar-presentation-special/:id', component: RadarPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_goradarpres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-satellites', component: SatellitesTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-adder/:id/:originStruct/:originStructFieldName', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-detail/:id', component: SatelliteDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-presentation/:id', component: SatellitePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-satellite-presentation-special/:id', component: SatellitePresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_gosatellitepres' },

	{ path: 'github_com_fullstack_lang_gongfly_go-scenarios', component: ScenariosTableComponent, outlet: 'github_com_fullstack_lang_gongfly_go_table' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-adder/:id/:originStruct/:originStructFieldName', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-detail/:id', component: ScenarioDetailComponent, outlet: 'github_com_fullstack_lang_gongfly_go_editor' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-presentation/:id', component: ScenarioPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongfly_go-scenario-presentation-special/:id', component: ScenarioPresentationComponent, outlet: 'github_com_fullstack_lang_gongfly_goscenariopres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

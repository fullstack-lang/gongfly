import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { DummyAgentsTableComponent } from './dummyagents-table/dummyagents-table.component'
import { DummyAgentDetailComponent } from './dummyagent-detail/dummyagent-detail.component'

import { EnginesTableComponent } from './engines-table/engines-table.component'
import { EngineDetailComponent } from './engine-detail/engine-detail.component'

import { EventsTableComponent } from './events-table/events-table.component'
import { EventDetailComponent } from './event-detail/event-detail.component'

import { GongsimCommandsTableComponent } from './gongsimcommands-table/gongsimcommands-table.component'
import { GongsimCommandDetailComponent } from './gongsimcommand-detail/gongsimcommand-detail.component'

import { GongsimStatussTableComponent } from './gongsimstatuss-table/gongsimstatuss-table.component'
import { GongsimStatusDetailComponent } from './gongsimstatus-detail/gongsimstatus-detail.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagents/:GONG__StackPath', component: DummyAgentsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-adder/:GONG__StackPath', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-dummyagent-detail/:id/:GONG__StackPath', component: DummyAgentDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-engines/:GONG__StackPath', component: EnginesTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-adder/:GONG__StackPath', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-engine-detail/:id/:GONG__StackPath', component: EngineDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-events/:GONG__StackPath', component: EventsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-adder/:GONG__StackPath', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-event-detail/:id/:GONG__StackPath', component: EventDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommands/:GONG__StackPath', component: GongsimCommandsTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-adder/:GONG__StackPath', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimcommand-detail/:id/:GONG__StackPath', component: GongsimCommandDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatuss/:GONG__StackPath', component: GongsimStatussTableComponent, outlet: 'github_com_fullstack_lang_gongsim_go_table' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-adder/:GONG__StackPath', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },
	{ path: 'github_com_fullstack_lang_gongsim_go-gongsimstatus-detail/:id/:GONG__StackPath', component: GongsimStatusDetailComponent, outlet: 'github_com_fullstack_lang_gongsim_go_editor' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

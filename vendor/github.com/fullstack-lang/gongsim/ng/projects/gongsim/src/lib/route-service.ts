import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

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


@Injectable({
    providedIn: 'root'
})
export class RouteService {
    private routes: Routes = []

    constructor(private router: Router) { }

    public addRoutes(newRoutes: Routes): void {
        const existingRoutes = this.router.config
        this.routes = this.router.config

        for (let newRoute of newRoutes) {
            if (!existingRoutes.includes(newRoute)) {
                this.routes.push(newRoute)
            }
        }
        this.router.resetConfig(this.routes)
    }

    getPathRoot(): string {
        return 'github_com_fullstack_lang_gongsim_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getDummyAgentTablePath(): string {
        return this.getPathRoot() + '-dummyagents/:GONG__StackPath'
    }
    getDummyAgentTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAgentTablePath(), component: DummyAgentsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDummyAgentAdderPath(): string {
        return this.getPathRoot() + '-dummyagent-adder/:GONG__StackPath'
    }
    getDummyAgentAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAgentAdderPath(), component: DummyAgentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDummyAgentAdderForUsePath(): string {
        return this.getPathRoot() + '-dummyagent-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDummyAgentAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAgentAdderForUsePath(), component: DummyAgentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDummyAgentDetailPath(): string {
        return this.getPathRoot() + '-dummyagent-detail/:id/:GONG__StackPath'
    }
    getDummyAgentDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDummyAgentDetailPath(), component: DummyAgentDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getEngineTablePath(): string {
        return this.getPathRoot() + '-engines/:GONG__StackPath'
    }
    getEngineTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEngineTablePath(), component: EnginesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getEngineAdderPath(): string {
        return this.getPathRoot() + '-engine-adder/:GONG__StackPath'
    }
    getEngineAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEngineAdderPath(), component: EngineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEngineAdderForUsePath(): string {
        return this.getPathRoot() + '-engine-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getEngineAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEngineAdderForUsePath(), component: EngineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEngineDetailPath(): string {
        return this.getPathRoot() + '-engine-detail/:id/:GONG__StackPath'
    }
    getEngineDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEngineDetailPath(), component: EngineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getEventTablePath(): string {
        return this.getPathRoot() + '-events/:GONG__StackPath'
    }
    getEventTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEventTablePath(), component: EventsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getEventAdderPath(): string {
        return this.getPathRoot() + '-event-adder/:GONG__StackPath'
    }
    getEventAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEventAdderPath(), component: EventDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEventAdderForUsePath(): string {
        return this.getPathRoot() + '-event-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getEventAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEventAdderForUsePath(), component: EventDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getEventDetailPath(): string {
        return this.getPathRoot() + '-event-detail/:id/:GONG__StackPath'
    }
    getEventDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getEventDetailPath(), component: EventDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongsimCommandTablePath(): string {
        return this.getPathRoot() + '-gongsimcommands/:GONG__StackPath'
    }
    getGongsimCommandTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimCommandTablePath(), component: GongsimCommandsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongsimCommandAdderPath(): string {
        return this.getPathRoot() + '-gongsimcommand-adder/:GONG__StackPath'
    }
    getGongsimCommandAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimCommandAdderPath(), component: GongsimCommandDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongsimCommandAdderForUsePath(): string {
        return this.getPathRoot() + '-gongsimcommand-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongsimCommandAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimCommandAdderForUsePath(), component: GongsimCommandDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongsimCommandDetailPath(): string {
        return this.getPathRoot() + '-gongsimcommand-detail/:id/:GONG__StackPath'
    }
    getGongsimCommandDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimCommandDetailPath(), component: GongsimCommandDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getGongsimStatusTablePath(): string {
        return this.getPathRoot() + '-gongsimstatuss/:GONG__StackPath'
    }
    getGongsimStatusTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimStatusTablePath(), component: GongsimStatussTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getGongsimStatusAdderPath(): string {
        return this.getPathRoot() + '-gongsimstatus-adder/:GONG__StackPath'
    }
    getGongsimStatusAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimStatusAdderPath(), component: GongsimStatusDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongsimStatusAdderForUsePath(): string {
        return this.getPathRoot() + '-gongsimstatus-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getGongsimStatusAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimStatusAdderForUsePath(), component: GongsimStatusDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getGongsimStatusDetailPath(): string {
        return this.getPathRoot() + '-gongsimstatus-detail/:id/:GONG__StackPath'
    }
    getGongsimStatusDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getGongsimStatusDetailPath(), component: GongsimStatusDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getDummyAgentTableRoute(stackPath),
            this.getDummyAgentAdderRoute(stackPath),
            this.getDummyAgentAdderForUseRoute(stackPath),
            this.getDummyAgentDetailRoute(stackPath),

            this.getEngineTableRoute(stackPath),
            this.getEngineAdderRoute(stackPath),
            this.getEngineAdderForUseRoute(stackPath),
            this.getEngineDetailRoute(stackPath),

            this.getEventTableRoute(stackPath),
            this.getEventAdderRoute(stackPath),
            this.getEventAdderForUseRoute(stackPath),
            this.getEventDetailRoute(stackPath),

            this.getGongsimCommandTableRoute(stackPath),
            this.getGongsimCommandAdderRoute(stackPath),
            this.getGongsimCommandAdderForUseRoute(stackPath),
            this.getGongsimCommandDetailRoute(stackPath),

            this.getGongsimStatusTableRoute(stackPath),
            this.getGongsimStatusAdderRoute(stackPath),
            this.getGongsimStatusAdderForUseRoute(stackPath),
            this.getGongsimStatusDetailRoute(stackPath),

        ])
    }
}

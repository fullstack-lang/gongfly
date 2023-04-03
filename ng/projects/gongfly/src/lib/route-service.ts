import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

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
        return 'github_com_fullstack_lang_gongfly_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getCivilianAirportTablePath(): string {
        return this.getPathRoot() + '-civilianairports/:GONG__StackPath'
    }
    getCivilianAirportTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCivilianAirportTablePath(), component: CivilianAirportsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCivilianAirportAdderPath(): string {
        return this.getPathRoot() + '-civilianairport-adder/:GONG__StackPath'
    }
    getCivilianAirportAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCivilianAirportAdderPath(), component: CivilianAirportDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCivilianAirportAdderForUsePath(): string {
        return this.getPathRoot() + '-civilianairport-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCivilianAirportAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCivilianAirportAdderForUsePath(), component: CivilianAirportDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCivilianAirportDetailPath(): string {
        return this.getPathRoot() + '-civilianairport-detail/:id/:GONG__StackPath'
    }
    getCivilianAirportDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCivilianAirportDetailPath(), component: CivilianAirportDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLinerTablePath(): string {
        return this.getPathRoot() + '-liners/:GONG__StackPath'
    }
    getLinerTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinerTablePath(), component: LinersTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLinerAdderPath(): string {
        return this.getPathRoot() + '-liner-adder/:GONG__StackPath'
    }
    getLinerAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinerAdderPath(), component: LinerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLinerAdderForUsePath(): string {
        return this.getPathRoot() + '-liner-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLinerAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinerAdderForUsePath(), component: LinerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLinerDetailPath(): string {
        return this.getPathRoot() + '-liner-detail/:id/:GONG__StackPath'
    }
    getLinerDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLinerDetailPath(), component: LinerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMessageTablePath(): string {
        return this.getPathRoot() + '-messages/:GONG__StackPath'
    }
    getMessageTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMessageTablePath(), component: MessagesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMessageAdderPath(): string {
        return this.getPathRoot() + '-message-adder/:GONG__StackPath'
    }
    getMessageAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMessageAdderPath(), component: MessageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMessageAdderForUsePath(): string {
        return this.getPathRoot() + '-message-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMessageAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMessageAdderForUsePath(), component: MessageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMessageDetailPath(): string {
        return this.getPathRoot() + '-message-detail/:id/:GONG__StackPath'
    }
    getMessageDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMessageDetailPath(), component: MessageDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getOpsLineTablePath(): string {
        return this.getPathRoot() + '-opslines/:GONG__StackPath'
    }
    getOpsLineTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOpsLineTablePath(), component: OpsLinesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getOpsLineAdderPath(): string {
        return this.getPathRoot() + '-opsline-adder/:GONG__StackPath'
    }
    getOpsLineAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOpsLineAdderPath(), component: OpsLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOpsLineAdderForUsePath(): string {
        return this.getPathRoot() + '-opsline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getOpsLineAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOpsLineAdderForUsePath(), component: OpsLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getOpsLineDetailPath(): string {
        return this.getPathRoot() + '-opsline-detail/:id/:GONG__StackPath'
    }
    getOpsLineDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getOpsLineDetailPath(), component: OpsLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getRadarTablePath(): string {
        return this.getPathRoot() + '-radars/:GONG__StackPath'
    }
    getRadarTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRadarTablePath(), component: RadarsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getRadarAdderPath(): string {
        return this.getPathRoot() + '-radar-adder/:GONG__StackPath'
    }
    getRadarAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRadarAdderPath(), component: RadarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRadarAdderForUsePath(): string {
        return this.getPathRoot() + '-radar-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getRadarAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRadarAdderForUsePath(), component: RadarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getRadarDetailPath(): string {
        return this.getPathRoot() + '-radar-detail/:id/:GONG__StackPath'
    }
    getRadarDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getRadarDetailPath(), component: RadarDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getSatelliteTablePath(): string {
        return this.getPathRoot() + '-satellites/:GONG__StackPath'
    }
    getSatelliteTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSatelliteTablePath(), component: SatellitesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getSatelliteAdderPath(): string {
        return this.getPathRoot() + '-satellite-adder/:GONG__StackPath'
    }
    getSatelliteAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSatelliteAdderPath(), component: SatelliteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSatelliteAdderForUsePath(): string {
        return this.getPathRoot() + '-satellite-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getSatelliteAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSatelliteAdderForUsePath(), component: SatelliteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getSatelliteDetailPath(): string {
        return this.getPathRoot() + '-satellite-detail/:id/:GONG__StackPath'
    }
    getSatelliteDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getSatelliteDetailPath(), component: SatelliteDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getScenarioTablePath(): string {
        return this.getPathRoot() + '-scenarios/:GONG__StackPath'
    }
    getScenarioTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getScenarioTablePath(), component: ScenariosTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getScenarioAdderPath(): string {
        return this.getPathRoot() + '-scenario-adder/:GONG__StackPath'
    }
    getScenarioAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getScenarioAdderPath(), component: ScenarioDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getScenarioAdderForUsePath(): string {
        return this.getPathRoot() + '-scenario-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getScenarioAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getScenarioAdderForUsePath(), component: ScenarioDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getScenarioDetailPath(): string {
        return this.getPathRoot() + '-scenario-detail/:id/:GONG__StackPath'
    }
    getScenarioDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getScenarioDetailPath(), component: ScenarioDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getCivilianAirportTableRoute(stackPath),
            this.getCivilianAirportAdderRoute(stackPath),
            this.getCivilianAirportAdderForUseRoute(stackPath),
            this.getCivilianAirportDetailRoute(stackPath),

            this.getLinerTableRoute(stackPath),
            this.getLinerAdderRoute(stackPath),
            this.getLinerAdderForUseRoute(stackPath),
            this.getLinerDetailRoute(stackPath),

            this.getMessageTableRoute(stackPath),
            this.getMessageAdderRoute(stackPath),
            this.getMessageAdderForUseRoute(stackPath),
            this.getMessageDetailRoute(stackPath),

            this.getOpsLineTableRoute(stackPath),
            this.getOpsLineAdderRoute(stackPath),
            this.getOpsLineAdderForUseRoute(stackPath),
            this.getOpsLineDetailRoute(stackPath),

            this.getRadarTableRoute(stackPath),
            this.getRadarAdderRoute(stackPath),
            this.getRadarAdderForUseRoute(stackPath),
            this.getRadarDetailRoute(stackPath),

            this.getSatelliteTableRoute(stackPath),
            this.getSatelliteAdderRoute(stackPath),
            this.getSatelliteAdderForUseRoute(stackPath),
            this.getSatelliteDetailRoute(stackPath),

            this.getScenarioTableRoute(stackPath),
            this.getScenarioAdderRoute(stackPath),
            this.getScenarioAdderForUseRoute(stackPath),
            this.getScenarioDetailRoute(stackPath),

        ])
    }
}

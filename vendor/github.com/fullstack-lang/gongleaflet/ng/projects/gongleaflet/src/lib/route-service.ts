import { Injectable } from '@angular/core';
import { Route, Router, Routes } from '@angular/router';

// insertion point for imports
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
        return 'github_com_fullstack_lang_gongleaflet_go'
    }
    getTableOutlet(stackPath: string): string {
        return this.getPathRoot() + '_table' + '_' + stackPath
    }
    getEditorOutlet(stackPath: string): string {
        return this.getPathRoot() + '_editor' + '_' + stackPath
    }
    // insertion point for per gongstruct route/path getters
    getCircleTablePath(): string {
        return this.getPathRoot() + '-circles/:GONG__StackPath'
    }
    getCircleTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleTablePath(), component: CirclesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getCircleAdderPath(): string {
        return this.getPathRoot() + '-circle-adder/:GONG__StackPath'
    }
    getCircleAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleAdderPath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCircleAdderForUsePath(): string {
        return this.getPathRoot() + '-circle-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getCircleAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleAdderForUsePath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getCircleDetailPath(): string {
        return this.getPathRoot() + '-circle-detail/:id/:GONG__StackPath'
    }
    getCircleDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getCircleDetailPath(), component: CircleDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getDivIconTablePath(): string {
        return this.getPathRoot() + '-divicons/:GONG__StackPath'
    }
    getDivIconTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDivIconTablePath(), component: DivIconsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getDivIconAdderPath(): string {
        return this.getPathRoot() + '-divicon-adder/:GONG__StackPath'
    }
    getDivIconAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDivIconAdderPath(), component: DivIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDivIconAdderForUsePath(): string {
        return this.getPathRoot() + '-divicon-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getDivIconAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDivIconAdderForUsePath(), component: DivIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getDivIconDetailPath(): string {
        return this.getPathRoot() + '-divicon-detail/:id/:GONG__StackPath'
    }
    getDivIconDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getDivIconDetailPath(), component: DivIconDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLayerGroupTablePath(): string {
        return this.getPathRoot() + '-layergroups/:GONG__StackPath'
    }
    getLayerGroupTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupTablePath(), component: LayerGroupsTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLayerGroupAdderPath(): string {
        return this.getPathRoot() + '-layergroup-adder/:GONG__StackPath'
    }
    getLayerGroupAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupAdderPath(), component: LayerGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLayerGroupAdderForUsePath(): string {
        return this.getPathRoot() + '-layergroup-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLayerGroupAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupAdderForUsePath(), component: LayerGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLayerGroupDetailPath(): string {
        return this.getPathRoot() + '-layergroup-detail/:id/:GONG__StackPath'
    }
    getLayerGroupDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupDetailPath(), component: LayerGroupDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getLayerGroupUseTablePath(): string {
        return this.getPathRoot() + '-layergroupuses/:GONG__StackPath'
    }
    getLayerGroupUseTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupUseTablePath(), component: LayerGroupUsesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getLayerGroupUseAdderPath(): string {
        return this.getPathRoot() + '-layergroupuse-adder/:GONG__StackPath'
    }
    getLayerGroupUseAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupUseAdderPath(), component: LayerGroupUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLayerGroupUseAdderForUsePath(): string {
        return this.getPathRoot() + '-layergroupuse-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getLayerGroupUseAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupUseAdderForUsePath(), component: LayerGroupUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getLayerGroupUseDetailPath(): string {
        return this.getPathRoot() + '-layergroupuse-detail/:id/:GONG__StackPath'
    }
    getLayerGroupUseDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getLayerGroupUseDetailPath(), component: LayerGroupUseDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMapOptionsTablePath(): string {
        return this.getPathRoot() + '-mapoptionss/:GONG__StackPath'
    }
    getMapOptionsTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMapOptionsTablePath(), component: MapOptionssTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMapOptionsAdderPath(): string {
        return this.getPathRoot() + '-mapoptions-adder/:GONG__StackPath'
    }
    getMapOptionsAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMapOptionsAdderPath(), component: MapOptionsDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMapOptionsAdderForUsePath(): string {
        return this.getPathRoot() + '-mapoptions-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMapOptionsAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMapOptionsAdderForUsePath(), component: MapOptionsDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMapOptionsDetailPath(): string {
        return this.getPathRoot() + '-mapoptions-detail/:id/:GONG__StackPath'
    }
    getMapOptionsDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMapOptionsDetailPath(), component: MapOptionsDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getMarkerTablePath(): string {
        return this.getPathRoot() + '-markers/:GONG__StackPath'
    }
    getMarkerTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMarkerTablePath(), component: MarkersTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getMarkerAdderPath(): string {
        return this.getPathRoot() + '-marker-adder/:GONG__StackPath'
    }
    getMarkerAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMarkerAdderPath(), component: MarkerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMarkerAdderForUsePath(): string {
        return this.getPathRoot() + '-marker-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getMarkerAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMarkerAdderForUsePath(), component: MarkerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getMarkerDetailPath(): string {
        return this.getPathRoot() + '-marker-detail/:id/:GONG__StackPath'
    }
    getMarkerDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getMarkerDetailPath(), component: MarkerDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getUserClickTablePath(): string {
        return this.getPathRoot() + '-userclicks/:GONG__StackPath'
    }
    getUserClickTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUserClickTablePath(), component: UserClicksTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getUserClickAdderPath(): string {
        return this.getPathRoot() + '-userclick-adder/:GONG__StackPath'
    }
    getUserClickAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUserClickAdderPath(), component: UserClickDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUserClickAdderForUsePath(): string {
        return this.getPathRoot() + '-userclick-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getUserClickAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUserClickAdderForUsePath(), component: UserClickDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getUserClickDetailPath(): string {
        return this.getPathRoot() + '-userclick-detail/:id/:GONG__StackPath'
    }
    getUserClickDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getUserClickDetailPath(), component: UserClickDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getVLineTablePath(): string {
        return this.getPathRoot() + '-vlines/:GONG__StackPath'
    }
    getVLineTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVLineTablePath(), component: VLinesTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getVLineAdderPath(): string {
        return this.getPathRoot() + '-vline-adder/:GONG__StackPath'
    }
    getVLineAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVLineAdderPath(), component: VLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVLineAdderForUsePath(): string {
        return this.getPathRoot() + '-vline-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getVLineAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVLineAdderForUsePath(), component: VLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVLineDetailPath(): string {
        return this.getPathRoot() + '-vline-detail/:id/:GONG__StackPath'
    }
    getVLineDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVLineDetailPath(), component: VLineDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }

    getVisualTrackTablePath(): string {
        return this.getPathRoot() + '-visualtracks/:GONG__StackPath'
    }
    getVisualTrackTableRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVisualTrackTablePath(), component: VisualTracksTableComponent, outlet: this.getTableOutlet(stackPath) }
        return route
    }
    getVisualTrackAdderPath(): string {
        return this.getPathRoot() + '-visualtrack-adder/:GONG__StackPath'
    }
    getVisualTrackAdderRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVisualTrackAdderPath(), component: VisualTrackDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVisualTrackAdderForUsePath(): string {
        return this.getPathRoot() + '-visualtrack-adder/:id/:originStruct/:originStructFieldName/:GONG__StackPath'
    }
    getVisualTrackAdderForUseRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVisualTrackAdderForUsePath(), component: VisualTrackDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }
    getVisualTrackDetailPath(): string {
        return this.getPathRoot() + '-visualtrack-detail/:id/:GONG__StackPath'
    }
    getVisualTrackDetailRoute(stackPath: string): Route {
        let route: Route =
            { path: this.getVisualTrackDetailPath(), component: VisualTrackDetailComponent, outlet: this.getEditorOutlet(stackPath) }
        return route
    }



    addDataPanelRoutes(stackPath: string) {

        this.addRoutes([
            // insertion point for all routes getter
            this.getCircleTableRoute(stackPath),
            this.getCircleAdderRoute(stackPath),
            this.getCircleAdderForUseRoute(stackPath),
            this.getCircleDetailRoute(stackPath),

            this.getDivIconTableRoute(stackPath),
            this.getDivIconAdderRoute(stackPath),
            this.getDivIconAdderForUseRoute(stackPath),
            this.getDivIconDetailRoute(stackPath),

            this.getLayerGroupTableRoute(stackPath),
            this.getLayerGroupAdderRoute(stackPath),
            this.getLayerGroupAdderForUseRoute(stackPath),
            this.getLayerGroupDetailRoute(stackPath),

            this.getLayerGroupUseTableRoute(stackPath),
            this.getLayerGroupUseAdderRoute(stackPath),
            this.getLayerGroupUseAdderForUseRoute(stackPath),
            this.getLayerGroupUseDetailRoute(stackPath),

            this.getMapOptionsTableRoute(stackPath),
            this.getMapOptionsAdderRoute(stackPath),
            this.getMapOptionsAdderForUseRoute(stackPath),
            this.getMapOptionsDetailRoute(stackPath),

            this.getMarkerTableRoute(stackPath),
            this.getMarkerAdderRoute(stackPath),
            this.getMarkerAdderForUseRoute(stackPath),
            this.getMarkerDetailRoute(stackPath),

            this.getUserClickTableRoute(stackPath),
            this.getUserClickAdderRoute(stackPath),
            this.getUserClickAdderForUseRoute(stackPath),
            this.getUserClickDetailRoute(stackPath),

            this.getVLineTableRoute(stackPath),
            this.getVLineAdderRoute(stackPath),
            this.getVLineAdderForUseRoute(stackPath),
            this.getVLineDetailRoute(stackPath),

            this.getVisualTrackTableRoute(stackPath),
            this.getVisualTrackAdderRoute(stackPath),
            this.getVisualTrackAdderForUseRoute(stackPath),
            this.getVisualTrackDetailRoute(stackPath),

        ])
    }
}

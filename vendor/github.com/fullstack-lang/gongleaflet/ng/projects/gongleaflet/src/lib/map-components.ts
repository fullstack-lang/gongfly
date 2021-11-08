// insertion point sub template for components imports 
  import { VisualCentersTableComponent } from './visualcenters-table/visualcenters-table.component'
  import { VisualCenterSortingComponent } from './visualcenter-sorting/visualcenter-sorting.component'
  import { VisualCirclesTableComponent } from './visualcircles-table/visualcircles-table.component'
  import { VisualCircleSortingComponent } from './visualcircle-sorting/visualcircle-sorting.component'
  import { VisualIconsTableComponent } from './visualicons-table/visualicons-table.component'
  import { VisualIconSortingComponent } from './visualicon-sorting/visualicon-sorting.component'
  import { VisualLayersTableComponent } from './visuallayers-table/visuallayers-table.component'
  import { VisualLayerSortingComponent } from './visuallayer-sorting/visuallayer-sorting.component'
  import { VisualLinesTableComponent } from './visuallines-table/visuallines-table.component'
  import { VisualLineSortingComponent } from './visualline-sorting/visualline-sorting.component'
  import { VisualMapsTableComponent } from './visualmaps-table/visualmaps-table.component'
  import { VisualMapSortingComponent } from './visualmap-sorting/visualmap-sorting.component'
  import { VisualTracksTableComponent } from './visualtracks-table/visualtracks-table.component'
  import { VisualTrackSortingComponent } from './visualtrack-sorting/visualtrack-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfVisualCentersComponents: Map<string, any> = new Map([["VisualCentersTableComponent", VisualCentersTableComponent],])
  export const MapOfVisualCenterSortingComponents: Map<string, any> = new Map([["VisualCenterSortingComponent", VisualCenterSortingComponent],])
  export const MapOfVisualCirclesComponents: Map<string, any> = new Map([["VisualCirclesTableComponent", VisualCirclesTableComponent],])
  export const MapOfVisualCircleSortingComponents: Map<string, any> = new Map([["VisualCircleSortingComponent", VisualCircleSortingComponent],])
  export const MapOfVisualIconsComponents: Map<string, any> = new Map([["VisualIconsTableComponent", VisualIconsTableComponent],])
  export const MapOfVisualIconSortingComponents: Map<string, any> = new Map([["VisualIconSortingComponent", VisualIconSortingComponent],])
  export const MapOfVisualLayersComponents: Map<string, any> = new Map([["VisualLayersTableComponent", VisualLayersTableComponent],])
  export const MapOfVisualLayerSortingComponents: Map<string, any> = new Map([["VisualLayerSortingComponent", VisualLayerSortingComponent],])
  export const MapOfVisualLinesComponents: Map<string, any> = new Map([["VisualLinesTableComponent", VisualLinesTableComponent],])
  export const MapOfVisualLineSortingComponents: Map<string, any> = new Map([["VisualLineSortingComponent", VisualLineSortingComponent],])
  export const MapOfVisualMapsComponents: Map<string, any> = new Map([["VisualMapsTableComponent", VisualMapsTableComponent],])
  export const MapOfVisualMapSortingComponents: Map<string, any> = new Map([["VisualMapSortingComponent", VisualMapSortingComponent],])
  export const MapOfVisualTracksComponents: Map<string, any> = new Map([["VisualTracksTableComponent", VisualTracksTableComponent],])
  export const MapOfVisualTrackSortingComponents: Map<string, any> = new Map([["VisualTrackSortingComponent", VisualTrackSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["VisualCenter", MapOfVisualCentersComponents],
      ["VisualCircle", MapOfVisualCirclesComponents],
      ["VisualIcon", MapOfVisualIconsComponents],
      ["VisualLayer", MapOfVisualLayersComponents],
      ["VisualLine", MapOfVisualLinesComponents],
      ["VisualMap", MapOfVisualMapsComponents],
      ["VisualTrack", MapOfVisualTracksComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["VisualCenter", MapOfVisualCenterSortingComponents],
      ["VisualCircle", MapOfVisualCircleSortingComponents],
      ["VisualIcon", MapOfVisualIconSortingComponents],
      ["VisualLayer", MapOfVisualLayerSortingComponents],
      ["VisualLine", MapOfVisualLineSortingComponents],
      ["VisualMap", MapOfVisualMapSortingComponents],
      ["VisualTrack", MapOfVisualTrackSortingComponents],
    ]
  )

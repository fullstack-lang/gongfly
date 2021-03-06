// insertion point sub template for components imports 
  import { CivilianAirportsTableComponent } from './civilianairports-table/civilianairports-table.component'
  import { CivilianAirportSortingComponent } from './civilianairport-sorting/civilianairport-sorting.component'
  import { LinersTableComponent } from './liners-table/liners-table.component'
  import { LinerSortingComponent } from './liner-sorting/liner-sorting.component'
  import { MessagesTableComponent } from './messages-table/messages-table.component'
  import { MessageSortingComponent } from './message-sorting/message-sorting.component'
  import { OpsLinesTableComponent } from './opslines-table/opslines-table.component'
  import { OpsLineSortingComponent } from './opsline-sorting/opsline-sorting.component'
  import { RadarsTableComponent } from './radars-table/radars-table.component'
  import { RadarSortingComponent } from './radar-sorting/radar-sorting.component'
  import { SatellitesTableComponent } from './satellites-table/satellites-table.component'
  import { SatelliteSortingComponent } from './satellite-sorting/satellite-sorting.component'
  import { ScenariosTableComponent } from './scenarios-table/scenarios-table.component'
  import { ScenarioSortingComponent } from './scenario-sorting/scenario-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfCivilianAirportsComponents: Map<string, any> = new Map([["CivilianAirportsTableComponent", CivilianAirportsTableComponent],])
  export const MapOfCivilianAirportSortingComponents: Map<string, any> = new Map([["CivilianAirportSortingComponent", CivilianAirportSortingComponent],])
  export const MapOfLinersComponents: Map<string, any> = new Map([["LinersTableComponent", LinersTableComponent],])
  export const MapOfLinerSortingComponents: Map<string, any> = new Map([["LinerSortingComponent", LinerSortingComponent],])
  export const MapOfMessagesComponents: Map<string, any> = new Map([["MessagesTableComponent", MessagesTableComponent],])
  export const MapOfMessageSortingComponents: Map<string, any> = new Map([["MessageSortingComponent", MessageSortingComponent],])
  export const MapOfOpsLinesComponents: Map<string, any> = new Map([["OpsLinesTableComponent", OpsLinesTableComponent],])
  export const MapOfOpsLineSortingComponents: Map<string, any> = new Map([["OpsLineSortingComponent", OpsLineSortingComponent],])
  export const MapOfRadarsComponents: Map<string, any> = new Map([["RadarsTableComponent", RadarsTableComponent],])
  export const MapOfRadarSortingComponents: Map<string, any> = new Map([["RadarSortingComponent", RadarSortingComponent],])
  export const MapOfSatellitesComponents: Map<string, any> = new Map([["SatellitesTableComponent", SatellitesTableComponent],])
  export const MapOfSatelliteSortingComponents: Map<string, any> = new Map([["SatelliteSortingComponent", SatelliteSortingComponent],])
  export const MapOfScenariosComponents: Map<string, any> = new Map([["ScenariosTableComponent", ScenariosTableComponent],])
  export const MapOfScenarioSortingComponents: Map<string, any> = new Map([["ScenarioSortingComponent", ScenarioSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["CivilianAirport", MapOfCivilianAirportsComponents],
      ["Liner", MapOfLinersComponents],
      ["Message", MapOfMessagesComponents],
      ["OpsLine", MapOfOpsLinesComponents],
      ["Radar", MapOfRadarsComponents],
      ["Satellite", MapOfSatellitesComponents],
      ["Scenario", MapOfScenariosComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["CivilianAirport", MapOfCivilianAirportSortingComponents],
      ["Liner", MapOfLinerSortingComponents],
      ["Message", MapOfMessageSortingComponents],
      ["OpsLine", MapOfOpsLineSortingComponents],
      ["Radar", MapOfRadarSortingComponents],
      ["Satellite", MapOfSatelliteSortingComponents],
      ["Scenario", MapOfScenarioSortingComponents],
    ]
  )

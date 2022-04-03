// insertion point sub template for components imports 
  import { AnotherDummyDatasTableComponent } from './anotherdummydatas-table/anotherdummydatas-table.component'
  import { AnotherDummyDataSortingComponent } from './anotherdummydata-sorting/anotherdummydata-sorting.component'
  import { CellsTableComponent } from './cells-table/cells-table.component'
  import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
  import { DummyDatasTableComponent } from './dummydatas-table/dummydatas-table.component'
  import { DummyDataSortingComponent } from './dummydata-sorting/dummydata-sorting.component'
  import { ElementsTableComponent } from './elements-table/elements-table.component'
  import { ElementSortingComponent } from './element-sorting/element-sorting.component'
  import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
  import { MarkdownContentSortingComponent } from './markdowncontent-sorting/markdowncontent-sorting.component'
  import { RowsTableComponent } from './rows-table/rows-table.component'
  import { RowSortingComponent } from './row-sorting/row-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfAnotherDummyDatasComponents: Map<string, any> = new Map([["AnotherDummyDatasTableComponent", AnotherDummyDatasTableComponent],])
  export const MapOfAnotherDummyDataSortingComponents: Map<string, any> = new Map([["AnotherDummyDataSortingComponent", AnotherDummyDataSortingComponent],])
  export const MapOfCellsComponents: Map<string, any> = new Map([["CellsTableComponent", CellsTableComponent],])
  export const MapOfCellSortingComponents: Map<string, any> = new Map([["CellSortingComponent", CellSortingComponent],])
  export const MapOfDummyDatasComponents: Map<string, any> = new Map([["DummyDatasTableComponent", DummyDatasTableComponent],])
  export const MapOfDummyDataSortingComponents: Map<string, any> = new Map([["DummyDataSortingComponent", DummyDataSortingComponent],])
  export const MapOfElementsComponents: Map<string, any> = new Map([["ElementsTableComponent", ElementsTableComponent],])
  export const MapOfElementSortingComponents: Map<string, any> = new Map([["ElementSortingComponent", ElementSortingComponent],])
  export const MapOfMarkdownContentsComponents: Map<string, any> = new Map([["MarkdownContentsTableComponent", MarkdownContentsTableComponent],])
  export const MapOfMarkdownContentSortingComponents: Map<string, any> = new Map([["MarkdownContentSortingComponent", MarkdownContentSortingComponent],])
  export const MapOfRowsComponents: Map<string, any> = new Map([["RowsTableComponent", RowsTableComponent],])
  export const MapOfRowSortingComponents: Map<string, any> = new Map([["RowSortingComponent", RowSortingComponent],])

// map of all ng components of the stacks
export const MapOfComponents: Map<string, any> =
  new Map(
    [
      // insertion point sub template for map of components 
      ["AnotherDummyData", MapOfAnotherDummyDatasComponents],
      ["Cell", MapOfCellsComponents],
      ["DummyData", MapOfDummyDatasComponents],
      ["Element", MapOfElementsComponents],
      ["MarkdownContent", MapOfMarkdownContentsComponents],
      ["Row", MapOfRowsComponents],
    ]
  )

// map of all ng components of the stacks
export const MapOfSortingComponents: Map<string, any> =
  new Map(
    [
    // insertion point sub template for map of sorting components 
      ["AnotherDummyData", MapOfAnotherDummyDataSortingComponents],
      ["Cell", MapOfCellSortingComponents],
      ["DummyData", MapOfDummyDataSortingComponents],
      ["Element", MapOfElementSortingComponents],
      ["MarkdownContent", MapOfMarkdownContentSortingComponents],
      ["Row", MapOfRowSortingComponents],
    ]
  )

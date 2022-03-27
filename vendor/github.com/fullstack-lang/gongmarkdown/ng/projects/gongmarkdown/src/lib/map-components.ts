// insertion point sub template for components imports 
  import { CellsTableComponent } from './cells-table/cells-table.component'
  import { CellSortingComponent } from './cell-sorting/cell-sorting.component'
  import { ElementsTableComponent } from './elements-table/elements-table.component'
  import { ElementSortingComponent } from './element-sorting/element-sorting.component'
  import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
  import { MarkdownContentSortingComponent } from './markdowncontent-sorting/markdowncontent-sorting.component'
  import { RowsTableComponent } from './rows-table/rows-table.component'
  import { RowSortingComponent } from './row-sorting/row-sorting.component'

// insertion point sub template for map of components per struct 
  export const MapOfCellsComponents: Map<string, any> = new Map([["CellsTableComponent", CellsTableComponent],])
  export const MapOfCellSortingComponents: Map<string, any> = new Map([["CellSortingComponent", CellSortingComponent],])
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
      ["Cell", MapOfCellsComponents],
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
      ["Cell", MapOfCellSortingComponents],
      ["Element", MapOfElementSortingComponents],
      ["MarkdownContent", MapOfMarkdownContentSortingComponents],
      ["Row", MapOfRowSortingComponents],
    ]
  )

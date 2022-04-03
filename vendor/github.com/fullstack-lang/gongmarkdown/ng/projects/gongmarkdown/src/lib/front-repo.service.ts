import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { AnotherDummyDataDB } from './anotherdummydata-db'
import { AnotherDummyDataService } from './anotherdummydata.service'

import { CellDB } from './cell-db'
import { CellService } from './cell.service'

import { DummyDataDB } from './dummydata-db'
import { DummyDataService } from './dummydata.service'

import { ElementDB } from './element-db'
import { ElementService } from './element.service'

import { MarkdownContentDB } from './markdowncontent-db'
import { MarkdownContentService } from './markdowncontent.service'

import { RowDB } from './row-db'
import { RowService } from './row.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  AnotherDummyDatas_array = new Array<AnotherDummyDataDB>(); // array of repo instances
  AnotherDummyDatas = new Map<number, AnotherDummyDataDB>(); // map of repo instances
  AnotherDummyDatas_batch = new Map<number, AnotherDummyDataDB>(); // same but only in last GET (for finding repo instances to delete)
  Cells_array = new Array<CellDB>(); // array of repo instances
  Cells = new Map<number, CellDB>(); // map of repo instances
  Cells_batch = new Map<number, CellDB>(); // same but only in last GET (for finding repo instances to delete)
  DummyDatas_array = new Array<DummyDataDB>(); // array of repo instances
  DummyDatas = new Map<number, DummyDataDB>(); // map of repo instances
  DummyDatas_batch = new Map<number, DummyDataDB>(); // same but only in last GET (for finding repo instances to delete)
  Elements_array = new Array<ElementDB>(); // array of repo instances
  Elements = new Map<number, ElementDB>(); // map of repo instances
  Elements_batch = new Map<number, ElementDB>(); // same but only in last GET (for finding repo instances to delete)
  MarkdownContents_array = new Array<MarkdownContentDB>(); // array of repo instances
  MarkdownContents = new Map<number, MarkdownContentDB>(); // map of repo instances
  MarkdownContents_batch = new Map<number, MarkdownContentDB>(); // same but only in last GET (for finding repo instances to delete)
  Rows_array = new Array<RowDB>(); // array of repo instances
  Rows = new Map<number, RowDB>(); // map of repo instances
  Rows_batch = new Map<number, RowDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private anotherdummydataService: AnotherDummyDataService,
    private cellService: CellService,
    private dummydataService: DummyDataService,
    private elementService: ElementService,
    private markdowncontentService: MarkdownContentService,
    private rowService: RowService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<AnotherDummyDataDB[]>,
    Observable<CellDB[]>,
    Observable<DummyDataDB[]>,
    Observable<ElementDB[]>,
    Observable<MarkdownContentDB[]>,
    Observable<RowDB[]>,
  ] = [ // insertion point sub template 
      this.anotherdummydataService.getAnotherDummyDatas(),
      this.cellService.getCells(),
      this.dummydataService.getDummyDatas(),
      this.elementService.getElements(),
      this.markdowncontentService.getMarkdownContents(),
      this.rowService.getRows(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            anotherdummydatas_,
            cells_,
            dummydatas_,
            elements_,
            markdowncontents_,
            rows_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var anotherdummydatas: AnotherDummyDataDB[]
            anotherdummydatas = anotherdummydatas_ as AnotherDummyDataDB[]
            var cells: CellDB[]
            cells = cells_ as CellDB[]
            var dummydatas: DummyDataDB[]
            dummydatas = dummydatas_ as DummyDataDB[]
            var elements: ElementDB[]
            elements = elements_ as ElementDB[]
            var markdowncontents: MarkdownContentDB[]
            markdowncontents = markdowncontents_ as MarkdownContentDB[]
            var rows: RowDB[]
            rows = rows_ as RowDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.AnotherDummyDatas_array = anotherdummydatas

            // clear the map that counts AnotherDummyData in the GET
            FrontRepoSingloton.AnotherDummyDatas_batch.clear()

            anotherdummydatas.forEach(
              anotherdummydata => {
                FrontRepoSingloton.AnotherDummyDatas.set(anotherdummydata.ID, anotherdummydata)
                FrontRepoSingloton.AnotherDummyDatas_batch.set(anotherdummydata.ID, anotherdummydata)
              }
            )

            // clear anotherdummydatas that are absent from the batch
            FrontRepoSingloton.AnotherDummyDatas.forEach(
              anotherdummydata => {
                if (FrontRepoSingloton.AnotherDummyDatas_batch.get(anotherdummydata.ID) == undefined) {
                  FrontRepoSingloton.AnotherDummyDatas.delete(anotherdummydata.ID)
                }
              }
            )

            // sort AnotherDummyDatas_array array
            FrontRepoSingloton.AnotherDummyDatas_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Cells_array = cells

            // clear the map that counts Cell in the GET
            FrontRepoSingloton.Cells_batch.clear()

            cells.forEach(
              cell => {
                FrontRepoSingloton.Cells.set(cell.ID, cell)
                FrontRepoSingloton.Cells_batch.set(cell.ID, cell)
              }
            )

            // clear cells that are absent from the batch
            FrontRepoSingloton.Cells.forEach(
              cell => {
                if (FrontRepoSingloton.Cells_batch.get(cell.ID) == undefined) {
                  FrontRepoSingloton.Cells.delete(cell.ID)
                }
              }
            )

            // sort Cells_array array
            FrontRepoSingloton.Cells_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.DummyDatas_array = dummydatas

            // clear the map that counts DummyData in the GET
            FrontRepoSingloton.DummyDatas_batch.clear()

            dummydatas.forEach(
              dummydata => {
                FrontRepoSingloton.DummyDatas.set(dummydata.ID, dummydata)
                FrontRepoSingloton.DummyDatas_batch.set(dummydata.ID, dummydata)
              }
            )

            // clear dummydatas that are absent from the batch
            FrontRepoSingloton.DummyDatas.forEach(
              dummydata => {
                if (FrontRepoSingloton.DummyDatas_batch.get(dummydata.ID) == undefined) {
                  FrontRepoSingloton.DummyDatas.delete(dummydata.ID)
                }
              }
            )

            // sort DummyDatas_array array
            FrontRepoSingloton.DummyDatas_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Elements_array = elements

            // clear the map that counts Element in the GET
            FrontRepoSingloton.Elements_batch.clear()

            elements.forEach(
              element => {
                FrontRepoSingloton.Elements.set(element.ID, element)
                FrontRepoSingloton.Elements_batch.set(element.ID, element)
              }
            )

            // clear elements that are absent from the batch
            FrontRepoSingloton.Elements.forEach(
              element => {
                if (FrontRepoSingloton.Elements_batch.get(element.ID) == undefined) {
                  FrontRepoSingloton.Elements.delete(element.ID)
                }
              }
            )

            // sort Elements_array array
            FrontRepoSingloton.Elements_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.MarkdownContents_array = markdowncontents

            // clear the map that counts MarkdownContent in the GET
            FrontRepoSingloton.MarkdownContents_batch.clear()

            markdowncontents.forEach(
              markdowncontent => {
                FrontRepoSingloton.MarkdownContents.set(markdowncontent.ID, markdowncontent)
                FrontRepoSingloton.MarkdownContents_batch.set(markdowncontent.ID, markdowncontent)
              }
            )

            // clear markdowncontents that are absent from the batch
            FrontRepoSingloton.MarkdownContents.forEach(
              markdowncontent => {
                if (FrontRepoSingloton.MarkdownContents_batch.get(markdowncontent.ID) == undefined) {
                  FrontRepoSingloton.MarkdownContents.delete(markdowncontent.ID)
                }
              }
            )

            // sort MarkdownContents_array array
            FrontRepoSingloton.MarkdownContents_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Rows_array = rows

            // clear the map that counts Row in the GET
            FrontRepoSingloton.Rows_batch.clear()

            rows.forEach(
              row => {
                FrontRepoSingloton.Rows.set(row.ID, row)
                FrontRepoSingloton.Rows_batch.set(row.ID, row)
              }
            )

            // clear rows that are absent from the batch
            FrontRepoSingloton.Rows.forEach(
              row => {
                if (FrontRepoSingloton.Rows_batch.get(row.ID) == undefined) {
                  FrontRepoSingloton.Rows.delete(row.ID)
                }
              }
            )

            // sort Rows_array array
            FrontRepoSingloton.Rows_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            anotherdummydatas.forEach(
              anotherdummydata => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )
            cells.forEach(
              cell => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Row.Cells redeeming
                {
                  let _row = FrontRepoSingloton.Rows.get(cell.Row_CellsDBID.Int64)
                  if (_row) {
                    if (_row.Cells == undefined) {
                      _row.Cells = new Array<CellDB>()
                    }
                    _row.Cells.push(cell)
                    if (cell.Row_Cells_reverse == undefined) {
                      cell.Row_Cells_reverse = _row
                    }
                  }
                }
              }
            )
            dummydatas.forEach(
              dummydata => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field DummyPointerToGongStruct redeeming
                {
                  let _anotherdummydata = FrontRepoSingloton.AnotherDummyDatas.get(dummydata.DummyPointerToGongStructID.Int64)
                  if (_anotherdummydata) {
                    dummydata.DummyPointerToGongStruct = _anotherdummydata
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            elements.forEach(
              element => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Element.SubElements redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(element.Element_SubElementsDBID.Int64)
                  if (_element) {
                    if (_element.SubElements == undefined) {
                      _element.SubElements = new Array<ElementDB>()
                    }
                    _element.SubElements.push(element)
                    if (element.Element_SubElements_reverse == undefined) {
                      element.Element_SubElements_reverse = _element
                    }
                  }
                }
              }
            )
            markdowncontents.forEach(
              markdowncontent => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming
                // insertion point for pointer field Root redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(markdowncontent.RootID.Int64)
                  if (_element) {
                    markdowncontent.Root = _element
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )
            rows.forEach(
              row => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Element.Rows redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(row.Element_RowsDBID.Int64)
                  if (_element) {
                    if (_element.Rows == undefined) {
                      _element.Rows = new Array<RowDB>()
                    }
                    _element.Rows.push(row)
                    if (row.Element_Rows_reverse == undefined) {
                      row.Element_Rows_reverse = _element
                    }
                  }
                }
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // AnotherDummyDataPull performs a GET on AnotherDummyData of the stack and redeem association pointers 
  AnotherDummyDataPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.anotherdummydataService.getAnotherDummyDatas()
        ]).subscribe(
          ([ // insertion point sub template 
            anotherdummydatas,
          ]) => {
            // init the array
            FrontRepoSingloton.AnotherDummyDatas_array = anotherdummydatas

            // clear the map that counts AnotherDummyData in the GET
            FrontRepoSingloton.AnotherDummyDatas_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            anotherdummydatas.forEach(
              anotherdummydata => {
                FrontRepoSingloton.AnotherDummyDatas.set(anotherdummydata.ID, anotherdummydata)
                FrontRepoSingloton.AnotherDummyDatas_batch.set(anotherdummydata.ID, anotherdummydata)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear anotherdummydatas that are absent from the GET
            FrontRepoSingloton.AnotherDummyDatas.forEach(
              anotherdummydata => {
                if (FrontRepoSingloton.AnotherDummyDatas_batch.get(anotherdummydata.ID) == undefined) {
                  FrontRepoSingloton.AnotherDummyDatas.delete(anotherdummydata.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // CellPull performs a GET on Cell of the stack and redeem association pointers 
  CellPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.cellService.getCells()
        ]).subscribe(
          ([ // insertion point sub template 
            cells,
          ]) => {
            // init the array
            FrontRepoSingloton.Cells_array = cells

            // clear the map that counts Cell in the GET
            FrontRepoSingloton.Cells_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            cells.forEach(
              cell => {
                FrontRepoSingloton.Cells.set(cell.ID, cell)
                FrontRepoSingloton.Cells_batch.set(cell.ID, cell)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Row.Cells redeeming
                {
                  let _row = FrontRepoSingloton.Rows.get(cell.Row_CellsDBID.Int64)
                  if (_row) {
                    if (_row.Cells == undefined) {
                      _row.Cells = new Array<CellDB>()
                    }
                    _row.Cells.push(cell)
                    if (cell.Row_Cells_reverse == undefined) {
                      cell.Row_Cells_reverse = _row
                    }
                  }
                }
              }
            )

            // clear cells that are absent from the GET
            FrontRepoSingloton.Cells.forEach(
              cell => {
                if (FrontRepoSingloton.Cells_batch.get(cell.ID) == undefined) {
                  FrontRepoSingloton.Cells.delete(cell.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // DummyDataPull performs a GET on DummyData of the stack and redeem association pointers 
  DummyDataPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.dummydataService.getDummyDatas()
        ]).subscribe(
          ([ // insertion point sub template 
            dummydatas,
          ]) => {
            // init the array
            FrontRepoSingloton.DummyDatas_array = dummydatas

            // clear the map that counts DummyData in the GET
            FrontRepoSingloton.DummyDatas_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            dummydatas.forEach(
              dummydata => {
                FrontRepoSingloton.DummyDatas.set(dummydata.ID, dummydata)
                FrontRepoSingloton.DummyDatas_batch.set(dummydata.ID, dummydata)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field DummyPointerToGongStruct redeeming
                {
                  let _anotherdummydata = FrontRepoSingloton.AnotherDummyDatas.get(dummydata.DummyPointerToGongStructID.Int64)
                  if (_anotherdummydata) {
                    dummydata.DummyPointerToGongStruct = _anotherdummydata
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear dummydatas that are absent from the GET
            FrontRepoSingloton.DummyDatas.forEach(
              dummydata => {
                if (FrontRepoSingloton.DummyDatas_batch.get(dummydata.ID) == undefined) {
                  FrontRepoSingloton.DummyDatas.delete(dummydata.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // ElementPull performs a GET on Element of the stack and redeem association pointers 
  ElementPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.elementService.getElements()
        ]).subscribe(
          ([ // insertion point sub template 
            elements,
          ]) => {
            // init the array
            FrontRepoSingloton.Elements_array = elements

            // clear the map that counts Element in the GET
            FrontRepoSingloton.Elements_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            elements.forEach(
              element => {
                FrontRepoSingloton.Elements.set(element.ID, element)
                FrontRepoSingloton.Elements_batch.set(element.ID, element)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Element.SubElements redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(element.Element_SubElementsDBID.Int64)
                  if (_element) {
                    if (_element.SubElements == undefined) {
                      _element.SubElements = new Array<ElementDB>()
                    }
                    _element.SubElements.push(element)
                    if (element.Element_SubElements_reverse == undefined) {
                      element.Element_SubElements_reverse = _element
                    }
                  }
                }
              }
            )

            // clear elements that are absent from the GET
            FrontRepoSingloton.Elements.forEach(
              element => {
                if (FrontRepoSingloton.Elements_batch.get(element.ID) == undefined) {
                  FrontRepoSingloton.Elements.delete(element.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // MarkdownContentPull performs a GET on MarkdownContent of the stack and redeem association pointers 
  MarkdownContentPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.markdowncontentService.getMarkdownContents()
        ]).subscribe(
          ([ // insertion point sub template 
            markdowncontents,
          ]) => {
            // init the array
            FrontRepoSingloton.MarkdownContents_array = markdowncontents

            // clear the map that counts MarkdownContent in the GET
            FrontRepoSingloton.MarkdownContents_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            markdowncontents.forEach(
              markdowncontent => {
                FrontRepoSingloton.MarkdownContents.set(markdowncontent.ID, markdowncontent)
                FrontRepoSingloton.MarkdownContents_batch.set(markdowncontent.ID, markdowncontent)

                // insertion point for redeeming ONE/ZERO-ONE associations
                // insertion point for pointer field Root redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(markdowncontent.RootID.Int64)
                  if (_element) {
                    markdowncontent.Root = _element
                  }
                }

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear markdowncontents that are absent from the GET
            FrontRepoSingloton.MarkdownContents.forEach(
              markdowncontent => {
                if (FrontRepoSingloton.MarkdownContents_batch.get(markdowncontent.ID) == undefined) {
                  FrontRepoSingloton.MarkdownContents.delete(markdowncontent.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // RowPull performs a GET on Row of the stack and redeem association pointers 
  RowPull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.rowService.getRows()
        ]).subscribe(
          ([ // insertion point sub template 
            rows,
          ]) => {
            // init the array
            FrontRepoSingloton.Rows_array = rows

            // clear the map that counts Row in the GET
            FrontRepoSingloton.Rows_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            rows.forEach(
              row => {
                FrontRepoSingloton.Rows.set(row.ID, row)
                FrontRepoSingloton.Rows_batch.set(row.ID, row)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Element.Rows redeeming
                {
                  let _element = FrontRepoSingloton.Elements.get(row.Element_RowsDBID.Int64)
                  if (_element) {
                    if (_element.Rows == undefined) {
                      _element.Rows = new Array<RowDB>()
                    }
                    _element.Rows.push(row)
                    if (row.Element_Rows_reverse == undefined) {
                      row.Element_Rows_reverse = _element
                    }
                  }
                }
              }
            )

            // clear rows that are absent from the GET
            FrontRepoSingloton.Rows.forEach(
              row => {
                if (FrontRepoSingloton.Rows_batch.get(row.ID) == undefined) {
                  FrontRepoSingloton.Rows.delete(row.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getAnotherDummyDataUniqueID(id: number): number {
  return 31 * id
}
export function getCellUniqueID(id: number): number {
  return 37 * id
}
export function getDummyDataUniqueID(id: number): number {
  return 41 * id
}
export function getElementUniqueID(id: number): number {
  return 43 * id
}
export function getMarkdownContentUniqueID(id: number): number {
  return 47 * id
}
export function getRowUniqueID(id: number): number {
  return 53 * id
}

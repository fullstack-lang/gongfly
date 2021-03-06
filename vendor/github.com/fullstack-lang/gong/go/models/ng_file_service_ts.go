package models

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const NgServiceTmpl = `// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs';
import { Observable, of } from 'rxjs';
import { catchError, map, tap } from 'rxjs/operators';

import { {{Structname}}DB } from './{{structname}}-db';

// insertion point for imports{{` + string(rune(NgServiceTsInsertionImports)) + `}}

@Injectable({
  providedIn: 'root'
})
export class {{Structname}}Service {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  {{Structname}}ServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private {{structname}}sUrl: string

  constructor(
    private http: HttpClient,
    private location: Location,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.{{structname}}sUrl = origin + '/api/{{PkgPathRoot}}/v1/{{structname}}s';
  }

  /** GET {{structname}}s from the server */
  get{{Structname}}s(): Observable<{{Structname}}DB[]> {
    return this.http.get<{{Structname}}DB[]>(this.{{structname}}sUrl)
      .pipe(
        tap(_ => this.log('fetched {{structname}}s')),
        catchError(this.handleError<{{Structname}}DB[]>('get{{Structname}}s', []))
      );
  }

  /** GET {{structname}} by id. Will 404 if id not found */
  get{{Structname}}(id: number): Observable<{{Structname}}DB> {
    const url = ` + "`" + `${this.{{structname}}sUrl}/${id}` + "`" + `;
    return this.http.get<{{Structname}}DB>(url).pipe(
      tap(_ => this.log(` + "`" + `fetched {{structname}} id=${id}` + "`" + `)),
      catchError(this.handleError<{{Structname}}DB>(` + "`" + `get{{Structname}} id=${id}` + "`" + `))
    );
  }

  //////// Save methods //////////

  /** POST: add a new {{structname}} to the server */
  post{{Structname}}({{structname}}db: {{Structname}}DB): Observable<{{Structname}}DB> {

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON){{` + string(rune(NgServiceTsInsertionPointerReset)) + `}}

    return this.http.post<{{Structname}}DB>(this.{{structname}}sUrl, {{structname}}db, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers{{` + string(rune(NgServiceTsInsertionPointerRestore)) + `}}
        this.log(` + "`" + `posted {{structname}}db id=${{{structname}}db.ID}` + "`" + `)
      }),
      catchError(this.handleError<{{Structname}}DB>('post{{Structname}}'))
    );
  }

  /** DELETE: delete the {{structname}}db from the server */
  delete{{Structname}}({{structname}}db: {{Structname}}DB | number): Observable<{{Structname}}DB> {
    const id = typeof {{structname}}db === 'number' ? {{structname}}db : {{structname}}db.ID;
    const url = ` + "`" + `${this.{{structname}}sUrl}/${id}` + "`" + `;

    return this.http.delete<{{Structname}}DB>(url, this.httpOptions).pipe(
      tap(_ => this.log(` + "`" + `deleted {{structname}}db id=${id}` + "`" + `)),
      catchError(this.handleError<{{Structname}}DB>('delete{{Structname}}'))
    );
  }

  /** PUT: update the {{structname}}db on the server */
  update{{Structname}}({{structname}}db: {{Structname}}DB): Observable<{{Structname}}DB> {
    const id = typeof {{structname}}db === 'number' ? {{structname}}db : {{structname}}db.ID;
    const url = ` + "`" + `${this.{{structname}}sUrl}/${id}` + "`" + `;

    // insertion point for reset of pointers and reverse pointers (to avoid circular JSON){{` + string(rune(NgServiceTsInsertionPointerReset)) + `}}

    return this.http.put<{{Structname}}DB>(url, {{structname}}db, this.httpOptions).pipe(
      tap(_ => {
        // insertion point for restoration of reverse pointers{{` + string(rune(NgServiceTsInsertionPointerRestore)) + `}}
        this.log(` + "`" + `updated {{structname}}db id=${{{structname}}db.ID}` + "`" + `)
      }),
      catchError(this.handleError<{{Structname}}DB>('update{{Structname}}'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error(error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(` + "`" + `${operation} failed: ${error.message}` + "`" + `);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {

  }
}
`

// Insertion points
// insertion points in the main template
type NgServiceTsInsertionPoint int

const (
	NgServiceTsInsertionPointerReset NgServiceTsInsertionPoint = iota
	NgServiceTsInsertionPointerRestore

	NgServiceTsInsertionImports

	NgServiceTsInsertionsNb
)

type NgServiceSubTemplate int

const (
	NgServiceTSPointerToGongStructImports NgServiceSubTemplate = iota
	NgServiceTSPointerToGongStructReset

	NgServiceTSSliceOfPointerToGongStructReset
	NgServiceTSSliceOfPointerToGongStructReversePointerReset
	NgServiceTSSliceOfPointerToGongStructReversePointerRestore

	NgServiceTSReversePointerToSliceOfGongStructImports
)

var NgServiceSubTemplateCode map[NgServiceSubTemplate]string = map[NgServiceSubTemplate]string{

	NgServiceTSPointerToGongStructImports: `
import { {{AssocStructName}}DB } from './{{assocStructName}}-db'`,

	NgServiceTSPointerToGongStructReset: `
    {{structname}}db.{{FieldName}} = new {{AssocStructName}}DB`,

	NgServiceTSSliceOfPointerToGongStructReset: `
    {{structname}}db.{{FieldName}} = []`,

	NgServiceTSSliceOfPointerToGongStructReversePointerReset: `
    let _{{AssocStructName}}_{{FieldName}}_reverse = {{structname}}db.{{AssocStructName}}_{{FieldName}}_reverse
    {{structname}}db.{{AssocStructName}}_{{FieldName}}_reverse = new {{AssocStructName}}DB`,

	NgServiceTSSliceOfPointerToGongStructReversePointerRestore: `
        {{structname}}db.{{AssocStructName}}_{{FieldName}}_reverse = _{{AssocStructName}}_{{FieldName}}_reverse`,

	NgServiceTSReversePointerToSliceOfGongStructImports: `
import { {{AssocStructName}}DB } from './{{assocStructName}}-db'`,
}

// MultiCodeGeneratorNgService generates the code for the
// services
func MultiCodeGeneratorNgService(
	mdlPkg *ModelPkg,
	PkgName,
	MatTargetPath,
	PkgGoPath string,
	apiPath string) {

	// have alphabetical order generation
	structList := []*GongStruct{}
	for _, _struct := range mdlPkg.GongStructs {
		structList = append(structList, _struct)
	}
	sort.Slice(structList[:], func(i, j int) bool {
		return structList[i].Name < structList[j].Name
	})

	for _, _struct := range structList {

		if !_struct.HasNameField() {
			continue
		}

		// generate the typescript file
		codeTS := NgServiceTmpl

		codeTS = strings.ReplaceAll(codeTS, "{{addr}}", apiPath)

		TSinsertions := make(map[NgServiceTsInsertionPoint]string)
		for insertion := NgServiceTsInsertionPoint(0); insertion < NgServiceTsInsertionsNb; insertion++ {
			TSinsertions[insertion] = ""
		}

		for _, field := range _struct.Fields {
			switch field := field.(type) {
			case *PointerToGongStructField:

				TSinsertions[NgServiceTsInsertionPointerReset] +=
					Replace2(NgServiceSubTemplateCode[NgServiceTSPointerToGongStructReset],
						"{{FieldName}}", field.Name,
						"{{AssocStructName}}", field.GongStruct.Name)

				var importToInsert = Replace2(NgServiceSubTemplateCode[NgServiceTSPointerToGongStructImports],
					"{{AssocStructName}}", field.GongStruct.Name,
					"{{assocStructName}}", strings.ToLower(field.GongStruct.Name))

				// cannot insert twice the same import
				// or import twice the DB
				if !strings.Contains(TSinsertions[NgServiceTsInsertionImports], importToInsert) &&
					_struct.Name != field.GongStruct.Name {
					TSinsertions[NgServiceTsInsertionImports] += importToInsert
				}

			case *SliceOfPointerToGongStructField:

				TSinsertions[NgServiceTsInsertionPointerReset] +=
					Replace1(NgServiceSubTemplateCode[NgServiceTSSliceOfPointerToGongStructReset],
						"{{FieldName}}", field.Name)
			}
		}

		//
		// Parse all fields from other structs that points to this struct
		//
		for _, __struct := range structList {
			for _, field := range __struct.Fields {
				switch field := field.(type) {
				case *SliceOfPointerToGongStructField:

					if field.GongStruct == _struct {

						TSinsertions[NgServiceTsInsertionPointerReset] +=
							Replace2(NgServiceSubTemplateCode[NgServiceTSSliceOfPointerToGongStructReversePointerReset],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						TSinsertions[NgServiceTsInsertionPointerRestore] +=
							Replace2(NgServiceSubTemplateCode[NgServiceTSSliceOfPointerToGongStructReversePointerRestore],
								"{{FieldName}}", field.Name,
								"{{AssocStructName}}", __struct.Name)

						var importToInsert = Replace2(NgServiceSubTemplateCode[NgServiceTSReversePointerToSliceOfGongStructImports],
							"{{AssocStructName}}", __struct.Name,
							"{{assocStructName}}", strings.ToLower(__struct.Name))

						// cannot insert twice the same import
						if !strings.Contains(TSinsertions[NgServiceTsInsertionImports], importToInsert) &&
							__struct.Name != _struct.Name {
							TSinsertions[NgServiceTsInsertionImports] += importToInsert
						}
					}
				}
			}
		}

		for insertion := NgServiceTsInsertionPoint(0); insertion < NgServiceTsInsertionsNb; insertion++ {
			toReplace := "{{" + string(rune(insertion)) + "}}"
			codeTS = strings.ReplaceAll(codeTS, toReplace, TSinsertions[insertion])
		}
		// final replacement
		codeTS = Replace6(codeTS,
			"{{PkgName}}", PkgName,
			"{{TitlePkgName}}", strings.Title(PkgName),
			"{{pkgname}}", strings.ToLower(PkgName),
			"{{PkgPathRoot}}", strings.ReplaceAll(PkgGoPath, "/models", ""),
			"{{Structname}}", _struct.Name,
			"{{structname}}", strings.ToLower(_struct.Name))

		{
			file, err := os.Create(filepath.Join(MatTargetPath, strings.ToLower(_struct.Name)+".service.ts"))
			if err != nil {
				log.Panic(err)
			}
			defer file.Close()
			fmt.Fprint(file, codeTS)
		}
	}
}

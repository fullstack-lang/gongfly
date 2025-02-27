// generated by ng_file_service_ts
import { Injectable, Component, Inject } from '@angular/core';
import { HttpParams } from '@angular/common/http';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { DOCUMENT, Location } from '@angular/common'

/*
 * Behavior subject
 */
import { BehaviorSubject } from 'rxjs'
import { Observable, of } from 'rxjs'
import { catchError, map, tap } from 'rxjs/operators'

import { VLineAPI } from './vline-api'
import { VLine, CopyVLineToVLineAPI } from './vline'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'

@Injectable({
  providedIn: 'root'
})
export class VLineService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  VLineServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private vlinesUrl: string

  constructor(
    private http: HttpClient,
    @Inject(DOCUMENT) private document: Document
  ) {
    // path to the service share the same origin with the path to the document
    // get the origin in the URL to the document
    let origin = this.document.location.origin

    // if debugging with ng, replace 4200 with 8080
    origin = origin.replace("4200", "8080")

    // compute path to the service
    this.vlinesUrl = origin + '/api/github.com/fullstack-lang/gongleaflet/go/v1/vlines';
  }

  /** GET vlines from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI[]> {
    return this.getVLines(GONG__StackPath, frontRepo)
  }
  getVLines(GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<VLineAPI[]>(this.vlinesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<VLineAPI[]>('getVLines', []))
      );
  }

  /** GET vline by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {
    return this.getVLine(id, GONG__StackPath, frontRepo)
  }
  getVLine(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.vlinesUrl}/${id}`;
    return this.http.get<VLineAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched vline id=${id}`)),
      catchError(this.handleError<VLineAPI>(`getVLine id=${id}`))
    );
  }

  // postFront copy vline to a version with encoded pointers and post to the back
  postFront(vline: VLine, GONG__StackPath: string): Observable<VLineAPI> {
    let vlineAPI = new VLineAPI
    CopyVLineToVLineAPI(vline, vlineAPI)
    const id = typeof vlineAPI === 'number' ? vlineAPI : vlineAPI.ID
    const url = `${this.vlinesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<VLineAPI>(url, vlineAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<VLineAPI>('postVLine'))
    );
  }
  
  /** POST: add a new vline to the server */
  post(vlinedb: VLineAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {
    return this.postVLine(vlinedb, GONG__StackPath, frontRepo)
  }
  postVLine(vlinedb: VLineAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<VLineAPI>(this.vlinesUrl, vlinedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted vlinedb id=${vlinedb.ID}`)
      }),
      catchError(this.handleError<VLineAPI>('postVLine'))
    );
  }

  /** DELETE: delete the vlinedb from the server */
  delete(vlinedb: VLineAPI | number, GONG__StackPath: string): Observable<VLineAPI> {
    return this.deleteVLine(vlinedb, GONG__StackPath)
  }
  deleteVLine(vlinedb: VLineAPI | number, GONG__StackPath: string): Observable<VLineAPI> {
    const id = typeof vlinedb === 'number' ? vlinedb : vlinedb.ID;
    const url = `${this.vlinesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<VLineAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted vlinedb id=${id}`)),
      catchError(this.handleError<VLineAPI>('deleteVLine'))
    );
  }

  // updateFront copy vline to a version with encoded pointers and update to the back
  updateFront(vline: VLine, GONG__StackPath: string): Observable<VLineAPI> {
    let vlineAPI = new VLineAPI
    CopyVLineToVLineAPI(vline, vlineAPI)
    const id = typeof vlineAPI === 'number' ? vlineAPI : vlineAPI.ID
    const url = `${this.vlinesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<VLineAPI>(url, vlineAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<VLineAPI>('updateVLine'))
    );
  }

  /** PUT: update the vlinedb on the server */
  update(vlinedb: VLineAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {
    return this.updateVLine(vlinedb, GONG__StackPath, frontRepo)
  }
  updateVLine(vlinedb: VLineAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<VLineAPI> {
    const id = typeof vlinedb === 'number' ? vlinedb : vlinedb.ID;
    const url = `${this.vlinesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<VLineAPI>(url, vlinedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated vlinedb id=${vlinedb.ID}`)
      }),
      catchError(this.handleError<VLineAPI>('updateVLine'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in VLineService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("VLineService" + error); // log to console instead

      // TODO: better job of transforming error for user consumption
      this.log(`${operation} failed: ${error.message}`);

      // Let the app keep running by returning an empty result.
      return of(result as T);
    };
  }

  private log(message: string) {
    console.log(message)
  }
}

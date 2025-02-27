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

import { LayerGroupUseAPI } from './layergroupuse-api'
import { LayerGroupUse, CopyLayerGroupUseToLayerGroupUseAPI } from './layergroupuse'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports
import { LayerGroupAPI } from './layergroup-api'

@Injectable({
  providedIn: 'root'
})
export class LayerGroupUseService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  LayerGroupUseServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private layergroupusesUrl: string

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
    this.layergroupusesUrl = origin + '/api/github.com/fullstack-lang/gongleaflet/go/v1/layergroupuses';
  }

  /** GET layergroupuses from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI[]> {
    return this.getLayerGroupUses(GONG__StackPath, frontRepo)
  }
  getLayerGroupUses(GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<LayerGroupUseAPI[]>(this.layergroupusesUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<LayerGroupUseAPI[]>('getLayerGroupUses', []))
      );
  }

  /** GET layergroupuse by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {
    return this.getLayerGroupUse(id, GONG__StackPath, frontRepo)
  }
  getLayerGroupUse(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.layergroupusesUrl}/${id}`;
    return this.http.get<LayerGroupUseAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched layergroupuse id=${id}`)),
      catchError(this.handleError<LayerGroupUseAPI>(`getLayerGroupUse id=${id}`))
    );
  }

  // postFront copy layergroupuse to a version with encoded pointers and post to the back
  postFront(layergroupuse: LayerGroupUse, GONG__StackPath: string): Observable<LayerGroupUseAPI> {
    let layergroupuseAPI = new LayerGroupUseAPI
    CopyLayerGroupUseToLayerGroupUseAPI(layergroupuse, layergroupuseAPI)
    const id = typeof layergroupuseAPI === 'number' ? layergroupuseAPI : layergroupuseAPI.ID
    const url = `${this.layergroupusesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LayerGroupUseAPI>(url, layergroupuseAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LayerGroupUseAPI>('postLayerGroupUse'))
    );
  }
  
  /** POST: add a new layergroupuse to the server */
  post(layergroupusedb: LayerGroupUseAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {
    return this.postLayerGroupUse(layergroupusedb, GONG__StackPath, frontRepo)
  }
  postLayerGroupUse(layergroupusedb: LayerGroupUseAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<LayerGroupUseAPI>(this.layergroupusesUrl, layergroupusedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted layergroupusedb id=${layergroupusedb.ID}`)
      }),
      catchError(this.handleError<LayerGroupUseAPI>('postLayerGroupUse'))
    );
  }

  /** DELETE: delete the layergroupusedb from the server */
  delete(layergroupusedb: LayerGroupUseAPI | number, GONG__StackPath: string): Observable<LayerGroupUseAPI> {
    return this.deleteLayerGroupUse(layergroupusedb, GONG__StackPath)
  }
  deleteLayerGroupUse(layergroupusedb: LayerGroupUseAPI | number, GONG__StackPath: string): Observable<LayerGroupUseAPI> {
    const id = typeof layergroupusedb === 'number' ? layergroupusedb : layergroupusedb.ID;
    const url = `${this.layergroupusesUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<LayerGroupUseAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted layergroupusedb id=${id}`)),
      catchError(this.handleError<LayerGroupUseAPI>('deleteLayerGroupUse'))
    );
  }

  // updateFront copy layergroupuse to a version with encoded pointers and update to the back
  updateFront(layergroupuse: LayerGroupUse, GONG__StackPath: string): Observable<LayerGroupUseAPI> {
    let layergroupuseAPI = new LayerGroupUseAPI
    CopyLayerGroupUseToLayerGroupUseAPI(layergroupuse, layergroupuseAPI)
    const id = typeof layergroupuseAPI === 'number' ? layergroupuseAPI : layergroupuseAPI.ID
    const url = `${this.layergroupusesUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<LayerGroupUseAPI>(url, layergroupuseAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<LayerGroupUseAPI>('updateLayerGroupUse'))
    );
  }

  /** PUT: update the layergroupusedb on the server */
  update(layergroupusedb: LayerGroupUseAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {
    return this.updateLayerGroupUse(layergroupusedb, GONG__StackPath, frontRepo)
  }
  updateLayerGroupUse(layergroupusedb: LayerGroupUseAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<LayerGroupUseAPI> {
    const id = typeof layergroupusedb === 'number' ? layergroupusedb : layergroupusedb.ID;
    const url = `${this.layergroupusesUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<LayerGroupUseAPI>(url, layergroupusedb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated layergroupusedb id=${layergroupusedb.ID}`)
      }),
      catchError(this.handleError<LayerGroupUseAPI>('updateLayerGroupUse'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in LayerGroupUseService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("LayerGroupUseService" + error); // log to console instead

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

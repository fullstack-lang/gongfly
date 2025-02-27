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

import { DivIconAPI } from './divicon-api'
import { DivIcon, CopyDivIconToDivIconAPI } from './divicon'

import { FrontRepo, FrontRepoService } from './front-repo.service';

// insertion point for imports

@Injectable({
  providedIn: 'root'
})
export class DivIconService {

  // Kamar Raïmo: Adding a way to communicate between components that share information
  // so that they are notified of a change.
  DivIconServiceChanged: BehaviorSubject<string> = new BehaviorSubject("");

  private diviconsUrl: string

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
    this.diviconsUrl = origin + '/api/github.com/fullstack-lang/gongleaflet/go/v1/divicons';
  }

  /** GET divicons from the server */
  // gets is more robust to refactoring
  gets(GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI[]> {
    return this.getDivIcons(GONG__StackPath, frontRepo)
  }
  getDivIcons(GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI[]> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    return this.http.get<DivIconAPI[]>(this.diviconsUrl, { params: params })
      .pipe(
        tap(),
        catchError(this.handleError<DivIconAPI[]>('getDivIcons', []))
      );
  }

  /** GET divicon by id. Will 404 if id not found */
  // more robust API to refactoring
  get(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {
    return this.getDivIcon(id, GONG__StackPath, frontRepo)
  }
  getDivIcon(id: number, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)

    const url = `${this.diviconsUrl}/${id}`;
    return this.http.get<DivIconAPI>(url, { params: params }).pipe(
      // tap(_ => this.log(`fetched divicon id=${id}`)),
      catchError(this.handleError<DivIconAPI>(`getDivIcon id=${id}`))
    );
  }

  // postFront copy divicon to a version with encoded pointers and post to the back
  postFront(divicon: DivIcon, GONG__StackPath: string): Observable<DivIconAPI> {
    let diviconAPI = new DivIconAPI
    CopyDivIconToDivIconAPI(divicon, diviconAPI)
    const id = typeof diviconAPI === 'number' ? diviconAPI : diviconAPI.ID
    const url = `${this.diviconsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DivIconAPI>(url, diviconAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DivIconAPI>('postDivIcon'))
    );
  }
  
  /** POST: add a new divicon to the server */
  post(divicondb: DivIconAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {
    return this.postDivIcon(divicondb, GONG__StackPath, frontRepo)
  }
  postDivIcon(divicondb: DivIconAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.post<DivIconAPI>(this.diviconsUrl, divicondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`posted divicondb id=${divicondb.ID}`)
      }),
      catchError(this.handleError<DivIconAPI>('postDivIcon'))
    );
  }

  /** DELETE: delete the divicondb from the server */
  delete(divicondb: DivIconAPI | number, GONG__StackPath: string): Observable<DivIconAPI> {
    return this.deleteDivIcon(divicondb, GONG__StackPath)
  }
  deleteDivIcon(divicondb: DivIconAPI | number, GONG__StackPath: string): Observable<DivIconAPI> {
    const id = typeof divicondb === 'number' ? divicondb : divicondb.ID;
    const url = `${this.diviconsUrl}/${id}`;

    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.delete<DivIconAPI>(url, httpOptions).pipe(
      tap(_ => this.log(`deleted divicondb id=${id}`)),
      catchError(this.handleError<DivIconAPI>('deleteDivIcon'))
    );
  }

  // updateFront copy divicon to a version with encoded pointers and update to the back
  updateFront(divicon: DivIcon, GONG__StackPath: string): Observable<DivIconAPI> {
    let diviconAPI = new DivIconAPI
    CopyDivIconToDivIconAPI(divicon, diviconAPI)
    const id = typeof diviconAPI === 'number' ? diviconAPI : diviconAPI.ID
    const url = `${this.diviconsUrl}/${id}`;
    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    }

    return this.http.put<DivIconAPI>(url, diviconAPI, httpOptions).pipe(
      tap(_ => {
      }),
      catchError(this.handleError<DivIconAPI>('updateDivIcon'))
    );
  }

  /** PUT: update the divicondb on the server */
  update(divicondb: DivIconAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {
    return this.updateDivIcon(divicondb, GONG__StackPath, frontRepo)
  }
  updateDivIcon(divicondb: DivIconAPI, GONG__StackPath: string, frontRepo: FrontRepo): Observable<DivIconAPI> {
    const id = typeof divicondb === 'number' ? divicondb : divicondb.ID;
    const url = `${this.diviconsUrl}/${id}`;


    let params = new HttpParams().set("GONG__StackPath", GONG__StackPath)
    let httpOptions = {
      headers: new HttpHeaders({ 'Content-Type': 'application/json' }),
      params: params
    };

    return this.http.put<DivIconAPI>(url, divicondb, httpOptions).pipe(
      tap(_ => {
        // this.log(`updated divicondb id=${divicondb.ID}`)
      }),
      catchError(this.handleError<DivIconAPI>('updateDivIcon'))
    );
  }

  /**
   * Handle Http operation that failed.
   * Let the app continue.
   * @param operation - name of the operation that failed
   * @param result - optional value to return as the observable result
   */
  private handleError<T>(operation = 'operation in DivIconService', result?: T) {
    return (error: any): Observable<T> => {

      // TODO: send the error to remote logging infrastructure
      console.error("DivIconService" + error); // log to console instead

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

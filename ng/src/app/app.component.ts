import { Component } from '@angular/core';

import { Router, RouterState } from '@angular/router';

import * as gongfly from 'gongfly'
import { combineLatest, Observable, timer } from 'rxjs'

import * as gongleaflet from 'gongleaflet'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'ng';

  view = 'Carto view'
  carto = 'Carto view'
  data = 'Data view'
  leaflet = 'Leaflet view'
  sim = 'Sim view'
  diagrams = 'Diagrams view'
  meta_diagrams = 'Meta diagrams view'
  meta = 'Meta view'
  markdown = 'Markdonw view'
  markdown_text = 'Markdonw text view'
  markdown_data = 'Markdown data view'
  views: string[] = [this.carto, this.data, this.leaflet, this.sim, this.diagrams,
  this.meta_diagrams, this.meta, this.markdown, this.markdown_text, this.markdown_data];

  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period
  currTime: number = 0

  constructor(
  ) {
  }

  ngOnInit(): void {
  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (updateDisplay: boolean): void => {
  }
}

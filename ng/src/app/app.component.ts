import { Component } from '@angular/core';

import { Router, RouterState } from '@angular/router';
import * as gongdoc from 'gongdoc'

import * as gongfly from 'gongfly'
import { combineLatest, Observable, timer } from 'rxjs'

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

  diagrams = 'UML view'
  meta_diagrams = 'Meta diagrams view'

  meta = 'Meta view'

  markdown = 'Markdown view'
  markdown_text = 'Markdonw text view'
  markdown_data = 'Markdown data view'

  charts = 'Charts view'
  charts_data = 'Charts data view'

  views: string[] = [this.carto, this.data, this.leaflet, this.sim, this.diagrams, this.meta_diagrams, this.meta,

  this.markdown, this.markdown_text, this.markdown_data, this.charts, this.charts_data];

  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period
  currTime: number = 0
  ''

  constructor(
    private gongstructSelectionService: gongfly.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (updateDisplay: boolean): void => {
  }
}

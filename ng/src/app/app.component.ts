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
  views: string[] = [this.carto, this.data, this.leaflet, this.sim, this.diagrams];

  obsTimer: Observable<number> = timer(1000, 500) // due time 1', period
  currTime: number = 0

  constructor(
    private MessageService: gongfly.MessageService,
    private LinerService: gongfly.LinerService,

    private VisualLineService: gongleaflet.VisualLineService,
    private VisualTrackService: gongleaflet.VisualTrackService,

    private router: Router,
  ) {
  }

  ngOnInit(): void {
  }

  // callbak function that is attached to the generic engine
  engineUpdatedCallbackFunction = (updateDisplay: boolean): void => {

    if (updateDisplay) {

      this.MessageService.MessageServiceChanged.next("update")
      this.LinerService.LinerServiceChanged.next("update")

      this.VisualLineService.VisualLineServiceChanged.next("update")
      this.VisualTrackService.VisualTrackServiceChanged.next("update")

    }
  }
}

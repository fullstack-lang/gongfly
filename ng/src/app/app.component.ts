import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongfly from 'gongfly'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  default = 'Gongfly Data/Model'
  dashboard = "Dashboard"
  simulation = 'Simulation Data/Model'
  view = this.dashboard

  views: string[] = [this.dashboard, this.default, this.simulation];

  GONG__StackPath = "github.com/fullstack-lang/gongfly/go/models"

  GONG__SIM__StackPath = "github.com/fullstack-lang/gongsim/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}

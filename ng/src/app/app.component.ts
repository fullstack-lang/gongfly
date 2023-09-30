import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongfly from 'gongfly'

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongtreeModule } from 'gongtree'
import { GongtreespecificModule } from 'gongtreespecific'

import { GongtableModule } from 'gongtable'
import { GongtablespecificModule } from 'gongtablespecific'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {


  dashboard = "Dashboard"
  gongflyProbe = 'Gongfly Probe'
  gongsimProbe = 'Simulation Probe'
  gongleafletProbe = 'Leaflet Probe'
  view = this.dashboard

  views: string[] = [this.dashboard, this.gongflyProbe, this.gongsimProbe, this.gongleafletProbe];


  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongfly"
  GongflyStackType = "github.com/fullstack-lang/gongfly/go/models"
  GongsimStackType = "github.com/fullstack-lang/gongsim/go/models"
  GongleafletStackType = "github.com/fullstack-lang/gongleaflet/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}

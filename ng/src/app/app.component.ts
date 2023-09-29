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

  default = 'Gongfly Data/Model'
  dashboard = "Dashboard"
  simulation = 'Simulation Data/Model'
  leaflet = 'Leaflet Data/Model'
  view = this.dashboard

  views: string[] = [this.dashboard, this.default, this.simulation, this.leaflet];


  scrollStyle = {
    'overflow- x': 'auto',
    'width': '100%',  // Ensure the div takes the full width of its parent container
  }

  StackName = "gongfly"
  StackType = "github.com/fullstack-lang/gongfly/go/models"

  constructor(
  ) {

  }

  ngOnInit(): void {
  }
}

import { Component, OnInit } from '@angular/core';

import { Observable, combineLatest, timer } from 'rxjs'

// for angular & angular material
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { MatRadioModule } from '@angular/material/radio';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';

import { AngularSplitModule } from 'angular-split';

import * as gongfly from '../../projects/gongfly/src/public-api'

import { TreeComponent } from '@vendored_components/github.com/fullstack-lang/gongtree/ng-github.com-fullstack-lang-gongtree/projects/gongtreespecific/src/public-api'
import { MaterialTableComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-table/material-table.component';
import { MaterialFormComponent } from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtablespecific/src/lib/material-form/material-form.component';
import * as gongtable from '@vendored_components/github.com/fullstack-lang/gongtable/ng-github.com-fullstack-lang-gongtable/projects/gongtable/src/public-api';
import { PanelComponent } from '@vendored_components/github.com/fullstack-lang/gongdoc/ng-github.com-fullstack-lang-gongdoc/projects/gongdocspecific/src/public-api'
import { GongflyDashboardComponent } from '../../projects/gongflyspecific/src/lib/gongfly-dashboard/gongfly-dashboard.component';

import { LeafletModule } from '@bluehalo/ngx-leaflet';
import { MatOptionModule } from '@angular/material/core'

@Component({
    selector: 'app-root',
    imports: [
        CommonModule,
        FormsModule,
        MatRadioModule,
        MatButtonModule,
        MatIconModule,
        AngularSplitModule,
        TreeComponent,
        MaterialTableComponent,
        MaterialFormComponent,
        PanelComponent,
        GongflyDashboardComponent,
        LeafletModule,
        MatOptionModule,
    ],
    templateUrl: './app.component.html'
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

  StackNames = gongfly.StacksNames

  constructor(
  ) {

  }

  ngOnInit(): void {
    console.log("")
  }
}

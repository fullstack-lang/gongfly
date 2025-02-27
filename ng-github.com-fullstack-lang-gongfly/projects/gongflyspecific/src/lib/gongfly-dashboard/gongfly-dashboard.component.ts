import { Component, Input, OnInit } from '@angular/core';

import { AngularSplitModule } from 'angular-split';


import { LeafletModule } from '@bluehalo/ngx-leaflet';
import { MatOptionModule } from '@angular/material/core'

import { EngineControlComponent } from '@vendored_components/github.com/fullstack-lang/gongsim/ng-github.com-fullstack-lang-gongsim/projects/gongsimspecific/src/public-api'
import { MapoptionsComponent } from '@vendored_components/github.com/fullstack-lang/gongleaflet/ng-github.com-fullstack-lang-gongleaflet/projects/gongleafletspecific/src/public-api'

@Component({
    selector: 'lib-gongfly-dashboard',
    templateUrl: './gongfly-dashboard.component.html',
    styleUrls: ['./gongfly-dashboard.component.css'],
    imports: [
        LeafletModule,
        MatOptionModule,
        AngularSplitModule,
        EngineControlComponent,
        MapoptionsComponent,
    ]
})
export class GongflyDashboardComponent implements OnInit {

  @Input() GONG__SIM__StackPath: string = ""
  @Input() GONG__LEAFLET__StackPath: string = ""

  loading: boolean = true

  ngOnInit(): void {
    this.loading = false
  }

}

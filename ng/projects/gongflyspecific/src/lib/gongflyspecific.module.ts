import { NgModule } from '@angular/core';

import { GongflyspecificComponent } from './gongflyspecific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongflyModule } from 'gongfly'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';
import { GongflyDashboardComponent } from './gongfly-dashboard/gongfly-dashboard.component';

// reusable simulation library from gong
import { GongsimspecificModule } from 'gongsimspecific'
import { GongsimModule } from 'gongsim'

// Leaflet
// import { GongleafletModule } from 'gongleaflet'
// import { GongleafletspecificModule } from 'gongleafletspecific'
// import { LeafletModule } from '@asymmetrik/ngx-leaflet';


@NgModule({
  declarations: [
    GongflyspecificComponent,
    DataModelPanelComponent,
    GongflyDashboardComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongflyModule,

    // gongsim stack
    GongsimspecificModule,
    GongsimModule,

    // // gongleaflet stack
    // GongleafletModule,
    // GongleafletspecificModule,
    // LeafletModule
    

  ],
  exports: [
    GongflyspecificComponent,
    DataModelPanelComponent,
    GongflyDashboardComponent
  ]
})
export class GongflyspecificModule { }

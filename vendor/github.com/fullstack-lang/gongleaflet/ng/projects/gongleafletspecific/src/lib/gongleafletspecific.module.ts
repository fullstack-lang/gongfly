import { NgModule } from '@angular/core';

import { GongleafletspecificComponent } from './gongleafletspecific.component';
import { DataModelPanelComponent } from './data-model-panel/data-model-panel.component';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongleafletModule } from 'gongleaflet'

import { MatRadioModule } from '@angular/material/radio';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { MapoptionsComponent } from './mapoptions/mapoptions.component';
// Leaflet
import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { CartoatcControlSettingsComponent } from './mapoptions/cartoatc-control-settings/cartoatc-control-settings.component';
// Material UI
import { MatButtonModule } from '@angular/material/button';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [
    GongleafletspecificComponent,
    DataModelPanelComponent,

    MapoptionsComponent,
    CartoatcControlSettingsComponent,
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    BrowserModule,
    BrowserAnimationsModule,
    CommonModule,
    LeafletModule,
    MatButtonModule,
    MatRadioModule,
    MatSlideToggleModule,
    MatIconModule,
    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongleafletModule,
  ],
  exports: [
    GongleafletspecificComponent,
    DataModelPanelComponent,
    MapoptionsComponent,
    CartoatcControlSettingsComponent,
  ]
})
export class GongleafletspecificModule { }

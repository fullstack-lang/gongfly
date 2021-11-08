import { NgModule } from '@angular/core';

import { BrowserModule } from '@angular/platform-browser';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';

import { CommonModule } from '@angular/common';

import { GongleafletspecificComponent } from './gongleafletspecific.component';

import { CartoatcComponent } from './cartoatc/cartoatc.component';
// Leaflet
import { LeafletModule } from '@asymmetrik/ngx-leaflet';
import { CartoatcCentersComponent } from './cartoatc/cartoatc-centers/cartoatc-centers.component';
import { CartoatcLinesComponent } from './cartoatc/cartoatc-lines/cartoatc-lines.component';
import { CartoatcControlSettingsComponent } from './cartoatc/cartoatc-control-settings/cartoatc-control-settings.component';
import { CartoatcCirclesComponent } from './cartoatc/cartoatc-circles/cartoatc-circles.component';
// Material UI
import { MatButtonModule } from '@angular/material/button';
import { MatRadioModule } from '@angular/material/radio';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';
import { MatIconModule } from '@angular/material/icon';

@NgModule({
  declarations: [
    GongleafletspecificComponent,

    CartoatcComponent,
    CartoatcCentersComponent,
    CartoatcLinesComponent,
    CartoatcControlSettingsComponent,
    CartoatcCirclesComponent,
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,
    CommonModule,
    LeafletModule,
    MatButtonModule,
    MatRadioModule,
    MatSlideToggleModule,
    MatIconModule,
  ],
  exports: [
    GongleafletspecificComponent,

    CartoatcComponent,
    CartoatcCentersComponent,
    CartoatcLinesComponent,
    CartoatcControlSettingsComponent,
    CartoatcCirclesComponent,
  ],
})
export class GongleafletspecificModule {}

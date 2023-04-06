import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { AppComponent } from './app.component';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';


// for angular material
import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatCardModule } from '@angular/material/card'
import { MatTooltipModule } from '@angular/material/tooltip';
import { MatRadioModule } from '@angular/material/radio';
import { MatSlideToggleModule } from '@angular/material/slide-toggle';

import { FormsModule } from '@angular/forms';

// to split the screen
import { AngularSplitModule } from 'angular-split';

import { GongdocModule } from 'gongdoc'
import { GongdocdiagramsModule } from 'gongdocdiagrams'

import { GongModule } from 'gong'

import { GongflyModule } from 'gongfly'
import { GongflyspecificModule } from 'gongflyspecific'
import { GongstructSelectionService } from 'gongfly'

// // Leaflet
import { GongleafletModule } from 'gongleaflet'
import { GongleafletspecificModule } from 'gongleafletspecific'
// import { LeafletModule } from '@asymmetrik/ngx-leaflet';

// reusable simulation library from gong
import { GongsimspecificModule } from 'gongsimspecific'
import { GongsimModule } from 'gongsim'

// mandatory
import { HttpClientModule } from '@angular/common/http';

@NgModule({
  declarations: [
    AppComponent
  ],
  imports: [
    BrowserModule,
    BrowserAnimationsModule,

    HttpClientModule,

    MatSliderModule,
    MatSelectModule,
    MatFormFieldModule,
    MatInputModule,
    MatDatepickerModule,
    MatTableModule,
    MatCheckboxModule,
    MatButtonModule,
    MatIconModule,
    MatToolbarModule,
    MatCardModule,
    MatTooltipModule,
    MatRadioModule,
    MatSlideToggleModule,
    FormsModule,

    AngularSplitModule,

    // gong stack (for analysis of gong code in the current stack)
    GongModule,

    // gongdoc stack (for displaying UML diagrams of the gong code in the current stack)
    GongdocModule,
    GongdocdiagramsModule,

    GongflyModule,
    GongflyspecificModule,

    // gongsim stack
    GongsimspecificModule,
    GongsimModule,

    // // gongleaflet stack
    GongleafletModule,
    GongleafletspecificModule,
    // LeafletModule

  ],
  providers: [
    GongstructSelectionService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }

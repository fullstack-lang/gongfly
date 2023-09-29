import { NgModule } from '@angular/core';

import { GongsimspecificComponent } from './gongsimspecific.component';

import { GongdocModule } from 'gongdoc'
import { GongdocspecificModule } from 'gongdocspecific'

import { GongsimModule } from 'gongsim'

import { MatRadioModule } from '@angular/material/radio';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatIconModule } from '@angular/material/icon';
import { MatButtonModule } from '@angular/material/button';
import { FormsModule } from '@angular/forms';
import { CommonModule } from '@angular/common';
import { AngularSplitModule } from 'angular-split';
import { EngineControlComponent } from './engine-control/engine-control.component';

@NgModule({
  declarations: [
    GongsimspecificComponent,
    EngineControlComponent
  ],
  imports: [
    GongdocModule,
    GongdocspecificModule,

    MatToolbarModule,
    MatIconModule,
    MatButtonModule,
    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongsimModule,
  ],
  exports: [
    GongsimspecificComponent,
    EngineControlComponent
  ]
})
export class GongsimspecificModule { }

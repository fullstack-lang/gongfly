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

@NgModule({
  declarations: [
    GongflyspecificComponent,
    DataModelPanelComponent
  ],
  imports: [
    GongdocModule,
    GongdocdiagramsModule,

    MatRadioModule,
    FormsModule,
    CommonModule,

    AngularSplitModule,

    GongflyModule,
  ],
  exports: [
    GongflyspecificComponent,
    DataModelPanelComponent
  ]
})
export class GongflyspecificModule { }

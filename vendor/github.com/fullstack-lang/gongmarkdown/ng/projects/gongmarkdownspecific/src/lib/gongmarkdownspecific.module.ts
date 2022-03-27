import { NgModule } from '@angular/core';
import { GongmarkdownspecificComponent } from './gongmarkdownspecific.component';
import { GongmarkdownspecificMarkdownComponent } from './gongmarkdownspecific-markdown/gongmarkdownspecific-markdown.component';

import { BrowserModule } from '@angular/platform-browser';
import { FormsModule } from '@angular/forms';

import { MarkdownModule, MarkedOptions } from 'ngx-markdown';
import { GongmarkdownspecificTextareaComponent } from './gongmarkdownspecific-textarea/gongmarkdownspecific-textarea.component';

@NgModule({
  declarations: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent,
    GongmarkdownspecificTextareaComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    MarkdownModule.forRoot()
  ],
  exports: [
    GongmarkdownspecificComponent,
    GongmarkdownspecificMarkdownComponent,
    GongmarkdownspecificTextareaComponent
  ]
})
export class GongmarkdownspecificModule { }

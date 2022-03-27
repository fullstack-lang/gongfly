import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { CellsTableComponent } from './cells-table/cells-table.component'
import { CellDetailComponent } from './cell-detail/cell-detail.component'
import { CellPresentationComponent } from './cell-presentation/cell-presentation.component'

import { ElementsTableComponent } from './elements-table/elements-table.component'
import { ElementDetailComponent } from './element-detail/element-detail.component'
import { ElementPresentationComponent } from './element-presentation/element-presentation.component'

import { MarkdownContentsTableComponent } from './markdowncontents-table/markdowncontents-table.component'
import { MarkdownContentDetailComponent } from './markdowncontent-detail/markdowncontent-detail.component'
import { MarkdownContentPresentationComponent } from './markdowncontent-presentation/markdowncontent-presentation.component'

import { RowsTableComponent } from './rows-table/rows-table.component'
import { RowDetailComponent } from './row-detail/row-detail.component'
import { RowPresentationComponent } from './row-presentation/row-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cells', component: CellsTableComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_table' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cell-adder', component: CellDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cell-adder/:id/:originStruct/:originStructFieldName', component: CellDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cell-detail/:id', component: CellDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cell-presentation/:id', component: CellPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-cell-presentation-special/:id', component: CellPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_gocellpres' },

	{ path: 'github_com_fullstack_lang_gongmarkdown_go-elements', component: ElementsTableComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_table' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-element-adder', component: ElementDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-element-adder/:id/:originStruct/:originStructFieldName', component: ElementDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-element-detail/:id', component: ElementDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-element-presentation/:id', component: ElementPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-element-presentation-special/:id', component: ElementPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_goelementpres' },

	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontents', component: MarkdownContentsTableComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_table' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontent-adder', component: MarkdownContentDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontent-adder/:id/:originStruct/:originStructFieldName', component: MarkdownContentDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontent-detail/:id', component: MarkdownContentDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontent-presentation/:id', component: MarkdownContentPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-markdowncontent-presentation-special/:id', component: MarkdownContentPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_gomarkdowncontentpres' },

	{ path: 'github_com_fullstack_lang_gongmarkdown_go-rows', component: RowsTableComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_table' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-row-adder', component: RowDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-row-adder/:id/:originStruct/:originStructFieldName', component: RowDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-row-detail/:id', component: RowDetailComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-row-presentation/:id', component: RowPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmarkdown_go-row-presentation-special/:id', component: RowPresentationComponent, outlet: 'github_com_fullstack_lang_gongmarkdown_gorowpres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }

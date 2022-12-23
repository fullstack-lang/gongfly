import { Component, OnInit } from '@angular/core';

import { Observable, timer } from 'rxjs';

import * as gongmarkdown from 'gongmarkdown'

@Component({
  selector: 'lib-gongmarkdownspecific-textarea',
  templateUrl: './gongmarkdownspecific-textarea.component.html',
  styleUrls: ['./gongmarkdownspecific-textarea.component.css']
})
export class GongmarkdownspecificTextareaComponent implements OnInit {

  markdownContentDB: gongmarkdown.MarkdownContentDB | undefined

  /**
 * the component is refreshed when modification are performed in the back repo 
 * 
 * the checkCommitNbTimer polls the commit number of the back repo
 * if the commit number has increased, it pulls the front repo and redraw the diagram
 */
  checkCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  constructor(
    private gongmarkdownCommitNbService: gongmarkdown.CommitNbFromBackService,
    private gongmarkdownPushFromFrontNbService: gongmarkdown.PushFromFrontNbService,
    private gongmarkdownMarkdownContentService: gongmarkdown.MarkdownContentService
  ) { }

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongmarkdownCommitNbService.getCommitNbFromBack().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )

        // see above for the explanation
        this.gongmarkdownPushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              this.refresh()
              this.lastPushFromFrontNb = pushFromFrontNb
            }
          }
        )
      }
    )


  }

  refresh(): void {
    // get the singloton
    this.gongmarkdownMarkdownContentService.getMarkdownContents().subscribe(
      markdownContentDBs => {
        if (markdownContentDBs.length == 1) {
          this.markdownContentDB = markdownContentDBs[0]
        } else {
          console.log("wrong number of markdown content " + markdownContentDBs.length)
        }
      }
    )
  }

}

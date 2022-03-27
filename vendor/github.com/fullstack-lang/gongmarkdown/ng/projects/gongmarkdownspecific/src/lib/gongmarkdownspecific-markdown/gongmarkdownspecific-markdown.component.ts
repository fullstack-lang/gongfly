import { Component, OnInit } from '@angular/core';

import * as gongmarkdown from 'gongmarkdown'
import { Observable, timer } from 'rxjs';

@Component({
  selector: 'lib-gongmarkdownspecific-markdown',
  templateUrl: './gongmarkdownspecific-markdown.component.html',
  styleUrls: ['./gongmarkdownspecific-markdown.component.css']
})
export class GongmarkdownspecificMarkdownComponent implements OnInit {

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
    private gongmarkdownCommitNbService: gongmarkdown.CommitNbService,
    private gongmarkdownPushFromFrontNbService: gongmarkdown.PushFromFrontNbService,
    private gongmarkdownMarkdownContentService: gongmarkdown.MarkdownContentService
  ) { }

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongmarkdownCommitNbService.getCommitNb().subscribe(
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

// generated by gong
import { Component, OnInit, Inject, Optional } from '@angular/core';
import { TypeofExpr } from '@angular/compiler';
import { CdkDragDrop, moveItemInArray } from '@angular/cdk/drag-drop';

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData } from '../front-repo.service'
import { SelectionModel } from '@angular/cdk/collections';

import { Router, RouterState } from '@angular/router';
import { NoteDB } from '../note-db'
import { NoteService } from '../note.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { NullInt64 } from '../null-int64'

@Component({
  selector: 'lib-note-sorting',
  templateUrl: './note-sorting.component.html',
  styleUrls: ['./note-sorting.component.css']
})
export class NoteSortingComponent implements OnInit {

  frontRepo: FrontRepo = new (FrontRepo)

  // array of Note instances that are in the association
  associatedNotes = new Array<NoteDB>();

  constructor(
    private noteService: NoteService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of note instances
    public dialogRef: MatDialogRef<NoteSortingComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {
    this.router.routeReuseStrategy.shouldReuseRoute = function () {
      return false;
    };
  }

  ngOnInit(): void {
    this.getNotes()
  }

  getNotes(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        let index = 0
        for (let note of this.frontRepo.Notes_array) {
          let ID = this.dialogData.ID
          let revPointerID = note[this.dialogData.ReversePointer as keyof NoteDB] as unknown as NullInt64
          let revPointerID_Index = note[this.dialogData.ReversePointer + "_Index" as keyof NoteDB] as unknown as NullInt64
          if (revPointerID.Int64 == ID) {
            if (revPointerID_Index == undefined) {
              revPointerID_Index = new NullInt64
              revPointerID_Index.Valid = true
              revPointerID_Index.Int64 = index++
            }
            this.associatedNotes.push(note)
          }
        }

        // sort associated note according to order
        this.associatedNotes.sort((t1, t2) => {
          let t1_revPointerID_Index = t1[this.dialogData.ReversePointer + "_Index" as keyof typeof t1] as unknown as NullInt64
          let t2_revPointerID_Index = t2[this.dialogData.ReversePointer + "_Index" as keyof typeof t2] as unknown as NullInt64
          if (t1_revPointerID_Index && t2_revPointerID_Index) {
            if (t1_revPointerID_Index.Int64 > t2_revPointerID_Index.Int64) {
              return 1;
            }
            if (t1_revPointerID_Index.Int64 < t2_revPointerID_Index.Int64) {
              return -1;
            }
          }
          return 0;
        });
      }
    )
  }

  drop(event: CdkDragDrop<string[]>) {
    moveItemInArray(this.associatedNotes, event.previousIndex, event.currentIndex);

    // set the order of Note instances
    let index = 0

    for (let note of this.associatedNotes) {
      let revPointerID_Index = note[this.dialogData.ReversePointer + "_Index" as keyof NoteDB] as unknown as NullInt64
      revPointerID_Index.Valid = true
      revPointerID_Index.Int64 = index++
    }
  }

  save() {

    this.associatedNotes.forEach(
      note => {
        this.noteService.updateNote(note)
          .subscribe(note => {
            this.noteService.NoteServiceChanged.next("update")
          });
      }
    )

    this.dialogRef.close('Sorting of ' + this.dialogData.ReversePointer +' done');
  }
}

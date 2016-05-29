import { Component, OnInit } from '@angular/core';
import { Note } from './note'
import { NoteDetailComponent } from './note-detail.component'
import { NoteService } from './note.service'

@Component({
  selector: 'note-list',
  template: require('./note-list.component.html'),
  directives: [NoteDetailComponent],
  providers: [NoteService]
})

export class NoteListComponent implements OnInit{
  title = 'Notes';
  notes: Note[];
  selectedNote: Note;

  constructor(private noteService: NoteService) { }
  getNotes() {
    this.noteService.getNotes().then(notes => this.notes = notes);
  }
  ngOnInit() {
    this.getNotes();
  }

  onSelect(note: Note){ this.selectedNote = note; }
}

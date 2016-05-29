import { Injectable } from '@angular/core';
import { Note } from './note';
import { NOTES } from './mock-notes';

@Injectable()
export class NoteService {
  getNotes() {
    return Promise.resolve(NOTES);
  }
  // See the "Take it slow" appendix
  getNotesSlowly() {
    return new Promise<Note[]>(resolve =>
      setTimeout(()=>resolve(NOTES), 2000) // 2 seconds
    );
  }
}

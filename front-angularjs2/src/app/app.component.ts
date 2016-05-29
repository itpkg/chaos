import { Component } from '@angular/core';
import '../../public/css/styles.css';

import { NoteListComponent } from './note-list.component'

@Component({
    selector: 'root-app',
    template: require('./app.component.html'),
    directives: [NoteListComponent]
})

export class AppComponent {

}


import {
  NOTE_LIST, NOTE_ADD, NOTE_DEL, NOTE_CHG
} from './actions'


function readingNotes(state=[], action){
  switch (action.type) {
    case NOTE_LIST:
      return action.notes
    case NOTE_ADD:
      state.unshift(action.note)
      return state.slice(0)
    case NOTE_CHG:
      var i=0;
      for(; i<state.length; i++){
        if(action.note.id == state[i].id){
          break
        }
      }
      state[i] = action.note
      return state.slice(0)
    case NOTE_DEL:
      for(var i=0; i<state.length; i++){
        if(state[i].id===action.id){
          state.splice(i, 1);
          break
        }
      }
      return state.slice(0);
    default:
      return state

  }
}


const reducers = {readingNotes}
export default reducers

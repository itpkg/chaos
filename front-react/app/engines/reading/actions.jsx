
export const NOTE_LIST = "reading.notices.list"
export const NOTE_ADD = "reading.notices.add"
export const NOTE_DEL = "reading.notices.delete"
export const NOTE_CHG = "reading.notices.update"


export function listNote(notes){
  return {type:NOTE_LIST, notes}
}

export function addNote(note){
  return {type:NOTE_ADD, note}
}

export function chgNote(note){
  return {type:NOTE_CHG, note}
}

export function delNote(id){
  return {type:NOTE_DEL, id}
}

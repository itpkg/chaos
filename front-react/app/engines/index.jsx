import { combineReducers } from 'redux'

import cms from './cms'
import hr from './hr'
import ops from './ops'
import platform from './platform'
import reading from './reading'
import team from './team'

const engines = [
  // cms:cms,
  // hr:hr,
  // ops:ops,
  // reading:reading,
  // team:team,
  platform
]

export default {
  routes(){
    return engines.reduce(function(obj, en) {
      return obj.concat(en.routes)
    }, []);    
  },
  reducers(){
    return {}
  }
}

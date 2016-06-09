import { combineReducers } from 'redux'

import cms from './cms'
import hr from './hr'
import ops from './ops'
import platform from './platform'
import reading from './reading'
import team from './team'

const engines = [
  cms,
  hr,
  ops,
  reading,
  team,
  platform
]

export default {
  routes(){
    return engines.reduce(function(obj, en) {
      return obj.concat(en.routes)
    }, []);
  },
  reducers(){
    return engines.reduce(function(obj, en) {
      return combineReducers(...obj, en.reducers)
    }, {});
  }
}

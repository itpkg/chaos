import cms from './cms'
import hr from './hr'
import ops from './ops'
import platform from './platform'
import reading from './reading'
import team from './team'

const engines = {
  cms,
  hr,
  ops,
  reading,
  team,
  platform
}

export default {
  routes(){
    return CHAOS_ENV.engines.reduce(function(obj, en) {
      return obj.concat(engines[en].routes)
    }, []);
  },
  reducers(){
    return CHAOS_ENV.engines.reduce(function(obj, en) {
      return Object.assign(obj, engines[en].reducers)
    }, {});
  }
}

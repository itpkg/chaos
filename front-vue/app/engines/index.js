// import cms from './cms'
// import hr from './hr'
// import ops from './ops'
// import reading from './reading'
// import team from './team'
import platform from './platform'

const engines = {
  // cms,
  // hr,
  // ops,
  // reading,
  // team,
  platform
}

export default {
  routes () {
    return engines.reduce(function (obj, en) {
      return obj.concat(engines[en].routes)
    }, [])
  },
  reducers () {
    return engines.reduce(function (obj, en) {
      return Object.assign(obj, engines[en].reducers)
    }, {})
  }
}

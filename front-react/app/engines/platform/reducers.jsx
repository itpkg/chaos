import { combineReducers } from 'redux'

import {AUTH_SIGN_IN, AUTH_SIGN_OUT, SITE_REFRESH} from './actions'

function currentUser(state = {}, action){
  switch (action.type) {
    case AUTH_SIGN_IN:
      console.log(action.token);
      return {name:"change-me", uid:"1234", roles:["admin", "root"]}
    case AUTH_SIGN_OUT:
      return {}
    default:
      return state
  }
}

function siteInfo(state = {}, action){
  switch (action.type) {
    case SITE_REFRESH:
      return action.info
    default:
      return state
  }
}

const reducers = combineReducers({currentUser, siteInfo})
export default reducers

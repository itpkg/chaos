import jwtDecode from 'jwt-decode'

import {
  AUTH_SIGN_IN, AUTH_SIGN_OUT,
  SITE_REFRESH,
  USER_INFO, USER_LOGS,
  ADMIN_SITE_INFO
} from './actions'

import {TOKEN} from '../../constants'

function parse(tkn) {
    try {
        return jwtDecode(tkn);
    } catch (e) {
        return {}
    }
}

const initCurrentUserState = parse(sessionStorage.getItem(TOKEN));

function currentUser(state = initCurrentUserState, action){
  switch (action.type) {
    case AUTH_SIGN_IN:
      sessionStorage.setItem(TOKEN, action.token)
      return parse(action.token);
    case AUTH_SIGN_OUT:
      sessionStorage.removeItem(TOKEN)
      return {}
    default:
      return state
  }
}

function siteInfo(state = {navLinks:[], oauth2:[]}, action){
  switch (action.type) {
    case SITE_REFRESH:
      return action.info
    default:
      return state
  }
}

function dashboard(state = {site:{author:{}, navLinks:[]}, logs:[], user:{}}, action){
  switch (action.type) {
    case ADMIN_SITE_INFO:
      return Object.assign({}, state, {site:action.info})
    case USER_INFO:
      return Object.assign({}, state, {user: action.info})
    case USER_LOGS:
      return Object.assign({}, state, {logs: action.logs})
    default:
      return state
  }
}

const reducers = {currentUser, siteInfo, dashboard}
export default reducers

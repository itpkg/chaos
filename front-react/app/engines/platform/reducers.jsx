import jwtDecode from 'jwt-decode'

import {AUTH_SIGN_IN, AUTH_SIGN_OUT, DASHBOARD, SITE_REFRESH} from './actions'

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

function siteInfo(state = {links:[], oauth2:[]}, action){
  switch (action.type) {
    case SITE_REFRESH:
      return action.info
    default:
      return state
  }
}

function dashboard(state = {site:{author:{}, links:[]}, logs:[], user:{}}, action){
  switch (action.type) {
    case DASHBOARD:      
      return action.info
    default:
      return state
  }
}

const reducers = {currentUser, siteInfo, dashboard}
export default reducers

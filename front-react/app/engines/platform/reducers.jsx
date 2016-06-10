import jwtDecode from 'jwt-decode'

import {AUTH_SIGN_IN, AUTH_SIGN_OUT, SITE_REFRESH} from './actions'

const key='token'

function parse(tkn) {
    try {
        return jwtDecode(tkn);
    } catch (e) {
        return {}
    }
}

const initCurrentUserState = parse(sessionStorage.getItem(key));

function currentUser(state = initCurrentUserState, action){
  switch (action.type) {
    case AUTH_SIGN_IN:
      sessionStorage.setItem(key, action.token)
      return parse(action.token);
    case AUTH_SIGN_OUT:
      sessionStorage.removeItem(key)
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

const reducers = {currentUser, siteInfo}
export default reducers

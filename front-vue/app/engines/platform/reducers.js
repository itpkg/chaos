import jwtDecode from 'jwt-decode'

import { AUTH_SIGN_IN, AUTH_SIGN_OUT, SITE_REFRESH, USER_INFO, USER_LOGS, ADMIN_SITE_INFO, NOTICE_LIST, NOTICE_ADD, NOTICE_DEL } from './actions'

const TOKEN = 'token'

function parse (tkn) {
  try {
    return jwtDecode(tkn)
  } catch (e) {
    return {}
  }
}

const initCurrentUserState = parse(window.sessionStorage.getItem(TOKEN))

function currentUser (state = initCurrentUserState, action) {
  switch (action.type) {
    case AUTH_SIGN_IN:
      window.sessionStorage.setItem(TOKEN, action.token)
      return parse(action.token)
    case AUTH_SIGN_OUT:
      window.sessionStorage.removeItem(TOKEN)
      return {}
    default:
      return state
  }
}

function siteInfo (state = {aboutUs: '', navLinks: [], oauth2: []}, action) {
  switch (action.type) {
    case SITE_REFRESH:
      return action.info
    default:
      return state
  }
}

function dashboard (state = {site: {author: {}, navLinks: []}, logs: [], user: {}}, action) {
  switch (action.type) {
    case ADMIN_SITE_INFO:
      return Object.assign({}, state, {site: action.info})
    case USER_INFO:
      return Object.assign({}, state, {user: action.info})
    case USER_LOGS:
      return Object.assign({}, state, {logs: action.logs})
    default:
      return state
  }
}

function notices (state = [], action) {
  switch (action.type) {
    case NOTICE_LIST:
      return action.notices
    case NOTICE_ADD:
      state.unshift(action.notice)
      return state.slice(0)
    case NOTICE_DEL:
      for (var i = 0; i < state.length; i++) {
        if (state[i].id === action.id) {
          state.splice(i, 1)
          break
        }
      }
      return state.slice(0)
    default:
      return state

  }
}

const reducers = {currentUser, siteInfo, dashboard, notices}
export default reducers

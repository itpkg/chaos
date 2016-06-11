export const AUTH_SIGN_IN = "platform.auth.sign_in"
export const AUTH_SIGN_OUT = "platform.auth.sign_out"

export const SITE_REFRESH = "platform.site.refresh"

export const USER_INFO = "platform.user_info"
export const USER_LOGS = "platform.user_logs"

export const ADMIN_SITE_INFO = "platform.admin_site_info"

export function signIn(token){
  return {type:AUTH_SIGN_IN, token}
}

export function signOut(){
  return {type:AUTH_SIGN_OUT}
}

export function refresh(info){
  return {type:SITE_REFRESH, info}
}

export function userInfo(info){
  return {type:USER_INFO, info}
}
export function userLogs(logs){
  return {type:USER_LOGS, logs}
}

export function adminSiteInfo(info){
  return {type:ADMIN_SITE_INFO, info}
}

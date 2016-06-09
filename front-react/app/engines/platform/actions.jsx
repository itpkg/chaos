export const AUTH_SIGN_IN = "platform.auth.sign_in"
export const AUTH_SIGN_OUT = "platform.auth.sign_out"
export const SITE_REFRESH = "platform.site.refresh"

export function signIn(token){
  return {type:AUTH_SIGN_IN, token}
}

export function signOut(){
  return {type:AUTH_SIGN_OUT}
}

export function refresh(info){  
  return {type:SITE_REFRESH, info}
}

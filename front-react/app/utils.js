import $ from 'jquery'

export function ajax(method, url, data, done, fail){
  $.ajax({
    url: CHAOS_ENV.backend+url,
    method: method,
    data: data,
    xhrFields: {
      withCredentials: true
    }
  }).then(done, fail);
}

export function isSignIn(user){
  var now = new Date().getTime()
  return !$.isEmptyObject(user) && user.nbf < now && user.exp > now
}

export function hasRole(user, name){
  return isSignIn(user) && $.inArray(name, user.roles)
}

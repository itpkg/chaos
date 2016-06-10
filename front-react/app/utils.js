import $ from 'jquery'

import {TOKEN} from './constants'

export function ajax(method, url, data, done, fail){
  if(!fail){
    fail= function(xhr){
      alert(xhr.responseText)
    }
  }
  $.ajax({
    url: CHAOS_ENV.backend+url,
    method: method,
    data: data,
    crossDomain : true,
    xhrFields: {
      withCredentials: true
    },
    beforeSend: function (xhr) {
      xhr.setRequestHeader('Authorization', 'Bearer '+sessionStorage.getItem(TOKEN));
    }
  }).then(done, fail);
}

export function isSignIn(user){
  var now = new Date().getTime()/1000
  return !$.isEmptyObject(user) && user.nbf < now && user.exp > now
}

export function hasRole(user, name){
  return isSignIn(user) && $.inArray(name, user.roles)!==-1
}

export function isAdmin(user){
  return hasRole(user, "admin")
}

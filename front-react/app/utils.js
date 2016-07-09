import $ from 'jquery'
import i18next from 'i18next'

import {
    TOKEN
} from './constants'

export function onDelete (url, done) {
  if (window.confirm(i18next.t('messages.are_you_sure'))) {
    ajax('delete', url, null, done)
  }
}

export function ajax (method, url, data, done, fail) {
  if (!done) {
    done = function (rst) {
      window.alert(i18next.t('messages.success'))
    }
  }
  if (!fail) {
    fail = function (xhr) {
      window.alert(xhr.responseText)
    }
  }
  $.ajax({
    url: process.env.CHAOS.backend + url,
    method: method,
    data: data,
    crossDomain: true,
    xhrFields: {
      withCredentials: true
    },
    beforeSend: function (xhr) {
      xhr.setRequestHeader('Authorization', 'Bearer ' + window.sessionStorage.getItem(TOKEN))
    }
  }).then(done, fail)
}

export function isSignIn (user) {
  var now = new Date().getTime() / 1000
  return !$.isEmptyObject(user) && user.nbf < now && user.exp > now
}

export function hasRole (user, name) {
  return isSignIn(user) && $.inArray(name, user.roles) !== -1
}

export function isAdmin (user) {
  return hasRole(user, 'admin')
}

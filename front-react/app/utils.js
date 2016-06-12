import $ from 'jquery'
import i18next from 'i18next'

import {
    TOKEN
} from './constants'

export function onDelete(url, done) {
    if (confirm(i18next.t('messages.are_you_sure'))) {
        ajax('delete', url, null, done)
    }
}

export function ajax(method, url, data, done, fail) {
    if (!done) {
        done = function(rst) {
            alert(i18next.t("messages.success"))
        }
    }
    if (!fail) {
        fail = function(xhr) {
            alert(xhr.responseText)
        }
    }
    $.ajax({
        url: CHAOS_ENV.backend + url,
        method: method,
        data: data,
        crossDomain: true,
        xhrFields: {
            withCredentials: true
        },
        beforeSend: function(xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + sessionStorage.getItem(TOKEN));
        }
    }).then(done, fail);
}

export function isSignIn(user) {
    var now = new Date().getTime() / 1000
    return !$.isEmptyObject(user) && user.nbf < now && user.exp > now
}

export function hasRole(user, name) {
    return isSignIn(user) && $.inArray(name, user.roles) !== -1
}

export function isAdmin(user) {
    return hasRole(user, "admin")
}
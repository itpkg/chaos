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

import AjaxService from 'ember-ajax/services/ajax';
import ENV from 'it-package/config/environment';

import $ from 'jquery';
$.ajaxSetup({
  xhrFields: {
    withCredentials: true
  }
});

export default AjaxService.extend({
  host: ENV.APP.API_HOST
});

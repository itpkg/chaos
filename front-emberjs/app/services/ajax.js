import Ember from 'ember';
import $ from 'jquery';
import ENV from 'it-package/config/environment';

export default Ember.Service.extend({
  host: null,
  init(){
    this.set("host", ENV.APP.API_HOST);
  },
  get(url, data, success){
    $.getJSON(this.host+url, data).done(success);
  }
});

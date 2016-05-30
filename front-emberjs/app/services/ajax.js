import Ember from 'ember';
import $ from 'jquery';

export default Ember.Service.extend({
  host: null,
  init(){
    this.set("host", "http://localhost:8080");
  },
  get(url, data, success){
    $.getJSON(this.host+url, data).done(success);

  }
});

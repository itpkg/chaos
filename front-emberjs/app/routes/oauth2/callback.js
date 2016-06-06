import Ember from 'ember';
//import {parseUrl, parseXhr} from '../../utils';
import { translationMacro as t } from "ember-i18n";

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  i18n: Ember.inject.service(),
  model(){
    console.log(t("messages.please_waiting"));
    return {
      alert: {
        style:"info",
        messages:[
          Ember.computed('i18n.locale', function(){
            return this.get('i18n').t("messages.please_waiting");
          })
        ],
        created:new Date()}
    };
  },
  init(){
    this._super(...arguments);
    //this.set("alert", {style:"danger", messages:"aaa", created:new Date()});
//    this.get('ajax')
//      .post('/oauth2/callback', {data: parseUrl()})
//      .then(
//        function(rst){
//          //this.set('item', rst);
//          console.log(rst);
//        }.bind(this),
//        function(xhr){
//          this.set("alert", {style:"danger", messages:parseXhr(xhr), created:new Date()});
//        }.bind(this));

  },
});

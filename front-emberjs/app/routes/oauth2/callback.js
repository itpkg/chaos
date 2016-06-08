import Ember from 'ember';
import {parseUrl, parseXhr} from '../../utils';
import { translationMacro as t } from "ember-i18n";

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  i18n: Ember.inject.service(),
  alert: null,
  // model(){
  //   this.store.push({
  //     data:{
  //       id:"alert",
  //         style:"info",
  //         messages:[this.get('i18n').t("messages.please_waiting")],
  //         created:new Date()
  //     }
  //   });
  // },
  init(){
    this._super(...arguments);

    this.set("alert", {style:"danger", messages:["aaa"], created:new Date()});
    console.log(this.get('alert'));
   this.get('ajax')
     .post('/oauth2/callback', {data: parseUrl()})
     .then(
       function(data){
         //this.set('item', rst);
         console.log(data);
       }.bind(this),
       function(jqXHR){
         this.set("alert", {style:"danger", messages:["aaa"], created:new Date()});
       }.bind(this));

  }
});

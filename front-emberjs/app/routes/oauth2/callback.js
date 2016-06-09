import Ember from 'ember';
import {parseUrl} from '../../utils';

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  i18n: Ember.inject.service(),
  auth: Ember.inject.service(),
  alertBox: Ember.inject.service(),
  init(){
    this._super(...arguments);
    this.get('alertBox').show("info", [this.get('i18n').t("messages.please_waiting")]);
    this.get('ajax')
       .post('/oauth2/callback', {data: parseUrl()})
       .then(
         function(rst){
           this.get('auth').signIn(rst.token);
           this.transitionTo('index');
         }.bind(this),
         function(jqXHR){
           this.get('alertBox').error(jqXHR);
         }.bind(this));

    }
});

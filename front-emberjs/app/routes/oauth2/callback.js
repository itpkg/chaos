import Ember from 'ember';
import {parseUrl, parseXhr} from '../../utils';

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  init(){
    this._super(...arguments);

    this.get('ajax')
      .post('/oauth2/callback', {data: parseUrl()})
      .then(
        function(rst){
          //this.set('item', rst);
          console.log(rst);
        }.bind(this),
        function(xhr){
          console.log(parseXhr(xhr));
        }.bind(this));

  },
});

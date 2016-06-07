import Ember from 'ember';

export default Ember.Service.extend({
  items: null,
  ajax: Ember.inject.service(),
  init() {
    this.get('ajax').request('/reading/notes').then(function(rst){ this.set('items', rst); }.bind(this));
  }
});

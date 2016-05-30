import Ember from 'ember';

export default Ember.Service.extend({
  items: null,
  ajax: Ember.inject.service(),
  init() {
    this.get('ajax').get('/reading/notes', null, function(rst){
      this.set('items', rst);
    }.bind(this));
  }
});

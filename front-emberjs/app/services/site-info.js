import Ember from 'ember';

export default Ember.Service.extend({
  item: null,
  ajax: Ember.inject.service(),
  init() {
    this.get('ajax').get('/info', null, function(rst){
      this.set('item', rst);
    }.bind(this));
  }
});

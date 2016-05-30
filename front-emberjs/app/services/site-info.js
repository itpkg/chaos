import Ember from 'ember';

export default Ember.Service.extend({
  item: null,
  ajax: Ember.inject.service(),
  init() {
    this.get('ajax').request('/info').then(function(rst){this.set('item', rst);}.bind(this));
  }
});

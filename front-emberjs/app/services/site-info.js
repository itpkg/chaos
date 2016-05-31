import Ember from 'ember';

export default Ember.Service.extend({
  i18n: Ember.inject.service(),
  item: null,
  ajax: Ember.inject.service(),
  init() {
    this.refresh(this.get('i18n.locale'));
  },
  refresh(lang){
    this.get('ajax')
      .request('/info', {data:{locale:lang.replace(/-/g, '_')}})
      .then(function(rst){this.set('item', rst);}.bind(this));
  }
});

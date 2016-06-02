import Ember from 'ember';

export default Ember.Component.extend({
  init(){
    this._super(...arguments);
    this.set('createdAt', new Date());
  },
  createdAt: null,
});

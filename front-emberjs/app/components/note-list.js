import Ember from 'ember';

export default Ember.Component.extend({
  model() {
    console.log("note-list");
    return ['111', '222', '333'];
  }
});

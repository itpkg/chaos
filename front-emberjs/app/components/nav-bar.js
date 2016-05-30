import Ember from 'ember';

export default Ember.Component.extend({
  siteInfo: Ember.inject.service(),
  i18n: Ember.inject.service(),
  doubleClick(lang){
    console.log(lang);
  }
});

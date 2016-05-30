import Ember from 'ember';

const key = "locale";

export default Ember.Component.extend({
  siteInfo: Ember.inject.service(),
  i18n: Ember.inject.service(),
  init(){
    this._super(...arguments);
    var lang = localStorage.getItem(key);
    if(lang != null){
      this.set('i18n.locale', lang);
    }
  },
  actions: {
    switchLang(lang){
      this.set('i18n.locale', lang);
      localStorage.setItem(key, lang);
    }
  }
});

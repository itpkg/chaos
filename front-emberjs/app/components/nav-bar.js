import Ember from 'ember';
import ENV from 'it-package/config/environment';

const key = "locale";

export default Ember.Component.extend({
  siteInfo: Ember.inject.service(),
  i18n: Ember.inject.service(),
  init(){
    this._super(...arguments);
    var lang = localStorage.getItem(key);
    if(lang == null){
      lang = ENV.i18n.defaultLocale;
      localStorage.setItem(key, lang);
    }
    this.set('i18n.locale', lang);
  },
  actions: {
    switchLang(lang){
      this.set('i18n.locale', lang);
      localStorage.setItem(key, lang);
    }
  }
});

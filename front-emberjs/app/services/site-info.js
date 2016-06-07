import Ember from 'ember';
import ENV from 'it-package/config/environment';

const key = "locale";

export default Ember.Service.extend({
  item: null,
  i18n: Ember.inject.service(),
  ajax: Ember.inject.service(),
  init() {
    //this._super(...arguments);
    var lang = localStorage.getItem(key);
    if(lang == null){
      lang = ENV.i18n.defaultLocale;
    }else{
      this.set('i18n.locale', lang);
    }
    this.refresh(lang);
  },
  refresh(lang){
    localStorage.setItem(key, lang);
    this.get('ajax')
      .request('/info', {data:{locale:lang.replace(/-/g, '_')}})
      .then(function(rst){this.set('item', rst);}.bind(this));
  }
});

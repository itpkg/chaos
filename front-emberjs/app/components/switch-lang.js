import Ember from 'ember';

export default Ember.Component.extend({
  tagName:'ul',
  classNames:['nav', 'navbar-nav', 'navbar-right'],
  siteInfo: Ember.inject.service(),
  i18n: Ember.inject.service(),
  actions: {
    switchLang(lang){
      this.set('i18n.locale', lang);
      this.get('siteInfo').refresh(lang);
    }
  }
});

import Ember from 'ember';

export default Ember.Component.extend({
  tagName:'li',
  classNames:['dropdown'],
  i18n: Ember.inject.service(),
  auth: Ember.inject.service(),
  siteInfo: Ember.inject.service(),
});

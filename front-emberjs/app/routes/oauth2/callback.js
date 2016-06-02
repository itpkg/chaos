import Ember from 'ember';

export default Ember.Route.extend({
  ajax: Ember.inject.service(),
  init(){
    this._super(...arguments);
    this.get('ajax')
      .post(
        '/oauth2/callback',
        {data:{href:window.location.href}})
      .then(
        function(rst){this.set('item', rst);
      }.bind(this));

  },
});

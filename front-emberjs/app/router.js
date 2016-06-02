import Ember from 'ember';
import config from './config/environment';

const Router = Ember.Router.extend({
  location: config.locationType
});

Router.map(function() {
  this.route('reading', function() {
    this.route('notes');
  });

  this.route('personal', function() {
    this.route('sign-in');
    this.route('sign-up');
    this.route('forgot-password');
    this.route('reset-password');
    this.route('confirm');
    this.route('unlock');
    this.route('profile');
    this.route('dashboard');
    this.route('oauth');
  });
});

export default Router;

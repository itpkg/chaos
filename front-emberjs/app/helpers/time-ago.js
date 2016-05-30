import Ember from 'ember';
//import moment from 'moment';
//var moment = require('moment');
//var moment = require('moment');
//import moment from 'moment';
//import moment from 'bower_components/moment/moment.js';

export function timeAgo(date) {
  return moment(date, moment.ISO_8601).fromNow();
}

export default Ember.Helper.helper(timeAgo);

import Ember from 'ember';
//import moment from 'moment';

export function timeAgo(date) {
  return moment(date, moment.ISO_8601).fromNow();
}

export default Ember.Helper.helper(timeAgo);

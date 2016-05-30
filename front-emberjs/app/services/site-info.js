import Ember from 'ember';
//import $ from 'jquery';

export default Ember.Service.extend({
  item: null,
  init() {
  this._super(...arguments);
    this.set('copyright', 'ccc');
    this.set('title', 'ttt');
    this.set('item', {
      title:'ttt',
      copyright:'cccc',
      languages: ['en_US', 'zh_CN'],
      links:[
        {href:'/', label:'Home'},
        {href:'/notes', label:'Notes'}
      ]
    });
  }
});

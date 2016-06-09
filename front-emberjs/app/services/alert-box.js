import Ember from 'ember';

export default Ember.Service.extend({
  item: null,
  hide(){
    this.set('item', null);
  },
  show(style, messages){
    this.set('item', {style:style, messages:messages, created:new Date()});
  },
  error(xhr){
    //console.log(xhr);
    this.set(
      'item',
      {
        style: "danger",
        messages: xhr.errors.map(function(e){return e.detail;}),
        created: new Date()
      }
    );
  }  
});

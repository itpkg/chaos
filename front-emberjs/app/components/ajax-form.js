import Ember from 'ember';

export default Ember.Component.extend({
  ajax: Ember.inject.service(),
  actions: {
    submit(){
      var fm = this.get("form");
      var data = {};
      fm.fields.forEach(f =>{
        return data[f.id] = f.value;
      });

      //console.log(data);
      this.get('ajax')
      .request(fm.action, {data:data, method:fm.method})
      .then(function(rst){
        //this.set('item', rst);
        console.log(rst);
      }.bind(this));

    }
  }
});

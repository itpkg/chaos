import Ember from 'ember';
import {parseXhr} from '../utils';

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
      .then(
        function(rst){
          //console.log(rst);
          if( typeof rst === 'string'){
            this.set("alert", {style:"success", messages:[rst], created:new Date()});
          }
        }.bind(this),
        function(xhr){
          this.set("alert", {style:"danger", messages:parseXhr(xhr), created:new Date()});
        }.bind(this)
      );
    }
  }
});

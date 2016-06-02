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
      .then(
        function(rst){
          //console.log(rst);
          if( typeof rst == 'string'){
            this.set("alert", {style:"success", messages:[rst], created:new Date()});
          }
        }.bind(this),
        function(xhr){
          console.log(xhr);
          var msg = xhr.errors.map(function(e){
            if(e.detail){
              return e.detail.message;
            }
            if(e.field && e.defaultMessage){
              return e.field+" "+e.defaultMessage;
            }
            return e;
          });
          this.set("alert", {style:"danger", messages:msg, created:new Date()});
        }.bind(this)
      );
    }
  }
});

import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      title: "core.pages.personal.confirm",
      action: "/personal/confirm",
      method: "POST",
      fields: [
        {
          id: "email",
          label: "forms.label.email",
          type: "email"
        }
      ],
      buttons: [
        {
          type: "submit"
        },
        {
          type: "reset"
        }
      ]
    };
  }
});

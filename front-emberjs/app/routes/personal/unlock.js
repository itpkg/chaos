import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      title: "core.pages.personal.unlock",
      action: "/personal/unlock",
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

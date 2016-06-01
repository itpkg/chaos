import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      title: "core.pages.personal.signIn",
      action: "/personal/signIn",
      method: "POST",
      fields: [
        {
          id: "email",
          label: "forms.label.email",
          type: "email"
        },
        {
          id: "password",
          label: "forms.label.password",
          type: "password"
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

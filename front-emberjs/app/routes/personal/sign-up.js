import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      title: "core.pages.personal.signUp",
      action: "/personal/signUp",
      method: "POST",
      fields: [
        {
          id: "name",
          label: "forms.label.username",
          type: "text"
        },
        {
          id: "email",
          label: "forms.label.email",
          type: "email"
        },
        {
          id: "password",
          label: "forms.label.password",
          type: "password"
        },
        {
          id: "passwordConfirm",
          label: "forms.label.passwordConfirm",
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

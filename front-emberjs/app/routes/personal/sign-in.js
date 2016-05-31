import Ember from 'ember';

export default Ember.Route.extend({
  model(){
    return {
      title: "title",
      action: "/personal/signIn",
      method: "POST",
      fields: [
        {
          id: "email",
          label: "Email",
          type: "email"
        },
        {
          id: "password",
          label: "Password",
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

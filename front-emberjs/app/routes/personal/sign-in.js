import Ember from 'ember';
//import { translationMacro as t } from "ember-i18n";

export default Ember.Route.extend({
  i18n: Ember.inject.service(),
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

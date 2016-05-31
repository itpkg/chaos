import Ember from 'ember';

const key = "token";

function parse(token){
  //todo
  return {username: token, email:'bbb', roles:['admin']};
}

export default Ember.Service.extend({
  user: null,
  init(){
    var token = sessionStorage.getItem(key);
    if(token !=null){
      this.set('user', parse(token));
    }
  },
  signIn(token){
    this.set('user', parse(token));
    sessionStorage.setItem(key, token);
  },
  signOut(){
    this.set('user', null);
    sessionStorage.removeItem(key);
  },
  isSignIn(){
    return this.get('user') != null;
  }
});

import Ember from 'ember';

const key = "token";

function parse(token){
  try{
      return JSON.parse(Base64.decode(token.split('.')[1]));
  }catch(e){
    console.log(e);
    return null;
  }
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
  }
});

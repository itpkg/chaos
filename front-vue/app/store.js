import Vue from 'vue'
import Revue from 'revue'
import {createStore} from 'redux'
// create the logic how you would update the todos
import root from './engines'

// create a redux store
const reduxStore = createStore(root.reducers)
// binding the store to Vue instance, actions are optional
const store = new Revue(Vue, reduxStore)
// expose the store for your component to use
export default store

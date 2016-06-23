import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Home from './Home'
import {Index as BookIndex} from './Books'
import Layout from './Layout'

export default [
  <Route key="reading" path="reading" component={Layout}>
    <Route path="books" component={BookIndex}/>
  </Route>
]

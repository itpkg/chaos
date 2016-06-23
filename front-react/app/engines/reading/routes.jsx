import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Home from './Home'

export default [
  <Route key="reading" path="reading">
    <IndexRoute component={Home}/>    
  </Route>
]

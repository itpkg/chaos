import React from 'react'
import {Route, IndexRoute} from 'react-router'

import Notes from './Notes'

export default [
  <Route key="reading" path="reading">
    <Route path="notes" component={Notes}/>
  </Route>
]

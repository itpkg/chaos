import React from 'react'
import {Route, IndexRoute} from 'react-router'

import {Index as BookIndex} from './books'
import {Index as BlogIndex, Show as BlogShow} from './blogs'
import Layout from './Layout'

export default [
  <Route key="reading" path="reading" component={Layout}>
    <Route path="books" component={BookIndex}/>
    <Route path="blogs" component={BlogIndex}/>
    <Route path="blog/**" component={BlogShow}/>
  </Route>
]

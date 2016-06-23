import React from 'react'
import {Route, IndexRoute} from 'react-router'

import {Index as BookIndex, Show as BookShow} from './books'
import {Index as BlogIndex, Show as BlogShow} from './blogs'
import Notes from './Notes'
import Layout from './Layout'

export default [
  <Route key="reading" path="reading" component={Layout}>
    <Route path="books" component={BookIndex}/>
    <Route path="book/**" component={BookShow}/>

    <Route path="blogs" component={BlogIndex}/>
    <Route path="blog/**" component={BlogShow}/>

    <Route path="notes" component={Notes}/>
  </Route>
]

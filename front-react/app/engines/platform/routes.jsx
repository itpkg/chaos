import React from 'react'
import {Route, IndexRoute} from 'react-router'

import {Index,AboutUs} from './home'
import {Index as NoticesIndex} from './notices'
import {Callback as Oauth2Callback} from './oauth2'

export default [
  <IndexRoute key="platform.index" component={Index}/>,
  <Route key="platform.about_us" path="about_us" component={AboutUs}/>,
  <Route key="platform.notices" path="notices" component={NoticesIndex}/>,
  <Route key="platform.oauth2_callback" path="oauth2/callback" component={Oauth2Callback}/>
]

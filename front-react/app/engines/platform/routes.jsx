import React from 'react'
import {Route, IndexRoute} from 'react-router'

import {Index,AboutUs} from './home'
import {Index as NoticesIndex} from './notices'
import {Callback as OauthCallback} from './oauth'

export default [
  <IndexRoute key="platform.index" component={Index}/>,
  <Route key="platform.about_us" path="about_us" component={AboutUs}/>,
  <Route key="platform.notices" path="notices" component={NoticesIndex}/>,
  <Route key="platform.oauth_callback" path="oauth/callback" component={OauthCallback}/>
]

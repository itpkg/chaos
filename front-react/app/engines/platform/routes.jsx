import React from 'react'
import {Route, IndexRoute} from 'react-router'

import {Index,AboutUs} from './home'
import {Index as NoticesIndex} from './notices'
import {Callback as Oauth2Callback} from './oauth2'
import {Profile as PersonalProfile} from './personal'

export default [
  <IndexRoute key="platform.index" component={Index}/>,
  <Route key="platform.notices" path="home" component={Index}/>,
  <Route key="platform.about_us" path="about-us" component={AboutUs}/>,
  <Route key="platform.notices" path="notices" component={NoticesIndex}/>,
  <Route key="platform.oauth2/callback" path="oauth2/callback" component={Oauth2Callback}/>,
  <Route key="platform.personal/profile" path="personal/profile" component={PersonalProfile}/>
]

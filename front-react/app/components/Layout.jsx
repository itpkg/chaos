import React from 'react'
import { connect } from 'react-redux'

import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'
import darkBaseTheme from 'material-ui/styles/baseThemes/darkBaseTheme'
import getMuiTheme from 'material-ui/styles/getMuiTheme'

import Header from './Header'
import Footer from './Footer'
// import {refresh} from '../engines/platform/actions'

const Widget = ({onRefresh, children}) => (
  <MuiThemeProvider muiTheme={getMuiTheme(darkBaseTheme)}>
    <div>
      <Header/>
      {children}
      <hr />
      <Footer/>
    </div>
  </MuiThemeProvider>
)

Widget.propTypes = {
  onRefresh: React.PropTypes.func.isRequired
  // children: PropTypes.node.isRequired
}

export default connect(
  state => ({info: state.siteInfo}),
  dispatch => ({
    onRefresh () {
      // TODO 刷新页面
      // ajax('get', "/site/info", null, function(ifo){
      //   dispatch(refresh(ifo));
      //   document.documentElement.lang = ifo.lang;
      //   document.title = ifo.title;
      // });
    }
  })
)(Widget)

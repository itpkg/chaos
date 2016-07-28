import { PropTypes } from 'react'
import { connect } from 'react-redux'
import MuiThemeProvider from 'material-ui/styles/MuiThemeProvider'

import Header from './Header'
import Footer from './Footer'
// import {refresh} from '../engines/platform/actions'

const Widget = ({onRefresh, children}) => (
  <MuiThemeProvider>
    <Header/>
    {this.props.children}
    <Footer/>
  </MuiThemeProvider>
)

Widget.propTypes = {
  onRefresh: PropTypes.func.isRequired,
  children: PropTypes.node.isRequired
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

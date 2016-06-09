import React, { PropTypes } from 'react'
import { connect } from 'react-redux'

import Header from './Header'
import Footer from './Footer'
import {refresh} from '../engines/platform/actions'

const Widget = React.createClass({
  componentDidMount: function(){
    const {onRefresh} = this.props;
    onRefresh();
  },
  render() {
    return (
      <div>
        <Header/>
        <div className="container-fluid">
            {this.props.children}
            <hr/>
            <Footer/>
        </div>
      </div>
    )
  }
})

Widget.propTypes = {
    onRefresh: PropTypes.func.isRequired
};

export default connect(
  state=>({info:state.siteInfo}),
  dispatch => ({
    onRefresh: function(){
      //TODO
      dispatch(refresh({title:'aaa', subTitle:'sub title', copyright:'ccc'}))
      // ajax("get", "/site/info", null, function(ifo){
      //   dispatch(refresh(ifo));
      //   document.documentElement.lang = ifo.lang;
      //   document.title = ifo.title;
      // });
    }
  })
)(Widget);

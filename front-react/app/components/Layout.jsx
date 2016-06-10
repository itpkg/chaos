import React, { PropTypes } from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'

import Header from './Header'
import Footer from './Footer'
import {refresh} from '../engines/platform/actions'
import {ajax} from '../utils'

const Widget = React.createClass({
  componentDidMount: function(){
    const {onRefresh} = this.props;
    onRefresh();
  },
  render() {
    return (
      <div>
        <Header/>
        <div className="container">
          <div className="row">
              {this.props.children}
          </div>          
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
      ajax("get", "/info", null, function(ifo){
        dispatch(refresh(ifo));
        document.documentElement.lang = ifo.lang;
        document.title = ifo.title;
      });
    }
  })
)(Widget);

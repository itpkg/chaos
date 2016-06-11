import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import {Link} from 'react-router'
import ReactMarkdown from 'react-markdown'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'

import {ajax} from '../../utils'

export const Index = React.createClass({
  getInitialState: function() {
    return {
      notices: []
    }
  },
  componentDidMount: function() {
    ajax('get', '/notices', null, function(rst){
      this.setState({notices:rst})
    }.bind(this))
  },
  render() {
    const {notices} = this.state
    return (
      <div className="col-md-10 col-md-offset-1">
        <h3>{i18next.t("platform.notices")}</h3>
        <hr/>
        {notices.map((n,i)=>{
          return <blockquote key={i}>
                  <ReactMarkdown source={n.content}/>
                  <footer><cite><TimeAgo date={n.created_at}/></cite></footer>
                 </blockquote>
        })}
      </div>
    )
  }
})


//-----------------------------------------------------------------------------

const AboutUsW = React.createClass({
  render() {
    const {info} = this.props
    return (
        <ReactMarkdown source={info.aboutUs} />
    )
  }
})


AboutUsW.propTypes = {
    info: PropTypes.object.isRequired
}

export const AboutUs = connect(
  state => ({ info: state.siteInfo })
)(AboutUsW)

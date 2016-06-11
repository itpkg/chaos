import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import {Link} from 'react-router'
import ReactMarkdown from 'react-markdown'

export const Index = React.createClass({
  render() {
    return (
      <div>
        <ul>
          <li><Link to="/notices">Notices</Link></li>
          <li><Link to="/oauth2/callback">Oauth Callback</Link></li>
          <li><Link to="/about-us">About us</Link></li>
          <li><Link to="/">Index</Link></li>
        </ul>
        index
      </div>
    )
  }
})

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

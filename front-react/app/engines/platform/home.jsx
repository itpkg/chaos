import React from 'react'
import {Link} from 'react-router'

export const Index = React.createClass({
  render() {
    return (
      <div>
        <ul>
          <li><Link to="/notices">Notices</Link></li>
          <li><Link to="/oauth2/callback">Oauth Callback</Link></li>
          <li><Link to="/about_us">About us</Link></li>
          <li><Link to="/">Index</Link></li>
        </ul>
        index
      </div>
    )
  }
})

export const AboutUs = React.createClass({
  render() {
    return (
      <div>
        about us
      </div>
    )
  }
})

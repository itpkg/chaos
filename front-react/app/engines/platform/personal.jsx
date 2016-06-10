import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import {Tabs, Tab, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'

import {isSignIn, hasRole} from '../../utils'

const Profile =  React.createClass({
  render() {
    console.log("profile")
    return (
      <div>
        <br/>
        profile
      </div>
    )
  }
})

//-----------------------------------------------------------------------------
const DashboardW = React.createClass({
  getInitialState() {
    return {
      key: "self"
    };
  },
  handleSelect(key) {
    event.preventDefault();
    this.setState({key});
    console.log(key)
  },
  render() {
    const {user} = this.props
    return (
        <Tabs activeKey={this.state.key} onSelect={this.handleSelect} id="personal-dashboard">
          <Tab eventKey={"self"} title="Profile"><Profile/></Tab>
          <Tab eventKey={"siteInfo"} title="Site info">Tab 2 content</Tab>
          <Tab eventKey={"logs"} title="Logs">Tab 3 content</Tab>
        </Tabs>
    )
  }
})


DashboardW.propTypes = {
    user: PropTypes.object.isRequired
}

export const Dashboard = connect(
  state => ({user:state.currentUser}),
  dispatch => ({
}))(DashboardW);

import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import {Tabs, Tab, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'

import {isSignIn, isAdmin, ajax} from '../../utils'
import NoMatch from '../../components/NoMatch'

const Profile =  React.createClass({
  render() {
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
      key: "profile"
    };
  },
  componentDidMount: function(){
    const {user} = this.props
    if(isSignIn(user)){
      ajax(
        "get",
        "/personal/dashboard",
        null,
        function(rst){
          console.log(rst)
        },
        function(xhr){
          alert(xhr.responseText)
        }
      )
    }
  },
  handleSelect(key) {
    this.setState({key});
  },
  render() {
    const {user} = this.props
    if(isSignIn(user)){
      var tabs = [
        (<Tab key="profile" eventKey={"profile"} title="Profile">
          <Profile/>
        </Tab>)
      ]
      if( isAdmin(user)){
        tabs.push(<Tab key="site.info" eventKey={"site.info"} title="Site info">
          Tab 2 content
        </Tab>)
      }
      tabs.push(<Tab key="logs" eventKey={"logs"} title="Logs">Tab 3 content</Tab>)
      return (
          <Tabs activeKey={this.state.key} onSelect={this.handleSelect} id="personal-dashboard">
            {tabs}
          </Tabs>
      )
    }
    return <NoMatch />
  }
})


DashboardW.propTypes = {
    user: PropTypes.object.isRequired
}

export const Dashboard = connect(
  state => ({user:state.currentUser}),
  dispatch => ({
}))(DashboardW);

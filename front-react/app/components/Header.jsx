import React, {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'
import {Navbar, Nav, NavItem, NavDropdown, MenuItem} from 'react-bootstrap'


const Widget = React.createClass({
  render() {
    const {info} = this.props
    return (
      <Navbar inverse fixedTop fluid>
        <Navbar.Header>
          <Navbar.Brand>
            <a href="#">{info.subTitle}</a>
          </Navbar.Brand>
          <Navbar.Toggle />
        </Navbar.Header>
        <Navbar.Collapse>
          <Nav>
            <NavItem eventKey={1} href="#">Link</NavItem>
            <NavItem eventKey={2} href="#">Link</NavItem>
            <NavDropdown eventKey={3} title="Dropdown" id="basic-nav-dropdown">
              <MenuItem eventKey={3.1}>Action</MenuItem>
              <MenuItem eventKey={3.2}>Another action</MenuItem>
              <MenuItem eventKey={3.3}>Something else here</MenuItem>
              <MenuItem divider />
              <MenuItem eventKey={3.3}>Separated link</MenuItem>
            </NavDropdown>
          </Nav>
          <Nav pullRight>
            <NavItem eventKey={1} href="#">English</NavItem>
            <NavItem eventKey={2} href="#">简体中文</NavItem>
          </Nav>
        </Navbar.Collapse>
      </Navbar>
    )
  }
})

Widget.propTypes = {
    info: PropTypes.object.isRequired,
}

export default connect(
  state => ({ info: state.siteInfo })
)(Widget)

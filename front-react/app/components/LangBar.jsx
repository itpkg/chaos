import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {Nav, NavItem} from 'react-bootstrap'

const Widget = React.createClass({
  render() {
    return (
    <Nav pullRight>
      <NavItem href="/?locale=en-US" target="_blank">English</NavItem>
      <NavItem href="/?locale=zh-CN" target="_blank">简体中文</NavItem>
    </Nav>
    )
  }
})

export default Widget

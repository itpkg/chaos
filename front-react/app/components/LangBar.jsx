import React from 'react'
import i18next from 'i18next'
import {Nav, NavItem} from 'react-bootstrap'

const Widget = React.createClass({
  render() {
    return (
    <Nav pullRight>
      <NavItem href="/?locale=en-US" target="_blank">{i18next.t("locales.en_US")}</NavItem>
      <NavItem href="/?locale=zh-CN" target="_blank">{i18next.t("locales.zh_Hans")}</NavItem>
    </Nav>
    )
  }
})

export default Widget

import React, {PropTypes} from 'react'
import {Link} from 'react-router'
import { connect } from 'react-redux'
import i18next from 'i18next'
import {Navbar, Nav, NavItem} from 'react-bootstrap'
import {IndexLinkContainer} from 'react-router-bootstrap'

import LangBar from './LangBar'
import PersonalBar from './PersonalBar'

const Widget = React.createClass({
  render() {
    const {info} = this.props
    return (
      <Navbar inverse fixedTop>
        <Navbar.Header>
          <Navbar.Brand>
            <Link to="/">{info.subTitle}</Link>
          </Navbar.Brand>
          <Navbar.Toggle />
        </Navbar.Header>
        <Navbar.Collapse>
          <Nav>
            {info.links.map((l,i) =>{
              return <IndexLinkContainer key={i} to={l.href}>
                      <NavItem>{i18next.t(l.label)}</NavItem>
                     </IndexLinkContainer>
            })}
            <PersonalBar />
          </Nav>
          <LangBar />
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

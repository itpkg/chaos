import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {NavDropdown, MenuItem} from 'react-bootstrap'
import {connect} from 'react-redux'
import $ from 'jquery'
import {IndexLinkContainer} from 'react-router-bootstrap'

const Widget = React.createClass({
    render() {
        const {user, info} = this.props

        return $.isEmptyObject(user)
            ? (
                <NavDropdown title={i18next.t('platform.sign_in_or_up')} id="personal-bar">
                    <MenuItem href={info.oauth2.google}>{i18next.t('platform.sign_in_with_google')}</MenuItem>
                </NavDropdown>
            )
            : (
                <NavDropdown title={i18next.t('platform.welcome', {name: user.name})} id="personal-bar">
                    <IndexLinkContainer to='/personal/profile'>
                        <MenuItem>{i18next.t('platform.dashboard')}</MenuItem>
                    </IndexLinkContainer>
                    <MenuItem divider/>
                    <MenuItem>{i18next.t('platform.sign_out')}</MenuItem>
                </NavDropdown>
            )
    }
})

Widget.propTypes = {
    user: PropTypes.object.isRequired,
    info: PropTypes.object.isRequired
}

export default connect(state => ({user: state.currentUser, info: state.siteInfo}))(Widget)

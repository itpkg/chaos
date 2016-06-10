import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import i18next from 'i18next'
import {Alert} from 'react-bootstrap'
import parse from 'url-parse'
import {browserHistory} from 'react-router'

import {signIn} from './actions'
import {ajax} from '../../utils'

const CallbackW = React.createClass({
    componentDidMount() {
        const {onSignIn} = this.props;
        onSignIn();
    },
    render() {
        return (
            <div className="row">
                <br/>
                <div className="col-md-offset-1 col-md-10">
                    <Alert bsStyle="warning">
                        <strong>{i18next.t("messages.please_waiting")}</strong>{new Date().toLocaleString()}
                    </Alert>
                </div>
            </div>
        )
    }
})

CallbackW.propTypes = {
    onSignIn: PropTypes.func.isRequired
}

export const Callback = connect(state => ({}), dispatch => ({
    onSignIn: function() {
        ajax(
          'post',
          '/oauth2/callback',
          parse(location.href, true).query,
          function(rst) {
            dispatch(signIn(rst.token));
            browserHistory.push('/');
          },
      )
    }
}))(CallbackW);

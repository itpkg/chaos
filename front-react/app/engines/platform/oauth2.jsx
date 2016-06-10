import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import i18next from 'i18next'
import {Alert} from 'react-bootstrap'
import parse from 'url-parse'
import {browserHistory} from 'react-router'

import {signIn} from './actions'
import {ajax} from '../../utils'

const CallbackW = React.createClass({
    getInitialState: function() {
        return {
            alert: {
                style: "info",
                message: i18next.t("messages.please_waiting"),
                created: new Date()
            }
        };
    },
    componentDidMount() {
        const {onSignIn} = this.props;
        onSignIn(function(xhr) {
            this.setState({
                alert: {
                    style: "danger",
                    message: xhr.responseText,
                    created: new Date()
                }
            })
        }.bind(this));
    },
    render() {
        var msg = this.state.alert;
        return (
            <div className="row">
                <br/>
                <div className="col-md-offset-1 col-md-10">
                    <Alert bsStyle={msg.style}>
                        <strong>{msg.created.toLocaleString()}:
                        </strong>{msg.message}
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
    onSignIn: function(showError) {
        ajax('post', '/oauth2/callback', parse(location.href, true).query, function(rst) {
            dispatch(signIn(rst.token));
            browserHistory.push('/');
        }, showError)
    }
}))(CallbackW);

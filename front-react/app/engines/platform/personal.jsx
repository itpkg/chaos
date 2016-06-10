import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import {Tabs, Tab, ListGroup, ListGroupItem, Jumbotron, Thumbnail} from 'react-bootstrap'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'

import {isSignIn, isAdmin, ajax} from '../../utils'
import NoMatch from '../../components/NoMatch'

const Profile = React.createClass({
    render() {
        const {user} = this.props
        return (
            <div className="col-md-offset-1 col-md-6">
              <br/>
                <Thumbnail src={user.logo} alt="242x200">
                  <h3>{user.name}</h3>

                    <ul>
                      <li>{i18next.t("platform.user.email")}: {user.email}</li>
                      <li>{i18next.t("platform.user.provider")}: {user.provider_id}@{user.provider_type}</li>
                      <li>{i18next.t("platform.user.last_sign_in")}: {user.last_sign_in}</li>
                      <li>{i18next.t("platform.user.sign_in_count")}: {user.sign_in_count}</li>
                      <li>{i18next.t("platform.user.confirmed_at")}: {user.confirmed_at}</li>
                    </ul>

                </Thumbnail>
            </div>
        )
    }
})
//-----------------------------------------------------------------------------
const Logs = React.createClass({
    render() {
        const {items} = this.props
        return (
            <div className="col-md-offset-1 col-md-10">
                <br/>
                <ListGroup>
                    {items.map((l, i) => {
                        return <ListGroupItem key={i}>
                            <TimeAgo date={l.created_at}/>: {l.message}
                        </ListGroupItem>
                    })}
                </ListGroup>
            </div>
        )
    }
})
//-----------------------------------------------------------------------------
const SiteInfo = React.createClass({
    render() {
        return (
            <div>
                site info
            </div>
        )
    }
})
//-----------------------------------------------------------------------------
const DashboardW = React.createClass({
    getInitialState() {
        return {
          key: "profile",
          logs: [],
          user: {},
          site: {}
        };
    },
    componentDidMount: function() {
        const {user} = this.props
        if (isSignIn(user)) {
            ajax("get", "/personal/dashboard", null, function(rst) {
                this.setState(rst)
            }.bind(this))
        }
    },
    handleSelect(key) {
        this.setState({key});
    },
    render() {
        const {user} = this.props
        if (isSignIn(user)) {
            var tabs = [(
                    <Tab key="profile" eventKey={"profile"} title={i18next.t('platform.profile')}>
                        <Profile user={this.state.user}/>
                    </Tab>
                )]
            if (isAdmin(user)) {
                tabs.push(
                    <Tab key="site.info" eventKey={"site.info"} title={i18next.t('platform.site_info')}>
                        <SiteInfo info={this.state.site}/>
                    </Tab>
                )
            }
            tabs.push(
                <Tab key="logs" eventKey={"logs"} title={i18next.t('platform.logs')}>
                    <Logs items={this.state.logs}/>
                </Tab>
            )
            return (
                <Tabs activeKey={this.state.key} onSelect={this.handleSelect} id="personal-dashboard">
                    {tabs}
                </Tabs>
            )
        }
        return <NoMatch/>
    }
})

DashboardW.propTypes = {
    user: PropTypes.object.isRequired
}

export const Dashboard = connect(state => ({user: state.currentUser}), dispatch => ({}))(DashboardW);

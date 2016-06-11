import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import {
    Tabs,
    Tab,
    ListGroup,
    ListGroupItem,
    Thumbnail,
    Button,
    Form,
    Col,
    FormGroup,
    ControlLabel,
    FormControl
} from 'react-bootstrap'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'

import {setDashboard} from './actions'
import {isSignIn, isAdmin, ajax} from '../../utils'
import NoMatch from '../../components/NoMatch'

const ProfileW = React.createClass({
    render() {
        const {user} = this.props
        return (
            <div className="col-md-offset-1 col-md-6">
                <br/>
                <Thumbnail src={user.logo}>
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

ProfileW.propTypes = {
    user: PropTypes.object.isRequired
}

const Profile = connect(state => ({user: state.dashboard.user}), dispatch => ({}))(ProfileW);
//-----------------------------------------------------------------------------
const LogsW = React.createClass({
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

LogsW.propTypes = {
    items: PropTypes.array.isRequired
}

const Logs = connect(state => ({items: state.dashboard.logs}), dispatch => ({}))(LogsW);
//-----------------------------------------------------------------------------
const SiteInfoFm = React.createClass({
    render() {
        return (
            <div>
                site info
            </div>
        )
    }
})
//-----------------------------------------------------------------------------
const AboutUsFm = React.createClass({
    render() {
        return (
            <div>
                abort info
            </div>
        )
    }
})
//-----------------------------------------------------------------------------
const NavLinksFmW = React.createClass({
    getInitialState: function() {
        const {value} = this.props
        return {value: value};
    },
    componentWillReceiveProps(props){
        const {value} = props
        this.setState({value: value});
    },
    handleChange: function(e) {
        this.setState({value: e.target.value});
    },
    handleSubmit: function(e) {
      e.preventDefault();
      ajax(
        "post",
        "/admin/site/navLinks",
        {value:this.state.value}
      )
    },
    render() {
        return (
            <form onSubmit={this.handleSubmit}>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.links")}</ControlLabel>
                    <FormControl rows={12} componentClass="textarea" value={this.state.value} onChange={this.handleChange}/>
                </FormGroup>
                <Button type="submit" bsStyle="primary">
                    {i18next.t("buttons.save")}
                </Button>
            </form>
        )
    }
})
NavLinksFmW.propTypesW = {
    value: PropTypes.string.isRequired
}

const NavLinksFm = connect(state => ({
    value: JSON.stringify(state.dashboard.site.links, null, 2)//state.dashboard.site.links.map(l => (l.href + ": " + l.label)).join("\n")
}), dispatch => ({}))(NavLinksFmW);
//-----------------------------------------------------------------------------
const DashboardW = React.createClass({
    getInitialState() {
        return {key: "profile"};
    },
    componentDidMount: function() {
        const {onRefresh, user} = this.props
        if (isSignIn(user)) {
            onRefresh()
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
                        <Profile/>
                    </Tab>
                )]
            if (isAdmin(user)) {
                tabs.push(
                    <Tab key="site.info" eventKey={"site.info"} title={i18next.t('platform.site_info')}>
                        <div className="col-md-offset-1 col-md-10">
                            <br/>
                            <SiteInfoFm/>
                            <br/>
                            <NavLinksFm/>
                            <br/>
                            <AboutUsFm/>
                        </div>
                    </Tab>
                )
            }
            tabs.push(
                <Tab key="logs" eventKey={"logs"} title={i18next.t('platform.logs')}>
                    <Logs/>
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
    user: PropTypes.object.isRequired,
    onRefresh: PropTypes.func.isRequired
}

export const Dashboard = connect(state => ({user: state.currentUser}), dispatch => ({
    onRefresh: function() {
        ajax("get", "/personal/dashboard", null, function(rst) {
            dispatch(setDashboard(rst));
        })
    }
}))(DashboardW);

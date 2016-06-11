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

import {userInfo, userLogs, adminSiteInfo} from './actions'
import {isSignIn, isAdmin, onDelete, ajax} from '../../utils'
import NoMatch from '../../components/NoMatch'
import Notices from './Notices'

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
const SiteInfoFmW = React.createClass({
    getInitialState: function() {
      return {
        title:'',
        subTitle: '',
        keywords: '',
        description: '',
        aboutUs:'',
        copyright: '',
        author:{},
        navLinks:[],

        links: '{}',
        authorName: '',
        authorEmail:''
      }
    },
    componentWillReceiveProps(props){
        const {info} = props
        info.links = JSON.stringify(info.navLinks, null, 2)
        info.authorEmail = info.author.email
        info.authorName = info.author.name
        this.setState(info);
    },
    handleChange: function(e) {
          var o = {}
          o[e.target.id]=e.target.value
          this.setState(o);
    },

    handleSubmit: function(e) {
      e.preventDefault();
      var data = Object.assign({}, this.state)
      data.navLinks = data.links
      delete data.links;
      delete data.author
      //console.log(data)
      ajax(
        "post",
        "/admin/site/info",
        data
      )
    },
    render: function() {
        return (
            <form onSubmit={this.handleSubmit}>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.title")}</ControlLabel>
                    <FormControl id="title" type="text" value={this.state.title} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.subTitle")}</ControlLabel>
                    <FormControl id="subTitle" type="text" value={this.state.subTitle} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.author.name")}</ControlLabel>
                    <FormControl id="authorName" type="text" value={this.state.authorName} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.author.email")}</ControlLabel>
                    <FormControl id="authorEmail" type="text" value={this.state.authorEmail} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.keywords")}</ControlLabel>
                    <FormControl id="keywords" type="text" value={this.state.keywords} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.description")}</ControlLabel>
                    <FormControl id="description" rows={12} componentClass="textarea" value={this.state.description} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.aboutUs")}</ControlLabel>
                    <FormControl id="aboutUs" rows={12} componentClass="textarea" value={this.state.aboutUs} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.copyright")}</ControlLabel>
                    <FormControl id="copyright" type="text" value={this.state.copyright} onChange={this.handleChange}/>
                </FormGroup>
                <FormGroup>
                    <ControlLabel>{i18next.t("platform.site.navLinks")}</ControlLabel>
                    <FormControl id="links" rows={12} componentClass="textarea" value={this.state.links} onChange={this.handleChange}/>
                </FormGroup>

                <Button type="submit" bsStyle="primary">
                    {i18next.t("buttons.save")}
                </Button>
            </form>
        )
    }
})
SiteInfoFmW.propTypesW = {
    info: PropTypes.object.isRequired
}

const SiteInfoFm = connect(state => ({
    info: state.dashboard.site
}), dispatch => ({}))(SiteInfoFmW);
//-----------------------------------------------------------------------------
const StatusPane = React.createClass({
  handleClearCache(){
    onDelete("/admin/cache")
  },
  render(){
    return <div className="col-md-offset-1 col-md-10">
      <br/>
      <div className="row">
        <div className="col-ms-3">
          <Button bsStyle="danger" onClick={this.handleClearCache}>{i18next.t("platform.clear_cache")}</Button>
        </div>
      </div>
    </div>
  }
})
//-----------------------------------------------------------------------------
const DashboardW = React.createClass({
    getInitialState() {
        return {key: "profile"};
    },
    componentDidMount: function() {
        const {onRefresh, user} = this.props
        onRefresh(user)
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
                        </div>
                    </Tab>
                )
                tabs.push(
                    <Tab key="site.notices" eventKey={"site.notices"} title={i18next.t('platform.notices')}>
                        <Notices/>
                    </Tab>
                )
                tabs.push(
                    <Tab key="site.status" eventKey={"site.status"} title={i18next.t('platform.site_status')}>
                        <StatusPane/>
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
    onRefresh: function(user) {
        if (!isSignIn(user)) {
          return
        }
        ajax("get", "/personal/self", null, function(rst) {
            dispatch(userInfo(rst));
        })
        ajax("get", "/personal/logs", null, function(rst) {
            dispatch(userLogs(rst));
        })
        if(isAdmin(user)){
          ajax("get", "/admin/site/info", null, function(rst) {
              dispatch(adminSiteInfo(rst));
          })
        }
    }
}))(DashboardW);

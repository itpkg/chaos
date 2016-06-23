import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import i18next from 'i18next'
import {Nav, NavItem} from 'react-bootstrap'
import {IndexLinkContainer} from 'react-router-bootstrap'

import Dict from './Dict'
import {isSignIn, ajax} from '../../utils'

const Widget = React.createClass({
    render() {
        const {user} = this.props;
        var links = ["books", "blogs"]
        if(isSignIn(user)){
          links.push("notes");
        }
        return (
            <div>
                <div className="col-md-3">
                    <Nav bsStyle="pills" stacked activeKey={0}>
                      {links.map((l,i)=>{
                          return (
                            <IndexLinkContainer key={i} to={"/reading/"+l}>
                            <NavItem  eventKey={i} >
                            {i18next.t("reading.pages."+l)}
                          </NavItem>
                          </IndexLinkContainer>
                        )
                      })}
                    </Nav>
                    <br/>
                    <Dict/>
                </div>
                <div className="col-md-9">
                    {this.props.children}
                </div>
                <hr/>
            </div>
        )
    }
});

Widget.propTypes = {
    user: PropTypes.object.isRequired
}

export default connect(state => ({notes: state.readingNotes, user: state.currentUser}), dispatch => ({}))(Widget);

import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {Link} from 'react-router'
import i18next from 'i18next'

import Dict from './Dict'

const Home = React.createClass({
    getInitialState: function() {
        return {}
    },
    componentDidMount: function() {
      console.log("init rading.page")
    },
    render() {
        return (
            <div>
                <div className="col-md-3">
                    <Dict/>
                </div>
                <div className="col-md-9"></div>
                <hr/>
            </div>
        )
    }
});

export default Home;

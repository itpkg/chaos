import React, {PropTypes} from 'react'
import {Link} from 'react-router'
import i18next from 'i18next'

import Dict from './Dict'
import Notes from './Notes'
import Books from './Books'

const Widget = React.createClass({
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
                    <br/>
                    <Notes/>
                </div>
                <div className="col-md-9">
                  <Books/>
                </div>
                <hr/>
            </div>
        )
    }
});

export default Widget;

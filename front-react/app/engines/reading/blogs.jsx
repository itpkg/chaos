import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {ListGroup, ListGroupItem} from 'react-bootstrap'
import ReactMarkdown from 'react-markdown'
import {IndexLinkContainer} from 'react-router-bootstrap'

import {ajax} from '../../utils'

export const Index = React.createClass({
    getInitialState() {
        return {items:{}};
    },
    componentDidMount() {
      ajax('get', '/reading/blogs', null, function(rst){
        this.setState({items:rst});
      }.bind(this))
    },
    render() {
        var items = this.state.items;
        return (
          <ListGroup>
            {Object.keys(items).map((v,i)=>{
              return (
              <IndexLinkContainer key={i} to={`/reading/blog/${v}`}>
                <ListGroupItem> {items[v]} </ListGroupItem>
              </IndexLinkContainer>
            )
            })}
          </ListGroup>
        )
    }
});


export const Show = React.createClass({
    getInitialState() {
        return {content:""};
    },
    componentDidMount() {
      const {params} = this.props;      
      ajax('get', '/reading/blog/'+params.splat, null, function(rst){
         this.setState({content:rst});
      }.bind(this))
    },
    render() {
        return <ReactMarkdown source={this.state.content}/>
    }
});

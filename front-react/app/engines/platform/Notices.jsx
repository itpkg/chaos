import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import {
    ListGroup,
    ListGroupItem,
    Button,
    FormGroup,
    ControlLabel,
    FormControl
} from 'react-bootstrap'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'

import {addNotice, delNotice, listNotice} from './actions'
import {isSignIn, isAdmin, ajax, onDelete} from '../../utils'
import NoMatch from '../../components/NoMatch'

const IndexW = React.createClass({
  getInitialState: function() {
    return {
      content: ''
    }
  },
  componentDidMount: function() {
    const {onListNotice} = this.props
    onListNotice()
  },
  handleChange: function(e) {
        var o = {}
        o[e.target.id]=e.target.value
        this.setState(o);
  },
  handleSubmit: function(e) {
    e.preventDefault();
    const {onAddNotice} = this.props
    ajax(
      'post',
      '/admin/notices',
      {content:this.state.content},
      function(n){
        onAddNotice(n)
        this.setState({content:''})
      }.bind(this)
    )
  },
  handleRemove: function(id){
    const {onDelNotice} = this.props
    onDelete('/admin/notices/'+id, function(){
      onDelNotice(id)
      this.forceUpdate() //FIXME
    }.bind(this)
    );
  },
  render() {
    const {notices} = this.props
    return (
      <div className="col-md-10 col-md-offset-1">
        <br/>
        <form onSubmit={this.handleSubmit}>
          <FormGroup>
              <ControlLabel>{i18next.t("platform.notice.content")}</ControlLabel>
              <FormControl id="content" rows={16} componentClass="textarea" value={this.state.content} onChange={this.handleChange}/>
          </FormGroup>
          <Button type="submit" bsStyle="primary">
              {i18next.t("buttons.new")}
          </Button>
        </form>
        <br/>
        <ListGroup>
          {notices.map((n,i)=>{
            return (<ListGroupItem key={i}>
              {n.created_at}: &nbsp;
              <Button bsStyle="danger" onClick={()=>this.handleRemove(n.id)} bsSize="small">{i18next.t("buttons.remove")}</Button>
              <br/>
              {n.content}
            </ListGroupItem>)
          })}
        </ListGroup>
      </div>
    )
  }
})

IndexW.propTypes = {
    notices: PropTypes.array.isRequired,
    onAddNotice: PropTypes.func.isRequired,
    onListNotice: PropTypes.func.isRequired,
    onDelNotice: PropTypes.func.isRequired
}

const Index = connect(state => ({notices: state.notices}), dispatch => ({
  onAddNotice: function(n){
    dispatch(addNotice(n))
  },
  onDelNotice: function(i){
    dispatch(delNotice(i))
  },
  onListNotice: function(){
    ajax('get', '/admin/notices', null, function(rst){dispatch(listNotice(rst))})
  }
}))(IndexW);
export default Index

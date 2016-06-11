import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'
import ReactMarkdown from 'react-markdown'

import {addNote, listNote, delNote} from './actions'
import {isSignIn, ajax, onDelete} from '../../utils'
import NoMatch from '../../components/NoMatch'

const Widget = React.createClass({
  componentDidMount: function() {
    const {user, onListNote} = this.props
    ajax('get', '/reading/notes', null, function(rst){
      onListNote(rst)
    })
  },
  render(){
    const {user, notes} = this.props
    return (<div className="contai">      
      {notes.map((n, i)=>{
        return (<div className="col-md-3" key={i}>
          <h2>{n.title}</h2>
          {n.body}
        </div>)
      })}
    </div>)
  }
})


Widget.propTypes = {
    user: PropTypes.object.isRequired,
    notes: PropTypes.array.isRequired,
    onAddNote: PropTypes.func.isRequired,
    onListNote: PropTypes.func.isRequired,
    onDelNote: PropTypes.func.isRequired
}

export default connect(
  state => ({notes: state.readingNotes, user:state.currentUser}),
  dispatch => ({
    onAddNote: function(n){
      dispatch(addNote(n))
    },
    onDelNote: function(i){
      dispatch(delNote(i))
    },
    onListNote: function(){
      ajax('get', '/reading/notes', null, function(rst){dispatch(listNote(rst))})
    }
}))(Widget);

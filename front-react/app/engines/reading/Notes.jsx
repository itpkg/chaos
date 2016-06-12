import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'
import ReactMarkdown from 'react-markdown'
import {Modal, Button, FormGroup, ControlLabel, FormControl} from 'react-bootstrap'

import {addNote, listNote, chgNote, delNote} from './actions'
import {isSignIn, ajax, onDelete} from '../../utils'
import NoMatch from '../../components/NoMatch'

const Widget = React.createClass({
    getInitialState() {
        return {showModal: false, title: '', body: '', id: null};
    },
    close() {
        this.setState({showModal: false});
    },
    open(id) {
        const {notes} = this.props
        if (id) {
            for (var i = 0; i < notes.length; i++) {
                var n = notes[i]
                if (n.id === id) {
                    this.setState({showModal: true, id: id, title: n.title, body: n.body})
                    break
                }
            }
        } else {
            this.setState({showModal: true, id: null, title: '', body: ''});
        }

    },
    handleChange: function(e) {
        var o = {}
        o[e.target.id] = e.target.value
        this.setState(o);
    },
    handleSubmit: function(e) {
        e.preventDefault();
        const {onChgNote, onAddNote} = this.props
        var id = this.state.id
        ajax('post', id
            ? "/reading/notes/" + id
            : "/reading/notes", {
            title: this.state.title,
            body: this.state.body
        }, function(rst) {
            this.setState({showModal: false, id: null, title: '', body: ''})
            if (id) {
                onChgNote(rst)
            } else {
                onAddNote(rst)
            }
        }.bind(this))
    },
    componentDidMount: function() {
        const {user, onListNote} = this.props
        ajax('get', '/reading/notes', null, function(rst) {
            onListNote(rst)
        })
    },
    render() {
        const {user, notes} = this.props

        var fm = (
            <div className="col-md-3">
                <Button bsStyle="info" bsSize="large" onClick={this.open.bind(this, null)}>
                    {i18next.t('buttons.new')}
                </Button>
                <form>
                    <Modal show={this.state.showModal} onHide={this.close}>
                        <Modal.Header closeButton>
                            <Modal.Title>{i18next.t(this.state.id
                                    ? 'buttons.edit'
                                    : 'buttons.new')}</Modal.Title>
                        </Modal.Header>
                        <Modal.Body>
                            <FormGroup>
                                <ControlLabel>{i18next.t("reading.note.title")}</ControlLabel>
                                <FormControl id="title" type="text" value={this.state.title} onChange={this.handleChange}/>
                            </FormGroup>
                            <FormGroup>
                                <ControlLabel>{i18next.t("reading.note.body")}</ControlLabel>
                                <FormControl id="body" rows={16} componentClass="textarea" value={this.state.body} onChange={this.handleChange}/>
                            </FormGroup>
                        </Modal.Body>
                        <Modal.Footer>
                            <Button onClick={this.handleSubmit} bsStyle="primary">
                                {i18next.t("buttons.save")}
                            </Button>
                            <Button onClick={this.close}>{i18next.t("buttons.close")}</Button>
                        </Modal.Footer>
                    </Modal>
                </form>
            </div>
        )

        return (
            <div className="container-fluid">
                <div className="row">
                    {isSignIn(user)
                        ? fm
                        : <br/>}
                    {notes.map((n, i) => {
                        return (
                            <div className="col-md-3" key={i}>
                                <h2>{n.title}</h2>
                                <ReactMarkdown source={n.body}/> {isSignIn(user) && n.user_id == user.id
                                    ? <p>
                                            <Button bsStyle='link' onClick={this.open.bind(this, n.id)}>{i18next.t('buttons.edit')}</Button>
                                        </p>
                                    : <br/>
}
                            </div>
                        )
                    })}
                </div>
            </div>
        )
    }
})

Widget.propTypes = {
    user: PropTypes.object.isRequired,
    notes: PropTypes.array.isRequired,
    onAddNote: PropTypes.func.isRequired,
    onListNote: PropTypes.func.isRequired,
    onDelNote: PropTypes.func.isRequired,
    onChgNote: PropTypes.func.isRequired
}

export default connect(state => ({notes: state.readingNotes, user: state.currentUser}), dispatch => ({
    onAddNote: function(n) {
        dispatch(addNote(n))
    },
    onDelNote: function(i) {
        dispatch(delNote(i))
    },
    onChgNote: function(n) {
        dispatch(chgNote(n))
    },
    onListNote: function() {
        ajax('get', '/reading/notes', null, function(rst) {
            dispatch(listNote(rst))
        })
    }
}))(Widget);

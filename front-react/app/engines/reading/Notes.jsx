import React, {PropTypes} from 'react'
import {connect} from 'react-redux'
import TimeAgo from 'react-timeago'
import i18next from 'i18next'
import ReactMarkdown from 'react-markdown'
import {Modal, Button,
  FormGroup, ControlLabel, FormControl,
  ListGroup, ListGroupItem
} from 'react-bootstrap'

import {addNote, listNote, chgNote, delNote} from './actions'
import {isSignIn, ajax, onDelete} from '../../utils'

const Widget = React.createClass({
    getInitialState() {
        return {showModal: false, title: '', body: '', id: null, items: []};
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
    handleRemove: function(id) {
        const {onDelNote} = this.props
        onDelete('/reading/notes/' + id, function(rst) {
            //TODO
            console.log("delete " + id);
        })
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
            //TODO
            console.log("save " + rst.id);
        }.bind(this))
    },
    componentDidMount: function() {
        const {user} = this.props;
        if (isSignIn(user)) {
            ajax("get", "/reading/notes", null, function(rst) {
                this.setState({items: rst});
            }.bind(this));
        }
    },
    render() {
        const {user} = this.props

        var fm = (
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
        )

        return isSignIn(user)
            ? (
                <fieldset>
                    <legend>{i18next.t("reading.pages.notes")}</legend>
                    <div className="pull-right">
                        <Button bsStyle="info" onClick={this.open.bind(this, null)}>
                            {i18next.t('buttons.new')}
                        </Button>
                    </div>
                    <br/>
                    {fm}
                    <br/>
                    <div>
                        <ListGroup>
                            {this.state.items.map((t, i) => {
                                return (
                                    <ListGroupItem key={i} onClick={this.open.bind(this, t)}>{t}</ListGroupItem>
                                )
                            })}
                        </ListGroup>
                    </div>
                </fieldset>
            )
            : (<br/>)
    }
})

Widget.propTypes = {
    user: PropTypes.object.isRequired
}

export default connect(state => ({notes: state.readingNotes, user: state.currentUser}), dispatch => ({}))(Widget);

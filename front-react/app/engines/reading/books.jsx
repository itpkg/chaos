import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {connect} from 'react-redux'
import {
  Button, Thumbnail,
  FormGroup, ControlLabel, FormControl} from 'react-bootstrap'
import {IndexLinkContainer} from 'react-router-bootstrap'
import {Link} from 'react-router'

import {ajax, isAdmin, onDelete} from '../../utils'

const IndexW = React.createClass({
    getInitialState() {
        return {items:[], keyword:''};
    },
    // handleChange: function(e) {
    //     var o = {
    //         result: ''
    //     }
    //     o[e.target.id] = e.target.value
    //     this.setState(o);
    // },
    handleRemove:function(id){
      onDelete("/reading/books/"+id, function(){
        var books = this.state.items;
        for (var i = 0; i < books.length; i++) {
            if (books[i].id === id) {
                books.splice(i, 1);
                break
            }
        }
        this.setState({items:books});
      }.bind(this));
    },
    componentDidMount() {
      ajax('get', '/reading/books', null, function(rst){
        this.setState({items:rst});
      }.bind(this))
    },
    render() {
      const {user} = this.props;
      // <div className="col-md-12">
      //   <FormGroup>
      //     <ControlLabel>{i18next.t("buttons.filter")}</ControlLabel>
      //     <FormControl id="keyword" value={this.state.keyword} onChange={this.handleChange}/>
      //   </FormGroup>
      // </div>
      // <br/>

        var showBook = function(b){
          return <Link className="btn btn-primary" to={`/reading/books/${b.id}`}>
            {i18next.t("buttons.more")}
          </Link>
        }
        return (
          <div className="row">
            <h3>{i18next.t('reading.pages.books')}</h3>
            <hr/>
            {this.state.items.map((b,i)=>{
              return (
                  <div key={i} className="col-md-3">
                    <Thumbnail>
                    <h4>{b.title}</h4>
                    <p>
                      {i18next.t("reading.book.creator")}: {b.creator}<br/>
                      {i18next.t("reading.book.subject")}: {b.subject}<br/>
                      {i18next.t("reading.book.version")}: {b.version}
                    </p>
                    <p>
                      {showBook(b)}
                      &nbsp;
                      {isAdmin(user) ? (<Button onClick={this.handleRemove.bind(this, b.id)} bsStyle="danger">{i18next.t("buttons.remove")}</Button>):(<span/>)}
                    </p>
                    </Thumbnail>
                  </div>
              )
            })}
          </div>
        )
    }
});


IndexW.propTypes = {
    user: PropTypes.object.isRequired,
}

export const Index = connect(
  state => ({user: state.currentUser}),
  dispatch => ({

  })
)(IndexW);

//-----------------------------------------------------------------------------

export const Show = React.createClass({
    getInitialState() {
        return {item:{
          opf:{
            metadata:{
              title:[''],
              creator:[{author:''}]
            }
          },
          ncx:{}
        }};
    },
    componentDidMount() {
      const {params} = this.props;
      ajax('get', '/reading/book/'+params.id, null, function(item){
        this.setState({item:item});
      }.bind(this));
    },
    render() {
      const {params} = this.props;
      const {item} = this.state;
      var show_point = function(p){
        if(p){
          return p.map((l, i)=>{
            return (<li key={i}>
              <a target="_blank"
                href={CHAOS_ENV.backend+'/reading/book/'+params.id+'/'+l.content.src}>
                {l.text}
              </a>
              <ol>
              {show_point(l.points)}
              </ol>
            </li>)
          })
        }
        return null
      }
      return (<fieldset>
      <legend>{item.opf.metadata.title[0]}-{item.opf.metadata.creator[0].author}</legend>
      <ol>
      {show_point(item.ncx.points)}
    </ol>
      </fieldset>)
        // const {params} = this.props;
        // const url = CHAOS_ENV.backend+'/reading/book/'+params.splat;
        // return (
        //   <object className="html-object" data={url}></object>
        // )
    }
});

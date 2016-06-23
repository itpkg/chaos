import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {
  Button, Thumbnail,
  FormGroup, ControlLabel, FormControl} from 'react-bootstrap'
import {IndexLinkContainer} from 'react-router-bootstrap'
import {Link} from 'react-router'

import {ajax} from '../../utils'

export const Index = React.createClass({
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
    componentDidMount() {
      ajax('get', '/reading/books', null, function(rst){
        this.setState({items:rst});
      }.bind(this))
    },
    render() {
      // <div className="col-md-12">
      //   <FormGroup>
      //     <ControlLabel>{i18next.t("buttons.filter")}</ControlLabel>
      //     <FormControl id="keyword" value={this.state.keyword} onChange={this.handleChange}/>
      //   </FormGroup>
      // </div>
      // <br/>

        // var showBook = function(b){
        //   return <Link className="btn btn-primary" to={`/reading/book/${b.id}/${b.home}`}>
        //     {i18next.t("buttons.more")}
        //   </Link>
        // }
        var showBook = function(b){
          return <a className="btn btn-primary" href={CHAOS_ENV.backend+'/reading/book/'+b.id+'/'+b.home} target='_blank'>
                        {i18next.t("buttons.more")}
                      </a>
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
                    </p>
                    </Thumbnail>
                  </div>
              )
            })}
          </div>
        )
    }
});


export const Show = React.createClass({
    componentDidMount() {
      //console.log(params);
    },
    render() {
        const {params} = this.props;
        const url = CHAOS_ENV.backend+'/reading/book/'+params.splat;
        return (
          <object className="html-object" data={url}></object>
        )
    }
});

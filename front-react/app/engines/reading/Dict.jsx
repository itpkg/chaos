import React, {PropTypes} from 'react'
import i18next from 'i18next'
import {Button, FormGroup, ControlLabel, FormControl} from 'react-bootstrap'
import {ajax} from '../../utils'

const Widget = React.createClass({
    getInitialState() {
        return {keyword: '', result: ''};
    },
    handleChange: function(e) {
        var o = {
            result: ''
        }
        o[e.target.id] = e.target.value
        this.setState(o);
    },
    handleSubmit: function(e) {
        e.preventDefault();
        ajax("post", "/reading/dict", {
            keyword: this.state.keyword
        }, function(rst) {
            //console.log(rst);
            this.setState({result: rst});
        }.bind(this))
    },
    render() {
        var rst = this.state.result
            ? (
                <FormGroup>
                    <pre>
                      {this.state.result}
                </pre>
                </FormGroup>
            )
            : (<FormGroup/>);
        return (
            <fieldset>
                <legend>{i18next.t("reading.pages.dict")}</legend>
                <form>
                    <FormGroup>
                        <ControlLabel>{i18next.t("reading.dict.keyword")}</ControlLabel>
                        <FormControl id="keyword" type="keyword" value={this.state.keyword} onChange={this.handleChange}/>
                    </FormGroup>
                    {rst}
                    <Button onClick={this.handleSubmit} bsStyle="primary">
                        {i18next.t("buttons.query")}
                    </Button>
                </form>
            </fieldset>
        )
    }
});

export default Widget;

import React from 'react'
import i18next from 'i18next'
import {Alert} from 'react-bootstrap'

const Widget = React.createClass({
  render() {
    return (
      <Alert bsStyle="danger">
          <strong>{new Date().toLocaleString()}:
          </strong>{i18next.t('')}
      </Alert>
    )
  }
})

export default Widget

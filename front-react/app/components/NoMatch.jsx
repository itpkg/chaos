import { PropTypes } from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'
import {Alert} from 'react-bootstrap'

const Widget = () => (
  <Alert bsStyle="danger">
      <strong>{new Date().toLocaleString()}: </strong>{i18next.t('platform.no_match')}
  </Alert>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({ info: state.siteInfo })
)(Widget)

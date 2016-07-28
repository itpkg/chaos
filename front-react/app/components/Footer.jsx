import {PropTypes} from 'react'
import { connect } from 'react-redux'
import i18next from 'i18next'

const Widget = ({info}) => (
  <footer>
    <p>
      {info.copyright}
      &nbsp;
      <span dangerouslySetInnerHTML={
      {
        __html: i18next.t(
          'platform.build_using',
          {
            link: 'https://github.com/itpkg/chaos'
          }
        )
      }
      } />
    </p>
  </footer>
)

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({ info: state.siteInfo })
)(Widget)

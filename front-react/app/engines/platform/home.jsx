import {PropTypes} from 'react'
import { connect } from 'react-redux'
import ReactMarkdown from 'react-markdown'
// import TimeAgo from 'react-timeago'
// import i18next from 'i18next'

const IndexW = ({info}) => (
  <ReactMarkdown source={info.aboutUs} />
)

IndexW.propTypes = {
  info: PropTypes.object.isRequired
}

export const Index = connect(
  state => ({ info: state.siteInfo })
)(IndexW)

// export const Index = React.createClass({
//   getInitialState () {
//     return {
//       notices: []
//     }
//   },
//   componentDidMount () {
//     ajax('get', '/notices', null, function (rst) {
//       this.setState({ notices: rst })
//     }.bind(this))
//   },
//   render () {
//     const {notices} = this.state
//     return (
//       <div className="col-md-10 col-md-offset-1">
//         <h3>{i18next.t("platform.notices")}</h3>
//         <hr/>
//         {notices.map((n,i)=>{
//           return <blockquote key={i}>
//                   <ReactMarkdown source={n.content}/>
//                   <footer><cite><TimeAgo date={n.created_at}/></cite></footer>
//                  </blockquote>
//         })}
//       </div>
//     )
//   }
// })

// ----------------------------------------------------------------------------

const AboutUsW = ({info}) => (
  <ReactMarkdown source={info.aboutUs} />
)

AboutUsW.propTypes = {
  info: PropTypes.object.isRequired
}

export const AboutUs = connect(
  state => ({ info: state.siteInfo })
)(AboutUsW)

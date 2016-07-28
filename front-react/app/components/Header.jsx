import {PropTypes} from 'react'
// import {Link} from 'react-router'
import { connect } from 'react-redux'
// import i18next from 'i18next'
// import {Navbar, Nav, NavItem} from 'react-bootstrap'
// import {IndexLinkContainer} from 'react-router-bootstrap'
import AppBar from 'material-ui/AppBar'

// import LangBar from './LangBar'
// import PersonalBar from './PersonalBar'

const Widget = ({info}) => (
  <AppBar
    title="Title"
    iconElementLeft={<IconButton><NavigationClose /></IconButton>}
    iconElementRight={
     <IconMenu
       iconButtonElement={
         <IconButton><MoreVertIcon /></IconButton>
       }
       targetOrigin={{horizontal: 'right', vertical: 'top'}}
       anchorOrigin={{horizontal: 'right', vertical: 'top'}}
     >
       <MenuItem primaryText="Refresh" />
       <MenuItem primaryText="Help" />
       <MenuItem primaryText="Sign out" />
     </IconMenu>
   }
  />
)
// const Widget = React.createClass({
//   render() {
//     const {info} = this.props
//     return (
//       <Navbar inverse fixedTop>
//         <Navbar.Header>
//           <Navbar.Brand>
//             <Link to="/">{info.subTitle}</Link>
//           </Navbar.Brand>
//           <Navbar.Toggle />
//         </Navbar.Header>
//         <Navbar.Collapse>
//           <Nav>
//             {info.navLinks.map((l,i) =>{
//               return <IndexLinkContainer key={i} to={l.href}>
//                       <NavItem>{i18next.t(l.label)}</NavItem>
//                      </IndexLinkContainer>
//             })}
//             <PersonalBar />
//           </Nav>
//           <LangBar />
//         </Navbar.Collapse>
//       </Navbar>
//     )
//   }
// })

Widget.propTypes = {
  info: PropTypes.object.isRequired
}

export default connect(
  state => ({ info: state.siteInfo })
)(Widget)

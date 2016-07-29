import React from 'react'
// import {Link} from 'react-router'
import { connect } from 'react-redux'
// import i18next from 'i18next'
// import {Navbar, Nav, NavItem} from 'react-bootstrap'
// import {IndexLinkContainer} from 'react-router-bootstrap'
import AppBar from 'material-ui/AppBar'
import IconButton from 'material-ui/IconButton'
import IconMenu from 'material-ui/IconMenu'
import MenuItem from 'material-ui/MenuItem'
import MoreVertIcon from 'material-ui/svg-icons/navigation/more-vert'
import NavigationMenu from 'material-ui/svg-icons/navigation/menu'
import Divider from 'material-ui/Divider'

// import LangBar from './LangBar'
// import PersonalBar from './PersonalBar'

const Widget = ({info}) => (
  <AppBar
    title="Title"
    iconElementLeft={<IconButton><NavigationMenu /></IconButton>}
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
       <Divider />
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
  info: React.PropTypes.object.isRequired
}

export default connect(
  state => ({ info: state.siteInfo })
)(Widget)

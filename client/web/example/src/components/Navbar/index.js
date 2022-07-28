import React from 'react'
import LogoutButton from '../Logout/LogoutButton'
import { Nav, NavHeader, NavLink, Bars, NavMenu } from './NavbarElements'

const Navbar = () => {
  return (
    <div>
        <Nav>
            <NavHeader to="/">
                <h1>SMART-HYBRID WORKFORCE MANAGER</h1>
            </NavHeader>
            <Bars />
            <NavMenu>
                <NavLink to="/" activeStyle>
                    HOME
                </NavLink>
                <NavLink to="/bookings" activeStyle>
                    BOOKINGS
                </NavLink>
                <NavLink to="/admin" activeStyle>
                    ADMIN
                </NavLink>
                <NavLink to="/profile" activeStyle>
                    PROFILE
                </NavLink>
                <LogoutButton/>
            </NavMenu>
        </Nav>
    </div>
  )
}

export default Navbar
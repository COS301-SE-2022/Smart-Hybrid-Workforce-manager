import React from 'react'
import LogoutButton from '../Logout/LogoutButton'
import { Nav, NavHeader, NavLink, Bars, NavMenu } from './NavbarElements'

const Navbar1 = () => {
  return (
    <div>
        <Nav>
            <NavHeader to="/">
                <h1>SMART-HYBRID WORKFORCE MANAGER</h1>
            </NavHeader>
            <Bars />
            <NavMenu>
                <NavLink to="/">
                    HOME
                </NavLink>
                <NavLink to="/bookings">
                    BOOKINGS
                </NavLink>
                <NavLink to="/admin">
                    ADMIN
                </NavLink>
                <NavLink to="/profile">
                    PROFILE
                </NavLink>
                <LogoutButton/>
            </NavMenu>
        </Nav>
    </div>
  )
}

export default Navbar1
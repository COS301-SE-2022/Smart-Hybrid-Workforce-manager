import React from 'react'
import { Nav, NavHeader, NavLink, Bars, NavMenu, NavBtn, NavBtnLink } from './NavbarElements'

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
                <NavLink to="/teams" activeStyle>
                    TEAMS
                </NavLink>
                <NavLink to="/meetings" activeStyle>
                    MEETINGS
                </NavLink>
            </NavMenu>
            <NavBtn>
                <NavBtnLink to="/sign-in">
                    SIGN IN
                </NavBtnLink>
            </NavBtn>
        </Nav>
    </div>
  )
}

export default Navbar
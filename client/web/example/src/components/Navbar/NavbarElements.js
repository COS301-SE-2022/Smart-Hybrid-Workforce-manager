import styled from "styled-components"
import { NavLink as Link } from "react-router-dom"
import { FaBars } from "react-icons/fa"

export const Nav =  styled.nav`
    background: #ffffff;
    height: 10vh;
    display: flex;
    justify-content: space-between;
    z-index: 10;
    padding-top: 3vh;
    box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
`

export const NavHeader =  styled.nav`
    color: #374146;
    display: flex;
    align-items: center;
    text-decoration: none;
    padding: 0 3vh;
    letter-spacing: 0.5vh;
    font-size: 1.8vh;
    border-style: none;
    height: 100%;
`

export const NavLink = styled(Link)`
    color: #374146;
    display: flex;
    align-items: center;
    text-decoration: none;
    padding: 0 3vh;
    cursor: pointer;
    letter-spacing: 0.5vh;
    font-size: 1.8vh;
    border-style: none;
    height: 100%;

    &.active
    {
        border-bottom: 0.5vh solid #09A4FB;
    }

    &:hover
    {
        color: #919191;
    }
`

export const Bars = styled(FaBars)`
    display: none;
    color: #fff;

    @media screen and (max-width: 768px)
    {
        display: block;
        position: absolute;
        top: 0;
        right: 0;
        transform: translate(-100%, 75%);
        font-size: 1.8rem;
        cursor: pointer;
    }
`

export const NavMenu = styled.div`
    display: flex;
    align-items: center;

    @media screen and (max-width: 768px)
    {
        display: none;
    }
`
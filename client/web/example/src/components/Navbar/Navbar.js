import React, { useContext, useEffect, useRef, useState } from 'react'
import { useNavigate } from "react-router-dom"
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie, FaDoorOpen } from 'react-icons/fa'
import { UserContext } from '../../App';

const Navbar = (props, ref) =>
{
    const homeRef = useRef(null);
    const bookingsRef = useRef(null);
    const dropdownRef = useRef(null);
    const deskRef = useRef(null);
    const meetingRef = useRef(null);
    const mapRef = useRef(null);
    const statisticsRef = useRef(null);

    const [currLocation, setCurrLocation] = useState("");
    const [dropDown, setDropDown] = useState();

    const navigate = useNavigate();

    const {userData,setUserData} = useContext(UserContext);

    const NavigateHome = () =>
    {
        navigate("/");
    }

    const ShowBookings = () =>
    {
        if(!dropDown)
        {
            dropdownRef.current.style.display = "block";
            setDropDown(true);
        }
        else
        {
            dropdownRef.current.style.display = "none";
            setDropDown(false);
        }
    }

    const NavigateDesk = () =>
    {
        navigate("/bookings-desk");
    }

    const NavigateMeeting = () =>
    {
        navigate("/bookings-meetingroom");
    }

    const NavigateMap = () =>
    {
        navigate("/map");
    }

    const NavigateStatistics = () =>
    {
        navigate("/statistics");
    }

    useEffect(() =>
    {
        if(currLocation === "/")
        {
            dropdownRef.current.style.display = "none";
            homeRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/bookings-desk")
        {
            bookingsRef.current.style.color = "#09a2fb";
            dropdownRef.current.style.display = "block";
            deskRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/bookings-meetingroom")
        {
            bookingsRef.current.style.color = "#09a2fb";
            dropdownRef.current.style.display = "block";
            meetingRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/map")
        {
            dropdownRef.current.style.display = "none";
            mapRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/statistics")
        {
            dropdownRef.current.style.display = "none";
            statisticsRef.current.style.color = "#09a2fb";
        }

    },[currLocation])

    useEffect(() =>
    {
        setCurrLocation(window.location.pathname);
    },[])

    return (
        <div ref={ref} className='navbar-container'>
            <div className='logo-container'>
                S.H.W.M
            </div>
            <div className='navlink-container'>
                <div ref={homeRef} className='navlink' onClick={NavigateHome}>
                    <FaCalendar />
                    &nbsp;
                    Calendar
                </div>
                <div ref={bookingsRef} className='navlink' onClick={ShowBookings}>
                    <FaTicketAlt />
                    &nbsp;
                    Bookings
                </div>
                <div ref={dropdownRef} className='navlink-dropdown-container'>
                    <div ref={deskRef} className='navlink-dropdown' onClick={NavigateDesk}>
                        Desk
                    </div>
                    <div ref={meetingRef} className='navlink-dropdown' onClick={NavigateMeeting}>
                        Meeting Room
                    </div>
                </div>
                <div ref={mapRef} className='navlink' onClick={NavigateMap}>
                    <FaMap />
                    &nbsp;
                    Office Map
                </div>
                <div ref={statisticsRef} className='navlink' onClick={NavigateStatistics}>
                    <FaChartPie />
                    &nbsp;
                    Statistics
                </div>
            </div>
        </div>
    )
}

export default React.forwardRef(Navbar)
import React, { useEffect, useRef, useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie } from 'react-icons/fa'

const Navbar = (props, ref) =>
{
    const homeRef = useRef(null);
    const bookingsRef = useRef(null);
    const mapRef = useRef(null);
    const statisticsRef = useRef(null);

    const [currLocation, setCurrLocation] = useState("");

    const navigate = useNavigate();

    const NavigateHome = () =>
    {
        setCurrLocation("/home")
        navigate("/");
    }

    const NavigateBookings = () =>
    {
        setCurrLocation("/bookings")
        navigate("/");
    }

    const NavigateMap = () =>
    {
        setCurrLocation("/map")
        navigate("/");
    }

    const NavigateStatistics = () =>
    {
        setCurrLocation("/statistics")
        navigate("/");
    }

    useEffect(() =>
    {
        if(homeRef.current)
        {
            if(currLocation === "/home")
            {
                homeRef.current.style.color = "#09a2fb";
            }
        }

        if(bookingsRef.current)
        {
            if(currLocation === "/bookings")
            {
                bookingsRef.current.style.color = "#09a2fb";
            }
        }

        if(mapRef.current)
        {
            if(currLocation === "/map")
            {
                mapRef.current.style.color = "#09a2fb";
            }
        }

        if(statisticsRef.current)
        {
            if(currLocation === "/statistics")
            {
                statisticsRef.current.style.color = "#09a2fb";
            }
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
                <div ref={bookingsRef} className='navlink' onClick={NavigateBookings}>
                    <FaTicketAlt />
                    &nbsp;
                    Bookings
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
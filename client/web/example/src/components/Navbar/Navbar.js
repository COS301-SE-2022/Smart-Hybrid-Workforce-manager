import React, { useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie } from 'react-icons/fa';
import styles from './navbar.module.css';

const Navbar = (props, ref) =>
{
    const homeRef = useRef(null);
    const bookingsRef = useRef(null);
    const dropdownRef = useRef(null);
    const deskRef = useRef(null);
    const meetingRef = useRef(null);
    const calendarRef = useRef(null);
    const statisticsRef = useRef(null);

    const [currLocation, setCurrLocation] = useState("");
    const [dropDown, setDropDown] = useState();

    const navigate = useNavigate();

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

    const NavigateCalendar = () =>
    {
        navigate("/calendar");
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

        if(currLocation === "/calendar")
        {
            dropdownRef.current.style.display = "none";
            calendarRef.current.style.color = "#09a2fb";
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
        <div ref={ref} className={styles.navbarContainer}>
            <div className={styles.logoContainer}>
                deskflow
            </div>
            <div className={styles.navlinkContainer}>
                <div ref={homeRef} className={styles.navlink} onClick={NavigateHome}>
                    <FaMap style={{verticalAlign: 'baseline'}}/>
                    &nbsp;
                    Home
                </div>
                <div ref={bookingsRef} className={styles.navlink} onClick={ShowBookings}>
                    <FaTicketAlt style={{verticalAlign: 'baseline'}}/>
                    &nbsp;
                    Bookings
                </div>
                <div ref={dropdownRef} className={styles.navlinkDropdownContainer}>
                    <div ref={deskRef} className={styles.navlinkDropdown} onClick={NavigateDesk}>
                        Desk
                    </div>
                    <div ref={meetingRef} className={styles.navlinkDropdown} onClick={NavigateMeeting}>
                        Meeting Room
                    </div>
                </div>
                <div ref={calendarRef} className={styles.navlink} onClick={NavigateCalendar}>
                    <FaCalendar style={{verticalAlign: 'baseline'}}/>
                    &nbsp;
                    Calendar
                </div>
                <div ref={statisticsRef} className={styles.navlink} onClick={NavigateStatistics}>
                    <FaChartPie style={{verticalAlign: 'baseline'}}/>
                    &nbsp;
                    Statistics
                </div>
            </div>
        </div>
    )
}

export default React.forwardRef(Navbar)
import React, { useEffect, useRef, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie, FaUserShield } from 'react-icons/fa';
import styles from './navbar.module.css';

const NavbarAdmin = (props, ref) =>
{
    const homeRef = useRef(null);
    const bookingsRef = useRef(null);
    const dropdownRef = useRef(null);
    const deskRef = useRef(null);
    const meetingRef = useRef(null);
    const calendarRef = useRef(null);
    const statisticsRef = useRef(null);
    const adminRef = useRef(null);
    const dropdownAdminRef = useRef(null);
    const teamRef = useRef(null);
    const resourceRef = useRef(null);

    const [currLocation, setCurrLocation] = useState("");
    const [dropDown, setDropDown] = useState();
    const [dropDownAdmin, setDropDownAdmin] = useState();

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

    const ShowAdminOptions = () =>
    {
        if(!dropDownAdmin)
        {
            dropdownAdminRef.current.style.display = "block";
            setDropDownAdmin(true);
        }
        else
        {
            dropdownAdminRef.current.style.display = "none";
            setDropDownAdmin(false);
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

    const NavigateAdmin = () =>
    {
        navigate("/admin");
    }

    const NavigateResources = () =>
    {
        navigate("/layout");
    }

    useEffect(() =>
    {
        if(currLocation === "/")
        {
            dropdownRef.current.style.display = "none";
            dropdownAdminRef.current.style.display = "none";
            homeRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/bookings-desk")
        {
            bookingsRef.current.style.color = "#09a2fb";
            dropdownRef.current.style.display = "block";
            dropdownAdminRef.current.style.display = "none";
            deskRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/bookings-meetingroom")
        {
            bookingsRef.current.style.color = "#09a2fb";
            dropdownRef.current.style.display = "block";
            dropdownAdminRef.current.style.display = "none";
            meetingRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/calendar")
        {
            dropdownRef.current.style.display = "none";
            dropdownAdminRef.current.style.display = "none";
            calendarRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/statistics")
        {
            dropdownRef.current.style.display = "none";
            dropdownAdminRef.current.style.display = "none";
            statisticsRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/admin")
        {
            dropdownRef.current.style.display = "none";
            dropdownAdminRef.current.style.display = "block";
            adminRef.current.style.color = "#09a2fb";
            teamRef.current.style.color = "#09a2fb";
        }

        if(currLocation === "/layout")
        {
            dropdownRef.current.style.display = "none";
            dropdownAdminRef.current.style.display = "block";
            adminRef.current.style.color = "#09a2fb";
            resourceRef.current.style.color = "#09a2fb";
        }

    },[currLocation])

    useEffect(() =>
    {
        setCurrLocation(window.location.pathname);
    },[])

    return (
        <div ref={ref} className={styles.navbarContainer}>
            <div className={styles.logoContainer}>
                S.H.W.M
            </div>
            <div className={styles.navlinkContainer}>
                <div ref={homeRef} className={styles.navlink} onClick={NavigateHome}>
                    <FaMap />
                    &nbsp;
                    Home
                </div>
                <div ref={bookingsRef} className={styles.navlink} onClick={ShowBookings}>
                    <FaTicketAlt />
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
                    <FaCalendar />
                    &nbsp;
                    Calendar
                </div>
                <div ref={statisticsRef} className={styles.navlink} onClick={NavigateStatistics}>
                    <FaChartPie />
                    &nbsp;
                    Statistics
                </div>
                <div ref={adminRef} className={styles.navlink} onClick={ShowAdminOptions}>
                    <FaUserShield />
                    &nbsp;
                    Admin
                </div>
                <div ref={dropdownAdminRef} className={styles.navlinkDropdownContainer}>
                    <div ref={teamRef} className={styles.navlinkDropdown} onClick={NavigateAdmin}>
                        Team Management
                    </div>
                    <div ref={resourceRef} className={styles.navlinkDropdown} onClick={NavigateResources}>
                        Office Creator
                    </div>
                </div>
            </div>
        </div>
    )
}

export default React.forwardRef(NavbarAdmin)
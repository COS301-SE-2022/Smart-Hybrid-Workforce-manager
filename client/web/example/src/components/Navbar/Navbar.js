import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"
import { FaCalendar, FaTicketAlt, FaMap, FaChartPie } from 'react-icons/fa'

const Navbar = (props, ref) =>
{
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endTime, setEndTime] = useState("");

    const navigate = useNavigate();

    return (
        <div ref={ref} className='navbar-container'>
            <div className='logo-container'>
                S.H.W.M
            </div>
            <div className='navlink-container'>
                <div className='navlink'>
                    <FaCalendar />
                    &nbsp;
                    Calendar
                </div>
                <div className='navlink'>
                    <FaTicketAlt />
                    &nbsp;
                    Bookings
                </div>
                <div className='navlink'>
                    <FaMap />
                    &nbsp;
                    Office Layout
                </div>
                <div className='navlink'>
                    <FaChartPie />
                    &nbsp;
                    Statistics
                </div>
            </div>
        </div>
    )
}

export default React.forwardRef(Navbar)
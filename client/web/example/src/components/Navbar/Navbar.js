import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"

const Navbar = (props, ref) =>
{
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endTime, setEndTime] = useState("");

    const navigate = useNavigate();

    return (
        <div ref={ref} className='navbar-container'>
            W
        </div>
    )
}

export default React.forwardRef(Navbar)
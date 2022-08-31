import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from "react-router-dom"

const DeskBooking = (props, ref) =>
{
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endTime, setEndTime] = useState("");

    const navigate = useNavigate();

    let handleSubmit = async (e) =>
    {
        e.preventDefault();
        try
        {
            let res = await fetch("http://localhost:8100/api/booking/create", 
            {
                method: "POST",
                body: JSON.stringify({
                    id: null,
                    user_id: window.sessionStorage.getItem("UserID"),
                    resource_type: "DESK",
                    resource_preference_id: null,
                    resource_id: null,
                    start: startDate + "T" + startTime + ":43.511Z",
                    end: startDate + "T" + endTime + ":43.511Z",
                    booked: false
                })
            });

            if(res.status === 200)
            {
                alert("Booking Successfully Created!");
                navigate("/");

                await fetch("http://localhost:8100/api/notification/send", 
                {
                    method: "POST",
                    body: JSON.stringify({
                        to: "archedevelop@gmail.com",
                        sDate: startDate,
                        sTime: startTime,
                        eDate: startDate,
                        eTime: endTime
                    })
                });
            }
        }
        catch(err)
        {
            console.log(err);
        }
    }; 

    return (
        <div ref={ref} className='form-container-desk-booking'>
            <div className='form-header'><h1>CREATE YOUR DESK BOOKING</h1>Please enter your booking details.</div>
            
            <Form className='form' onSubmit={handleSubmit}>
                <Form.Group className='form-group' controlId="formBasicName">
                <Form.Label className='form-label'>Date<br></br></Form.Label>
                <Form.Control className='form-input' type="date" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
                </Form.Group>

                <Form.Group className='form-group' controlId="formBasicName">
                <Form.Label className='form-label'>Start Time<br></br></Form.Label>
                <Form.Control className='form-input' type="time" placeholder="hh:mm" value={startTime} onChange={(e) => setStartTime(e.target.value)} />
                </Form.Group>
                
                <Form.Group className='form-group' controlId="formBasicName">
                <Form.Label className='form-label'>End Time<br></br></Form.Label>
                <Form.Control className='form-input' type="time" placeholder="hh:mm" value={endTime} onChange={(e) => setEndTime(e.target.value)} />
                </Form.Group>

                <Button className='button-submit' variant='primary' type='submit'>Create Booking</Button>
            </Form>
        </div>
    )
}

export default React.forwardRef(DeskBooking)
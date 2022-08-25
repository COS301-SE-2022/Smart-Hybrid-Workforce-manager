import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React, { useContext, useState, useRef, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { UserContext } from "../App"
import { useNavigate } from "react-router-dom"
import Background from "../components/Shapes/Background"

function BookingsDesk()
{
    const [startDate, setStartDate] = useState("");
    const [startTime, setStartTime] = useState("");
    const [endTime, setEndTime] = useState("");

    const {userData} = useContext(UserContext)
    const navigate = useNavigate();

    const circleRef = useRef(null);

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

                let res = await fetch("http://localhost:8100/api/notification/send", 
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

    
    useEffect(() =>
    {
        const Parallax = (e) =>
        {
            const speed = 10;
            const x = (window.innerWidth - e.pageX*speed)/100;
            const y = (window.innerHeight - e.pageY*speed)/100;

            circleRef.current.style.transform = `translateX(${x}px) translateY(${y}px)`;
        }

       //window.addEventListener('mousemove', Parallax);
    }, []);

    return (
        <div className='page-container'>
            <div className='content'>
                <Navbar />
                <div ref={circleRef} className="circle-container">
                    <Background />
                </div>
                <div className='form-container-desk-booking'>
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
            </div>
            <Footer />
        </div>
    )
}

export default BookingsDesk
import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import emailjs from '@emailjs/browser';
import '../App.css'

function BookingsDesk()
{
  const [startDate, setStartDate] = useState("");
  const [startTime, setStartTime] = useState("");
  const [endDate, setEndDate] = useState("");
  const [endTime, setEndTime] = useState("");

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
          user_id: "11111111-1111-4a06-9983-8b374586e459",
          resource_type: "DESK",
          resource_preference_id: null,
          resource_id: null,
          start: startDate + "T" + startTime + ":43.511Z",
          end: endDate + "T" + endTime + ":43.511Z",
          booked: false
        })
      });

      if(res.status === 200)
      {
        var data =
        {
          toEmail: "archecapstoneteam@gmail.com",
          sDate: startDate,
          sTime: startTime,
          eDate: endDate,
          eTime: endTime
        };

        emailjs.send('service_o88tkbb', 'template_xtvztfr', data, 'cKZXC1eO8lC78jvzV').then((result) =>
        {
          console.log(result.text);
          alert("Booking Successfully Created!");
          window.location.assign("./");
        }, (error) =>
        {
          console.log(error.text);
        });

        //alert("Booking Successfully Created!");
        //window.location.assign("./");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>CREATE YOUR DESK BOOKING</h1>Please enter your booking details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Start Date<br></br></Form.Label>
              <Form.Control name="sDate" className='form-input' type="text" placeholder="yyyy-mm-dd" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Start Time<br></br></Form.Label>
              <Form.Control name="sTime" className='form-input' type="text" placeholder="hh:mm" value={startTime} onChange={(e) => setStartTime(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>End Date<br></br></Form.Label>
              <Form.Control name="eDate" className='form-input' type="text" placeholder="yyyy-mm-dd" value={endDate} onChange={(e) => setEndDate(e.target.value)} />
            </Form.Group>
            
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>End Time<br></br></Form.Label>
              <Form.Control name="eTime" className='form-input' type="text" placeholder="hh:mm" value={endTime} onChange={(e) => setEndTime(e.target.value)} />
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
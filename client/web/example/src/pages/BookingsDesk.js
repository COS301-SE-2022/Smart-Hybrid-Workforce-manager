import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function BookingsDesk()
{
  const [startDate, setStartDate] = useState("");
  const [startTime, setStartTime] = useState("");
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
          end: startDate + "T" + endTime + ":43.511Z",
          booked: false
        })
      });

      if(res.status === 200)
      {
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

        if(res.status === 200)
        {
          alert("Booking Successfully Created!");
          window.location.assign("./");
        }
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
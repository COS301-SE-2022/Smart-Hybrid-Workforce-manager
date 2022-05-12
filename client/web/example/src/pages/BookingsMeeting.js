import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function BookingsMeeting()
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
          id: "33333333-dc08-4a06-9983-8b374586e453",
          user_id: "11111111-dc08-4a06-9983-8b374586e459",
          resource_type: "MEETINGROOM",
          resource_preference_id: null,
          start: startDate + "T" + startTime + ":43.511Z",
          end: endDate + "T" + endTime + ":43.511Z",
          booked: false
        })
      });

      if(res.status === 200)
      {
        alert("Booking Successfully Created!");
        window.location.reload();
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
          <p className='form-header'><h1>MEETING ROOM BOOKING</h1>Please enter your booking details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Start Date<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="yyyy-mm-dd" value={startDate} onChange={(e) => setStartDate(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Start Time<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="hh:mm" value={startTime} onChange={(e) => setStartTime(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>End Date<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="yyyy-mm-dd" value={endDate} onChange={(e) => setEndDate(e.target.value)} />
            </Form.Group>
            
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>End Time<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="hh:mm" value={endTime} onChange={(e) => setEndTime(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Booking</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default BookingsMeeting
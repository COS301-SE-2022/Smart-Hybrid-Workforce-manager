import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { UserContext } from "../App"
import { useNavigate } from 'react-router-dom';

function BookingsDeskEdit()
{
  const [startDate, setStartDate] = useState("");
  const [startTime, setStartTime] = useState("");
  const [endDate, setEndDate] = useState("");
  const [endTime, setEndTime] = useState("");

  const {userData} = UserContext(UserContext);
  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/booking/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("BookingID"),
          user_id: userData.user_id,
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
        let res = await fetch("http://localhost:8080/api/notification/send", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            to: "archedevelop@gmail.com",
            sDate: startDate,
            sTime: startTime,
            eDate: endDate,
            eTime: endTime
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        });

        if(res.status === 200)
        {
          alert("Booking Successfully Edited!");
          navigate("/");
        }
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    setStartDate(window.sessionStorage.getItem("StartDate"));
    setStartTime(window.sessionStorage.getItem("StartTime"));
    setEndDate(window.sessionStorage.getItem("EndDate"));
    setEndTime(window.sessionStorage.getItem("EndTime"));
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT YOUR DESK BOOKING</h1>Please enter your new booking details.</p>
          
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

            <Button className='button-submit' variant='primary' type='submit'>Edit Booking</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default BookingsDeskEdit
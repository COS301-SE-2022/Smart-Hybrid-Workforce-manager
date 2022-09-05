import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';

const EditMeetingRoom = () =>
{
  const [meetingRoomName, setMeetingRoomName] = useState("");
  const [meetingRoomLocation, setMeetingRoomLocation] = useState("");
  const [meetingRoomCapacity, setMeetingRoomCapacity] = useState("");

  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/resource/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("MeetingRoomID"),
          room_id: window.sessionStorage.getItem("RoomID"),
          name: meetingRoomName,
          location: meetingRoomLocation,
          role_id: null,
          resource_type: 'MEETINGROOM'
        })
      });

      if(res.status === 200)
      {
        alert("MeetingRoom Successfully Updated!");
        navigate("/resources");
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
    setMeetingRoomName(window.sessionStorage.getItem("MeetingRoomName"));
    setMeetingRoomLocation(window.sessionStorage.getItem("MeetingRoomLocation"));
    setMeetingRoomCapacity(window.sessionStorage.getItem("MeetingRoomCapacity"));
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT MEETING ROOM</h1>Please update the meeting room details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={meetingRoomName} onChange={(e) => setMeetingRoomName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Location<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Location" value={meetingRoomLocation} onChange={(e) => setMeetingRoomLocation(e.target.value)} />
            </Form.Group>
                      
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Meeting Room Capacity<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Capacity" value={meetingRoomCapacity} onChange={(e) => setMeetingRoomCapacity(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Meeting Room</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditMeetingRoom
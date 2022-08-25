import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';

const CreateRoom = () =>
{
  const [roomName, SetRoomName] = useState("");
  const [roomLocation, SetRoomLocation] = useState("");
  const [roomDimensions, SetRoomDimensions] = useState("");

  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/resource/room/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: null,
          building_id: window.sessionStorage.getItem("BuildingID"),
          name: roomName,
          location: roomLocation,
          dimension: roomDimensions
        })
      });

      if(res.status === 200)
      {
        alert("Room Successfully Created!");
        navigate("/resources");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  

  /*const RoomInput = () =>
  {
    let arr = [];
    for(let i = 0; i < numberOfRooms; i++)
    {
      arr.push(
        <Form.Group className='form-group' controlId="formBasicName">
          <Form.Label className='form-label'>{'Number of Desks in Room ' + (i+1)}<br></br></Form.Label>
          <Form.Control name={'room' + (i+1)} className='form-input' type="text" placeholder="Number" value={numberOfDesks} onChange={(e) => SetNumberOfDesks(e.target.value)} />
        </Form.Group>
      );
    }

    return arr;
  }*/

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>CREATE YOUR ROOM</h1>Please enter your room details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Name<br></br></Form.Label>
              <Form.Control name="rName" className='form-input' type="text" placeholder="Name" value={roomName} onChange={(e) => SetRoomName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Location<br></br></Form.Label>
              <Form.Control name="rLocation" className='form-input' type="text" placeholder="Location" value={roomLocation} onChange={(e) => SetRoomLocation(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Room Dimensions<br></br></Form.Label>
              <Form.Control name="rDimensions" className='form-input' type="text" placeholder="10x10" value={roomDimensions} onChange={(e) => SetRoomDimensions(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Room</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateRoom
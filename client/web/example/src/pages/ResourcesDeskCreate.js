import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';

const CreateDesk = () =>
{
  const [deskName, SetDeskName] = useState("");
  const [deskLocation, SetDeskLocation] = useState("");

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
          id: null,
          room_id: window.sessionStorage.getItem("RoomID"),
          name: deskName,
          location: deskLocation,
          role_id: null,
          resource_type: 'DESK',
          decorations: '{"computer": true}',
        })
      });

      if(res.status === 200)
      {
        alert("Desk Successfully Created!");
        navigate("/resources");
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
          <p className='form-header'><h1>CREATE YOUR DESK</h1>Please enter your desk details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Desk Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={deskName} onChange={(e) => SetDeskName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Desk Location<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Location" value={deskLocation} onChange={(e) => SetDeskLocation(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Desk</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateDesk
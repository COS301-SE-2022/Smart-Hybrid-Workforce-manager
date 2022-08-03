import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';

const EditDesk = () =>
{
  const [deskName, setDeskName] = useState("");
  const [deskLocation, setDeskLocation] = useState("");

  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/resource/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("DeskID"),
          room_id: window.sessionStorage.getItem("RoomID"),
          name: deskName,
          location: deskLocation,
          role_id: null,
          resource_type: 'DESK'
        })
      });

      if(res.status === 200)
      {
        alert("Desk Successfully Updated!");
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
    setDeskName(window.sessionStorage.getItem("DeskName"));
    setDeskLocation(window.sessionStorage.getItem("DeskLocation"));
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT DESK</h1>Please update the desk details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Desk Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={deskName} onChange={(e) => setDeskName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Desk Location<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="Location" value={deskLocation} onChange={(e) => setDeskLocation(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Desk</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditDesk
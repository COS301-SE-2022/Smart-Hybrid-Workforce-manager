import React, { useState } from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function Teams()
{
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [capacity, setCapacity] = useState("");
  const [picture, setPicture] = useState("");

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/team/create", 
      {
        method: "POST",
        body: JSON.stringify({
          name: name,
          description: description,
          capacity: parseInt(capacity),
          picture: picture
        })
      });

      if(res.status === 200)
      {
        alert("Team Successfully Created!");
        window.location.assign("./team");
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
          <p className='form-header'><h1>CREATE YOUR TEAM</h1>Please enter your team details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Team Name<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your team name" value={name} onChange={(e) => setName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Description<br></br></Form.Label>
              <Form.Control className='form-input-textarea' as="textarea" rows='5' placeholder="Enter your team description" value={description} onChange={(e) => setDescription(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicEmail">
              <Form.Label className='form-label'>Capacity<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your team capacity" value={capacity} onChange={(e) => setCapacity(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formFile">
              <Form.Label className='form-label'>Team Picture<br></br></Form.Label>
              <Form.Control className='form-input-file' type="file" value={picture} onChange={(e) => setPicture(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Team</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default Teams
import React, { useState } from 'react'
import Navbar from '../components/Navbar/Navbar.js'
import Footer from '../components/Footer'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { useNavigate } from 'react-router-dom'

function CreateUser()
{
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const navigate=useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/user/register", 
      {
        method: "POST",
        body: JSON.stringify({
          first_name: firstName,
          last_name: lastName,
          email: email,
          password: ""
        })
      });

      if(res.status === 200)
      {
        alert("User Successfully Created!");
        navigate("/users");
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
          <p className='form-header'><h1>CREATE USER</h1>Please enter user details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
                <Form.Label className='form-label'>First Name<br></br></Form.Label>
                <Form.Control className='form-input' type="text" placeholder="Enter your first name" value={firstName} onChange={(e) => setFirstName(e.target.value)} />
              </Form.Group>

              <Form.Group className='form-group' controlId="formBasicName">
                <Form.Label className='form-label'>Surname<br></br></Form.Label>
                <Form.Control className='form-input' type="text" placeholder="Enter your surname" value={lastName} onChange={(e) => setLastName(e.target.value)} />
              </Form.Group>

              <Form.Group className='form-group' controlId="formBasicEmail">
                <Form.Label className='form-label'>Email<br></br></Form.Label>
                <Form.Control className='form-input' type="text" placeholder="Enter your email" value={email} onChange={(e) => setEmail(e.target.value)} />
              </Form.Group>

              <Button className='button-submit' variant='primary' type='submit'>Create Account</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateUser
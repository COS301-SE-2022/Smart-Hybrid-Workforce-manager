import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom'

function Signup()
{
  const [firstName, setFirstName] = useState("");
  const [lastName, setLastName] = useState("");
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");

  const navigate=useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/user/register", 
      {
        method: "POST",
        body: JSON.stringify({
          id: "",
          first_name: firstName,
          last_name: lastName,
          email: email,
          password: password
        })
      });

      if(res.status === 200)
      {
        alert("Account Successfully Created!\nPlease verify your login details");
        navigate("/login");
        // window.location.assign("./bookings");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };  

  return (
    <div className='page-container'>
      <div className='content-login'>
        <div className='login-grid'>
          <div className='form-container-signup'>
            <p className='form-header'><h1>CREATE AN ACCOUNT</h1>Please enter your details.</p>
            
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

              <Form.Group className='form-group' controlId="formBasicPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control className='form-input' type="password" placeholder="Enter your password" value={password} onChange={(e) => setPassword(e.target.value)} />
              </Form.Group>

              <Button className='button-submit' variant='primary' type='submit'>Create Account</Button>
            </Form>
          </div>

          <div className='image-container'>
            <img className='login-image' src='https://i.pinimg.com/originals/3b/79/c7/3b79c7a4a275b5ee1dbb76731f9736b8.png' alt='office'></img>
          </div>
        </div>

      </div>
    </div>
  )
}

export default Signup
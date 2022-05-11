import Footer from "../components/Footer"
import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'
import Form from 'react-bootstrap/Form'

function Login()
{
  let navigate = useNavigate();
  const routeSignup = () =>
  {
    let path = "/signup";
    navigate(path);
  }

  return (
    <div>
      <div className='content'>
        <div className='form-container'>
          <p><h1>WELCOME BACK</h1>Please enter your details.</p>
          
          <Form className='form'>
            <Form.Group className='form-group' controlId="formBasicEmail">
              <Form.Label className='form-label'>Email<br></br></Form.Label>
              <Form.Control className='form-input' type="email" placeholder="Enter email" />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicPassword">
              <Form.Label className='form-label'>Password<br></br></Form.Label>
              <Form.Control className='form-input' type="password" placeholder="Enter password" />
            </Form.Group>
          </Form>
        </div>

      </div>
    </div>
  )
}

export default Login
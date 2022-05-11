import React from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'

function Login()
{
  return (
    <div>
      <div className='content'>
        <div className='login-grid'>
          <div className='form-container'>
            <p className='form-header'><h1>WELCOME BACK</h1>Please enter your details.</p>
            
            <Form className='form'>
              <Form.Group className='form-group' controlId="formBasicEmail">
                <Form.Label className='form-label'>Email<br></br></Form.Label>
                <Form.Control className='form-input' type="email" placeholder="Enter your email" />
              </Form.Group>

              <Form.Group className='form-group' controlId="formBasicPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control className='form-input' type="password" placeholder="Enter your password" />
              </Form.Group>

              <Button className='button-submit' variant='primary' type='submit'>Sign In</Button>
            </Form>
            <p className='signup-prompt'>Don't have an account? <a className='signup-link' href='/signup'>Sign up for free!</a></p>
          </div>

          <div className='image-container'>
            <img className='login-image' src='https://i.pinimg.com/originals/43/90/d7/4390d72e6a6cb6086c73e570bb6c439d.jpg' alt='office'></img>
          </div>
        </div>

      </div>
    </div>
  )
}

export default Login
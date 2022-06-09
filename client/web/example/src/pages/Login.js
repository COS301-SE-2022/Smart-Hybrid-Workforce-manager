import React from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'

function Login()
{
  // let handleSubmit = async (e) =>
  // {
  //   e.preventDefault();
  //   try
  //   {
  //     let res = await fetch("http://localhost:8100/api/booking/create", 
  //     {
  //       method: "POST",
  //       body: JSON.stringify({
  //         id: "33333333-dc08-4a06-9983-8b374586e453",
  //         user_id: "11111111-dc08-4a06-9983-8b374586e459",
  //         resource_type: "DESK",
  //         resource_preference_id: null,
  //         resource_id: null,
  //         start: startDate + "T" + startTime + ":43.511Z",
  //         end: endDate + "T" + endTime + ":43.511Z",
  //         booked: false
  //       })
  //     });

  //     if(res.status === 200)
  //     {
  //       alert("Booking Successfully Created!");
  //       window.location.reload();
  //     }
  //   }
  //   catch(err)
  //   {
  //     console.log(err);
  //   }
  // };
  
  return (
    <div className='page-container'>
      <div className='content-login'>
        <div className='login-grid'>
          <div className='form-container-login'>
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
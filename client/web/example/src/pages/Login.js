import React, { useContext, useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { UserContext } from '../App';
import { useLocation, useNavigate } from 'react-router-dom';

export default function Login()
{
  const [identifier, setIdentifier] = useState("");
  const [secret, setSecret] = useState("");
  const auth = sessionStorage.getItem("auth_data");
  const {setUserData}=useContext(UserContext)
  const navigate=useNavigate();
  const location=useLocation();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    console.log(JSON.stringify({
      identifier:identifier,
      secret:secret,
      id: null,
      type:null,
      active:null
    }))
    fetch("http://localhost:8080/api/user/login", 
    {
      method: "POST",
      body: JSON.stringify({
        identifier:identifier,
        secret:secret,
        id: null,
        type:null,
        active:null
      })
    }).then((res) => {
      if(res.status === 200){
        alert("Successfully Logged In!")
        return res.json();        
      }
      else{
        console.log(res)
        alert("Failed login");
      }        
    }).then((data) => {
      setUserData(data)
      localStorage.setItem("auth_data", data);
      if (location.state?.from) {
        navigate(location.state.from);
      }else{
        navigate("/");
      }
    }).catch((err) => {
      console.error(err);
    })
    
  };  

  return (
    <div className='page-container'>
      <div className='content-login'>
        <div className='login-grid'>
          <div className='form-container-login'>
            <p className='form-header'><h1>WELCOME BACK</h1>Please enter your details.</p>
            {auth===undefined?console.log("logged In"):console.log("not logged in")}
            <Form className='form' onSubmit={handleSubmit}>
              <Form.Group className='form-group' controlId="formBasicEmail">
                <Form.Label className='form-label'>Email<br></br></Form.Label>
                <Form.Control className='form-input' type="email" placeholder="Enter your email" value={identifier} onChange={(e) => setIdentifier(e.target.value)} />
              </Form.Group>

              <Form.Group className='form-group' controlId="formBasicPassword">
                <Form.Label className='form-label'>Password<br></br></Form.Label>
                <Form.Control className='form-input' type="password" placeholder="Enter your password" value={secret} onChange={(e) => setSecret(e.target.value)} />
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

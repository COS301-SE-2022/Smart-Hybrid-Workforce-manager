import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

const Teams = () => {
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>CREATE YOUR TEAM</h1>Please enter your team details.</p>
          
          <Form className='form'>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Team Name<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your team name" />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Description<br></br></Form.Label>
              <Form.Control className='form-input-textarea' as="textarea" rows='5' placeholder="Enter your team description" />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicEmail">
              <Form.Label className='form-label'>Capacity<br></br></Form.Label>
              <Form.Control className='form-input' type="text" placeholder="Enter your email" />
            </Form.Group>

            <Form.Group className='form-group' controlId="formFile">
              <Form.Label className='form-label'>Team Picture<br></br></Form.Label>
              <Form.Control className='form-input-file' type="file" />
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
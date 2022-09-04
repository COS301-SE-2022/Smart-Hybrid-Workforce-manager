import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import React, { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { useNavigate } from 'react-router-dom';

const CreateBuilding = () =>
{
  const [buildingName, SetBuildingName] = useState("");
  const [buildingLocation, SetBuildingLocation] = useState("");
  const [buildingDimensions, SetBuildingDimensions] = useState("");

  const navigate = useNavigate();

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/resource/building/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: null,
          name: buildingName,
          location: buildingLocation,
          dimension: buildingDimensions
        })
      });

      if(res.status === 200)
      {
        alert("Building Successfully Created!");
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
          <p className='form-header'><h1>CREATE YOUR BUILDING</h1>Please enter your building details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Name<br></br></Form.Label>
              <Form.Control name="bName" className='form-input' type="text" placeholder="Name" value={buildingName} onChange={(e) => SetBuildingName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Location<br></br></Form.Label>
              <Form.Control name="bLocation" className='form-input' type="text" placeholder="Location" value={buildingLocation} onChange={(e) => SetBuildingLocation(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Dimensions<br></br></Form.Label>
              <Form.Control name="bDimensions" className='form-input' type="text" placeholder="10x10" value={buildingDimensions} onChange={(e) => SetBuildingDimensions(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Building</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateBuilding
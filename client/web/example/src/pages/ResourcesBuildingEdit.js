import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import { useNavigate } from 'react-router-dom';

const BuildingEdit = () =>
{
  const [buildingName, setBuildingName] = useState("");
  const [buildingLocation, setBuildingLocation] = useState("");
  const [buildingDimensions, setBuildingDimensions] = useState("");

  const navigate = useNavigate();

  //POST request
  const FetchBuilding = () =>
  {
    fetch("http://localhost:8080/api/resource/building/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id: window.sessionStorage.getItem("BuildingID")
          })
        }).then((res) => res.json()).then(data => 
          {
            setBuildingName(data[0].name);
            setBuildingLocation(data[0].location);
            setBuildingDimensions(data[0].dimension);
          });
  }

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8080/api/resource/building/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("BuildingID"),
          name: buildingName,
          location: buildingLocation,
          dimension: buildingDimensions
        })
      });

      if(res.status === 200)
      {
        alert("Building Successfully Updated!");
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
    FetchBuilding();
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT BUILDING</h1>Please update the building details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Name<br></br></Form.Label>
              <Form.Control name="bName" className='form-input' type="text" placeholder="Name" value={buildingName} onChange={(e) => setBuildingName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Location<br></br></Form.Label>
              <Form.Control name="bLocation" className='form-input' type="text" placeholder="Location" value={buildingLocation} onChange={(e) => setBuildingLocation(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Building Dimensions<br></br></Form.Label>
              <Form.Control name="bDimensions" className='form-input' type="text" placeholder="10x10" value={buildingDimensions} onChange={(e) => setBuildingDimensions(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Building</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default BuildingEdit
import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'

const CreateRole = () =>
{
  const [roleName, setRoleName] = useState("");
  const [roleColor, setRoleColor] = useState("");

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/role/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: null,
          role_name: roleName
        })
      });

      if(res.status === 200)
      {
        alert("Role Successfully Created!");
        window.location.assign("./role");
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
          <p className='form-header'><h1>CREATE ROLE</h1>Please enter role details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Role Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={roleName} onChange={(e) => setRoleName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Role Color<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="#111111" value={roleColor} onChange={(e) => setRoleColor(e.target.value)} />
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Create Role</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default CreateRole
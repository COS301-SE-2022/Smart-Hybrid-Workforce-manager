import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import Button from 'react-bootstrap/Button'
import { useState, useEffect } from 'react';
import RoleListItem from '../components/Role/RoleListItem';

function Roles()
{
  const [roles, SetRoles] = useState([])

  //POST request
  const FetchRoles = () =>
  {
    fetch("http://localhost:8100/api/role/information", 
        {
          method: "POST",
          body: JSON.stringify({
          })
        }).then((res) => res.json()).then(data => 
          {
            SetRoles(data);
          });
  }

  const AddRole = () =>
  {
    window.location.assign("./role-create");
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchRoles()
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='resources-map'>
          {roles.length > 0 && (
            roles.map(role => 
            {
              return <RoleListItem id={role.id} name={'Role ' + role.role_name} lead={ role.role_lead_id } />
            }
          )
          )}
        </div>

        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddRole}>Add Role</Button>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Roles
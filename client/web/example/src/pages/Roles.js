import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react';
import RoleListItem from '../components/Role/RoleListItem';

function Roles()
{
  const [roles, SetRoles] = useState([])
  const [currRole, SetCurrRole] = useState("")

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
              return <RoleListItem id={role.id} name={role.role_name}/>
            }
          )
          )}
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Roles
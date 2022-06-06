import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react';
import RoleListItem from '../components/Role/RoleListItem';

function Users()
{
  const [users, SetUsers] = useState([])

  //POST request
  const FetchUsers = () =>
  {
    fetch("http://localhost:8100/api/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
          })
        }).then((res) => res.json()).then(data => 
          {
            SetUsers(data);
          });
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchUsers()
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='resources-map'>
          {users.length > 0 && (
            users.map(user => 
            {
              return <RoleListItem id={user.id} name={user.first_name + " " + user.last_name}/>
            }
          )
          )}
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Users
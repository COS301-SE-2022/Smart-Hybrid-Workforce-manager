import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react';
import UserListItem from '../components/Profile/UserListItem';
import Button from 'react-bootstrap/Button'

function Users()
{
  const [users, setUsers] = useState([])

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
            setUsers(data);
          });
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    FetchUsers()
  }, [])

  const AddUser = () =>
  {
    window.location.assign("./user-create");
  }

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='resources-map'>
          {users.length > 0 && (
            users.map(user => 
            {
              return <UserListItem id={user.id} name={user.first_name + " " + user.last_name} email = {user.email}/>
            }
          )
          )}
        </div>
        <div className='button-resource-container'>
          <Button className='button-resource' variant='primary' onClick={AddUser}>Add User</Button>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Users
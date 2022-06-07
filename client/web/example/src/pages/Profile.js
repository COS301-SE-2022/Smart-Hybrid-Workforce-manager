import { useState, useEffect } from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import '../App.css'
import Button from 'react-bootstrap/Button'
import RoleUserList from '../components/Role/RoleUserList'
import TeamUserList from '../components/Team/TeamUserList'

function Profile()
{
  const [firstName, SetFirstName] = useState("")
  const [lastName, SetLastName] = useState("")
  const [email, SetEmail] = useState("")

  const [roles, SetRoles] = useState([])
  const [teams, SetTeams] = useState([])
  
  //POST request
  const FetchUser = () =>
  {
    fetch("http://localhost:8100/api/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id:"11111111-1111-4a06-9983-8b374586e459"
          })
        }).then((res) => res.json()).then(data => 
          {
            SetFirstName(data[0].first_name);
            SetLastName(data[0].last_name);
            SetEmail(data[0].email);
          });
  }

  //POST request
  const FetchUserRoles = () =>
  {
    fetch("http://localhost:8100/api/role/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id:"11111111-1111-4a06-9983-8b374586e459"
          })
        }).then((res) => res.json()).then(data => 
          {
            SetRoles(data);
          });
  }

  //POST request
  const FetchUserTeams = () =>
  {
    fetch("http://localhost:8100/api/team/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id:"11111111-1111-4a06-9983-8b374586e459"
          })
        }).then((res) => res.json()).then(data => 
          {
            SetTeams(data);
          });
  }

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    FetchUser();
    FetchUserRoles();
    FetchUserTeams();
  }, [])

  const ProfileConfiguration = () =>
  {
    window.location.assign("./profile-configuration");
  }

  const LogOut = () =>
  {
    window.location.assign("./signup");
  }

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='user-container'>
          <div className="user">
            <div className="user-image"></div>
            <div className="user-text">
              <p className="user-text-name">{firstName + " " + lastName}</p>
              <p className="user-text-email">{email}</p>              
            </div>
            <div className="user-roles">
              {roles.length > 0 && (
                roles.map(role => (
                  <RoleUserList id={role.role_id}/>
                  
                ))
              )}
            </div>
            <div className="user-teams">
              {teams.length > 0 && (
                teams.map(team => (
                  <TeamUserList id={team.team_id} />
                ))
              )}
            </div>
            <Button className='button-user-profile' variant='primary' onClick={ProfileConfiguration}>Profile Configuration</Button>
            <Button className='button-user-profile' variant='primary' onClick={LogOut}>Log Out</Button>
          </div>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default Profile
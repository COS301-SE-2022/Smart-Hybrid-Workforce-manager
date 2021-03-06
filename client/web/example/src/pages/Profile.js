import { useState, useEffect, useContext } from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import '../App.css'
import Button from 'react-bootstrap/Button'
import RoleUserList from '../components/Role/RoleUserList'
import TeamUserList from '../components/Team/TeamUserList'
import { UserContext } from '../App'
import { useNavigate } from 'react-router-dom'

function Profile()
{
  const [identifier, SetIdentifier] = useState("")
  const [firstName, SetFirstName] = useState("")
  const [lastName, SetLastName] = useState("")
  const [email, SetEmail] = useState("")
  const [picture, SetPicture] = useState("")
  const [dateCreated, SetDateCreated] = useState("")
  const [workFromHome, SetWorkFromHome] = useState("")
  const [parking, SetParking] = useState("")
  const [officeDays, SetOfficeDays] = useState("")
  const [startTime, SetStartTime] = useState("")
  const [endTime, SetEndTime] = useState("")

  const [roles, SetRoles] = useState([])
  const [teams, SetTeams] = useState([])
  
  const {userData}=useContext(UserContext)

  const navigate=useNavigate();

  //POST request
  const FetchUser = () =>
  {
    fetch("http://localhost:8100/api/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            identifier:userData.user_id.substring(6)
          })
        }).then((res) => res.json()).then(data => 
        {
          sessionStorage.setItem("UserID",data[0].id)
          SetIdentifier(data[0].identifier)
          SetFirstName(data[0].first_name);
          SetLastName(data[0].last_name);
          SetEmail(data[0].email);
          SetPicture(data[0].picture);
          SetDateCreated(data[0].date_created);
          SetWorkFromHome(data[0].work_from_home);
          SetParking(data[0].parking);
          SetOfficeDays(data[0].office_days);
          SetStartTime(data[0].preferred_start_time);
          SetEndTime(data[0].preferred_end_time);
        });
  }

  console.log(window.sessionStorage.getItem("UserID"))
  console.log(userData)
  console.log(userData.user_id.substring(6))
  //POST request
  const FetchUserRoles = () =>
  {
    fetch("http://localhost:8100/api/role/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            identifier:userData.user_id.substring(6)
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
            identifier:userData.user_id.substring(6)
          })
        }).then((res) => res.json()).then(data => 
          {
            SetTeams(data);
          });
  }

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    // window.sessionStorage.setItem("UserID", "11111111-1111-4a06-9983-8b374586e459");
    FetchUser();
    FetchUserRoles();
    FetchUserTeams();
  }, [])

  const ProfileConfiguration = () =>
  {
    // window.sessionStorage.setItem("UserID", "11111111-1111-4a06-9983-8b374586e459");
    window.sessionStorage.setItem("Identifier", identifier);
    window.sessionStorage.setItem("FirstName", firstName);
    window.sessionStorage.setItem("LastName", lastName);
    window.sessionStorage.setItem("Email", email);
    window.sessionStorage.setItem("Picture", picture);
    window.sessionStorage.setItem("DateCreated", dateCreated);
    window.sessionStorage.setItem("WorkFromHome", workFromHome);
    window.sessionStorage.setItem("Parking", parking);
    window.sessionStorage.setItem("OfficeDays", officeDays);
    window.sessionStorage.setItem("StartTime", startTime);
    window.sessionStorage.setItem("EndTime", endTime);
    navigate("/profile-configuration")
    // window.location.assign("./profile-configuration");
  }

  const LogOut = () =>
  {

    window.location.assign("./login");
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
              <h3>Roles</h3>
              <div className='list'>
                {roles.length > 0 && (
                  roles.map(role => (
                    <RoleUserList id={role.role_id}/>
                    
                  ))
                )}
              </div>
            </div>
            <div className="user-teams">
              <h3>Teams</h3>
              <div className='list'>
                {teams.length > 0 && (
                  teams.map(team => (
                    <TeamUserList teamId={team.team_id} />
                  ))
                )}
              </div>
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
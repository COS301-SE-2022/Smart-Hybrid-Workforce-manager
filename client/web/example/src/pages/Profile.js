import { useState, useEffect, useContext } from 'react'
import Navbar from '../components/Navbar/Navbar.js'
import Footer from '../components/Footer'
import '../App.css'
import Button from 'react-bootstrap/Button'
import RoleUserList from '../components/Role/RoleUserList'
import TeamUserList from '../components/Team/TeamUserList'
import { UserContext } from '../App'
import { useNavigate } from 'react-router-dom'
import ProfileBar from '../components/Navbar/ProfileBar';
import { FaWheelchair, FaHouseUser } from 'react-icons/fa';
import NavbarAdmin from '../components/Navbar/NavbarAdmin.js'

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
  
  const {userData,setUserData}=useContext(UserContext)

  const navigate = useNavigate();

  /*console.log(window.sessionStorage.getItem("UserID"))
  console.log(userData)
  console.log(userData.user_id.substring(6))*/
  

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
    useEffect(() =>
    {
        //POST requests
        const FetchUser = () =>
        {
            fetch("http://localhost:8080/api/user/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    identifier : userData.user_identifier
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) => res.json()).then(data => 
            {
                console.log(data);
                sessionStorage.setItem("UserID", data[0].id);
                SetIdentifier(data[0].identifier);
                SetFirstName(data[0].first_name);
                SetLastName(data[0].last_name);
                SetEmail(data[0].email);
                SetPicture(data[0].picture);
                SetDateCreated(data[0].date_created);
                SetWorkFromHome(data[0].work_from_home);
                SetParking(data[0].parking);
                SetOfficeDays(data[0].office_days);
                SetStartTime(data[0].preferred_start_time.substring(11,16));
                SetEndTime(data[0].preferred_end_time.substring(11,16));
            });
        };

        const FetchUserRoles = () =>
        {
            fetch("http://localhost:8080/api/role/user/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    identifier: userData.user_identifier
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) => res.json()).then(data => 
            {
                SetRoles(data);
            });
        };

        const FetchUserTeams = () =>
        {
            fetch("http://localhost:8080/api/team/user/information", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    identifier: userData.user_identifier
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
            }).then((res) => res.json()).then(data => 
            {
                SetTeams(data);
            });
        }

        // window.sessionStorage.setItem("UserID", "11111111-1111-4a06-9983-8b374586e459");
        FetchUser();
        FetchUserRoles();
        FetchUserTeams();
  }, [userData])

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
  }

    const renderWheelchair = () =>
    {
        if(parking === 'DISABLED')
        {
            return <FaWheelchair />
        }
    }

    const renderHome = () =>
    {
        if(workFromHome === 'true')
        {
            return <FaHouseUser />
        }
    }

    const showNavbar = () =>
    {
        if(!userData.user_identifier.includes("admin"))
        {
            return <Navbar />;
        }
        else
        {
            return <NavbarAdmin />;
        }
    };

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                {showNavbar()}
                <div className='main-container'>
                    <div className='profile-container'>
                        
                        <div className='personal-information-container'>
                            <div className='profile-name'>{firstName} &nbsp; {lastName}</div>
                            <div className='profile-email'>{email}</div>
                            <div className='profile-icons'>
                                {renderWheelchair()}
                                {renderHome()}
                            </div>
                        </div>

                        <div className='preferences-container'>
                            <div className='profile-days'>{officeDays} office days per week</div>
                        <div className='profile-time'>Preferred times: {startTime} - {endTime}</div>
                            <div className='profile-roles'>
                                {roles.length > 0 && (
                                    roles.map(role =>
                                    (
                                        role.role_id 
                                    ))
                                )}
                            </div>
                        </div>

                        <div className='profile-teams-container'>
                            <div className='profile-teams-title'>Teams</div>
                            <div className='profile-teams-carousel'>
                                {teams.length > 0 && (
                                    teams.map(team =>
                                    (
                                        <div className='profile-team'>{team.team_id}</div>
                                    ))
                                )}
                            </div>
                        </div>

                        <div className="profile-image-container"></div>

                    </div>
                </div>
            </div>
        </div>
    )
}

export default Profile
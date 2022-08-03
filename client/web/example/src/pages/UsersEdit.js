import { useState, useEffect, useContext } from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import '../App.css'
import RoleUserList from '../components/Role/RoleUserList'
import TeamUserList from '../components/Team/TeamUserList'
import Button from 'react-bootstrap/Button'
import { UserContext } from '../App'
import { useNavigate } from 'react-router-dom'

function EditUser()
{
    const [name, SetName] = useState("")
    const [email, SetEmail] = useState("")

    const [roles, SetRoles] = useState([])
    const [teams, SetTeams] = useState([])

    const [currRole, SetCurrRole] = useState("")
    const [currTeam, SetCurrTeam] = useState("")

    const [userRoles, SetUserRoles] = useState([])
    const [userTeams, SetUserTeams] = useState([])

    const {userData} = useContext(UserContext)
    const navigate = useNavigate();

  //POST request
  const FetchUserRoles = () =>
  {
    fetch("http://localhost:8100/api/role/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            user_id: window.sessionStorage.getItem("UserID")
          })
        }).then((res) => res.json()).then(data => 
          {
            SetUserRoles(data);
          });
  }

  //POST request
  const FetchUserTeams = () =>
  {
    fetch("http://localhost:8100/api/team/user/information", 
        {
          method: "POST",
          body: JSON.stringify({
            user_id: window.sessionStorage.getItem("UserID")
          })
        }).then((res) => res.json()).then(data => 
          {
            SetUserTeams(data);
          });
    }

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

    const UpdateRole = (e) =>
    {
        SetCurrRole(e.target.value);
    }
    
    let AddRole = async (e) =>
    {
        e.preventDefault();
        try
        {
            let res = await fetch("http://localhost:8100/api/role/user/create", 
            {
                method: "POST",
                body: JSON.stringify({
                    role_id: currRole,
                    user_id: window.sessionStorage.getItem("UserID")
                })
            });

            if(res.status === 200)
            {
              alert("Role Successfully Added!");
              navigate(0);
            }
        }
        catch (err)
        {
            console.log(err);
        }
    }

    //POST request
    const FetchTeams = () =>
    {
        fetch("http://localhost:8100/api/team/information", 
            {
            method: "POST",
            body: JSON.stringify({
            })
            }).then((res) => res.json()).then(data => 
            {
                SetTeams(data);
            });
    }

    const UpdateTeam = (e) =>
    {
        SetCurrTeam(e.target.value);
    }
    
    let AddTeam = async (e) =>
    {
        e.preventDefault();
        try
        {
            let res = await fetch("http://localhost:8100/api/team/user/create", 
            {
                method: "POST",
                body: JSON.stringify({
                    team_id: currTeam,
                    user_id: window.sessionStorage.getItem("UserID")
                })
            });

            if(res.status === 200)
            {
                alert("Team Successfully Added!");
                navigate(0);
            }
        }
        catch (err)
        {
            console.log(err);
        }
    }
  
  const PermissionConfiguration = () =>
  {
    navigate("/user-permissions");
  }

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    SetName(window.sessionStorage.getItem("UserName"));
    SetEmail(window.sessionStorage.getItem("UserEmail"));
    
    FetchRoles();
    FetchTeams();
    FetchUserRoles();
    FetchUserTeams();
  }, [])

    
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='user-container'>
          <div className="user">
            <div className="user-image"></div>
            <div className="user-text">
              <p className="user-text-name">{name}</p>
              <p className="user-text-email">{email}</p>              
            </div>
            <div className="user-roles">
                <h3>Roles</h3>
                <select className='combo-box-profile' name='role' onChange={UpdateRole.bind(this)}>
                <option value='' disabled selected id='RoleDefault'>--Select a Role--</option>
                {roles.length > 0 && (
                    roles.map(role => (
                        <option value={role.id}>{role.role_name}</option>
                    ))
                )}
                </select>
                <Button className='button-profile-add' variant='primary' onClick={AddRole}>Add Role</Button>
                <div className='list'>
                    {userRoles.length > 0 && (
                        userRoles.map(userRole => (
                            <RoleUserList id={userRole.role_id}/>
                        ))
                    )}
                </div>
            </div>
            <div className="user-teams">
                <h3>Teams</h3>
                <select className='combo-box-profile' name='role' onChange={UpdateTeam.bind(this)}>
                <option value='' disabled selected id='TeamDefault'>--Select a Team--</option>
                {teams.length > 0 && (
                    teams.map(team => (
                        <option value={team.id}>{team.name}</option>
                    ))
                )}
                </select>
                <Button className='button-profile-add' variant='primary' onClick={AddTeam}>Add Team</Button>
              <div className='list'>
                {userTeams.length > 0 && (
                  userTeams.map(userTeam => (
                    <TeamUserList teamId={userTeam.team_id} />
                  ))
                )}
              </div>
            </div>
            <div style={{
                        'width': 'inherit',
                        'text-align': 'center',
                        }}>
                <Button className='button-submit' variant='primary' onClick={PermissionConfiguration}>Configure Permissions</Button>
            </div>              
          </div>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditUser
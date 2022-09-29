import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import UserRoleList from '../components/Role/UserRoleList'
import RoleLeadOption from '../components/Role/RoleLeadOption'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../App.js'

const EditRole = () =>
{
  const [roleName, setRoleName] = useState("");
  const [roleColor, setRoleColor] = useState("");
  const [roleLead, setRoleLead] = useState(window.sessionStorage.getItem("RoleLead"));

  const [roleUsers, SetRoleUsers] = useState([]);
  const navigate = useNavigate();

  const {userData} = useContext(UserContext);

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://deskflow.co.za:8080/api/role/create", 
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
          id: window.sessionStorage.getItem("RoleID"),
          role_name: roleName,
          role_lead_id: roleLead === "null" ? null : roleLead
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      });

      if(res.status === 200)
      {
        alert("Role Successfully Updated!");
        navigate("/role");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

    //POST request
  const FetchRoleUsers = () =>
  {
    fetch("http://deskflow.co.za:8080/api/role/user/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            role_id:window.sessionStorage.getItem("RoleID")
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        }).then((res) => res.json()).then(data => 
        {
          SetRoleUsers(data);
        });
  }

  const PermissionConfiguration = () =>
  {
    navigate("/role-permissions");
  }

  //Using useEffect hook. This will ste the default values of the form once the components are mounted
  useEffect(() =>
  {
    setRoleName(window.sessionStorage.getItem("RoleName").substring(5, window.sessionStorage.getItem("RoleName").length));
    setRoleColor(window.sessionStorage.getItem("RoleColor"));
    setRoleLead(window.sessionStorage.getItem("RoleLead"));

    FetchRoleUsers();
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>EDIT ROLE</h1>Please update the role details.</p>
          
          <Form className='form' onSubmit={handleSubmit}>
            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Role Name<br></br></Form.Label>
              <Form.Control name="dName" className='form-input' type="text" placeholder="Name" value={roleName} onChange={(e) => setRoleName(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <Form.Label className='form-label'>Role Color<br></br></Form.Label>
              <Form.Control name="dLocation" className='form-input' type="text" placeholder="#111111" value={roleColor} onChange={(e) => setRoleColor(e.target.value)} />
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicRoleLead">
              <Form.Label className='form-label'>Role Lead<br></br></Form.Label>
              <select className='combo-box' name='rolelead' value={roleLead} onChange={(e) => setRoleLead(e.target.value)}>
                <option value="null">--none--</option>
                {roleUsers.length > 0 && (
                  roleUsers.map(roleUser => (
                    <RoleLeadOption id={roleUser.user_id} roleLeadId={roleLead} />
                  ))
                )}
              </select>
            </Form.Group>

             <Form.Group className='form-group' controlId="formRoleMembers">
              <Form.Label className='form-label'>Role Members<br></br></Form.Label>
              {roleUsers.length > 0 && (
                  roleUsers.map(roleUser => (
                    <UserRoleList id={roleUser.user_id} />
                  ))
                )}
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Role</Button>
            <Button className='button-submit' variant='primary' onClick={PermissionConfiguration}>Configure Permissions</Button>
          </Form>
        </div>
      </div>
      <Footer />
    </div>
  )
}

export default EditRole
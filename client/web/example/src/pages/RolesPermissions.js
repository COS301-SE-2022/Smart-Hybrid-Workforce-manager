import Navbar from '../components/Navbar/Navbar.js'
import Footer from "../components/Footer"
import { useState, useEffect, useCallback, useContext } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../App.js'

function RolePermissions() {
  const [roleName, setRoleName] = useState(window.sessionStorage.getItem("RoleName").substring(5, window.sessionStorage.getItem("RoleName").length));

  const {userData} = useContext(UserContext);

  // Bookings
  const [createBookingIdentifier, SetCreateBookingIdentifier] = useState("") // allows users of a role to update the Booking for everyone
  const [createBookingIdentifierId, SetCreateBookingIdentifierId] = useState("")
  const [viewBookingIdentifier, SetViewBookingIdentifier] = useState("") // allows users of a role to view the Booking for everyone
  const [viewBookingIdentifierId, SetViewBookingIdentifierId] = useState("")
  const [deleteBookingIdentifier, SetDeleteBookingIdentifier] = useState("") // allows users of a role to delete the Booking for everyone
  const [deleteBookingIdentifierId, SetDeleteBookingIdentifierId] = useState("")

  // Permissions
  const [createPermissionIdentifier, SetCreatePermissionIdentifier] = useState("") // allows users of a role to update the Permission for everyone
  const [createPermissionIdentifierId, SetCreatePermissionIdentifierId] = useState("")
  const [viewPermissionIdentifier, SetViewPermissionIdentifier] = useState("") // allows users of a role to view the Permission for everyone
  const [viewPermissionIdentifierId, SetViewPermissionIdentifierId] = useState("")
  const [deletePermissionIdentifier, SetDeletePermissionIdentifier] = useState("") // allows users of a role to delete the Permission for everyone
  const [deletePermissionIdentifierId, SetDeletePermissionIdentifierId] = useState("")

  // Roles
  const [createRoleIdentifier, SetCreateRoleIdentifier] = useState("") // allows users of a role to update the Role for everyone
  const [createRoleIdentifierId, SetCreateRoleIdentifierId] = useState("")
  const [viewRoleIdentifier, SetViewRoleIdentifier] = useState("") // allows users of a role to view the Role for everyone
  const [viewRoleIdentifierId, SetViewRoleIdentifierId] = useState("")
  const [deleteRoleIdentifier, SetDeleteRoleIdentifier] = useState("") // allows users of a role to delete the Role for everyone
  const [deleteRoleIdentifierId, SetDeleteRoleIdentifierId] = useState("")

  // Teams
  const [createTeamIdentifier, SetCreateTeamIdentifier] = useState("") // allows users of a role to update the Team for everyone
  const [createTeamIdentifierId, SetCreateTeamIdentifierId] = useState("")
  const [viewTeamIdentifier, SetViewTeamIdentifier] = useState("") // allows users of a role to view the Team for everyone
  const [viewTeamIdentifierId, SetViewTeamIdentifierId] = useState("")
  const [deleteTeamIdentifier, SetDeleteTeamIdentifier] = useState("") // allows users of a role to delete the Team for everyone
  const [deleteTeamIdentifierId, SetDeleteTeamIdentifierId] = useState("")

  // Resource
  const [createResourceIdentifier, SetCreateResourceIdentifier] = useState("") // allows users of a role to update the Resource for everyone
  const [createResourceIdentifierId, SetCreateResourceIdentifierId] = useState("")
  const [viewResourceIdentifier, SetViewResourceIdentifier] = useState("") // allows users of a role to view the Resource for everyone
  const [viewResourceIdentifierId, SetViewResourceIdentifierId] = useState("")
  const [deleteResourceIdentifier, SetDeleteResourceIdentifier] = useState("") // allows users of a role to delete the Resource for everyone
  const [deleteResourceIdentifierId, SetDeleteResourceIdentifierId] = useState("")

  // Room
  const [createRoomIdentifier, SetCreateRoomIdentifier] = useState("") // allows users of a role to update the Room for everyone
  const [createRoomIdentifierId, SetCreateRoomIdentifierId] = useState("")
  const [viewRoomIdentifier, SetViewRoomIdentifier] = useState("") // allows users of a role to view the Room for everyone
  const [viewRoomIdentifierId, SetViewRoomIdentifierId] = useState("")
  const [deleteRoomIdentifier, SetDeleteRoomIdentifier] = useState("") // allows users of a role to delete the Room for everyone
  const [deleteRoomIdentifierId, SetDeleteRoomIdentifierId] = useState("")

  // Building
  const [createBuildingIdentifier, SetCreateBuildingIdentifier] = useState("") // allows users of a role to update the Building for everyone
  const [createBuildingIdentifierId, SetCreateBuildingIdentifierId] = useState("")
  const [viewBuildingIdentifier, SetViewBuildingIdentifier] = useState("") // allows users of a role to view the Building for everyone
  const [viewBuildingIdentifierId, SetViewBuildingIdentifierId] = useState("")
  const [deleteBuildingIdentifier, SetDeleteBuildingIdentifier] = useState("") // allows users of a role to delete the Building for everyone
  const [deleteBuildingIdentifierId, SetDeleteBuildingIdentifierId] = useState("")

  const navigate = useNavigate();

  let handleSubmit = async (e) => {
    e.preventDefault();

    if (createBookingIdentifier && createBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "BOOKING", "ROLE", null);
    }
    if (!createBookingIdentifier && createBookingIdentifierId != null) {
      RemovePermission(createBookingIdentifierId);
    }

    if (viewBookingIdentifier && viewBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "BOOKING", "ROLE", null);
    }
    if (!viewBookingIdentifier && viewBookingIdentifierId != null) {
      RemovePermission(viewBookingIdentifierId);
    }

    if (deleteBookingIdentifier && deleteBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "BOOKING", "ROLE", null);
    }
    if (!deleteBookingIdentifier && deleteBookingIdentifierId != null) {
      RemovePermission(deleteBookingIdentifierId);
    }

    // Permissions
    if (createPermissionIdentifier && createPermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "PERMISSION", "IDENTIFIER", null);
    }
    if (!createPermissionIdentifier && createPermissionIdentifierId != null) {
      RemovePermission(createPermissionIdentifierId);
    }

    if (viewPermissionIdentifier && viewPermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "PERMISSION", "IDENTIFIER", null);
    }
    if (!viewPermissionIdentifier && viewPermissionIdentifierId != null) {
      RemovePermission(viewPermissionIdentifierId);
    }

    if (deletePermissionIdentifier && deletePermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "PERMISSION", "IDENTIFIER", null);
    }
    if (!deletePermissionIdentifier && deletePermissionIdentifierId != null) {
      RemovePermission(deletePermissionIdentifierId);
    }

    // Role
    if (createRoleIdentifier && createRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "ROLE", "IDENTIFIER", null);
    }
    if (!createRoleIdentifier && createRoleIdentifierId != null) {
      RemovePermission(createRoleIdentifierId);
    }

    if (viewRoleIdentifier && viewRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "ROLE", "IDENTIFIER", null);
    }
    if (!viewRoleIdentifier && viewRoleIdentifierId != null) {
      RemovePermission(viewRoleIdentifierId);
    }

    if (deleteRoleIdentifier && deleteRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "ROLE", "IDENTIFIER", null);
    }
    if (!deleteRoleIdentifier && deleteRoleIdentifierId != null) {
      RemovePermission(deleteRoleIdentifierId);
    }

    // Team
    if (createTeamIdentifier && createTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "TEAM", "IDENTIFIER", null);
    }
    if (!createTeamIdentifier && createTeamIdentifierId != null) {
      RemovePermission(createTeamIdentifierId);
    }

    if (viewTeamIdentifier && viewTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "TEAM", "IDENTIFIER", null);
    }
    if (!viewTeamIdentifier && viewTeamIdentifierId != null) {
      RemovePermission(viewTeamIdentifierId);
    }

    if (deleteTeamIdentifier && deleteTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "TEAM", "IDENTIFIER", null);
    }
    if (!deleteTeamIdentifier && deleteTeamIdentifierId != null) {
      RemovePermission(deleteTeamIdentifierId);
    }

    // Resource
    if (createResourceIdentifier && createResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "RESOURCE", "IDENTIFIER", null);
    }
    if (!createResourceIdentifier && createResourceIdentifierId != null) {
      RemovePermission(createResourceIdentifierId);
    }

    if (viewResourceIdentifier && viewResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "RESOURCE", "IDENTIFIER", null);
    }
    if (!viewResourceIdentifier && viewResourceIdentifierId != null) {
      RemovePermission(viewResourceIdentifierId);
    }

    if (deleteResourceIdentifier && deleteResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "RESOURCE", "IDENTIFIER", null);
    }
    if (!deleteResourceIdentifier && deleteResourceIdentifierId != null) {
      RemovePermission(deleteResourceIdentifierId);
    }

    // Room
    if (createRoomIdentifier && createRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "RESOURCE", "ROOM", null);
    }
    if (!createRoomIdentifier && createRoomIdentifierId != null) {
      RemovePermission(createRoomIdentifierId);
    }

    if (viewRoomIdentifier && viewRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "RESOURCE", "ROOM", null);
    }
    if (!viewRoomIdentifier && viewRoomIdentifierId != null) {
      RemovePermission(viewRoomIdentifierId);
    }

    if (deleteRoomIdentifier && deleteRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "RESOURCE", "ROOM", null);
    }
    if (!deleteRoomIdentifier && deleteRoomIdentifierId != null) {
      RemovePermission(deleteRoomIdentifierId);
    }

    // Building
    if (createBuildingIdentifier && createBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "CREATE", "RESOURCE", "BUILDING", null);
    }
    if (!createBuildingIdentifier && createBuildingIdentifierId != null) {
      RemovePermission(createBuildingIdentifierId);
    }

    if (viewBuildingIdentifier && viewBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "VIEW", "RESOURCE", "BUILDING", null);
    }
    if (!viewBuildingIdentifier && viewBuildingIdentifierId != null) {
      RemovePermission(viewBuildingIdentifierId);
    }

    if (deleteBuildingIdentifier && deleteBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("RoleID"), "ROLE", "DELETE", "RESOURCE", "BUILDING", null);
    }
    if (!deleteBuildingIdentifier && deleteBuildingIdentifierId != null) {
      RemovePermission(deleteBuildingIdentifierId);
    }

    alert("Permissions successfully updated.")
    navigate("/role-permissions");
  };

  async function AddPermission(id, idType, type, category, tenant, tenant_id) {
    try {
      let res = await fetch("http://deskflow.co.za:8080/api/permission/create",
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            permission_id: id,
            permission_id_type: idType,
            permission_type: type,
            permission_category: category,
            permission_tenant: tenant,
            permission_tenant_id: tenant_id
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        });

      if (res.status === 200) {
        
      }
      if (res.status === 401) {
        alert("Unauthorized")
      }
    }
    catch (err) {
      console.log(err);
    }
  };

  async function RemovePermission(id) {
    try {
      let res = await fetch("http://deskflow.co.za:8080/api/permission/remove",
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            id: id,
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
        });

      if (res.status === 200) {
      }
    }
    catch (err) {
      console.log(err);
    }
  };

  //POST request
  const FetchRolePermissions = useCallback(() => {
    fetch("http://deskflow.co.za:8080/api/permission/information",
      {
        method: "POST",
        mode: "cors",
        body: JSON.stringify({
          permission_id: window.sessionStorage.getItem("RoleID"),
          permission_id_type: "ROLE",
        }),
        headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
        }
      }).then((res) => res.json()).then(data => {
        data.forEach(setPermissionStates);
      });
  },[]);

  function setPermissionStates(permission) {
    // Booking
    if (permission.permission_category === 'BOOKING'){
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "ROLE") {
          SetCreateBookingIdentifier(true);
          SetCreateBookingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "ROLE") {
          SetViewBookingIdentifier(true);
          SetViewBookingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "ROLE") {
          SetDeleteBookingIdentifier(true);
          SetDeleteBookingIdentifierId(permission.id);
        }
      }
    }

    // Permission
    if (permission.permission_category === 'PERMISSION') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetCreatePermissionIdentifier(true);
          SetCreatePermissionIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetViewPermissionIdentifier(true);
          SetViewPermissionIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetDeletePermissionIdentifier(true);
          SetDeletePermissionIdentifierId(permission.id);
        }
      }
    }

    // Role
    if (permission.permission_category === 'ROLE') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetCreateRoleIdentifier(true);
          SetCreateRoleIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetViewRoleIdentifier(true);
          SetViewRoleIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetDeleteRoleIdentifier(true);
          SetDeleteRoleIdentifierId(permission.id);
        }
      }
    }

    // Team
    if (permission.permission_category === 'TEAM') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetCreateTeamIdentifier(true);
          SetCreateTeamIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetViewTeamIdentifier(true);
          SetViewTeamIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetDeleteTeamIdentifier(true);
          SetDeleteTeamIdentifierId(permission.id);
        }
      }
    }

    // Resource
    if (permission.permission_category === 'RESOURCE') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetCreateResourceIdentifier(true);
          SetCreateResourceIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetViewResourceIdentifier(true);
          SetViewResourceIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetDeleteResourceIdentifier(true);
          SetDeleteResourceIdentifierId(permission.id);
        }
      }
    }
    
    // Room
    if (permission.permission_category === 'RESOURCE') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "ROOM") {
          SetCreateRoomIdentifier(true);
          SetCreateRoomIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "ROOM") {
          SetViewRoomIdentifier(true);
          SetViewRoomIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "ROOM") {
          SetDeleteRoomIdentifier(true);
          SetDeleteRoomIdentifierId(permission.id);
        }
      }
    }

    // Building
    if (permission.permission_category === 'RESOURCE') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "BUILDING") {
          SetCreateBuildingIdentifier(true);
          SetCreateBuildingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "BUILDING") {
          SetViewBuildingIdentifier(true);
          SetViewBuildingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "BUILDING") {
          SetDeleteBuildingIdentifier(true);
          SetDeleteBuildingIdentifierId(permission.id);
        }
      }
    }
  }

  useEffect(() => {
    FetchRolePermissions();

    setRoleName(window.sessionStorage.getItem("RoleName").substring(5, window.sessionStorage.getItem("RoleName").length));
  }, [FetchRolePermissions])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>Role {roleName}</h1> Configure Permissions</p>

          <Form className='form' onSubmit={handleSubmit}>
            <h2 className='permission-category'>Booking</h2>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit bookings for everyone
                <input className='checkbox' type="checkbox" checked={createBookingIdentifier} onChange={(e) => SetCreateBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view bookings for everyone
                <input className='checkbox' type="checkbox" checked={viewBookingIdentifier} onChange={(e) => SetViewBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete bookings for everyone
                <input className='checkbox' type="checkbox" checked={deleteBookingIdentifier} onChange={(e) => SetDeleteBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Permissions</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit permissions
                <input className='checkbox' type="checkbox" checked={createPermissionIdentifier} onChange={(e) => SetCreatePermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view permissions
                <input className='checkbox' type="checkbox" checked={viewPermissionIdentifier} onChange={(e) => SetViewPermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete permissions
                <input className='checkbox' type="checkbox" checked={deletePermissionIdentifier} onChange={(e) => SetDeletePermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Roles</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit roles
                <input className='checkbox' type="checkbox" checked={createRoleIdentifier} onChange={(e) => SetCreateRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view roles
                <input className='checkbox' type="checkbox" checked={viewRoleIdentifier} onChange={(e) => SetViewRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete roles
                <input className='checkbox' type="checkbox" checked={deleteRoleIdentifier} onChange={(e) => SetDeleteRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Resources</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit resources
                <input className='checkbox' type="checkbox" checked={createResourceIdentifier} onChange={(e) => SetCreateResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view resources
                <input className='checkbox' type="checkbox" checked={viewResourceIdentifier} onChange={(e) => SetViewResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete resources
                <input className='checkbox' type="checkbox" checked={deleteResourceIdentifier} onChange={(e) => SetDeleteResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Rooms</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit rooms
                <input className='checkbox' type="checkbox" checked={createRoomIdentifier} onChange={(e) => SetCreateRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view rooms
                <input className='checkbox' type="checkbox" checked={viewRoomIdentifier} onChange={(e) => SetViewRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete rooms
                <input className='checkbox' type="checkbox" checked={deleteRoomIdentifier} onChange={(e) => SetDeleteRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Buildings</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to create or edit buildings
                <input className='checkbox' type="checkbox" checked={createBuildingIdentifier} onChange={(e) => SetCreateBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to view buildings
                <input className='checkbox' type="checkbox" checked={viewBuildingIdentifier} onChange={(e) => SetViewBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow users of the role to delete buildings
                <input className='checkbox' type="checkbox" checked={deleteBuildingIdentifier} onChange={(e) => SetDeleteBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update Role Permissions</Button>
          </Form>

        </div>
      </div>
      <Footer />
    </div>
  )
}

export default RolePermissions
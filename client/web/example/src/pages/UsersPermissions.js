import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect, useCallback } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function UserPermissions() {
  const [userName, setUserName] = useState(window.sessionStorage.getItem("UserName"));
  //const [userPermissions, SetUserPermissions] = useState([]);

  // Bookings
  const [createBookingIdentifierUser, SetCreateBookingIdentifierUser] = useState("") // allows a user to update the Booking for themselves
  const [createBookingIdentifierUserId, SetCreateBookingIdentifierUserId] = useState("")
  const [viewBookingIdentifierUser, SetViewBookingIdentifierUser] = useState("") // allows a user to view the Booking for themselves
  const [viewBookingIdentifierUserId, SetViewBookingIdentifierUserId] = useState("")
  const [deleteBookingIdentifierUser, SetDeleteBookingIdentifierUser] = useState("") // allows a user to delete the Booking for themselves
  const [deleteBookingIdentifierUserId, SetDeleteBookingIdentifierUserId] = useState("")

  const [createBookingIdentifier, SetCreateBookingIdentifier] = useState("") // allows a user to update the Booking for everyone
  const [createBookingIdentifierId, SetCreateBookingIdentifierId] = useState("")
  const [viewBookingIdentifier, SetViewBookingIdentifier] = useState("") // allows a user to view the Booking for everyone
  const [viewBookingIdentifierId, SetViewBookingIdentifierId] = useState("")
  const [deleteBookingIdentifier, SetDeleteBookingIdentifier] = useState("") // allows a user to delete the Booking for everyone
  const [deleteBookingIdentifierId, SetDeleteBookingIdentifierId] = useState("")

  // Permissions
  const [createPermissionIdentifier, SetCreatePermissionIdentifier] = useState("") // allows a user to update the Permission for everyone
  const [createPermissionIdentifierId, SetCreatePermissionIdentifierId] = useState("")
  const [viewPermissionIdentifier, SetViewPermissionIdentifier] = useState("") // allows a user to view the Permission for everyone
  const [viewPermissionIdentifierId, SetViewPermissionIdentifierId] = useState("")
  const [deletePermissionIdentifier, SetDeletePermissionIdentifier] = useState("") // allows a user to delete the Permission for everyone
  const [deletePermissionIdentifierId, SetDeletePermissionIdentifierId] = useState("")

  // Roles
  const [createRoleIdentifier, SetCreateRoleIdentifier] = useState("") // allows a user to update the Role for everyone
  const [createRoleIdentifierId, SetCreateRoleIdentifierId] = useState("")
  const [viewRoleIdentifier, SetViewRoleIdentifier] = useState("") // allows a user to view the Role for everyone
  const [viewRoleIdentifierId, SetViewRoleIdentifierId] = useState("")
  const [deleteRoleIdentifier, SetDeleteRoleIdentifier] = useState("") // allows a user to delete the Role for everyone
  const [deleteRoleIdentifierId, SetDeleteRoleIdentifierId] = useState("")

  // Teams
  const [createTeamIdentifier, SetCreateTeamIdentifier] = useState("") // allows a user to update the Team for everyone
  const [createTeamIdentifierId, SetCreateTeamIdentifierId] = useState("")
  const [viewTeamIdentifier, SetViewTeamIdentifier] = useState("") // allows a user to view the Team for everyone
  const [viewTeamIdentifierId, SetViewTeamIdentifierId] = useState("")
  const [deleteTeamIdentifier, SetDeleteTeamIdentifier] = useState("") // allows a user to delete the Team for everyone
  const [deleteTeamIdentifierId, SetDeleteTeamIdentifierId] = useState("")

  // Resource
  const [createResourceIdentifier, SetCreateResourceIdentifier] = useState("") // allows a user to update the Resource for everyone
  const [createResourceIdentifierId, SetCreateResourceIdentifierId] = useState("")
  const [viewResourceIdentifier, SetViewResourceIdentifier] = useState("") // allows a user to view the Resource for everyone
  const [viewResourceIdentifierId, SetViewResourceIdentifierId] = useState("")
  const [deleteResourceIdentifier, SetDeleteResourceIdentifier] = useState("") // allows a user to delete the Resource for everyone
  const [deleteResourceIdentifierId, SetDeleteResourceIdentifierId] = useState("")

  // Room
  const [createRoomIdentifier, SetCreateRoomIdentifier] = useState("") // allows a user to update the Room for everyone
  const [createRoomIdentifierId, SetCreateRoomIdentifierId] = useState("")
  const [viewRoomIdentifier, SetViewRoomIdentifier] = useState("") // allows a user to view the Room for everyone
  const [viewRoomIdentifierId, SetViewRoomIdentifierId] = useState("")
  const [deleteRoomIdentifier, SetDeleteRoomIdentifier] = useState("") // allows a user to delete the Room for everyone
  const [deleteRoomIdentifierId, SetDeleteRoomIdentifierId] = useState("")

  // Building
  const [createBuildingIdentifier, SetCreateBuildingIdentifier] = useState("") // allows a user to update the Building for everyone
  const [createBuildingIdentifierId, SetCreateBuildingIdentifierId] = useState("")
  const [viewBuildingIdentifier, SetViewBuildingIdentifier] = useState("") // allows a user to view the Building for everyone
  const [viewBuildingIdentifierId, SetViewBuildingIdentifierId] = useState("")
  const [deleteBuildingIdentifier, SetDeleteBuildingIdentifier] = useState("") // allows a user to delete the Building for everyone
  const [deleteBuildingIdentifierId, SetDeleteBuildingIdentifierId] = useState("")

  let handleSubmit = async (e) => {
    e.preventDefault();

    // Booking
    if (createBookingIdentifierUser && createBookingIdentifierUserId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "BOOKING", "USER", window.sessionStorage.getItem("UserID"));
    }
    if (!createBookingIdentifierUser && createBookingIdentifierUserId != null) {
      RemovePermission(createBookingIdentifierUserId);
    }

    if (viewBookingIdentifierUser && viewBookingIdentifierUserId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "BOOKING", "USER", window.sessionStorage.getItem("UserID"));
    }
    if (!viewBookingIdentifierUser && viewBookingIdentifierUserId != null) {
      RemovePermission(viewBookingIdentifierUserId);
    }

    if (deleteBookingIdentifierUser && deleteBookingIdentifierUserId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "BOOKING", "USER", window.sessionStorage.getItem("UserID"));
    }
    if (!deleteBookingIdentifierUser && deleteBookingIdentifierUserId != null) {
      RemovePermission(deleteBookingIdentifierUserId);
    }

    if (createBookingIdentifier && createBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "BOOKING", "USER", null);
    }
    if (!createBookingIdentifier && createBookingIdentifierId != null) {
      RemovePermission(createBookingIdentifierId);
    }

    if (viewBookingIdentifier && viewBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "BOOKING", "USER", null);
    }
    if (!viewBookingIdentifier && viewBookingIdentifierId != null) {
      RemovePermission(viewBookingIdentifierId);
    }

    if (deleteBookingIdentifier && deleteBookingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "BOOKING", "USER", null);
    }
    if (!deleteBookingIdentifier && deleteBookingIdentifierId != null) {
      RemovePermission(deleteBookingIdentifierId);
    }

    // Permissions
    if (createPermissionIdentifier && createPermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "PERMISSION", "IDENTIFIER", null);
    }
    if (!createPermissionIdentifier && createPermissionIdentifierId != null) {
      RemovePermission(createPermissionIdentifierId);
    }

    if (viewPermissionIdentifier && viewPermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "PERMISSION", "IDENTIFIER", null);
    }
    if (!viewPermissionIdentifier && viewPermissionIdentifierId != null) {
      RemovePermission(viewPermissionIdentifierId);
    }

    if (deletePermissionIdentifier && deletePermissionIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "PERMISSION", "IDENTIFIER", null);
    }
    if (!deletePermissionIdentifier && deletePermissionIdentifierId != null) {
      RemovePermission(deletePermissionIdentifierId);
    }

    // Role
    if (createRoleIdentifier && createRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "ROLE", "IDENTIFIER", null);
    }
    if (!createRoleIdentifier && createRoleIdentifierId != null) {
      RemovePermission(createRoleIdentifierId);
    }

    if (viewRoleIdentifier && viewRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "ROLE", "IDENTIFIER", null);
    }
    if (!viewRoleIdentifier && viewRoleIdentifierId != null) {
      RemovePermission(viewRoleIdentifierId);
    }

    if (deleteRoleIdentifier && deleteRoleIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "ROLE", "IDENTIFIER", null);
    }
    if (!deleteRoleIdentifier && deleteRoleIdentifierId != null) {
      RemovePermission(deleteRoleIdentifierId);
    }

    // Team
    if (createTeamIdentifier && createTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "TEAM", "IDENTIFIER", null);
    }
    if (!createTeamIdentifier && createTeamIdentifierId != null) {
      RemovePermission(createTeamIdentifierId);
    }

    if (viewTeamIdentifier && viewTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "TEAM", "IDENTIFIER", null);
    }
    if (!viewTeamIdentifier && viewTeamIdentifierId != null) {
      RemovePermission(viewTeamIdentifierId);
    }

    if (deleteTeamIdentifier && deleteTeamIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "TEAM", "IDENTIFIER", null);
    }
    if (!deleteTeamIdentifier && deleteTeamIdentifierId != null) {
      RemovePermission(deleteTeamIdentifierId);
    }

    // Resource
    if (createResourceIdentifier && createResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "RESOURCE", "IDENTIFIER", null);
    }
    if (!createResourceIdentifier && createResourceIdentifierId != null) {
      RemovePermission(createResourceIdentifierId);
    }

    if (viewResourceIdentifier && viewResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "RESOURCE", "IDENTIFIER", null);
    }
    if (!viewResourceIdentifier && viewResourceIdentifierId != null) {
      RemovePermission(viewResourceIdentifierId);
    }

    if (deleteResourceIdentifier && deleteResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "RESOURCE", "IDENTIFIER", null);
    }
    if (!deleteResourceIdentifier && deleteResourceIdentifierId != null) {
      RemovePermission(deleteResourceIdentifierId);
    }

    // Room
    if (createRoomIdentifier && createRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "RESOURCE", "ROOM", null);
    }
    if (!createRoomIdentifier && createRoomIdentifierId != null) {
      RemovePermission(createRoomIdentifierId);
    }

    if (viewRoomIdentifier && viewRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "RESOURCE", "ROOM", null);
    }
    if (!viewRoomIdentifier && viewRoomIdentifierId != null) {
      RemovePermission(viewRoomIdentifierId);
    }

    if (deleteRoomIdentifier && deleteRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "RESOURCE", "ROOM", null);
    }
    if (!deleteRoomIdentifier && deleteRoomIdentifierId != null) {
      RemovePermission(deleteRoomIdentifierId);
    }

    // Building
    if (createBuildingIdentifier && createBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "RESOURCE", "BUILDING", null);
    }
    if (!createBuildingIdentifier && createBuildingIdentifierId != null) {
      RemovePermission(createBuildingIdentifierId);
    }

    if (viewBuildingIdentifier && viewBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "RESOURCE", "BUILDING", null);
    }
    if (!viewBuildingIdentifier && viewBuildingIdentifierId != null) {
      RemovePermission(viewBuildingIdentifierId);
    }

    if (deleteBuildingIdentifier && deleteBuildingIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "RESOURCE", "BUILDING", null);
    }
    if (!deleteBuildingIdentifier && deleteBuildingIdentifierId != null) {
      RemovePermission(deleteBuildingIdentifierId);
    }

    alert("Permissions successfully updated.")
  };

  async function AddPermission(id, idType, type, category, tenant, tenant_id) {
    try {
      let res = await fetch("http://localhost:8100/api/permission/create",
        {
          method: "POST",
          body: JSON.stringify({
            permission_id: id,
            permission_id_type: idType,
            permission_type: type,
            permission_category: category,
            permission_tenant: tenant,
            permission_tenant_id: tenant_id
          })
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
      let res = await fetch("http://localhost:8100/api/permission/remove",
        {
          method: "POST",
          body: JSON.stringify({
            id: id,
          })
        });

      if (res.status === 200) {
      }
    }
    catch (err) {
      console.log(err);
    }
  };

  //POST request
  const FetchUserPermissions = useCallback(() => {
    fetch("http://localhost:8100/api/permission/information",
      {
        method: "POST",
        body: JSON.stringify({
          permission_id: window.sessionStorage.getItem("UserID"),
          permission_id_type: "USER",
        })
      }).then((res) => res.json()).then(data => {
        //SetUserPermissions(data); [Uncomment when used]
        data.forEach(setPermissionStates);
      });
  },[]);

  function setPermissionStates(permission) {
    // Booking
    if (permission.permission_category === 'BOOKING' && permission.permission_tenant_id === window.sessionStorage.getItem("UserID")) {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "USER") {
          SetCreateBookingIdentifierUser(true);
          SetCreateBookingIdentifierUserId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "USER") {
          SetViewBookingIdentifierUser(true);
          SetViewBookingIdentifierUserId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "USER") {
          SetDeleteBookingIdentifierUser(true);
          SetDeleteBookingIdentifierUserId(permission.id);
        }
      }
    } else if (permission.permission_category === 'BOOKING'){
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "USER") {
          SetCreateBookingIdentifier(true);
          SetCreateBookingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "USER") {
          SetViewBookingIdentifier(true);
          SetViewBookingIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (permission.permission_tenant === "USER") {
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
    FetchUserPermissions();

    setUserName(window.sessionStorage.getItem("UserName"));
  }, [FetchUserPermissions])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>User {userName}</h1> Configure Permissions</p>

          <Form className='form' onSubmit={handleSubmit}>
            <h2 className='permission-category'>Booking</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit bookings for themselves
                <input className='checkbox' type="checkbox" checked={createBookingIdentifierUser} onChange={(e) => SetCreateBookingIdentifierUser(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view bookings for themselves
                <input className='checkbox' type="checkbox" checked={viewBookingIdentifierUser} onChange={(e) => SetViewBookingIdentifierUser(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete bookings for themselves
                <input className='checkbox' type="checkbox" checked={deleteBookingIdentifierUser} onChange={(e) => SetDeleteBookingIdentifierUser(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <br></br>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit bookings for everyone
                <input className='checkbox' type="checkbox" checked={createBookingIdentifier} onChange={(e) => SetCreateBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view bookings for everyone
                <input className='checkbox' type="checkbox" checked={viewBookingIdentifier} onChange={(e) => SetViewBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete bookings for everyone
                <input className='checkbox' type="checkbox" checked={deleteBookingIdentifier} onChange={(e) => SetDeleteBookingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Permissions</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit permissions
                <input className='checkbox' type="checkbox" checked={createPermissionIdentifier} onChange={(e) => SetCreatePermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view permissions
                <input className='checkbox' type="checkbox" checked={viewPermissionIdentifier} onChange={(e) => SetViewPermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete permissions
                <input className='checkbox' type="checkbox" checked={deletePermissionIdentifier} onChange={(e) => SetDeletePermissionIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Roles</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit roles
                <input className='checkbox' type="checkbox" checked={createRoleIdentifier} onChange={(e) => SetCreateRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view roles
                <input className='checkbox' type="checkbox" checked={viewRoleIdentifier} onChange={(e) => SetViewRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete roles
                <input className='checkbox' type="checkbox" checked={deleteRoleIdentifier} onChange={(e) => SetDeleteRoleIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Resources</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit resources
                <input className='checkbox' type="checkbox" checked={createResourceIdentifier} onChange={(e) => SetCreateResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view resources
                <input className='checkbox' type="checkbox" checked={viewResourceIdentifier} onChange={(e) => SetViewResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete resources
                <input className='checkbox' type="checkbox" checked={deleteResourceIdentifier} onChange={(e) => SetDeleteResourceIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Rooms</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit rooms
                <input className='checkbox' type="checkbox" checked={createRoomIdentifier} onChange={(e) => SetCreateRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view rooms
                <input className='checkbox' type="checkbox" checked={viewRoomIdentifier} onChange={(e) => SetViewRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete rooms
                <input className='checkbox' type="checkbox" checked={deleteRoomIdentifier} onChange={(e) => SetDeleteRoomIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <h2 className='permission-category'>Buildings</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create or edit buildings
                <input className='checkbox' type="checkbox" checked={createBuildingIdentifier} onChange={(e) => SetCreateBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to view buildings
                <input className='checkbox' type="checkbox" checked={viewBuildingIdentifier} onChange={(e) => SetViewBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to delete buildings
                <input className='checkbox' type="checkbox" checked={deleteBuildingIdentifier} onChange={(e) => SetDeleteBuildingIdentifier(e.target.checked)} />
                <span className="checkmark"></span>
              </label>
            </Form.Group>

            <Button className='button-submit' variant='primary' type='submit'>Update User Permissions</Button>
          </Form>

        </div>
      </div>
      <Footer />
    </div>
  )
}

export default UserPermissions
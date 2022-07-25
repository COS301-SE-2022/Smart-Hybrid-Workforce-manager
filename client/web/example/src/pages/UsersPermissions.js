import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import { useState, useEffect } from 'react'
import Form from 'react-bootstrap/Form'
import Button from 'react-bootstrap/Button'
import '../App.css'

function UserPermissions() {
  const [userName, setUserName] = useState(window.sessionStorage.getItem("UserName"));
  const [userPermissions, SetUserPermissions] = useState([]);

  // Resource
  const [createResourceIdentifier, SetCreateResourceIdentifier] = useState("") // allows a user to update the Resource
  const [createResourceIdentifierId, SetCreateResourceIdentifierId] = useState("")
  const [viewResourceIdentifier, SetViewResourceIdentifier] = useState("") // allows a user to view the Resource
  const [viewResourceIdentifierId, SetViewResourceIdentifierId] = useState("")
  const [deleteResourceIdentifier, SetDeleteResourceIdentifier] = useState("") // allows a user to delete the Resource
  const [deleteResourceIdentifierId, SetDeleteResourceIdentifierId] = useState("")

  // Room
  const [createRoomIdentifier, SetCreateRoomIdentifier] = useState("") // allows a user to update the Room
  const [createRoomIdentifierId, SetCreateRoomIdentifierId] = useState("")
  const [viewRoomIdentifier, SetViewRoomIdentifier] = useState("") // allows a user to view the Room
  const [viewRoomIdentifierId, SetViewRoomIdentifierId] = useState("")
  const [deleteRoomIdentifier, SetDeleteRoomIdentifier] = useState("") // allows a user to delete the Room
  const [deleteRoomIdentifierId, SetDeleteRoomIdentifierId] = useState("")

  let handleSubmit = async (e) => {
    e.preventDefault();

    // Resource
    if (createResourceIdentifier && createResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "RESOURCE", "IDENTIFIER");
    }
    if (!createResourceIdentifier && createResourceIdentifierId != null) {
      RemovePermission(createResourceIdentifierId);
    }

    if (viewResourceIdentifier && viewResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "RESOURCE", "IDENTIFIER");
    }
    if (!viewResourceIdentifier && viewResourceIdentifierId != null) {
      RemovePermission(viewResourceIdentifierId);
    }

    if (deleteResourceIdentifier && deleteResourceIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "RESOURCE", "IDENTIFIER");
    }
    if (!deleteResourceIdentifier && deleteResourceIdentifierId != null) {
      RemovePermission(deleteResourceIdentifierId);
    }

    // Room
    if (createRoomIdentifier && createRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "CREATE", "ROOM", "IDENTIFIER");
    }
    if (!createRoomIdentifier && createRoomIdentifierId != null) {
      RemovePermission(createRoomIdentifierId);
    }

    if (viewRoomIdentifier && viewRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "VIEW", "ROOM", "IDENTIFIER");
    }
    if (!viewRoomIdentifier && viewRoomIdentifierId != null) {
      RemovePermission(viewRoomIdentifierId);
    }

    if (deleteRoomIdentifier && deleteRoomIdentifierId === "") {
      AddPermission(window.sessionStorage.getItem("UserID"), "USER", "DELETE", "ROOM", "IDENTIFIER");
    }
    if (!deleteRoomIdentifier && deleteRoomIdentifierId != null) {
      RemovePermission(deleteRoomIdentifierId);
    }

    alert("Permissions successfully updated.")
  };

  async function AddPermission(id, idType, type, category, tenant) {
    try {
      let res = await fetch("http://localhost:8100/api/permission/create",
        {
          method: "POST",
          body: JSON.stringify({
            permission_id: id,
            permission_id_type: idType,
            permission_type: type,
            permission_category: category,
            permission_tenant: tenant
          })
        });

      if (res.status === 200) {
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
  const FetchUserPermissions = () => {
    fetch("http://localhost:8100/api/permission/information",
      {
        method: "POST",
        body: JSON.stringify({
          permission_id: window.sessionStorage.getItem("UserID"),
          permission_id_type: "USER",
        })
      }).then((res) => res.json()).then(data => {
        SetUserPermissions(data);
        data.forEach(setPermissionStates);
      });
  }

  function setPermissionStates(permission) {
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
        if (ppermission.permission_tenant === "IDENTIFIER") {
          SetDeleteResourceIdentifier(true);
          SetDeleteResourceIdentifierId(permission.id);
        }
      }
    }
    
    // Room
    if (permission.permission_category === 'ROOM') {
      if (permission.permission_type === "CREATE") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetCreateRoomIdentifier(true);
          SetCreateRoomIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "VIEW") {
        if (permission.permission_tenant === "IDENTIFIER") {
          SetViewRoomIdentifier(true);
          SetViewRoomIdentifierId(permission.id);
        }
      }
      if (permission.permission_type === "DELETE") {
        if (ppermission.permission_tenant === "IDENTIFIER") {
          SetDeleteRoomIdentifier(true);
          SetDeleteRoomIdentifierId(permission.id);
        }
      }
    }
  }

  useEffect(() => {
    FetchUserPermissions();

    setUserName(window.sessionStorage.getItem("UserName"));
  }, [])

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='form-container-team'>
          <p className='form-header'><h1>User {userName}</h1> Configure Permissions</p>

          <Form className='form' onSubmit={handleSubmit}>
            <h2 className='permission-category'>Resources</h2>
            <Form.Group className='form-group' controlId="formBasicName">
              <label className="container">
                Allow user to create resources
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
                Allow user to create rooms
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

            <Button className='button-submit' variant='primary' type='submit'>Update User Permissions</Button>
          </Form>

        </div>
      </div>
      <Footer />
    </div>
  )
}

export default UserPermissions
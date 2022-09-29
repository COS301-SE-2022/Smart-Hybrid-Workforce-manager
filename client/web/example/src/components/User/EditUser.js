import Button from 'react-bootstrap/Button';
import styles from './user.module.css';
import { useContext, useEffect, useState } from 'react';
import { UserContext } from '../../App';

const EditUser = ({userID, userName, userPicture, userRoles, allRoles, edited}) =>
{
    const [id, setID] = useState(userID);
    const [name, setName] = useState(userName);
    const [picture, setPicture] = useState(userPicture);
    const [activeRoles, setActiveRoles] = useState({});

    const {userData} = useContext(UserContext);

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

    const [roles, setRoles] = useState(allRoles);

    const EditActiveRoles = (role) =>
    {
        if(document.getElementById(role.id).checked)
        {
            fetch("http://deskflow.co.za:8080/api/role/user/create", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    role_id: role.id,
                    user_id: userID
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) =>
            {
                if(res.status === 200)
                {
                    setActiveRoles({
                        ...activeRoles,
                        [role.id]:
                        {
                            id: role.id
                        }
                    });

                    edited(true);
                }

                if(res.status !== 200)
                {
                    alert('An error occured when updating the team');
                }
            });
        }   
        else
        {
            fetch("http://deskflow.co.za:8080/api/role/user/remove", 
            {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    role_id: role.id,
                    user_id: userID
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) =>
            {
                if(res.status === 200)
                {
                    const newRoles = activeRoles;
                    delete newRoles[role.id];
                    setActiveRoles(newRoles);
                    edited(true);
                }

                if(res.status !== 200)
                {
                    alert('An error occured when updating the team');
                }
            });
        }
    }

    async function AddPermission(id, idType, type, category, tenant, tenant_id) 
    {
        try 
        {
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
                    'Authorization': `bearer ${userData.token}`
                }
            });

            if (res.status === 401)
            {
                alert("Unauthorized")
            }
        }
        catch (err)
        {
            console.log(err);
        }
    };
    
    async function RemovePermission(id) 
    {
        if(id !== '')
        {
            try
            {
                let res = await fetch("http://deskflow.co.za:8080/api/permission/remove",
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        id: id,
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}`
                    }
                });
            
                if (res.status === 401)
                {
                    alert("Unauthorized")
                }
            }
            catch (err)
            {
                console.log(err);
            }
        }
    };


    const UpdatePermissions = () =>
    {
        // Booking
        if (createBookingIdentifierUser && createBookingIdentifierUserId === "") 
        {
            AddPermission(id, "USER", "CREATE", "BOOKING", "USER", id);
        }
        if (!createBookingIdentifierUser && createBookingIdentifierUserId !== null) 
        {
            RemovePermission(createBookingIdentifierUserId);
        }
    
        if (viewBookingIdentifierUser && viewBookingIdentifierUserId === "") 
        {
            AddPermission(id, "USER", "VIEW", "BOOKING", "USER", id);
        }
        if (!viewBookingIdentifierUser && viewBookingIdentifierUserId !== null) 
        {
            RemovePermission(viewBookingIdentifierUserId);
        }
    
        if (deleteBookingIdentifierUser && deleteBookingIdentifierUserId === "") 
        {
            AddPermission(id, "USER", "DELETE", "BOOKING", "USER", id);
        }
        if (!deleteBookingIdentifierUser && deleteBookingIdentifierUserId !== null) 
        {
            RemovePermission(deleteBookingIdentifierUserId);
        }
    
        if (createBookingIdentifier && createBookingIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "BOOKING", "USER", null);
        }
        if (!createBookingIdentifier && createBookingIdentifierId !== null) 
        {
            RemovePermission(createBookingIdentifierId);
        }
    
        if (viewBookingIdentifier && viewBookingIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "BOOKING", "USER", null);
        }
        if (!viewBookingIdentifier && viewBookingIdentifierId !== null) 
        {
            RemovePermission(viewBookingIdentifierId);
        }
    
        if (deleteBookingIdentifier && deleteBookingIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "BOOKING", "USER", null);
        }
        if (!deleteBookingIdentifier && deleteBookingIdentifierId !== null) 
        {
            RemovePermission(deleteBookingIdentifierId);
        }
    
        // Permissions
        if (createPermissionIdentifier && createPermissionIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "PERMISSION", "IDENTIFIER", null);
        }
        if (!createPermissionIdentifier && createPermissionIdentifierId !== null) 
        {
            RemovePermission(createPermissionIdentifierId);
        }
    
        if (viewPermissionIdentifier && viewPermissionIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "PERMISSION", "IDENTIFIER", null);
        }
        if (!viewPermissionIdentifier && viewPermissionIdentifierId !== null) 
        {
            RemovePermission(viewPermissionIdentifierId);
        }
    
        if (deletePermissionIdentifier && deletePermissionIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "PERMISSION", "IDENTIFIER", null);
        }
        if (!deletePermissionIdentifier && deletePermissionIdentifierId !== null) 
        {
            RemovePermission(deletePermissionIdentifierId);
        }
    
        // Role
        if (createRoleIdentifier && createRoleIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "ROLE", "IDENTIFIER", null);
        }
        if (!createRoleIdentifier && createRoleIdentifierId !== null) 
        {
            RemovePermission(createRoleIdentifierId);
        }
    
        if (viewRoleIdentifier && viewRoleIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "ROLE", "IDENTIFIER", null);
        }
        if (!viewRoleIdentifier && viewRoleIdentifierId !== null) 
        {
            RemovePermission(viewRoleIdentifierId);
        }
    
        if (deleteRoleIdentifier && deleteRoleIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "ROLE", "IDENTIFIER", null);
        }
        if (!deleteRoleIdentifier && deleteRoleIdentifierId !== null) 
        {
            RemovePermission(deleteRoleIdentifierId);
        }
    
        // Team
        if (createTeamIdentifier && createTeamIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "TEAM", "IDENTIFIER", null);
        }
        if (!createTeamIdentifier && createTeamIdentifierId !== null) 
        {
            RemovePermission(createTeamIdentifierId);
        }
    
        if (viewTeamIdentifier && viewTeamIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "TEAM", "IDENTIFIER", null);
        }
        if (!viewTeamIdentifier && viewTeamIdentifierId !== null) 
        {
            RemovePermission(viewTeamIdentifierId);
        }
    
        if (deleteTeamIdentifier && deleteTeamIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "TEAM", "IDENTIFIER", null);
        }
        if (!deleteTeamIdentifier && deleteTeamIdentifierId !== null) 
        {
            RemovePermission(deleteTeamIdentifierId);
        }
    
        // Resource
        if (createResourceIdentifier && createResourceIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "RESOURCE", "IDENTIFIER", null);
        }
        if (!createResourceIdentifier && createResourceIdentifierId !== null) 
        {
            RemovePermission(createResourceIdentifierId);
        }
    
        if (viewResourceIdentifier && viewResourceIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "RESOURCE", "IDENTIFIER", null);
        }
        if (!viewResourceIdentifier && viewResourceIdentifierId !== null) 
        {
            RemovePermission(viewResourceIdentifierId);
        }
    
        if (deleteResourceIdentifier && deleteResourceIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "RESOURCE", "IDENTIFIER", null);
        }
        if (!deleteResourceIdentifier && deleteResourceIdentifierId !== null) 
        {
            RemovePermission(deleteResourceIdentifierId);
        }
    
        // Room
        if (createRoomIdentifier && createRoomIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "RESOURCE", "ROOM", null);
        }
        if (!createRoomIdentifier && createRoomIdentifierId !== null) 
        {
            RemovePermission(createRoomIdentifierId);
        }
    
        if (viewRoomIdentifier && viewRoomIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "RESOURCE", "ROOM", null);
        }
        if (!viewRoomIdentifier && viewRoomIdentifierId !== null) 
        {
            RemovePermission(viewRoomIdentifierId);
        }
    
        if (deleteRoomIdentifier && deleteRoomIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "RESOURCE", "ROOM", null);
        }
        if (!deleteRoomIdentifier && deleteRoomIdentifierId !== null) 
        {
            RemovePermission(deleteRoomIdentifierId);
        }
    
        // Building
        if (createBuildingIdentifier && createBuildingIdentifierId === "") 
        {
            AddPermission(id, "USER", "CREATE", "RESOURCE", "BUILDING", null);
        }
        if (!createBuildingIdentifier && createBuildingIdentifierId !== null) 
        {
            RemovePermission(createBuildingIdentifierId);
        }
    
        if (viewBuildingIdentifier && viewBuildingIdentifierId === "") 
        {
            AddPermission(id, "USER", "VIEW", "RESOURCE", "BUILDING", null);
        }
        if (!viewBuildingIdentifier && viewBuildingIdentifierId !== null) 
        {
            RemovePermission(viewBuildingIdentifierId);
        }
    
        if (deleteBuildingIdentifier && deleteBuildingIdentifierId === "") 
        {
            AddPermission(id, "USER", "DELETE", "RESOURCE", "BUILDING", null);
        }
        if (!deleteBuildingIdentifier && deleteBuildingIdentifierId !== null) 
        {
            RemovePermission(deleteBuildingIdentifierId);
        }
    }

    useEffect(() =>
    {
        setName(userName);
    }, [userName]);

    useEffect(() =>
    {
        setPicture(userPicture);
    }, [userPicture]);

    useEffect(() =>
    {
        if(userRoles)
        {
            var roles = {};
            for(let i = 0; i < userRoles.length; i++)
            {
                roles = 
                {
                    ...roles,
                    [userRoles[i].id]:
                    {
                        id: userRoles[i].id
                    }
                }
            }

            setActiveRoles(roles);
        }        
    }, [userRoles]);

    useEffect(() =>
    {
        setRoles(allRoles);
    }, [allRoles]);

    useEffect(() =>
    {
        if(activeRoles && roles)
        {
            for(let i = 0; i < roles.length; i++)
            {
                if(activeRoles[roles[i].id])
                {
                    document.getElementById(roles[i].id).checked = true;
                }
                else
                {
                    document.getElementById(roles[i].id).checked = false;
                }
            }
        }
    }, [roles, activeRoles]);

    useEffect(() =>
    {
        const setPermissionStates = (permission) =>
        {
            // Booking
            if (permission.permission_category === 'BOOKING' && permission.permission_tenant_id === userID)
            {
                if (permission.permission_type === "CREATE")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetCreateBookingIdentifierUser(true);
                        SetCreateBookingIdentifierUserId(permission.id);
                    }
                }
                if (permission.permission_type === "VIEW")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetViewBookingIdentifierUser(true);
                        SetViewBookingIdentifierUserId(permission.id);
                    }
                }
                if (permission.permission_type === "DELETE")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetDeleteBookingIdentifierUser(true);
                        SetDeleteBookingIdentifierUserId(permission.id);
                    }
                }
            }
            else if (permission.permission_category === 'BOOKING')
            {
                if (permission.permission_type === "CREATE")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetCreateBookingIdentifier(true);
                        SetCreateBookingIdentifierId(permission.id);
                    }
                }
                if (permission.permission_type === "VIEW")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetViewBookingIdentifier(true);
                        SetViewBookingIdentifierId(permission.id);
                    }
                }
                if (permission.permission_type === "DELETE")
                {
                    if (permission.permission_tenant === "USER")
                    {
                        SetDeleteBookingIdentifier(true);
                        SetDeleteBookingIdentifierId(permission.id);
                    }
                }
            }
        
            // Permission
            if(permission.permission_category === 'PERMISSION')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetCreatePermissionIdentifier(true);
                        SetCreatePermissionIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetViewPermissionIdentifier(true);
                        SetViewPermissionIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetDeletePermissionIdentifier(true);
                        SetDeletePermissionIdentifierId(permission.id);
                    }
                }
            }
        
            // Role
            if(permission.permission_category === 'ROLE')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetCreateRoleIdentifier(true);
                        SetCreateRoleIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetViewRoleIdentifier(true);
                        SetViewRoleIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetDeleteRoleIdentifier(true);
                        SetDeleteRoleIdentifierId(permission.id);
                    }
                }
            }
        
            // Team
            if (permission.permission_category === 'TEAM')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetCreateTeamIdentifier(true);
                        SetCreateTeamIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetViewTeamIdentifier(true);
                        SetViewTeamIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetDeleteTeamIdentifier(true);
                        SetDeleteTeamIdentifierId(permission.id);
                    }
                }
            }
        
            // Resource
            if (permission.permission_category === 'RESOURCE')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetCreateResourceIdentifier(true);
                        SetCreateResourceIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetViewResourceIdentifier(true);
                        SetViewResourceIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "IDENTIFIER")
                    {
                        SetDeleteResourceIdentifier(true);
                        SetDeleteResourceIdentifierId(permission.id);
                    }
                }
            }
            
            // Room
            if(permission.permission_category === 'RESOURCE')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "ROOM")
                    {
                        SetCreateRoomIdentifier(true);
                        SetCreateRoomIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "ROOM")
                    {
                        SetViewRoomIdentifier(true);
                        SetViewRoomIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "ROOM")
                    {
                        SetDeleteRoomIdentifier(true);
                        SetDeleteRoomIdentifierId(permission.id);
                    }
                }
            }
        
            // Building
            if(permission.permission_category === 'RESOURCE')
            {
                if(permission.permission_type === "CREATE")
                {
                    if(permission.permission_tenant === "BUILDING")
                    {
                        SetCreateBuildingIdentifier(true);
                        SetCreateBuildingIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "VIEW")
                {
                    if(permission.permission_tenant === "BUILDING")
                    {
                        SetViewBuildingIdentifier(true);
                        SetViewBuildingIdentifierId(permission.id);
                    }
                }
                if(permission.permission_type === "DELETE")
                {
                    if(permission.permission_tenant === "BUILDING")
                    {
                        SetDeleteBuildingIdentifier(true);
                        SetDeleteBuildingIdentifierId(permission.id);
                    }
                }
            }
        }

        setID(userID);
        if(userID !== '')
        {
            fetch("http://deskflow.co.za:8080/api/permission/information", 
            {
                method: "POST",
                mode: 'cors',
                body: JSON.stringify({
                    permission_id: userID,
                    permission_id_type: 'USER'
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}`
                }
            }).then((res) => res.json()).then(data => 
            {
                data.forEach(setPermissionStates);
            });
        }
    }, [userID, userData.token]);

    return (
        <div className={styles.editUserContainer}>
            <div className={styles.headerContainer}>
                <div className={styles.pictureContainer}>
                    <img className={styles.picture} src={picture} alt='User'></img>
                </div>

                <div className={styles.userName}>{name}</div>
            </div>

            <div className={styles.rolesContainer}>
                Roles
                {roles.map((role) =>
                {
                    return (
                        <div key={role.id}>
                            <input type='checkbox' id={role.id} onChange={() => EditActiveRoles(role)}></input>
                            <label className={styles.roleLabel}>{role.name}</label>
                        </div>
                    );
                })}
            </div>

            <div className={styles.rolesContainer}>
                Permissions
                <div>
                    <input type="checkbox" checked={createBookingIdentifierUser} onChange={(e) => SetCreateBookingIdentifierUser(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit bookings for themselves</label>
                </div>

                <div>
                    <input type="checkbox" checked={viewBookingIdentifierUser} onChange={(e) => SetViewBookingIdentifierUser(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view bookings for themselves</label>
                </div>

                <div>
                    <input  type="checkbox" checked={deleteBookingIdentifierUser} onChange={(e) => SetDeleteBookingIdentifierUser(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete bookings for themselves</label>
                </div>
            
                <div>
                    <input  type="checkbox" checked={createBookingIdentifier} onChange={(e) => SetCreateBookingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit bookings for everyone</label>
                </div>

                <div>
                    <input  type="checkbox" checked={viewBookingIdentifier} onChange={(e) => SetViewBookingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view bookings for everyone</label>
                </div>
                           
                <div>
                    <input  type="checkbox" checked={deleteBookingIdentifier} onChange={(e) => SetDeleteBookingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete bookings for everyone</label>
                </div>

                <div>
                    <input  type="checkbox" checked={createPermissionIdentifier} onChange={(e) => SetCreatePermissionIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit permissions</label>
                </div>

                <div>
                    <input  type="checkbox" checked={viewPermissionIdentifier} onChange={(e) => SetViewPermissionIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view permissions</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={deletePermissionIdentifier} onChange={(e) => SetDeletePermissionIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete permissions</label>
                </div>

                <div>
                    <input  type="checkbox" checked={createRoleIdentifier} onChange={(e) => SetCreateRoleIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit roles</label>
                </div>

                <div>
                    <input  type="checkbox" checked={viewRoleIdentifier} onChange={(e) => SetViewRoleIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view roles</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={deleteRoleIdentifier} onChange={(e) => SetDeleteRoleIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete roles</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={createResourceIdentifier} onChange={(e) => SetCreateResourceIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit resources</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={viewResourceIdentifier} onChange={(e) => SetViewResourceIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view resources</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={deleteResourceIdentifier} onChange={(e) => SetDeleteResourceIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete resources</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={createRoomIdentifier} onChange={(e) => SetCreateRoomIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit rooms</label>
                </div>

                <div>
                    <input  type="checkbox" checked={viewRoomIdentifier} onChange={(e) => SetViewRoomIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view rooms</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={deleteRoomIdentifier} onChange={(e) => SetDeleteRoomIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete rooms</label>
                </div>

                <div>
                    <input  type="checkbox" checked={createBuildingIdentifier} onChange={(e) => SetCreateBuildingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to create or edit buildings</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={viewBuildingIdentifier} onChange={(e) => SetViewBuildingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to view buildings</label>
                </div>
                
                <div>
                    <input  type="checkbox" checked={deleteBuildingIdentifier} onChange={(e) => SetDeleteBuildingIdentifier(e.target.checked)} />
                    <label className={styles.roleLabel}>Allow user to delete buildings</label>
                </div>
            </div>
            <button className={styles.submit} onClick={() => UpdatePermissions()}>Update</button>
        </div>
    );

}

export {EditUser as EditUserPanel}
import React from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { MdPermIdentity } from 'react-icons/md'

const RoleListItem = ({id, name}) => {

    let EditRole = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("RoleId", id);
        window.location.assign("./role-edit");
    }

    let DeleteRole = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this role?"))
        {
            try
            {
                let res = await fetch("http://localhost:8100/api/role/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                    id: id
                    })
                });

                if(res.status === 200)
                {
                    alert("Role Successfully Deleted!");
                    window.location.assign("./roles");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    const AddRole = () =>
    {
        window.sessionStorage.setItem("RoleID", currRoom);
        window.location.assign("./add-role");
    }

    return (
        <div>
            <div className="resource">
                <div className='resource-container'>
                    <div className='resource-name'>{'Role ' + name}</div>
                    <MdPermIdentity className='resource-icon' size={50} />
                </div>                
                <div className='resource-popup'>
                    <div className='resource-edit'><MdEdit size={30} className="resource-edit-icon" onClick={EditRole}/></div>
                    <div className='resource-delete'><MdDelete size={30} className="resource-delete-icon" onClick={DeleteRole}/></div>
                </div>
            </div>
        </div>
    )
}

export default RoleListItem
import { useEffect } from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { MdPermIdentity } from 'react-icons/md'
import { useNavigate } from 'react-router-dom'

const RoleListItem = ({id, name, color, lead}) =>
{
    const navigate=useNavigate();
    let EditRole = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("RoleID", id);
        window.sessionStorage.setItem("RoleName", name);
        window.sessionStorage.setItem("RoleColor", color);
        window.sessionStorage.setItem("RoleLead", lead);
        navigate("/role-edit");
    }

    let DeleteRole = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this role?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/role/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                    id: id
                    })
                });

                if(res.status === 200)
                {
                    alert("Role Successfully Deleted!");
                    navigate("/role");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    useEffect(() =>
    {
        window.sessionStorage.removeItem("RoleID");
        window.sessionStorage.removeItem("RoleName");
        window.sessionStorage.removeItem("RoleColor");
        window.sessionStorage.removeItem("RoleLead");
    }, [])

    return (
        <div>
            <div className="resource">
                <div className='resource-container'>
                    <div className='resource-name'>{name}</div>
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
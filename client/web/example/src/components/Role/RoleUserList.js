import { useState, useEffect } from 'react';
import { MdDelete } from 'react-icons/md'

const RoleUserList = ({id}) =>
{
  const [roleName, SetRoleName] = useState("")
    
    useEffect(() =>
    {
        fetch("http://localhost:8100/api/role/information", 
        {
        method: "POST",
        body: JSON.stringify({
            id: id
        })
        }).then((res) => res.json()).then(data => 
        {
            SetRoleName(data[0].role_name);
        }).catch((err) => console.log(err));
    }, [])

    let DeleteRole = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to remove this role?"))
        {
            try
            {
                let res = await fetch("http://localhost:8100/api/role/user/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                        role_id: id,
                        user_id: window.sessionStorage.getItem("UserID")
                    })
                });

                if(res.status === 200)
                {
                    alert("Role Successfully Removed!");
                    window.location.reload();
                }
            }
            catch (err)
            {
                console.log(err);
            }
        }
    }

    return (
        <div className='list-item'>
            <div className='list-item-content'>
                {roleName}
            </div>
            <div className='list-item-popup'>
                <div className='list-item-delete'><MdDelete size={20} className="list-item-delete-icon" onClick={DeleteRole}/></div>
            </div>
        </div>
    )
}

export default RoleUserList
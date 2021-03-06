import { useState, useEffect, useCallback } from 'react';
import { MdDelete } from 'react-icons/md'

const UserRoleList = ({id}) =>
{  
    const [name, setName] = useState("error");
    
  let DeleteUser = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to remove this user from this role?"))
        {
            try
            {
                let res = await fetch("http://localhost:8100/api/role/user/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                        user_id: id,
                        role_id: window.sessionStorage.getItem("RoleID")
                    })
                });

                if(res.status === 200)
                {
                    alert("User Successfully Removed!");
                    window.location.reload();
                }
            }
            catch (err)
            {
                console.log(err);
            }
        }
  }
    
  //POST request
  const getName = useCallback(() =>
  {
    fetch("http://localhost:8100/api/user/information", 
        {
          method: "POST",
            body: JSON.stringify({
              id: id
          })
        }).then((res) => res.json()).then(data => 
        {
          setName(data[0].first_name + " " + data[0].last_name);
        });
  },[id]);
    
  //Using useEffect hook. This will set the default values of the form once the components are mounted
  useEffect(() =>
  {
      getName();
  }, [getName])

    return (
        <div className='list-item'>
            <div className='list-item-content'>
                { name }
            </div>
            <div className='list-item-popup'>
                <div className='list-item-delete'><MdDelete size={20} className="list-item-delete-icon" onClick={DeleteUser}/></div>
            </div>
        </div>
    )
}

export default UserRoleList
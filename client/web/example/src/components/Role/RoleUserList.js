import { useState, useEffect } from 'react';

const RoleUserList = ({id}) =>
{
  const [roleName, SetRoleName] = useState("")

//POST request
  const FetchRoleInformation = () =>
  {
    fetch("http://localhost:8100/api/role/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id:id
          })
        }).then((res) => res.json()).then(data => 
          {
            SetRoleName(data[0].role_name);
          });
    }
    
    useEffect(() =>
    {
        FetchRoleInformation();
    }, [])

    return (
        <div>
            {roleName}
        </div>
    )
}

export default RoleUserList
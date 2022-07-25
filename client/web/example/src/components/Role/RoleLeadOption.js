import { useState, useEffect } from 'react';

const RoleLeadOption = ({id, roleLeadId}) =>
{  
    const [name, setName] = useState("error");
    
  //POST request
  const getName = () =>
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
  }
    
  //Using useEffect hook. This will set the default values of the form once the components are mounted
  useEffect(() =>
  {
      getName();
  }, [])

    return (
        <option value={id} selected={id === roleLeadId? "" : "selected"}>{name}</option>
    )
}

export default RoleLeadOption
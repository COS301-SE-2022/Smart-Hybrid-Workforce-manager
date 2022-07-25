import { useState, useEffect, useCallback } from 'react';

const RoleLeadOption = ({id, roleLeadId}) =>
{  
    const [name, setName] = useState("error");
    
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
        <option value={id} selected={id === roleLeadId? "" : "selected"}>{name}</option>
    )
}

export default RoleLeadOption
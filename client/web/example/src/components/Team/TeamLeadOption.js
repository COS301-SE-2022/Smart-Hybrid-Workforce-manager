import { useState, useEffect } from 'react';
import { MdDelete } from 'react-icons/md'

const TeamLeadOption = ({id, teamLeadId}) =>
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
        <option value={id} selected={id == teamLeadId? "" : "selected"}>{name}</option>
    )
}

export default TeamLeadOption
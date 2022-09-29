import { useState, useEffect, useCallback, useContext } from 'react';
import { UserContext } from '../../App';

const RoleLeadOption = ({id, roleLeadId}) =>
{  
    const [name, setName] = useState("error");
    const {userData} = useContext(UserContext);
  //POST request
  const getName = useCallback(() =>
  {
    fetch("http://deskflow.co.za:8080/api/user/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            id: id
          }),
          headers:{
              'Content-Type': 'application/json',
              'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
          }
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
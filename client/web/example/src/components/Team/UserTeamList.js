import { useState, useEffect, useCallback, useContext } from 'react';
import { MdDelete } from 'react-icons/md';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../../App';

const UserTeamList = ({id}) =>
{  
    const [name, setName] = useState("error");
    const navigate = useNavigate();    
    const {userData} = useContext(UserContext);

  let DeleteUser = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to remove this user from this team?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/team/user/remove", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        user_id: id,
                        team_id: window.sessionStorage.getItem("TeamID")
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                    }
                });

                if(res.status === 200)
                {
                    alert("User Successfully Removed!");
                    navigate(0);
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
    fetch("http://localhost:8080/api/user/information", 
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
  }, [getName, id])

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

export default UserTeamList
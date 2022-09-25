import { useState, useEffect, useContext } from 'react';
import { MdDelete } from 'react-icons/md';
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../../App';

const TeamUserList = ({teamId}) =>
{
  const [teamName, SetTeamName] = useState("");
  const navigate = useNavigate();
  const {userData} = useContext(UserContext);
    useEffect(() =>
    {
        fetch("http://localhost:8080/api/team/information", 
        {
        method: "POST",
        body: JSON.stringify({
            id: teamId
        })
        }).then((res) => res.json()).then(data => 
        {
            SetTeamName(data[0].name);
        });
    }, [teamId])
  
  let DeleteTeam = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to remove this team?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/team/user/remove", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                        team_id: teamId,
                        user_id: window.sessionStorage.getItem("UserID")
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                    }
                });

                if(res.status === 200)
                {
                    alert("Team Successfully Removed!");
                    navigate(0);
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
                { teamName }
            </div>
            <div className='list-item-popup'>
                <div className='list-item-delete'><MdDelete size={20} className="list-item-delete-icon" onClick={DeleteTeam}/></div>
            </div>
        </div>
    )
}

export default TeamUserList
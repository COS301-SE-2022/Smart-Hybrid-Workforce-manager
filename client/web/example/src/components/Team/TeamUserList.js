import { useState, useEffect } from 'react';
import { MdDelete } from 'react-icons/md';
import { useNavigate } from 'react-router-dom';

const TeamUserList = ({teamId}) =>
{
  const [teamName, SetTeamName] = useState("");
  const navigate = useNavigate();

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
                    body: JSON.stringify({
                        team_id: teamId,
                        user_id: window.sessionStorage.getItem("UserID")
                    })
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
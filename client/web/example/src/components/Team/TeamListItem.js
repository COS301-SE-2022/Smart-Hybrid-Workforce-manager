import React, { useContext } from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { MdSupervisedUserCircle } from 'react-icons/md'
import { useNavigate } from 'react-router-dom';
import { UserContext } from '../../App';

const TeamListItem = ({id, name, description, capacity, picture, lead, priority}) =>
{
    const navigate = useNavigate();
    const {userData} = useContext(UserContext);
    let EditTeam = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("TeamID", id);
        window.sessionStorage.setItem("TeamName", name);
        window.sessionStorage.setItem("TeamDescription", description);
        window.sessionStorage.setItem("TeamCapacity", capacity);
        window.sessionStorage.setItem("TeamPriority", priority);
        window.sessionStorage.setItem("TeamPicture", picture);
        window.sessionStorage.setItem("TeamLead", lead);
        navigate("/team-edit");
    }

    let DeleteTeam = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this team?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/team/remove", 
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
                });

                if(res.status === 200)
                {
                    alert("Team Successfully Deleted!");
                    navigate("/team");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    return (
        <div>
            <div className="resource">
                <div className='resource-container'>
                    <div className='resource-name'>{name}</div>
                    <MdSupervisedUserCircle className='resource-icon' size={50} />
                </div>                
                <div className='resource-popup'>
                    <div className='resource-edit'><MdEdit size={30} className="resource-edit-icon" onClick={EditTeam}/></div>
                    <div className='resource-delete'><MdDelete size={30} className="resource-delete-icon" onClick={DeleteTeam}/></div>
                </div>
            </div>
        </div>
    )
}

export default TeamListItem
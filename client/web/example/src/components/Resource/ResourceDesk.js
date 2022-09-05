import React, { useContext } from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { MdDesktopWindows } from 'react-icons/md'
import { useNavigate } from 'react-router-dom'
import { UserContext } from '../../App'

const ResourceDesk = ({id, name, location, roomId}) => {
const navigate=useNavigate();

const {userData} = useContext(UserContext);

    let EditResource = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("DeskID", id);
        window.sessionStorage.setItem("DeskName", name);
        window.sessionStorage.setItem("DeskLocation", location);
        window.sessionStorage.setItem("RoomID", roomId);
        navigate("/resources-desk-edit")
    }

    let DeleteResource = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this desk?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/resource/remove", 
                {
                    method: "POST",
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
                    alert("Resource Successfully Deleted!");
                    navigate("/resources");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    return (
        <div className="resource">
            <div className='resource-container'>
                <div className='resource-name'>{'Desk ' + name}</div>
                <MdDesktopWindows className='resource-icon' size={50} />
            </div>                
            <div className='resource-popup'>
                <div className='resource-edit'><MdEdit size={30} className="resource-edit-icon" onClick={EditResource}/></div>
                <div className='resource-delete'><MdDelete size={30} className="resource-delete-icon" onClick={DeleteResource}/></div>
            </div>
        </div>
    )
}

export default ResourceDesk
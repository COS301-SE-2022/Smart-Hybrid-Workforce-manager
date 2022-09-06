import React from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { MdSupervisorAccount } from 'react-icons/md'
import { useNavigate } from 'react-router-dom'

const ResourceMeetingRoom = ({id, name, location, capacity, roomId}) => {

    const navigate=useNavigate();

    let EditResource = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("MeetingRoomID", id);
        window.sessionStorage.setItem("MeetingRoomName", name);
        window.sessionStorage.setItem("MeetingRoomLocation", location);
        window.sessionStorage.setItem("MeetingRoomCapacity", capacity);
        window.sessionStorage.setItem("RoomID", roomId);
        navigate("/resources-meeting-room-edit");
    }

    let DeleteResource = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this meeting room?"))
        {
            try
            {
                let res = await fetch("http://localhost:8080/api/resource/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                    id: id
                    })
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
        <div>
            <div className="resource">
                <div className='resource-container'>
                    <div className='resource-name'>{'Room ' + name}</div>
                    <MdSupervisorAccount className='resource-icon' size={50} />
                </div>                
                <div className='resource-popup'>
                    <div className='resource-edit'><MdEdit size={30} className="resource-edit-icon" onClick={EditResource}/></div>
                    <div className='resource-delete'><MdDelete size={30} className="resource-delete-icon" onClick={DeleteResource}/></div>
                </div>
            </div>
        </div>
    )
}

export default ResourceMeetingRoom
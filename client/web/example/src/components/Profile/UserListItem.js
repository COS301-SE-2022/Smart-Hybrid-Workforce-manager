import React from 'react'
import { MdEdit } from 'react-icons/md'
import { MdAccountCircle } from 'react-icons/md'

const UserListItem = ({id, name, email}) => {

    let EditUser = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("UserID", id);
        window.sessionStorage.setItem("UserName", name);
        window.sessionStorage.setItem("UserEmail", email);
        window.location.assign("./user-edit");
    }

    return (
        <div>
            <div className="resource">
                <div className='resource-container'>
                    <div className='resource-name'>{name}</div>
                    <MdAccountCircle className='resource-icon' size={50} />
                </div>                
                <div className='user-popup'>
                    <div className='resource-edit'><MdEdit size={30} className="resource-edit-icon" onClick={EditUser}/></div>
                </div>
            </div>
        </div>
    )
}

export default UserListItem
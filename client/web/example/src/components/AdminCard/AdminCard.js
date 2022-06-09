import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'
import { MdAccountCircle, MdSupervisedUserCircle, MdHomeFilled, MdPermIdentity } from 'react-icons/md'

const AdminCard = ({name, description, path, type}) =>
{
    let navigate = useNavigate();
    const route = () =>
    {
        navigate(path);
    }

    function icon()
    {
        if (type === 'Users')
            return <MdAccountCircle className='admin-icon' size={50} />
        if (type === 'Teams')
            return <MdSupervisedUserCircle className='admin-icon' size={50} />
        if (type === 'Resources')
            return <MdHomeFilled className='admin-icon' size={50} />
        return <MdPermIdentity className='admin-icon' size={50} />
    }


    return (
        <div>
            <div className="admin-card" onClick={route}>
                <div className="admin-card-image">
                    {icon()}
                </div>
                
                <div className="admin-card-text">
                    <h2>{name}</h2>
                    <p>{description}</p>
                </div>
                <div className="admin-card-arrow">
                    <FaLongArrowAltRight size={50}/>
                </div>
            </div>
        </div>
    )
}

export default AdminCard
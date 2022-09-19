import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'
import { MdAccountCircle, MdSupervisedUserCircle, MdHomeFilled, MdPermIdentity } from 'react-icons/md'

const AdminCard = ({name, description, path, type}) =>
{
    let navigate = useNavigate();
    const route = () =>
    {
        if(path === '/resources')
        {
            window.open(path);
        }
        else
        {
            navigate(path);
        }
    }

    function icon()
    {
        if (type === 'Users')
            return <MdAccountCircle className='admin-icon' size={50} data-testid='admin-icon-users' />
        if (type === 'Teams')
            return <MdSupervisedUserCircle className='admin-icon' size={50} data-testid='admin-icon-teams' />
        if (type === 'Resources')
            return <MdHomeFilled className='admin-icon' size={50} data-testid='admin-icon-resources' />
        return <MdPermIdentity className='admin-icon' size={50} data-testid='admin-icon-default' />
    }


    return (
        <div>
            <div className="admin-card" onClick={route} data-testid='admin-card'>
                <div className="admin-card-image">
                    {icon()}
                </div>
                
                <div className="admin-card-text">
                    <h2 data-testid='admin-card-text-header'>{name}</h2>
                    <p data-testid='admin-card-text-body'>{description}</p>
                </div>
                <div className="admin-card-arrow">
                    <FaLongArrowAltRight size={50}/>
                </div>
            </div>
        </div>
    )
}

export default AdminCard
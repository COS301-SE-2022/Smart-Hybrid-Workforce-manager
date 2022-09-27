import ProfileBar from '../components/Navbar/ProfileBar.js';
import Navbar from '../components/Navbar/Navbar.js';
import NavbarAdmin from '../components/Navbar/NavbarAdmin.js';
import { useContext } from 'react';
import { UserContext } from '../App.js';
import Kanban from '../components/Kanban/Kanban.js';

const Admin = () =>
{
    const {userData} = useContext(UserContext);

    const showNavbar = () =>
    {
        if(userData !== null && !userData.user_identifier.includes("admin"))
        {
            return <Navbar />;
        }
        else
        {
            return <NavbarAdmin />;
        }
    };

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                {showNavbar()}

                <div className='main-container'>
                    <Kanban />
                </div>
                
                {/*<div className='admin-card-container'>
                    <AdminCard name='Users' description='Create and manage users.' path='/users' type='Users'/>
                
                    <AdminCard name='Teams' description='Create and manage teams.' path='/team' type='Teams'/>
                </div>
                <div className='admin-card-container'>
                    <AdminCard name='Resources' description='Create and manage resources.' path='/resources' type='Resources'/>
                
                    <AdminCard name='Roles' description='Create and manage roles.' path='/role' type='Roles'/>
                            </div>*/}

                
            </div>
        </div>
    )
}

export default Admin
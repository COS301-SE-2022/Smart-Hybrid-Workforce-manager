import { useContext } from 'react';
import Navbar from '../components/Navbar/Navbar.js';
import { UserContext } from '../App';
import ProfileBar from '../components/Navbar/ProfileBar';
import NavbarAdmin from '../components/Navbar/NavbarAdmin.js';
import ProfileComponent from '../components/Profile/ProfileComponent.js';

function Profile()
{  
    const {userData, setUserData} = useContext(UserContext)

    const showNavbar = () =>
    {
        if(!userData.user_identifier.includes("admin"))
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
                    <ProfileComponent />
                </div>
            </div>
        </div>
    )
}

export default Profile
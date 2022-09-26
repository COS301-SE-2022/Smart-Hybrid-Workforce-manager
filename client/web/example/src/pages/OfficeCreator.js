import { useContext } from 'react';
import { UserContext } from '../App';
import ProfileBar from '../components/Navbar/ProfileBar.js';
import Navbar from '../components/Navbar/Navbar.js';
import NavbarAdmin from '../components/Navbar/NavbarAdmin';
import Creator from '../components/Map/Creator';

const OfficeCreator = () =>
{
    const {userData} = useContext(UserContext);
    
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
                    <Creator />
                </div>
            </div>  
        </div>
    )
}

export default OfficeCreator
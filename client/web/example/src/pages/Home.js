import { useContext } from 'react';
import { UserContext } from '../App';
import ProfileBar from '../components/Navbar/ProfileBar.js';
import Navbar from '../components/Navbar/Navbar.js';
import NavbarAdmin from '../components/Navbar/NavbarAdmin';
import Map from '../components/Map/Map';

const Home = () =>
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
                    <Map />
                </div>
            </div>  
        </div>
    )
}

export default Home
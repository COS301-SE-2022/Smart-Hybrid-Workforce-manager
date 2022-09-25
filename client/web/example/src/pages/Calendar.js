import Navbar from '../components/Navbar/Navbar.js'
import ProfileBar from '../components/Navbar/ProfileBar.js';
import NavbarAdmin from '../components/Navbar/NavbarAdmin.js';
import CalendarComponent from '../components/Calendar/CalendarComponent.js';
import { UserContext } from '../App.js';
import { useContext } from 'react';

const Calendar = () =>
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
                    <CalendarComponent />
                </div>
            </div>
        </div>
    )
}

export default Calendar
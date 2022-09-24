import Navbar from '../components/Navbar/Navbar.js';
import NavbarAdmin from '../components/Navbar/NavbarAdmin.js';
import { useContext } from 'react';
import { UserContext } from '../App.js';
import StatisticsComponent from '../components/StatisticsComponent/StatisticsComponent.js';

const Statistics = () =>
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
                {showNavbar()} 

                <div className='main-container'>
                    <StatisticsComponent />
                </div>
            </div>
        </div>
    )
}

export default Statistics
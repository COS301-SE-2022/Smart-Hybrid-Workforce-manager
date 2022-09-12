import Navbar from "../components/Navbar/Navbar.js"
import DeskBooking from "../components/BookingForm/DeskBooking"
import { useContext, useRef } from "react"
import ProfileBar from "../components/Navbar/ProfileBar.js";
import NavbarAdmin from "../components/Navbar/NavbarAdmin.js";
import { UserContext } from '../App';

const BookingsDesk = () =>
{
    const deskRef = useRef(null);
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
                <div className="main-container">
                    <DeskBooking ref={deskRef}/>
                </div>
            </div>  
        </div>
    )
}

export default BookingsDesk
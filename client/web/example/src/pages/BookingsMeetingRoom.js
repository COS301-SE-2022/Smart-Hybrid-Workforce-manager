import Navbar from "../components/Navbar/Navbar.js"
import { useContext, useEffect, useRef } from "react"
import MeetingRoomBooking from "../components/BookingForm/MeetingRoomBooking"
import ProfileBar from "../components/Navbar/ProfileBar.js";
import NavbarAdmin from "../components/Navbar/NavbarAdmin.js";
import { UserContext } from '../App';;

const BookingsMeetingRoom = () =>
{
    const deskRef = useRef(null);
    const mainRef = useRef(null);
    const {userData} = useContext(UserContext);

    useEffect(() =>
    {
        mainRef.current.style.overflowY = 'scroll';
    },[])

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
                <div ref={mainRef} className="main-container">
                    <MeetingRoomBooking ref={deskRef}/>
                </div>
            </div>  
        </div>
    )
}

export default BookingsMeetingRoom
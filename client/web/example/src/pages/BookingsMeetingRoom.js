import Navbar from "../components/Navbar/Navbar.js"
import { useEffect, useRef } from "react"
import MeetingRoomBooking from "../components/BookingForm/MeetingRoomBooking"
import ProfileBar from "../components/Navbar/ProfileBar.js";

const BookingsMeetingRoom = () =>
{
    const deskRef = useRef(null);
    const mainRef = useRef(null);

    useEffect(() =>
    {
        mainRef.current.style.overflowY = 'scroll';
    },[])

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                <Navbar />
                <div ref={mainRef} className="main-container">
                    <MeetingRoomBooking ref={deskRef}/>
                </div>
            </div>  
        </div>
    )
}

export default BookingsMeetingRoom
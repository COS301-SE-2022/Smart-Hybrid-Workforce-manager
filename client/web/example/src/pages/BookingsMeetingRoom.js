import Navbar from "../components/Navbar/Navbar.js"
import { useRef } from "react"
import MeetingRoomBooking from "../components/BookingForm/MeetingRoomBooking"

const BookingsMeetingRoom = () =>
{
    const deskRef = useRef(null);

    return (
        <div className='page-container'>
            <div className='content'>
                <Navbar />
                <div className="main-container">
                    <MeetingRoomBooking ref={deskRef}/>
                </div>
            </div>  
        </div>
    )
}

export default BookingsMeetingRoom
import Navbar from "../components/Navbar/Navbar.js"
import DeskBooking from "../components/BookingForm/DeskBooking"
import { useRef } from "react"
import ProfileBar from "../components/Navbar/ProfileBar.js";

const BookingsDesk = () =>
{
    const deskRef = useRef(null);

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
                <Navbar />
                <div className="main-container">
                    <DeskBooking ref={deskRef}/>
                </div>
            </div>  
        </div>
    )
}

export default BookingsDesk
import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import Background from "../components/Shapes/Background"
import DeskBooking from "../components/BookingForm/DeskBooking"
import { FaLongArrowAltRight, FaLongArrowAltLeft } from 'react-icons/fa'
import { useRef, useEffect } from "react"
import MeetingRoomBooking from "../components/BookingForm/MeetingRoomBooking"

const Bookings = () =>
{
    const deskRef = useRef(null);
    const meetingRoomRef = useRef(null);
    const deskCardRef = useRef(null);
    const meetingRoomCardRef = useRef(null);
    const backRef = useRef(null);

    const ShowDeskForm = () =>
    {
        if(deskRef.current != null)
        {
            deskRef.current.style.display = 'block';
        }

        if(backRef.current != null)
        {
            backRef.current.style.display = 'block';
        }

        if(deskCardRef.current != null)
        {
            deskCardRef.current.style.display = 'none';
        }

        if(meetingRoomCardRef.current != null)
        {
            meetingRoomCardRef.current.style.display = 'none';
        }
    }

    const ShowMeetingRoomForm = () =>
    {
        if(meetingRoomRef.current != null)
        {
            meetingRoomRef.current.style.display = 'block';
        }

        if(backRef.current != null)
        {
            backRef.current.style.display = 'block';
        }

        if(deskCardRef.current != null)
        {
            deskCardRef.current.style.display = 'none';
        }

        if(meetingRoomCardRef.current != null)
        {
            meetingRoomCardRef.current.style.display = 'none';
        }
    }

    const GoBack = () =>
    {
        if(deskRef.current != null)
        {
            deskRef.current.style.display = 'none';
        }

        if(meetingRoomRef.current != null)
        {
            meetingRoomRef.current.style.display = 'none';
        }

        if(backRef.current != null)
        {
            backRef.current.style.display = 'none';
        }

        if(deskCardRef.current != null)
        {
            deskCardRef.current.style.display = 'grid';
        }

        if(meetingRoomCardRef.current != null)
        {
            meetingRoomCardRef.current.style.display = 'grid';
        }
    }

    useEffect(() =>
    {
        if(deskRef.current != null)
        {
            deskRef.current.style.display = 'none';
        }

        if(meetingRoomRef.current != null)
        {
            meetingRoomRef.current.style.display = 'none';
        }

        if(backRef.current != null)
        {
            backRef.current.style.display = 'none';
        }

        if(deskCardRef.current != null)
        {
            deskCardRef.current.style.display = 'grid';
        }

        if(meetingRoomCardRef.current != null)
        {
            meetingRoomCardRef.current.style.display = 'grid';
        }
    }, [])

    return (
        <div className='page-container'>
            <div className='content'>
                <Navbar />


                <div className='card-container'>
                    <div ref={deskCardRef} className="card" onClick={ShowDeskForm}>
                        <div className="card-image" 
                            style={{
                                gridArea : 'image',
                                background : 'linear-gradient(#fff0 0%, #fff0 70%, #1d1d1d 100%), url(https://introducingsa.co.za/wp-content/uploads/sites/142/2022/03/Home-office.png)',
                                'backgroundSize': 'cover',
                                'borderTopLeftRadius': '4vh',
                                'borderTopRightRadius': '4vh'
                                }}>
                        </div>
                        
                        <div className="card-text">
                            <h2>Desk</h2>
                            <p>Never arrive to the office and not have a desk to work at. Book a desk and let the Smart Schedular pair you with your team.</p>
                        </div>

                        <div className="card-arrow">
                            <FaLongArrowAltRight size={50}/>
                        </div>
                    </div>

                    <div ref={meetingRoomCardRef} className="card" onClick={ShowMeetingRoomForm}>
                        <div className="card-image" 
                            style={{
                                gridArea : 'image',
                                background : 'linear-gradient(#fff0 0%, #fff0 70%, #1d1d1d 100%), url(https://synivate.com/wp-content/uploads/conference-room-meetings-1-400x249.jpg)',
                                'backgroundSize': 'cover',
                                'borderTopLeftRadius': '4vh',
                                'borderTopRightRadius': '4vh'
                                }}>
                        </div>
                        
                        <div className="card-text">
                            <h2>Meeting Room</h2>
                            <p>Secure your meeting room and prevent those dreaded delays or reschedules. Or choose automation and let the Smart Schedular automate your bookings.</p>
                        </div>
                        
                        <div className="card-arrow">
                            <FaLongArrowAltRight size={50}/>
                        </div>
                    </div>
                </div>
                <DeskBooking ref={deskRef}/>
                <MeetingRoomBooking ref={meetingRoomRef} />

                <div ref={backRef} className='back-button' onClick={GoBack}>
                    <FaLongArrowAltLeft size={50} color={'#374146'}/>
                </div>
            </div>  
            <Footer />
        </div>
    )
}

export default Bookings
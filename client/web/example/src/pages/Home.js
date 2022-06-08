import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import BookingTicket from '../components/BookingTicket/BookingTicket';
import { useState, useEffect } from 'react';

const Home = () =>
{
  const [bookings, setBookings] = useState([])

  //POST request
  const fetchData = () =>
  {
    fetch("http://localhost:8100/api/booking/information", 
        {
          method: "POST",
          body: JSON.stringify({
            id: null,
            user_id: null
          })
        }).then((res) => res.json()).then(data => 
          {
            setBookings(data);
            window.sessionStorage.removeItem("BookingID");
            window.sessionStorage.removeItem("StartDate");
            window.sessionStorage.removeItem("StartTime");
            window.sessionStorage.removeItem("EndDate");
            window.sessionStorage.removeItem("EndTime");
          });
  }

  //Using useEffect hook. This will send the POST request once the component is mounted
  useEffect(() =>
  {
    fetchData()
  }, [])
  

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='booking-container'>
          {bookings.length > 0 && (
            bookings.map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked}/>
            ))
          )}
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Home
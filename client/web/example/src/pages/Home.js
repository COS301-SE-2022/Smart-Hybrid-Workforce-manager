import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import Button from 'react-bootstrap/Button';
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
            user_id: "11111111-dc08-4a06-9983-8b374586e459"
          })
        }).then((res) => res.json()).then(data => 
          {
            console.log(data);
            setBookings(data);
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
        <Button className='button-booking' variant='primary'>Refresh Bookings</Button>
        <div className='booking-container'>
          {bookings.length > 0 && (
            bookings.map(booking => (
              <BookingTicket startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked}/>
            ))
          )}

        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Home
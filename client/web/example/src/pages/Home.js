import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import BookingTicket from '../components/BookingTicket/BookingTicket';
import { useState, useEffect, useContext } from 'react';
import { UserContext } from '../App';

const Home = () =>
{
  const [bookings, setBookings] = useState([])
  const {userData} = useContext(UserContext);

  const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];

  let DayOne = new Date()
  DayOne.setHours(23, 59, 59, 59)

  let DayTwo = new Date()
  DayTwo.setHours(23, 59, 59, 59)
  DayTwo.setDate(DayTwo.getDate() + 1);

  let DayThree = new Date()
  DayThree.setHours(23, 59, 59, 59)
  DayThree.setDate(DayThree.getDate() + 2);

  let DayFour = new Date()
  DayFour.setHours(23, 59, 59, 59)
  DayFour.setDate(DayFour.getDate() + 3);

  let DayFive = new Date()
  DayFive.setHours(23, 59, 59, 59)
  DayFive.setDate(DayFive.getDate() + 4);

  let DaySix = new Date()
  DaySix.setHours(23, 59, 59, 59)
  DaySix.setDate(DaySix.getDate() + 5);

  let DaySeven = new Date()
  DaySeven.setHours(23, 59, 59, 59)
  DaySeven.setDate(DaySeven.getDate() + 6);

  //POST request
  const fetchData = () =>
  {
    let startDate = new Date()    
    startDate.setHours(0, 0, 0, 0)

    let endDate = new Date()
    endDate.setHours(26, 0, 0, 0)
    endDate.setDate(endDate.getDate() + 1 * 7);

    fetch("http://localhost:8100/api/booking/information", 
        {
          method: "POST",
          mode: "cors",
          body: JSON.stringify({
            start: startDate.toISOString(),
            end: endDate.toISOString()
          }),
          headers:{
            'Authorization': `bearer ${userData.token}`
          }
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
          <h2 className='white'>{DayOne.getDate() + " " + monthNames[DayOne.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DayOne.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DayTwo.getDate() + " " + monthNames[DayTwo.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DayTwo.toISOString() && booking.start > DayOne.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DayThree.getDate() + " " + monthNames[DayThree.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DayThree.toISOString() && booking.start > DayTwo.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DayFour.getDate() + " " + monthNames[DayFour.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DayFour.toISOString() && booking.start > DayThree.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DayFive.getDate() + " " + monthNames[DayFive.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DayFive.toISOString() && booking.start > DayFour.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DaySix.getDate() + " " + monthNames[DaySix.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DaySix.toISOString() && booking.start > DayFive.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}

          <br></br>
          <h2 className='white'>{DaySeven.getDate() + " " + monthNames[DaySeven.getMonth()]}</h2> <hr></hr>
          {bookings.length > 0 && (
            bookings.filter(booking => (booking.start < DaySeven.toISOString() && booking.start > DaySix.toISOString())).map(booking => (
              <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endDate={booking.end.substring(0,10)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type}/>
            ))
          )}
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Home
import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import BookingCard from "../components/BookingCard/BookingCard"

import React, { useState, useEffect, useContext } from "react";
import userContext from '../store/userContext';

function Bookings()
{
  // const { logout, isLoggedIn } = useContext(userContext);
  // if(!isLoggedIn){
  //   // alert("not logged in");
  //   // window.location.assign("./login");
  // }

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='card-container'>
          <BookingCard name='Desk' description='Never arrive to the office and not have a desk to work at. Book a desk and let the Smart Schedular pair you with your team.' 
          path='/bookings-desk' image='https://introducingsa.co.za/wp-content/uploads/sites/142/2022/03/Home-office.png'/>

          <BookingCard name='Meeting Room' description='Secure your meeting room and prevent those dreaded delays or reschedules. Or choose automation and let the Smart Schedular automate your bookings.' 
          path='/bookings-meeting' image='https://synivate.com/wp-content/uploads/conference-room-meetings-1-400x249.jpg'/>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Bookings
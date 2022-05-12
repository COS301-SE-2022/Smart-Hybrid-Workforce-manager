import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'

function Home()
{
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='booking-container'>
          <div className="booking">
            <div className="card-text">
              <h2>Desk</h2>
              <p>Never arrive to the office and not have a desk to work at. Book a desk and let the Smart Schedular pair you with your team.</p>
            </div>
          </div>

          <div className="booking">
            <div className="card-text">
              <h2>Meeting Room</h2>
              <p>Secure your meeting room and prevent those dreaded delays or reschedules. Or choose automation and let the Smart Schedular automate your bookings.</p>
            </div>
          </div>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Home
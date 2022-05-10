import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'

const Bookings = () => {
  return (
    <div className='content'>
      <div className='card-container'>
        <div className="card">
          <div className="card-image"></div>
          <div className="card-text">
            <h2>Desk</h2>
            <p>Never arrive to the office and not have a desk to work. Book a desk and let the Smart Schedular pair you with your team.</p>
          </div>
          <div className="card-arrow">
            <FaLongArrowAltRight size={50}/>
          </div>
        </div>

        <div className="card">
          <div className="card-image2"></div>
          <div className="card-text">
            <h2>Meeting Room</h2>
            <p>Secure your meeting room and prevent those dreaded delays or reschedules. Let the Smart Schedular access and calendar and automate the bookings.</p>
          </div>
          <div className="card-arrow">
            <FaLongArrowAltRight size={50}/>
          </div>
        </div>
      </div>
    </div>
  )
}

export default Bookings
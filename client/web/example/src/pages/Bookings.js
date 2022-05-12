import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'

function Bookings()
{
  let navigate = useNavigate();
  const routeDesk = () =>
  {
    let path = "/bookings-desk";
    navigate(path);
  }

  const routeMeeting = () =>
  {
    let path = "/bookings-meeting";
    navigate(path);
  }

  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <div className='card-container'>
          <div className="card" onClick={routeDesk}>
            <div className="card-image"></div>
            <div className="card-text">
              <h2>Desk</h2>
              <p>Never arrive to the office and not have a desk to work at. Book a desk and let the Smart Schedular pair you with your team.</p>
            </div>
            <div className="card-arrow">
              <FaLongArrowAltRight size={50}/>
            </div>
          </div>

          <div className="card" onClick={routeMeeting}>
            <div className="card-image2"></div>
            <div className="card-text">
              <h2>Meeting Room</h2>
              <p>Secure your meeting room and prevent those dreaded delays or reschedules. Or choose automation and let the Smart Schedular automate your bookings.</p>
            </div>
            <div className="card-arrow">
              <FaLongArrowAltRight size={50}/>
            </div>
          </div>
        </div>
      </div>  
      <Footer />
    </div>
  )
}

export default Bookings
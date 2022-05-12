import Navbar from "../components/Navbar"
import Footer from "../components/Footer"
import React from 'react'
import { FaLongArrowAltRight } from 'react-icons/fa'
import { useNavigate } from 'react-router-dom'

function Bookings()
{
  let navigate = useNavigate();
  /*const routeDesk = () =>
  {
    let path = "/bookings-desk";
    navigate(path);
  }*/

  let handleSubmit = async (e) =>
  {
    e.preventDefault();
    try
    {
      let res = await fetch("http://localhost:8100/api/booking/create", 
      {
        method: "POST",
        body: JSON.stringify({
          id: "33333333-dc08-4a06-9983-8b374586e453",
          user_id: "11111111-dc08-4a06-9983-8b374586e459",
          resource_type: "DESK",
          resource_preference_id: null,
          start: "2012-04-23T18:25:43.511Z",
          end: "2012-04-23T18:25:43.511Z",
          booked: false
        })
      });

      if(res.status === 200)
      {
        alert("Booking Successfully Created!");
      }
    }
    catch(err)
    {
      console.log(err);
    }
  };

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
          <div className="card" onClick={handleSubmit}>
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
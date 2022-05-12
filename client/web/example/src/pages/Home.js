import React from 'react'
import Navbar from '../components/Navbar'
import Footer from '../components/Footer'
import Button from 'react-bootstrap/Button';
import { Component } from 'react';

class Home extends Component
{
  constructor(props)
  {
    super(props);
    this.state = 
    {
      sDate1: 0,
      sTime1: 0,
      eDate1: 0,
      eTime1: 0,

      sDate2: 0,
      sTime2: 0,
      eDate2: 0,
      eTime2: 0
    }
  }

  refresh = () =>
  {
    try
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
          this.setState({sDate1: data[0].start.substring(0,10),
                        sTime1: data[0].start.substring(11,16),
                        eDate1: data[0].end.substring(0,10),
                        eTime1: data[0].end.substring(11,16),
                        sDate2: data[1].start.substring(0,10),
                        sTime2: data[1].start.substring(11,16),
                        eDate2: data[1].end.substring(0,10),
                        eTime2: data[1].end.substring(11,16)});
        });
    }
    catch(err)
    {
      console.log(err);
    }
  }

  render()
  {
  return (
    <div className='page-container'>
      <div className='content'>
        <Navbar />
        <Button className='button-booking' variant='primary' onClick={this.refresh}>Refresh Bookings</Button>
        <div className='booking-container'>
          <div className="booking">
            <div className="booking-image"></div>
            <div className="booking-text">
              <p>Start Date: { this.state.sDate1 }</p>
              <p>Start Time: { this.state.sTime1 }</p>
              <p>End Date: { this.state.eDate1 }</p>
              <p>End Time: { this.state.eTime1 }</p>
              <p>Confirmed: Pending</p>
            </div>
          </div>

          <div className="booking">
            <div className="booking-image"></div>
            <div className="booking-text">
              <p>Start Date: { this.state.sDate2 }</p>
              <p>Start Time: { this.state.sTime2 }</p>
              <p>End Date: { this.state.eDate2 }</p>
              <p>End Time: { this.state.eTime2 }</p>
              <p>Confirmed: Pending</p>
            </div>
          </div>
        </div>
      </div>  
      <Footer />
    </div>
  )
  }
}

export default Home
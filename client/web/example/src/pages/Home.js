import Navbar from '../components/Navbar/Navbar.js'
import Footer from '../components/Footer'
import BookingTicket from '../components/BookingTicket/BookingTicket';
import { useState, useEffect, useContext, useRef } from 'react';
import { UserContext } from '../App';

import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';

const Home = () =>
{
    const [bookings, setBookings] = useState([])
    const {userData} = useContext(UserContext);

    const monthSelectorRef = useRef(null);
    const weekSelectorRef = useRef(null);
    const [currentContext, setContext] = useState("week");

    const SelectMonth = () =>
    {
        setContext("month");
    }

    const SelectWeek = () =>
    {
        setContext("week");
    }

    const MouseOverMonth = () =>
    {
        monthSelectorRef.current.style.backgroundColor = "#09a2fb";
        monthSelectorRef.current.style.color = "#ffffff";
    }

    const MouseLeaveMonth = () =>
    {
        if(currentContext !== "month")
        {
            monthSelectorRef.current.style.backgroundColor = "#ffffff";
            monthSelectorRef.current.style.color = "#09a2fb";
        }
    }

    const MouseOverWeek = () =>
    {
        weekSelectorRef.current.style.backgroundColor = "#09a2fb";
        weekSelectorRef.current.style.color = "#ffffff";
    }

    const MouseLeaveWeek = () =>
    {
        if(currentContext !== "week")
        {
            weekSelectorRef.current.style.backgroundColor = "#ffffff";
            weekSelectorRef.current.style.color = "#09a2fb";
        }
    }

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
            user_id: window.sessionStorage.getItem("UserID"),
            start: startDate.toISOString(),
            end: endDate.toISOString()
          }),
          /*headers:{
            'Content-Type': 'application/json',
            'Authorization': `bearer ${userData.token}`
          }*/
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
        fetchData();
    }, [])

    useEffect(() =>
    {
        if(currentContext === "month")
        {
            monthSelectorRef.current.style.backgroundColor = "#09a2fb";
            monthSelectorRef.current.style.color = "#ffffff";
            weekSelectorRef.current.style.backgroundColor = "#ffffff";
            weekSelectorRef.current.style.color = "#09a2fb";
        }
        else
        {
            weekSelectorRef.current.style.backgroundColor = "#09a2fb";
            weekSelectorRef.current.style.color = "#ffffff";
            monthSelectorRef.current.style.backgroundColor = "#ffffff";
            monthSelectorRef.current.style.color = "#09a2fb";
        }
    }, [currentContext])
  

    return (
        <div className='page-container'>
            <div className='content'>
                <Navbar />

                <div className='main-container'>
                    <div className='calendar-container'>
                        <div className='top-bar'>
                            <div className='calendar-title'>
                                August 2022
                            </div>

                            <div className='context-container'>
                                <div ref={monthSelectorRef} className='month-selector' onClick={SelectMonth} onMouseOver={MouseOverMonth} onMouseLeave={MouseLeaveMonth}>
                                    Month
                                </div>
                                <div ref={weekSelectorRef}  className='week-selector' onClick={SelectWeek} onMouseOver={MouseOverWeek} onMouseLeave={MouseLeaveWeek}>
                                    Week
                                </div>
                            </div>

                            <div className='nav-container'>
                                <div className='prev'>
                                    <IoIosArrowBack />
                                </div>
                                <div className='next'>
                                    <IoIosArrowForward />
                                </div>
                            </div>
                        </div>


                    </div>
                </div>
            </div>
        </div>
    )
}

export default Home
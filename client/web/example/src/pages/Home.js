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

    const [date, setDate] = useState(new Date());
    const [month, setMonth] = useState("");
    const [monthIndex, setMonthIndex] = useState();
    const [year, setYear] = useState("");

    const titleRef = useRef(null);
    const monthSelectorRef = useRef(null);
    const weekSelectorRef = useRef(null);
    const prevRef = useRef(null);
    const prevWeekRef = useRef(null);
    const prevMonthRef = useRef(null);
    const nextRef = useRef(null);
    const nextWeekRef = useRef(null);
    const nextMonthRef = useRef(null);
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

    const MouseOverPrev = () =>
    {
        prevRef.current.style.backgroundColor = "#e8e8e8";

        if(currentContext === "month")
        {
            prevMonthRef.current.style.display = "inline-block";
        }
        else
        {
            prevWeekRef.current.style.display = "inline-block";
        }
    }

    const MouseLeavePrev = () =>
    {
        prevRef.current.style.backgroundColor = "transparent";

        if(currentContext === "month")
        {
            prevMonthRef.current.style.display = "none";
        }
        else
        {
            prevWeekRef.current.style.display = "none";
        }
    }

    const PrevClick = () =>
    {
        if(currentContext === "month")
        {
            if(monthIndex === 0)
            {
                setYear(year - 1);
                setMonthIndex(11);
            }
            else
            {
                setMonthIndex((monthIndex - 1) % 12);
            }
        }
        else
        {

        }
    }

    const MouseOverNext = () =>
    {
        nextRef.current.style.backgroundColor = "#e8e8e8";

        if(currentContext === "month")
        {
            nextMonthRef.current.style.display = "inline-block";
        }
        else
        {
            nextWeekRef.current.style.display = "inline-block";
        }
    }

    const MouseLeaveNext = () =>
    {
        nextRef.current.style.backgroundColor = "transparent";

        if(currentContext === "month")
        {
            nextMonthRef.current.style.display = "none";
        }
        else
        {
            nextWeekRef.current.style.display = "none";
        }
    }

    const NextClick = () =>
    {
        if(currentContext === "month")
        {
            if(monthIndex === 11)
            {
                setYear(year + 1);
            }
            setMonthIndex((monthIndex + 1) % 12);
        }
        else
        {

        }
    }

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
        setMonthIndex(date.getMonth());
        setYear(date.getFullYear());
    }, [date])

    useEffect(() =>
    {
        const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
        setMonth(monthNames[monthIndex]);
    },[monthIndex]);

    /*useEffect(() =>
    {
        if(titleRef.current !== null)
        {
            titleRef.current.innerHTML = date.toLocaleDateString('en-GB', {
                month: 'long',
                year: 'numeric'
            });
        }
    },[date]);*/

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
                    <div className='top-bar'>
                        <div className='calendar-title'>
                            <div className='month'>{month}</div>
                            <div className='year'>{year}</div>
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
                            <div ref={prevRef} className='prev' onMouseEnter={MouseOverPrev} onMouseLeave={MouseLeavePrev} onClick={PrevClick}>
                                <IoIosArrowBack />
                                <div ref={prevWeekRef} className='tooltip-prev'>
                                    Previous Week
                                </div>
                                <div ref={prevMonthRef} className='tooltip-prev'>
                                    Previous Month
                                </div>
                            </div>

                            <div ref={nextRef} className='next' onMouseEnter={MouseOverNext} onMouseLeave={MouseLeaveNext} onClick={NextClick}>
                                <IoIosArrowForward />
                                <div ref={nextWeekRef} className='tooltip-next'>
                                    Next Week
                                </div>
                                <div ref={nextMonthRef} className='tooltip-next'>
                                    Next Month
                                </div>
                            </div>
                        </div>
                    </div>

                    <div className='calendar-content-week'>
                        <div className='days-of-week'>
                            <div className='timezone'>
                                GMT+02
                            </div>

                            <div className='day-date'>
                                <p className='day'>Sun</p>
                                <p className='date'>28</p>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Mon</p>
                                <p className='date'>29</p>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Tue</p>
                                <p className='date'>30</p>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Wed</p>
                                <div className='datee'>31</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Thu</p>
                                <p className='date'>01</p>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Fri</p>
                                <p className='date'>02</p>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Sat</p>
                                <p className='date'>03</p>
                            </div>
                        </div>

                        <div className='days-of-week-borders'>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                            <div className='day-date-border'></div>
                        </div>


                        <div className='column-container'>

                        </div>
                    </div>

                </div>
            </div>
        </div>
    )
}

export default Home
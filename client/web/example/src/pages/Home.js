import Navbar from '../components/Navbar/Navbar.js'
import BookingTicket from '../components/BookingTicket/BookingTicket';
import { useState, useEffect, useContext, useRef } from 'react';
import { UserContext } from '../App';

import { IoIosArrowBack, IoIosArrowForward } from 'react-icons/io';
import Borders from '../components/Calendar/Borders.js';
import Times from '../components/Calendar/Times.js';
import ProfileBar from '../components/Navbar/ProfileBar.js';

const Home = () =>
{
    const [bookings, setBookings] = useState([])
    const {userData} = useContext(UserContext);

    const [month, setMonth] = useState("");
    const [monthIndex, setMonthIndex] = useState();
    const [dualMonth, setDualMonth] = useState(false);
    const [year, setYear] = useState("");
    const [days, setDays] = useState([
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        },
        {
            day: -1,
            date: -1,
            month: -1,
            year: -1,
        }
    ]);
    const [currDate, setCurrDate] = useState();

    const sunRef = useRef(null);
    const monRef = useRef(null);
    const tueRef = useRef(null);
    const wedRef = useRef(null);
    const thuRef = useRef(null);
    const friRef = useRef(null);
    const satRef = useRef(null);

    const monthSelectorRef = useRef(null);
    const weekSelectorRef = useRef(null);
    const todaySelectorRef = useRef(null);
    const prevRef = useRef(null);
    const prevWeekRef = useRef(null);
    const prevMonthRef = useRef(null);
    const nextRef = useRef(null);
    const nextWeekRef = useRef(null);
    const nextMonthRef = useRef(null);
    const [currentContext, setContext] = useState("week");
    const [isToday, setToday] = useState(true);

    const indicatorRef = useRef(null);

    

    const SelectMonth = () =>
    {
        setContext("month");
    }

    const SelectWeek = () =>
    {
        setContext("week");
    }

    const SelectToday = () =>
    {
        setToday(true);
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

    const MouseOverToday = () =>
    {
        todaySelectorRef.current.style.backgroundColor = "#09a2fb";
        todaySelectorRef.current.style.color = "#ffffff";
    }

    const MouseLeaveToday = () =>
    {
        if(!isToday)
        {
            todaySelectorRef.current.style.backgroundColor = "#ffffff";
            todaySelectorRef.current.style.color = "#09a2fb";
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
            const arr = [];
            arr[6] = {
                day: 6,
                date: days[0].date - 1,
                month: days[0].month,
                year: days[0].year
            };

            const lastDay = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31];

            //Populate temp array with days
            for(var i = 6; i > -1; i--)
            {
                if(i !== 6)
                {
                    arr[i] = {
                        day: arr[i+1].day - 1,
                        date: arr[i+1].date - 1,
                        month: arr[i+1].month,
                        year: arr[i+1].year
                    };
                }

                if(arr[i].date === 0)
                {
                    arr[i].month = (arr[i].month - 1) % 12;

                    if(arr[i].month === -1)
                    {
                        arr[i].month = 11;
                        arr[i].year = arr[i].year - 1;
                    }

                    arr[i].date = lastDay[arr[i].month];

                    if((arr[i].year % 4 === 0) && arr[i].month === 1)
                    {
                        arr[i].date = 29;
                    }
                }
            }

            //Set month and year to display. Uses lower month for dual cases
            setMonthIndex(arr[0].month);
            setYear(arr[0].year);

            //Set the days array
            setDays(arr);

            //Highlight current day
            indicatorRef.current.style.display = 'none';
            setToday(false);
            const refArray = [sunRef, monRef, tueRef, wedRef, thuRef, friRef, satRef];
            for(i = 0; i < 7; i++)
            {
                if((arr[i].date === currDate.date) && (arr[i].month === currDate.month) && (arr[i].year === currDate.year))
                {
                    refArray[i].current.style.backgroundColor = '#09a2fb';
                    refArray[i].current.style.color = '#ffffff';
                    indicatorRef.current.style.display = 'block';
                    setToday(true);
                }
                else
                {
                    refArray[i].current.style.backgroundColor = 'transparent';
                    refArray[i].current.style.color = '#374146';
                }
            }

            //Check for dual month
            for(i = 1; i < 7; i++)
            {
                if(arr[i].month !== arr[0].month)
                {
                    setDualMonth(true);
                    break;
                }
                else
                {
                    setDualMonth(false);
                }
            }
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
            const arr = [];
            arr[0] = days[6] + 1;
            arr[0] = {
                day: 0,
                date: days[6].date + 1,
                month: days[6].month,
                year: days[6].year
            };

            const thirty = [3, 5, 8, 10];
            const thirtyOne = [0, 2, 4, 6, 7, 9, 11];

            var lastDay;
            if(thirty.includes(monthIndex))
            {
                lastDay = 30;
            }
            else if(thirtyOne.includes(monthIndex))
            {
                lastDay = 31;
            }
            else if(monthIndex === 1 && (year % 4 === 0))
            {
                lastDay = 29;
            }
            else if(monthIndex === 1 && (year % 4 !== 0))
            {
                lastDay = 28
            }

            //Populate temp array with days
            for(var i = 0; i < 7; i++)
            {
                if(i !== 0)
                {
                    arr[i] = {
                        day: arr[i-1].day + 1,
                        date: arr[i-1].date + 1,
                        month: arr[i-1].month,
                        year: arr[i-1].year
                    };
                }

                if(arr[i].date === lastDay + 1)
                {
                    arr[i].date = 1;
                    arr[i].month = (arr[i].month + 1) % 12;

                    if(arr[i].month === 0)
                    {
                        arr[i].year = arr[i].year + 1;
                    }
                }
            }

            //Set month and year to display. Uses lower month for dual cases
            setMonthIndex(arr[0].month);
            setYear(arr[0].year);

            //Set the days array
            setDays(arr);

            //Highlight current day
            indicatorRef.current.style.display = 'none';
            setToday(false);
            const refArray = [sunRef, monRef, tueRef, wedRef, thuRef, friRef, satRef];
            for(i = 0; i < 7; i++)
            {
                if((arr[i].date === currDate.date) && (arr[i].month === currDate.month) && (arr[i].year === currDate.year))
                {
                    refArray[i].current.style.backgroundColor = '#09a2fb';
                    refArray[i].current.style.color = '#ffffff';
                    indicatorRef.current.style.display = 'block';
                    setToday(true);
                }
                else
                {
                    refArray[i].current.style.backgroundColor = 'transparent';
                    refArray[i].current.style.color = '#374146';
                }
            }

            //Check for dual month
            for(i = 1; i < 7; i++)
            {
                if(arr[i].month !== arr[0].month)
                {
                    setDualMonth(true);
                    break;
                }
                else
                {
                    setDualMonth(false);
                }
            }
        }
    }

    const loadToday = () =>
    {
        const date = new Date();

        //Create tuple for current date
        setCurrDate({
            day: date.getDay(),
            date: date.getDate(),
            month: date.getMonth(),
            year: date.getFullYear()
        });

        //Add current date to temp array
        const arr = [];
        arr[date.getDay()] = {
            day: date.getDay(),
            date: date.getDate(),
            month: date.getMonth(),
            year: date.getFullYear()
        };

        //Arrays to track how many days in each month
        const thirty = [3, 5, 8, 10];
        const thirtyOne = [0, 2, 4, 6, 7, 9, 11];
        const currMonth = date.getMonth();

        var prevMonth = currMonth - 1;
        if(prevMonth < 0)
        {
            prevMonth = 11;
        }

        var lastDay;
        if(thirty.includes(currMonth))
        {
            lastDay = 30;
        }
        else if(thirtyOne.includes(currMonth))
        {
            lastDay = 31;
        }
        else if(currMonth === 1 && (date.getFullYear() % 4 === 0))
        {
            lastDay = 29;
        }
        else if(currMonth === 1 && (date.getFullYear() % 4 !== 0))
        {
            lastDay = 28
        }

        //Add days to array after current date
        for(var i = date.getDay() + 1; i < 7; i++)
        {
            arr[i] = {
                day: arr[i-1].day + 1,
                date: arr[i-1].date + 1,
                month: arr[i-1].month,
                year: arr[i-1].year
            };
            
            if(arr[i].date === lastDay + 1)
            {
                arr[i].date = 1;
                arr[i].month = (arr[i].month + 1) % 12;

                if(arr[i].month === 0)
                {
                    arr[i].year = arr[i].year + 1;
                }
            }
        }

        var lastDayPrev;
        if(thirty.includes(prevMonth))
        {
            lastDayPrev = 30;
        }
        else if(thirtyOne.includes(prevMonth))
        {
            lastDayPrev = 31;
        }
        else if(prevMonth === 1 && (date.getFullYear() % 4 === 0))
        {
            lastDayPrev = 29;
        }
        else if(prevMonth === 1 && (date.getFullYear() % 4 !== 0))
        {
            lastDayPrev = 28
        }

        //Add days to array before the current date
        for(i = date.getDay() - 1; i > -1; i--)
        {
            arr[i] = {
                day: arr[i+1].day - 1,
                date: arr[i+1].date - 1,
                month: arr[i+1].month,
                year: arr[i+1].year
            }

            if(arr[i+1].date === 1)
            {
                arr[i].date = lastDayPrev;
                arr[i].month = arr[i].month - 1;
                
                if(arr[i].month === -1)
                {
                    arr[i].month = 11;
                    arr[i].year = arr[i].year - 1;
                }
            }
        }

        //Set month and year to display. Uses lower month for dual cases
        setMonthIndex(arr[0].month);
        setYear(arr[0].year);

        //Set the days array
        setDays(arr);

        //Highlight current day
        const refArray = [sunRef, monRef, tueRef, wedRef, thuRef, friRef, satRef];
        for(i = 0; i < 7; i++)
        {
            if(arr[i].date === date.getDate())
            {
                refArray[i].current.style.backgroundColor = '#09a2fb';
                refArray[i].current.style.color = '#ffffff';
                indicatorRef.current.style.display = 'block';
            }
            else
            {
                refArray[i].current.style.backgroundColor = 'transparent';
                refArray[i].current.style.color = '#374146';
            }
        }

        //Check for dual month
        for(i = 1; i < 7; i++)
        {
            if(arr[i].month !== arr[0].month)
            {
                setDualMonth(true);
                break;
            }
            else
            {
                setDualMonth(false);
            }
        }

        if(indicatorRef.current)
        {
            indicatorRef.current.style.top = (date.getHours())*7.9 + (date.getMinutes()/60)*7.9 + "vh";
            indicatorRef.current.style.left = (date.getDay())*11.3 + "vw";
        }
    }

    //Using useEffect hook. This will send the POST request once the component is mounted
    useEffect(() =>
    {
        //POST request
        const fetchData = () =>
        {

            fetch("http://localhost:8080/api/booking/information", 
                {
                method: "POST",
                mode: "cors",
                body: JSON.stringify({
                    user_id: userData.user_id
                }),
                headers:{
                    'Content-Type': 'application/json',
                    'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                }
                }).then((res) => res.json()).then(data => 
                {
                    setBookings(data);
                    console.log(data);
                    window.sessionStorage.removeItem("BookingID");
                    window.sessionStorage.removeItem("StartDate");
                    window.sessionStorage.removeItem("StartTime");
                    window.sessionStorage.removeItem("EndDate");
                    window.sessionStorage.removeItem("EndTime");
                });
        }

        fetchData();
        loadToday();
    }, [userData]);

    const refreshTime = () =>
    {
        const date = new Date();
        if(indicatorRef.current)
        {
            indicatorRef.current.style.top = (date.getHours())*7.9 + (date.getMinutes()/60)*7.9 + "vh";
            indicatorRef.current.style.left = (date.getDay())*11.3 + "vw";
        }
    }

    setInterval(refreshTime, 60000*5);

    useEffect(() =>
    {
        if(!dualMonth)
        {
            const monthNames = ["January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"];
            setMonth(monthNames[monthIndex]);
        }
        else
        {
            const monthNames = ["Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sept", "Oct", "Nov", "Dec"];
            setMonth(monthNames[monthIndex] + "-" + monthNames[(monthIndex + 1) % 12]);
        }
        
    },[monthIndex, dualMonth]);

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
    }, [currentContext]);

    useEffect(() =>
    {
        if(isToday)
        {
            todaySelectorRef.current.style.backgroundColor = '#09a2fb';
            todaySelectorRef.current.style.color = '#ffffff';
            loadToday();
        }
        else
        {
            todaySelectorRef.current.style.backgroundColor = '#ffffff';
            todaySelectorRef.current.style.color = '#09a2fb';
        }
    },[isToday]);
  

    return (
        <div className='page-container'>
            <div className='content'>
                <ProfileBar />
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

                        <div className='today-container'>
                            <div ref={todaySelectorRef} className='today-selector' onClick={SelectToday} onMouseOver={MouseOverToday} onMouseLeave={MouseLeaveToday}>
                                Today
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
                                <div ref={sunRef} className='date'>{days[0].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Mon</p>
                                <div ref={monRef} className='date'>{days[1].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Tue</p>
                                <div ref={tueRef} className='date'>{days[2].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Wed</p>
                                <div ref={wedRef} className='date'>{days[3].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Thu</p>
                                <div ref={thuRef} className='date'>{days[4].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Fri</p>
                                <div ref={friRef} className='date'>{days[5].date}</div>
                            </div>

                            <div className='day-date'>
                                <p className='day'>Sat</p>
                                <div ref={satRef} className='date'>{days[6].date}</div>
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
                            <Times />

                            <Borders />
                            <Borders />
                            <Borders />
                            <Borders />
                            <Borders />
                            <Borders />
                            <Borders />
                            
                            <div ref={indicatorRef} className='time-indicator-container'>
                                <div className='time-indicator-circle'></div>
                                <div className='time-indicator-line'></div>
                            </div>

                            <div className='bookings-container'>
                                {bookings.length > 0 && (
                                    bookings.map((booking, i) => (
                                        <BookingTicket id={booking.id} startDate={booking.start.substring(0,10)} startTime={booking.start.substring(11,16)} endTime={booking.end.substring(11,16)} confirmed={booking.booked} type={booking.resource_type} days={days} />
                                    ))
                                )}
                            </div>
                        </div>
                    </div>

                </div>
            </div>
        </div>
    )
}

export default Home
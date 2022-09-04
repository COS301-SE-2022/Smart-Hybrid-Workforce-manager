import React, { useEffect, useRef, useState } from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'

const BookingTicket = ({id, startDate, startTime, endTime, confirmed, type, days}) => 
{
    const [year, setYear] = useState(""); 
    const [month, setMonth] = useState("");
    const [day, setDay] = useState(""); 
    const [startHours, setStartHours] = useState(""); 
    const [startMins, setStartMins] = useState(""); 
    const [endHours, setEndHours] = useState(""); 
    const [endMins, setEndMins] = useState(""); 

    const ticketRef = useRef(null);

    let EditBooking = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("BookingID", id);
        window.sessionStorage.setItem("StartTime", startTime);
        window.sessionStorage.setItem("EndTime", endTime);
        window.location.assign("./bookings-desk-edit");
    }

    let DeleteBooking = async (e) =>
    {
        e.preventDefault();
        if(window.confirm("Are you sure you want to delete this booking?"))
        {
            try
            {
                let res = await fetch("http://localhost:8100/api/booking/remove", 
                {
                    method: "POST",
                    body: JSON.stringify({
                    id: id
                    })
                });

                if(res.status === 200)
                {
                    alert("Booking Successfully Deleted!");
                    window.location.assign("./");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    useEffect(() =>
    {
        setYear(startDate.substring(0,4));
        setMonth(startDate.substring(5,7));
        setDay(startDate.substring(8,10));
        setStartHours(startTime.substring(0,2));
        setStartMins(startTime.substring(3,5));
        setEndHours(endTime.substring(0,2));
        setEndMins(endTime.substring(3,5));      
    },[startDate, startTime, endTime]);

    useEffect(() =>
    {
        ticketRef.current.style.top = startHours*8 + (startMins/60)*8 + "vh";
        ticketRef.current.style.height = (endHours-startHours)*8 + (startMins/60)*8 + "vh";
    },[startHours, startMins, endHours, endMins]);

    useEffect(() =>
    {
        ticketRef.current.style.display  = 'none';
        for(var i = 0; i < days.length; i++)
        {
            if((parseInt(days[i].date) === parseInt(day)) && (parseInt(days[i].month) === parseInt(month-1)) && (parseInt(days[i].year) === parseInt(year)))
            {
                ticketRef.current.style.display  = 'block';
                ticketRef.current.style.left = i*11.3 + "vw";
                break;
            }
        }
    },[days, day, month, year])


    return (
        <div>
            <div ref={ticketRef} className="booking-ticket">
                <div className="booking-tex">
                    <p>Start Time: {startHours}{startMins}</p>
                    <p>End Time: {endHours}{endMins}</p>
                    <p>Type: {type}</p>
                    <p>Date: {year}{month}{day}</p>
                    {{confirmed} ? (
                        <p>Acceptance: Pending</p>
                    ) : (
                        <p>Acceptance: Approved</p>
                    )}                
                </div>

            </div>
        </div>
    )
}

export default BookingTicket
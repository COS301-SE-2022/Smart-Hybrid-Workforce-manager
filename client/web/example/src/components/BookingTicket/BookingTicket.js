import React, { useContext, useEffect, useRef, useState } from 'react';
import { MdEdit, MdDelete } from 'react-icons/md';
import { GiDesk, GiRoundTable } from 'react-icons/gi';
import { UserContext } from '../../App';

const BookingTicket = ({id, startDate, startTime, endTime, confirmed, type, days}) => 
{
    const [year, setYear] = useState(""); 
    const [month, setMonth] = useState("");
    const [day, setDay] = useState(""); 
    const [startHours, setStartHours] = useState(""); 
    const [startMins, setStartMins] = useState(""); 
    const [endHours, setEndHours] = useState(""); 
    const [endMins, setEndMins] = useState(""); 

    const {userData} = useContext(UserContext);

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
                let res = await fetch("http://localhost:8080/api/booking/remove", 
                {
                    method: "POST",
                    mode: "cors",
                    body: JSON.stringify({
                    id: id
                    }),
                    headers:{
                        'Content-Type': 'application/json',
                        'Authorization': `bearer ${userData.token}` //Changed for frontend editing .token
                    }
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
    },[startDate, startTime, endTime, type]);

    useEffect(() =>
    {
        ticketRef.current.style.top = startHours*8 + (startMins/60)*8 + "vh";
        ticketRef.current.style.height = (endHours-startHours)*8 + (startMins/60)*8 + "vh";
        ticketRef.current.style.paddingTop = ((endHours-startHours)*8 + (startMins/60)*8)/2 - 2.5 + "vh";
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

    const renderIcon = () =>
    {
        if(type === 'Desk')
        {
            return <GiDesk />;
        }
        else
        {
            return <GiRoundTable />;
        }
    }

    return (
        <div>
            <div ref={ticketRef} className="booking-ticket">
                <div className="booking-text">
                    {renderIcon()}       
                </div>
            </div>
        </div>
    )
}

export default BookingTicket
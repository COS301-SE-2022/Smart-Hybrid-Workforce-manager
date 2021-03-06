import React from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'
import { useNavigate } from 'react-router-dom'

const BookingTicket = ({id, startDate, startTime, endDate, endTime, confirmed, type}) => {
    const navigate = useNavigate();

    let EditBooking = async (e) =>
    {
        e.preventDefault();
        window.sessionStorage.setItem("BookingID", id);
        window.sessionStorage.setItem("StartDate", startDate);
        window.sessionStorage.setItem("StartTime", startTime);
        window.sessionStorage.setItem("EndDate", endDate);
        window.sessionStorage.setItem("EndTime", endTime);
        navigate("/bookings-desk-edit");
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
                    navigate("/");
                }
            }
            catch (err)
            {
                console.log(err);    
            }
        }
    }

    return (
        <div>
            <div className="booking">
                <div className="booking-image"></div>
                <div className="booking-text">
                <p>Start Time: {startTime}</p>
                <p>End Time: {endTime}</p>
                <p>Type: {type}</p>
                {{confirmed} ? (
                    <p>Acceptance: Pending</p>
                ) : (
                    <p>Acceptance: Approved</p>
                )}                
                </div>
                <div className='booking-popup'>
                    <div className='booking-edit'><MdEdit size={80} className="booking-edit-icon" onClick={EditBooking}/></div>
                    <div className='booking-delete'><MdDelete size={80} className="booking-delete-icon" onClick={DeleteBooking}/></div>
                </div>
            </div>
        </div>
    )
}

export default BookingTicket
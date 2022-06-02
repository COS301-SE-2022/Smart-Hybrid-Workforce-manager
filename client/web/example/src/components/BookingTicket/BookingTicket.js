import React from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'

const BookingTicket = ({id, startDate, startTime, endDate, endTime, confirmed}) => {

    let EditBooking = async (e) =>
    {
        e.preventDefault();
        window.location.assign("./bookings-desk-edit");
        window.sessionStorage.setItem("BookingID", id);
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

    return (
        <div>
            <div className="booking">
                <div className="booking-image"></div>
                <div className="booking-text">
                <p>Start Date: {startDate}</p>
                <p>Start Time: {startTime}</p>
                <p>End Date: {endDate}</p>
                <p>End Time: {endTime}</p>
                {{confirmed} ? (
                    <p>Confirmed: Pending</p>
                ) : (
                    <p>Confirmed: Approved</p>
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
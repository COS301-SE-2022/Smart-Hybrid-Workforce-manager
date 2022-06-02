import React from 'react'
import { MdEdit, MdDelete } from 'react-icons/md'

const BookingTicket = ({id, startDate, startTime, endDate, endTime, confirmed}) => {

    let EditBooking = async (e) =>
    {
        e.preventDefault();
        window.location.assign("./bookings-desk-edit");
        window.sessionStorage.setItem("BookingID", id);
        /*try
        {
            let res = await fetch("http://localhost:8100/api/booking/create", 
            {
                method: "POST",
                body: JSON.stringify({
                id: null,
                user_id: "11111111-1111-4a06-9983-8b374586e459",
                resource_type: "DESK",
                resource_preference_id: null,
                resource_id: null,
                start: startDate + "T" + startTime + ":43.511Z",
                end: endDate + "T" + endTime + ":43.511Z",
                booked: false
                })
            });
        }
        catch (err)
        {
            console.log(err);    
        }*/
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
                    <div className='booking-edit' onClick={EditBooking}><MdEdit size={80}/></div>
                    <div className='booking-delete'><MdDelete size={80}/></div>
                </div>
            </div>
        </div>
    )
}

export default BookingTicket
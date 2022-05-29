import React from 'react'

const BookingTicket = ({startDate, startTime, endDate, endTime, confirmed}) => {
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
            </div>
        </div>
    )
}

export default BookingTicket
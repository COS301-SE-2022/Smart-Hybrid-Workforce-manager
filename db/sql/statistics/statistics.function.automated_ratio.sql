CREATE OR REPLACE FUNCTION statistics.automated_ratio 
(
	_start_date DATE DEFAULT current_date - interval '7 years',
	_end_date 	DATE DEFAULT current_date
)
RETURNS TABLE (
	automated_bookings  bigint,
	manual_bookings	    bigint
) AS
$$
	BEGIN
		RETURN QUERY
			SELECT (SELECT COUNT(*) 
				    FROM booking.identifier b 
				    WHERE (b.start::date BETWEEN _start_date AND _end_date) AND b.booked = TRUE AND b.automated = TRUE
				   ) AS automated_bookings,
				   (SELECT COUNT(*) 
				    FROM booking.identifier b 
				    WHERE (b.start::date BETWEEN _start_date AND _end_date) AND b.booked = TRUE AND b.automated = FALSE
				   ) AS manual_bookings;
	
		RETURN;
	END
$$ LANGUAGE plpgsql;
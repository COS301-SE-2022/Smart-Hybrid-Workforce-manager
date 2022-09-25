CREATE OR REPLACE FUNCTION statistics.current_occupency ()
RETURNS TABLE (
	current_occupied    bigint,
	total               bigint
) AS
$$
	BEGIN
		RETURN QUERY
			SELECT (SELECT COUNT(DISTINCT(b.resource_id)) 
				    FROM booking.identifier b 
				    WHERE DATE_TRUNC('day', b.start::date) = DATE_TRUNC('day', current_date)  AND b.booked = TRUE
				   ) AS current_occupied,
				   (SELECT COUNT(DISTINCT(b.resource_id)) 
				    FROM resource.identifier b
				   ) AS total;
		RETURN;
	END
$$ LANGUAGE plpgsql;
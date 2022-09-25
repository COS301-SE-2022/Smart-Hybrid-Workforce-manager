CREATE OR REPLACE FUNCTION statistics.utilisation 
(
	_start_date DATE DEFAULT current_date - interval '7 days',
	_end_date 	DATE DEFAULT current_date
)
RETURNS TABLE (
	resource_id uuid,
	utilisation_percent	float
) AS
$$
	DECLARE
		num_days	INTEGER;
	BEGIN
		RETURN QUERY
			SELECT r.id AS resource_id, COALESCE(((SELECT COUNT(DISTINCT b.start::date) 
										 		   FROM booking.identifier b 
										 		   WHERE b.resource_id = r.id  AND (b.start::date BETWEEN _start_date AND _end_date) AND b.booked = TRUE 
										 		   GROUP BY b.start
												   )/(_end_date::date - _start_date::date + 1)::float)*100, 0)::float AS utilisation_percent
			FROM resource.identifier r;
	
		RETURN;
	END
$$ LANGUAGE plpgsql;

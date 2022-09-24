CREATE OR REPLACE FUNCTION statistics.average_utilisation 
(
	_start_date DATE DEFAULT current_date - interval '7 days',
	_end_date 	DATE DEFAULT current_date
)
RETURNS float AS
$$
	DECLARE
		num_days	INTEGER;
	BEGIN
		RETURN
			(SELECT AVG(COALESCE(((SELECT COUNT(DISTINCT b.start::date) 
                                   FROM booking.identifier b 
								   WHERE b.resource_preference_id = r.id  AND (b.start::date BETWEEN _start_date AND _end_date)
								   GROUP BY b.start
								   )/(_end_date::date - _start_date::date + 1)::float)*100, 0)::float) 
			FROM resource.identifier r);
	END
$$ LANGUAGE plpgsql;

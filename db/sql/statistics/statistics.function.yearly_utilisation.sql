CREATE OR REPLACE FUNCTION statistics.yearly_utilisation ()
RETURNS TABLE (
	month_ 				DATE,
	utilisation_percent	float
) AS
$$
	DECLARE
		_start_date DATE;
		_end_date 	DATE;
	BEGIN
		_start_date := DATE_TRUNC('month', (current_date - interval '11 months'));
		_end_date 	:= DATE_TRUNC('month', (current_date + interval '1 month'));
		RETURN QUERY
			SELECT DATE_TRUNC('month', b.start::date)::date, statistics.average_utilisation(DATE_TRUNC('month', b.start::date)::date, (DATE_TRUNC('month', b.start::date) + interval '1 month')::date)
			FROM booking.identifier b
			GROUP BY DATE_TRUNC('month', b.start::date)
			HAVING DATE_TRUNC('month', b.start::date)::date BETWEEN (current_date-interval '1 year') AND current_date
			ORDER BY DATE_TRUNC('month', b.start::date) ASC;
		RETURN;
	END
$$ LANGUAGE plpgsql;
CREATE OR REPLACE FUNCTION example.test()
RETURNS TABLE(id uuid) AS $$

SELECT id
FROM example.example

$$ LANGUAGE sql;
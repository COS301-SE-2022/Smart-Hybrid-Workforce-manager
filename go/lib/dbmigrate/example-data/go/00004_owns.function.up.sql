CREATE OR REPLACE FUNCTION pets.get_ownership()
	RETURNS TABLE(pname VARCHAR, dname VARCHAR)
AS
$$
BEGIN
  RETURN QUERY
  SELECT p.name, d.name
	FROM pets.owns AS o
	INNER JOIN pets.person AS p ON o.pid = p.id
	INNER JOIN pets.dog AS d ON o.did = d.id;
END
$$
LANGUAGE plpgsql;
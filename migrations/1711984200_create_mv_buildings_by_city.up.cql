-- creating a materialized view (table) based on buildings table
-- this allows to have a different data output and allow searching building by city
-- MV uses a city and name as PK when original table has only name as PK
-- As a rule to a MV all the components of the original PK must appear in a MV PK
-- Here we selecting all the table rows, bit it's possible to select only a specific
-- It's not allowed to delete data from MV, data should be deleted from the original table

CREATE MATERIALIZED VIEW buildings_by_city AS
 SELECT * FROM buildings
 WHERE city IS NOT NULL AND name IS NOT NULL
 PRIMARY KEY((city), name);

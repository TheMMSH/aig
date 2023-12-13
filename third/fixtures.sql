-- making a fixture to test the solution
CREATE TABLE IF NOT EXISTS  seat (
     id SERIAL PRIMARY KEY,
     student VARCHAR
);

INSERT INTO seat VALUES (1, 'Abbot'); -- explicitly have ids to prevent duplicate insertion
INSERT INTO seat VALUES (2, 'Doris');
INSERT INTO seat VALUES (3, 'Emerson');
INSERT INTO seat VALUES (4, 'Green');
INSERT INTO seat VALUES (5, 'James');
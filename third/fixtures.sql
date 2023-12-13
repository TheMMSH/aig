-- making a fixture to test the solution
CREATE TABLE IF NOT EXISTS  seat (
     id SERIAL PRIMARY KEY,
     student VARCHAR
);

INSERT INTO seat (student) VALUES ('Abbot');
INSERT INTO seat (student) VALUES ('Doris');
INSERT INTO seat (student) VALUES ('Emerson');
INSERT INTO seat (student) VALUES ('Green');
INSERT INTO seat (student) VALUES ('James');
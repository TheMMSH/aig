SELECT (
    CASE
       WHEN id % 2 != 0 AND sc != id THEN id + 1
    WHEN id % 2 != 0 AND sc = id THEN id
    ELSE id - 1
    END
) AS id, student
FROM seat, (SELECT COUNT(*) AS sc FROM seat) ORDER BY id;
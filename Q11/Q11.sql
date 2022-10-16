--- PostgreSQL
--- DDL (Data Definition Language) --- CREATE ---
CREATE TABLE items (
order_id INTEGER,
position TEXT,
price REAL
);
--- DML (Data Manipulation Language) --- INSERT ---
INSERT INTO items VALUES (1, 'Товар 1', 3.0);
INSERT INTO items VALUES (1, 'Товар 1', 4.0);
INSERT INTO items VALUES (1, 'Товар 1', 7.0);
INSERT INTO items VALUES (1, 'Товар 1', 1.0);
INSERT INTO items VALUES (1, 'Товар 2', 9.0);
--- DML (Data Manipulation Language) --- SELECT ---
SELECT
sum(price)
FROM items
GROUP BY position
HAVING count(position) > 3

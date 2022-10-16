--- SQLite
--- DDL (Data Definition Language) --- CREATE ---
CREATE TABLE customers (
customer_id INTEGER PRIMARY KEY,
customer_name TEXT
);
CREATE TABLE orders (
order_num INTEGER PRIMARY KEY,
customer_id INTEGER,
FOREIGN KEY(customer_id) REFERENCES customers(customer_id)
); 
--- DML (Data Manipulation Language) --- INSERT ---
INSERT INTO customers VALUES (0, 'Александр Петров');
INSERT INTO customers VALUES (1, 'Иван Иванов');
INSERT INTO customers VALUES (2, 'Алексей Сидоров');  
INSERT INTO orders VALUES (100, 0);
INSERT INTO orders VALUES (101, 1);
INSERT INTO orders VALUES (102, 0);
INSERT INTO orders VALUES (104, 2);
--- DML (Data Manipulation Language) --- SELECT ---
SELECT
*
FROM orders
LEFT JOIN customers
ON orders.customer_id == customers.customer_id
WHERE customers.customer_name LIKE 'Алекс%';

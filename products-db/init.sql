CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price NUMERIC(10, 2)
);

INSERT INTO products (name, price) VALUES
('Laptop', 999.99),
('Smartphone', 799.99),
('Headphones', 199.99);

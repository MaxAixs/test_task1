CREATE TABLE sellers (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
    phone VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
    description TEXT,
     price DECIMAL(10,2) NOT NULL,
     seller_id INT REFERENCES sellers(id) ON DELETE CASCADE
);

CREATE TABLE customers (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255) NOT NULL,
     phone VARCHAR(20) UNIQUE NOT NULL
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
     customer_id INT REFERENCES customers(id) ON DELETE CASCADE
);

CREATE TABLE order_products (
     order_id INT REFERENCES orders(id) ON DELETE CASCADE,
    product_id INT REFERENCES products(id) ON DELETE CASCADE,
    PRIMARY KEY (order_id, product_id)
);
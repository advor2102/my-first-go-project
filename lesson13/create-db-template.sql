-- Active: 1757069669177@@127.0.0.1@5432@postgres
CREATE TABLE Users 
(    
    id SERIAL,
    name VARCHAR NOT NULL,
    email VARCHAR UNIQUE NOT NULL,
    age INT,
    is_active BOOLEAN,
    created_at DATE,
    PRIMARY KEY (id)
);

CREATE TABLE Orders
(
    id SERIAL,
    user_id INT,
    product VARCHAR,
    amount INT,
    created_at DATE,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES Users (id)
)

DROP Table Users

DROP Table Orders
CREATE TABLE IF NOT EXISTS users (
     id SERIAL PRIMARY KEY,
     name VARCHAR(255),
     phone_number VARCHAR(20) UNIQUE NOT NULL,
     otp VARCHAR(4),
     otp_expiration_time TIMESTAMP
);
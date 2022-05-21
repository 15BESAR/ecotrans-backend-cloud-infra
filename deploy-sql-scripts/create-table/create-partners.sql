USE capstone;

CREATE TABLE partners(
    partner_id CHAR(36) NOT NULL PRIMARY KEY,
    partner_name VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(100)
);
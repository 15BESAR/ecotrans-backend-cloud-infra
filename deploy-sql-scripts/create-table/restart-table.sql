USE capstone;
DROP TABLE purchases,vouchers,partners,journeys,users;
CREATE TABLE users(
    user_id CHAR(36) NOT NULL PRIMARY KEY,
    email VARCHAR(50) UNIQUE,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(100),
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    birth_date DATE,
    age TINYINT,
    gender CHAR(1),
    job VARCHAR(25),
    points INT,
    voucher_interest VARCHAR(500),
    domicile VARCHAR(50),
    education VARCHAR(50),
    marriage_status BOOLEAN,
    income INT,
    vehicle VARCHAR(25)
);

CREATE TABLE journeys(
    journey_id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    start_time DATETIME NOT NULL,
    finish_time DATETIME NOT NULL,
    origin VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL,
    distance_travelled FLOAT NOT NULL,
    emission_saved FLOAT NOT NULL,
    reward INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

CREATE TABLE partners(
    partner_id CHAR(36) NOT NULL PRIMARY KEY,
    partner_name VARCHAR(50),
    email VARCHAR(50) UNIQUE,
    username VARCHAR(50) UNIQUE,
    password VARCHAR(100)
);

CREATE TABLE vouchers(
    voucher_id CHAR(36) NOT NULL PRIMARY KEY,
    partner_id CHAR(36) NOT NULL,
    voucher_name VARCHAR(50) NOT NULL,
    voucher_desc VARCHAR(500) NOT NULL,
    category VARCHAR(25) NOT NULL,
    image_url VARCHAR(100) NOT NULL,
    stock INT NOT NULL,
    price INT NOT NULL,
    FOREIGN KEY (partner_id) REFERENCES partners(partner_id)
);


CREATE TABLE purchases(
    voucher_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    buy_date DATETIME NOT NULL,
    buy_quantity INT NOT NULL,
    PRIMARY KEY (user_id, voucher_id)
);
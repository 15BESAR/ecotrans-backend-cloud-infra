USE capstone;
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
    marriage_status VARCHAR(20),
    income INT,
    vehicle VARCHAR(25)
);
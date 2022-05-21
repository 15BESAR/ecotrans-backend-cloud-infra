USE capstone-instance;
CREATE TABLE users(
    user_id CHAR(36) NOT NULL PRIMARY KEY,
    first_name VARCHAR(50),
    last_name VARCHAR(50),
    gender CHAR(1),
    voucher_interest VARCHAR(500),
    domicile VARCHAR(50),
    birth_date DATE,
    age TINYINT,
    education VARCHAR(50),
    points INT,
    marriage_status VARCHAR(20),
    job VARCHAR(25),
    income INT,
    vehicle VARCHAR(25),
    username VARCHAR(50) UNIQUE,
    email VARCHAR(50) UNIQUE,
    password VARCHAR(100) UNIQUE
);
USE capstone-instance;
CREATE TABLE journeys(
    journey_id CHAR(36) NOT NULL PRIMARY KEY,
    user_id FOREIGN KEY REFERENCES users(user_id)
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    origin VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL,
    distance_travelled VARCHAR(100) NOT NULL,
    emission_saved FLOAT NOT NULL,
    reward INT NOT NULL
);
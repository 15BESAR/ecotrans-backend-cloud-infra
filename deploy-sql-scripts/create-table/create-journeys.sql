USE capstone;

CREATE TABLE journeys(
    journey_id CHAR(36) NOT NULL PRIMARY KEY,
    user_id CHAR(36) NOT NULL,
    start_date DATETIME NOT NULL,
    end_date DATETIME NOT NULL,
    origin VARCHAR(100) NOT NULL,
    destination VARCHAR(100) NOT NULL,
    distance_travelled FLOAT NOT NULL,
    emission_saved FLOAT NOT NULL,
    reward INT NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

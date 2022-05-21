USE capstone-instance;
CREATE TABLE purchases(
    user_id CHAR(36) NOT NULL,
    voucher_id CHAR(36) NOT NULL,
    buy_date DATETIME NOT NULL,
    buy_quantity INT NOT NULL,
    PRIMARY KEY (user_id, voucher_id)
);
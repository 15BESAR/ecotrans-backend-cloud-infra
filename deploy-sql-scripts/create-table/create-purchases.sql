USE capstone;

CREATE TABLE purchases(
    voucher_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    buy_date DATETIME NOT NULL,
    buy_quantity INT NOT NULL,
    PRIMARY KEY (user_id, voucher_id)
);
USE capstone;

CREATE TABLE purchases(
    purchase_id CHAR(36) NOT NULL PRIMARY KEY, 
    voucher_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    buy_date DATETIME NOT NULL,
    buy_quantity INT NOT NULL
);
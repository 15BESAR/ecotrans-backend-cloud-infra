USE capstone-instance;
CREATE TABLE vouchers(
    voucher_id CHAR(36) NOT NULL PRIMARY KEY,
    partner_id FOREIGN KEY REFERENCES partners(partner_id) NOT NULL,
    voucher_name VARCHAR(50) NOT NULL,
    voucher_desc VARCHAR(500) NOT NULL,
    category VARCHAR(25) NOT NULL,
    image_url VARCHAR(100) NOT NULL,
    stock INT NOT NULL,
    price INT NOT NULL
);
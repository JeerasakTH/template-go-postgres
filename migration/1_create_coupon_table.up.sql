CREATE TABLE coupon_detail (
    id serial PRIMARY KEY,
    name VARCHAR,
    start_date TIMESTAMP,
    end_date TIMESTAMP,
    status INT,
    coupon_count INT,
    coupon_type INT,
    reward NUMERIC
);
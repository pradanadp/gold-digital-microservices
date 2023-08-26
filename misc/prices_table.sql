CREATE TABLE prices (
    price_id       VARCHAR(255) PRIMARY KEY,
    admin_id       VARCHAR(255) NOT NULL,
    topup_price    FLOAT8 NOT NULL,
    buyback_price  FLOAT8 NOT NULL,
    created_at     TIMESTAMP DEFAULT NOW() NOT NULL
);
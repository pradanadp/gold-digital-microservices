CREATE TABLE transactions (
    transaction_id    VARCHAR(255) PRIMARY KEY,
    account_number    VARCHAR(255) NOT NULL,
    quantity          NUMERIC(12, 6) NOT NULL,
    topup_price       NUMERIC(12, 2) NOT NULL,
    buyback_price     NUMERIC(12, 2) NOT NULL,
    type              VARCHAR(10) NOT NULL CHECK (TYPE IN ('topup', 'buyback')),
    balance           NUMERIC(12, 2) NOT NULL,
    created_at        TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
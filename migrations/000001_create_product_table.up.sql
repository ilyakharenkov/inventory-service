CREATE TABLE IF NOT EXISTS product_t
(
    id         SERIAL PRIMARY KEY,
    sku        VARCHAR(100) UNIQUE NOT NULL,
    name       VARCHAR(255)        NOT NULL,
    quantity   INT                 NOT NULL DEFAULT 0,
    reserved   INT                 NOT NULL DEFAULT 0,
    price      DECIMAL(10, 2)      NOT NULL,
    created_at TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP                    DEFAULT CURRENT_TIMESTAMP
);
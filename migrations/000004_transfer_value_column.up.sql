UPDATE product_t
SET free = quantity - reserved
WHERE quantity >= reserved;
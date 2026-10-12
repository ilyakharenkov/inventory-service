UPDATE product_t
SET free = 0
WHERE quantity >= reserved;
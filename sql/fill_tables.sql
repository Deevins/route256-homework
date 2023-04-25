INSERT INTO users (username, email)
SELECT 'user ' || i,
       'email' || i || '@gmail.com'
FROM generate_series(1, 50) as s(i);


INSERT INTO orders (user_id, product_name, quantity)
SELECT i,
       'product-' || i,
       (random() * 10)::numeric(10, 2)
FROM generate_series(1, 40) as s(i);
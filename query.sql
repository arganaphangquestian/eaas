-- GetCustomers :many
SELECT * FROM customers;

-- AddCustomer :one
INSERT INTO customers(name, email, password, address, balance)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
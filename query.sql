-- name: GetCustomers :many
SELECT * FROM customers;

-- name: AddCustomer :one
INSERT INTO customers(name, email, password, address, balance)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;
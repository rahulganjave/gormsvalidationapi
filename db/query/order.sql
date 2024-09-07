-- name: CreateOrder :one
INSERT INTO orders
    (
    order_id,
    client_id,
    trading_acct,
    exchange,
    product,
    instrument,
    side,
    status,
    order_type,
    order_qty,
    order_price,
    time_inforce,
    trading_session_id,
    expire_date,
    exec_qty,
    exec_price,
    cash,
    im,
    exposure,
    est_cost_to_close,
    order_time
    )
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21
)
RETURNING *;

-- name: GetOrder :one
SELECT *
FROM orders
WHERE order_id = $1
LIMIT 1;

-- name: GetOrdersForUpdate :one
/*SELECT *
FROM orders
WHERE order_id = $1
LIMIT 1
FOR NO KEY
UPDATE;*/

-- name: ListOrders :many
/*SELECT *
FROM orders
WHERE order_id = $1
ORDER BY id
LIMIT $2
OFFSET
$3;*/

-- name: UpdateOrder :one
UPDATE orders
SET order_qty = $2,
exec_qty = $3
WHERE order_id = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE order_id = $1;
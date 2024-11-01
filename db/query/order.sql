-- name: CreateOrder :one
INSERT INTO orders (
    "order_id",
    "client_id",
    "trading_acct",
    "exchange_code",
    "product_code",
    "instrument_code",
    "order_side",
    "order_status",
    "order_type",
    "order_qty",
    "order_price",
    "time_inforce",
    "trading_session_id",
    "expire_date",
    "exec_qty",
    "exec_price",
    "cash",
    "Initial_margin",
    "exposure",
    "est_cost_to_close",
    "order_time",
    "underlying_code",
    "open_close",
    "put_or_call",
    "contract_size",
    "created_at",
    "updated_at"
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27
)
RETURNING *;

-- name: GetOrder :one
SELECT *
FROM orders
WHERE "order_id" = $1
LIMIT 1;

-- name: GetOrderForUpdate :one
SELECT *
FROM orders
WHERE "order_id" = $1
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListOrders :many
SELECT *
FROM orders
WHERE "client_id" = $1
ORDER BY "order_time"
LIMIT $2
OFFSET $3;

-- name: UpdateOrder :one
UPDATE orders
SET "order_qty" = $2,
    "exec_qty" = $3
WHERE "order_id" = $1
RETURNING *;

-- name: DeleteOrder :exec
DELETE FROM orders
WHERE "order_id" = $1;

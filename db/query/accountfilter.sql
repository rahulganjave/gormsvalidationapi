-- name: CreateAccount :one
INSERT INTO accountfilter
    (
    client_id,
    account,
    trader,
    cash_balance_bf,
    cash_balance,
    est_trans_cost,
    equity_realized_pl,
    equity_unrealized_pl,
    margin_realized_pl,
    margin_unrealized_pl,
    est_closing_cost,
    transaction_not_booked,
    equity_order,
    account_value,
    deposit_withdraw,
    collateral_haircut,
    equity_position_value,
    margin_requirement,
    currency,
    buy_limit,
    sell_limit,
    buy_exposure,
    sell_exposure,
    loan_limit,
    margin_factor,
    authorized_order_type,
    calc_mode,
    limit_mode,
    action,
    netting_mode,
    max_ord_qty,
    max_ord_val,
    no_of_transaction,
    transaction_interval
    )
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34
)
RETURNING *;

-- name: GetAccount :one
SELECT *
FROM accountfilter
WHERE client_id = $1
    AND account = $2
LIMIT 1;

-- name: GetAccountForUpdate :one
SELECT *
FROM accountfilter
WHERE client_id = $1
AND account = $2
LIMIT 1
FOR NO KEY
UPDATE;

-- name: ListAccount :many
SELECT *
FROM accountfilter
WHERE client_id = $1
AND account = $2
LIMIT $2
OFFSET
$3;

-- name: UpdateAccount :one
UPDATE accountfilter
SET buy_exposure = $2,
sell_exposure = $3
WHERE account = $1
RETURNING *;

-- name: DeleteAccount :exec
DELETE FROM accountfilter
WHERE account = $1;
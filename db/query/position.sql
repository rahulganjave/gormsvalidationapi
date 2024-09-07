-- name: CreatePosition :one
INSERT INTO Positions
    (
    position_id,
    ref_position_id,
    trading_acct,
    client_id,
    product,
    exchange,
    instrument,
    side,
    status,
    position_type,
    settlement_status,
    calc_type,
    open_qty,
    open_px,
    closed_qty,
    closed_px,
    trading_curr,
    settle_curr,
    mark_price,
    mark_value,
    exchange_rate,
    unrealized_pl,
    realized_pl,
    mm_and_cash,
    collateral_haircut_ratio,
    collateral_haircut,
    open_date,
    closed_date,
    cost_to_close,
    trade_date,
    trading_session_id,
    underlying,
    open_close,
    put_or_call,
    contract_size
    )
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34, $35
)
RETURNING *;

-- name: GetPosition :one
SELECT *
FROM Positions
WHERE position_id = $1
LIMIT 1;

-- name: GetPositionsForUpdate :one
SELECT *
FROM Positions
WHERE position_id = $1
LIMIT 1
FOR NO KEY
UPDATE;

-- name: ListPositions :many
SELECT *
FROM Positions
WHERE position_id = $1
ORDER BY position_id
LIMIT $2
OFFSET
$3;

-- name: UpdatePosition :one
UPDATE Positions
SET open_qty = $2,
open_px = $3
WHERE position_id = $1
RETURNING *;

-- name: DeletePosition :exec
DELETE FROM Positions
WHERE position_id = $1;
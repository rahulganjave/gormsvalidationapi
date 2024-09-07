-- name: CreateInstrumentfilter :one
INSERT INTO instrumentfilter
    (
    product,
    exchange,
    instrument,
    underlying,
    filter,
    max_order_qty,
    max_order_price,
    max_order_value,
    buy_lmt,
    sell_lmt,
    exposure_factor,
    buy_exposure,
    sell_exposure,
    buy_qty_lmt,
    buy_qty_exposure,
    sell_qty_lmt,
    sell_qty_exposure,
    net_qty_lmt,
    net_qty_exposure,
    margin_type,
    im,
    mm,
    marginable_ratio,
    margin_curr,
    commission_type,
    commission,
    min_commission,
    commission_curr,
    authorized_order_type,
    trading_curr,
    price_code,
    ref_px,
    ref_px_type,
    contract_size
    )
VALUES
    (
        $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34
)
RETURNING *;

-- name: GetInstrumentfilter :one
SELECT *
FROM instrumentfilter
WHERE instrument = $1
AND exchange = $2
AND product = $3
LIMIT 1;

-- name: GetInstrumentfilterForUpdate :one
SELECT *
FROM instrumentfilter
WHERE instrument = $1
AND exchange = $2
AND product = $3
LIMIT 1
FOR NO KEY
UPDATE;

-- name: ListInstrumentfilter :many
SELECT *
FROM instrumentfilter
WHERE instrument = $1
AND exchange = $2
AND product = $3
LIMIT $4
OFFSET
$5;

-- name: UpdateInstrumentfilter :one
UPDATE instrumentfilter
SET buy_exposure = $4,
sell_exposure = $5
WHERE instrument = $1
AND exchange = $2
AND product = $3
RETURNING *;

-- name: DeleteInstrumentfilter :exec
DELETE FROM instrumentfilter
WHERE instrument = $1
AND exchange = $2
AND product = $3;
-- name: CreateInstrument :one
INSERT INTO instruments (
    "product_code",
    "exchange_code",
    "instrument_code",
    "underlying_code",
    "instrument_desc",
    "maturity_monthyear",
    "maturity_day",
    "put_or_call",
    "strike_price",
    "symbol",
    "symbol_sfx",
    "cusip",
    "isin",
    "ric",
    "sedol",
    "blpsyn",
    "cficode",
    "price_code",
    "effective_date",
    "minqty",
    "sector",
    "contract_size",
    "corporate_action",
    "trading_session_id",
    "trading_curr",
    "price_dec",
    "handl_inst",
    "tick_id",
    "market_slippage",
    "ref_px",
    "ref_px_type",
    "tplusone",
    "created_at",
    "updated_at"
)
VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29, $30, $31, $32, $33, $34
)
RETURNING *;

-- name: GetInstrument :one
SELECT *
FROM instruments
WHERE "product_code" = $1 AND "exchange_code" = $2 AND "instrument_code" = $3
LIMIT 1;

-- name: GetInstrumentForUpdate :one
SELECT *
FROM instruments
WHERE "product_code" = $1 AND "exchange_code" = $2 AND "instrument_code" = $3
LIMIT 1
FOR NO KEY UPDATE;

-- name: ListInstruments :many
SELECT *
FROM instruments
ORDER BY "product_code", "exchange_code", "instrument_code"
LIMIT $1
OFFSET $2;

-- name: UpdateInstrument :one
UPDATE instruments
SET "underlying_code" = $2,
    "instrument_desc" = $3,
    "maturity_monthyear" = $4,
    "maturity_day" = $5,
    "put_or_call" = $6,
    "strike_price" = $7,
    "symbol" = $8,
    "symbol_sfx" = $9,
    "cusip" = $10,
    "isin" = $11,
    "ric" = $12,
    "sedol" = $13,
    "blpsyn" = $14,
    "cficode" = $15,
    "price_code" = $16,
    "effective_date" = $17,
    "minqty" = $18,
    "sector" = $19,
    "contract_size" = $20,
    "corporate_action" = $21,
    "trading_session_id" = $22,
    "trading_curr" = $23,
    "price_dec" = $24,
    "handl_inst" = $25,
    "tick_id" = $26,
    "market_slippage" = $27,
    "ref_px" = $28,
    "ref_px_type" = $29,
    "tplusone" = $30,
    "updated_at" = NOW()
WHERE "product_code" = $1 AND "exchange_code" = $2 AND "instrument_code" = $3
RETURNING *;

-- name: DeleteInstrument :exec
DELETE FROM instruments
WHERE "product_code" = $1 AND "exchange_code" = $2 AND "instrument_code" = $3;
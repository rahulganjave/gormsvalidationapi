-- name: CreateCurrency :one
INSERT INTO currencies
    (
    currency_pair,
    bid,
    ask
    )
VALUES
    (
        $1, $2, $3
)
RETURNING *;

-- name: GetCurrency :one
SELECT *
FROM currencies
WHERE currency_pair = $1
LIMIT 1;

-- name: GetCurrencyForUpdate :one
/*SELECT *
FROM currencies
WHERE currency_pair = $1
LIMIT 1
FOR NO KEY
UPDATE;*/

-- name: ListCurrency :many
/*SELECT *
FROM currencies
WHERE currency_pair = $1
LIMIT $2
OFFSET
$3;*/

-- name: UpdateCurrency :one
UPDATE currencies
SET bid = $2,
ask = $3
WHERE currency_pair = $1
RETURNING *;

-- name: DeleteCurrency :exec
DELETE FROM currencies
WHERE currency_pair = $1;
// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	CreateAccount(ctx context.Context, arg CreateAccountParams) (Accountfilter, error)
	CreateCurrency(ctx context.Context, arg CreateCurrencyParams) (Currency, error)
	CreateInstrumentfilter(ctx context.Context, arg CreateInstrumentfilterParams) (Instrumentfilter, error)
	CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error)
	CreatePosition(ctx context.Context, arg CreatePositionParams) (Position, error)
	DeleteAccount(ctx context.Context, account string) error
	DeleteCurrency(ctx context.Context, currencyPair string) error
	DeleteInstrumentfilter(ctx context.Context, arg DeleteInstrumentfilterParams) error
	DeleteOrder(ctx context.Context, orderID string) error
	DeletePosition(ctx context.Context, positionID string) error
	GetAccount(ctx context.Context, arg GetAccountParams) (Accountfilter, error)
	GetAccountForUpdate(ctx context.Context, arg GetAccountForUpdateParams) (Accountfilter, error)
	GetCurrency(ctx context.Context, currencyPair string) (Currency, error)
	GetCurrencyForUpdate(ctx context.Context, arg GetCurrencyForUpdateParams) (Currency, error)
	GetInstrumentfilter(ctx context.Context, arg GetInstrumentfilterParams) (Instrumentfilter, error)
	GetInstrumentfilterForUpdate(ctx context.Context, arg GetInstrumentfilterForUpdateParams) (Instrumentfilter, error)
	GetOrder(ctx context.Context, orderID string) (Order, error)
	GetOrdersForUpdate(ctx context.Context, arg GetOrdersForUpdateParams) (Order, error)
	GetPosition(ctx context.Context, positionID string) (Position, error)
	GetPositionsForUpdate(ctx context.Context, positionID string) (Position, error)
	ListAccount(ctx context.Context, arg ListAccountParams) ([]Accountfilter, error)
	ListInstrumentfilter(ctx context.Context, arg ListInstrumentfilterParams) ([]Instrumentfilter, error)
	ListPositions(ctx context.Context, arg ListPositionsParams) ([]Position, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Accountfilter, error)
	UpdateInstrumentfilter(ctx context.Context, arg UpdateInstrumentfilterParams) (Instrumentfilter, error)
	UpdatePosition(ctx context.Context, arg UpdatePositionParams) (Position, error)
}

var _ Querier = (*Queries)(nil)

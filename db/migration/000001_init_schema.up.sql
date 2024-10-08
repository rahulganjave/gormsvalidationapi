CREATE TABLE "orders"
(
    "order_id" varchar(20) NOT NULL,
    "client_id" varchar(10) NOT NULL,
    "trading_acct" varchar(25) NOT NULL,
    "exchange" varchar(6) NOT NULL,
    "product" varchar(3) NOT NULL,
    "instrument" varchar(32) NOT NULL,
    "side" char(1),
    "status" char(1),
    "order_type" char(1),
    "order_qty" decimal(24,9) DEFAULT 0,
    "order_price" decimal(24,9) DEFAULT 0,
    "time_inforce" char(1),
    "trading_session_id" varchar(10),
    "expire_date" date,
    "exec_qty" decimal(24,9) DEFAULT 0,
    "exec_price" decimal(24,9) DEFAULT 0,
    "cash" decimal(24,9) DEFAULT 0,
    "im" decimal(24,9) DEFAULT 0,
    "exposure" decimal(24,9) DEFAULT 0,
    "est_cost_to_close" decimal(24,9) DEFAULT 0,
    "order_time" timestamptz,
    "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "positions"
(
    "position_id" varchar(20) NOT NULL,
    "ref_position_id" varchar(20),
    "trading_acct" varchar(25) NOT NULL,
    "client_id" varchar(10) NOT NULL,
    "product" varchar(3) NOT NULL,
    "exchange" varchar(6) NOT NULL,
    "instrument" varchar(32) NOT NULL,
    "side" char(1),
    "status" char(1),
    "position_type" char(1),
    "settlement_status" char(1),
    "calc_type" char(1),
    "open_qty" decimal(24,9) DEFAULT 0,
    "open_px" decimal(24,9) DEFAULT 0,
    "closed_qty" decimal(24,9) DEFAULT 0,
    "closed_px" decimal(24,9) DEFAULT 0,
    "trading_curr" varchar(3),
    "settle_curr" varchar(3),
    "mark_price" decimal(24,9) DEFAULT 0,
    "mark_value" decimal(24,9) DEFAULT 0,
    "exchange_rate" decimal(24,9) DEFAULT 0,
    "unrealized_pl" decimal(24,9) DEFAULT 0,
    "realized_pl" decimal(24,9) DEFAULT 0,
    "mm_and_cash" decimal(24,9) DEFAULT 0,
    "collateral_haircut_ratio" decimal(24,9) DEFAULT 0,
    "collateral_haircut" decimal(24,9) DEFAULT 0,
    "open_date" timestamptz,
    "closed_date" timestamptz,
    "cost_to_close" decimal(24,9) DEFAULT 0,
    "trade_date" date,
    "trading_session_id" varchar(10),
    "underlying" varchar(32),
    "open_close" char(1),
    "put_or_call" char(1),
    "contract_size" decimal(24,9) DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    UNIQUE("position_id", "trading_acct", "client_id")
);

CREATE TABLE "accountfilter"
(
    --"line" BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    --"parent_line" bigint NOT NULL,
    --"path" varchar(256) NOT NULL,
    "client_id" varchar(10),
    "account" varchar(25) PRIMARY KEY NOT NULL,
    "trader" varchar(10),
    "cash_balance_bf" decimal(24,9) DEFAULT 0,
    "cash_balance" decimal(24,9) DEFAULT 0,
    "est_trans_cost" decimal(24,9) DEFAULT 0,
    "equity_realized_pl" decimal(24,9) DEFAULT 0,
    "equity_unrealized_pl" decimal(24,9) DEFAULT 0,
    "margin_realized_pl" decimal(24,9) DEFAULT 0,
    "margin_unrealized_pl" decimal(24,9) DEFAULT 0,
    "est_closing_cost" decimal(24,9) DEFAULT 0,
    "transaction_not_booked" decimal(24,9) DEFAULT 0,
    "equity_order" decimal(24,9) DEFAULT 0,
    "account_value" decimal(24,9) DEFAULT 0,
    "deposit_withdraw" decimal(24,9) DEFAULT 0,
    "collateral_haircut" decimal(24,9) DEFAULT 0,
    "equity_position_value" decimal(24,9) DEFAULT 0,
    "margin_requirement" decimal(24,9) DEFAULT 0,
    "currency" VARCHAR(3),
    "buy_limit" decimal(30,9) DEFAULT 0,
    "sell_limit" decimal(30,9) DEFAULT 0,
    "buy_exposure" decimal(30,9) DEFAULT 0,
    "sell_exposure" decimal(30,9) DEFAULT 0,
    "loan_limit" decimal(30,9) DEFAULT 0,
    "margin_factor" decimal(24,9) DEFAULT 0,
    "authorized_order_type" bigint DEFAULT 0,
    "calc_mode" char(1) DEFAULT '0',
    "limit_mode" int DEFAULT 0,
    "action" char(1),
    "netting_mode" char(1) NOT NULL DEFAULT '0',
    "max_ord_qty" decimal(24,9) NOT NULL DEFAULT 0,
    "max_ord_val" decimal(24,9) NOT NULL DEFAULT 0,
    "no_of_transaction" int NOT NULL DEFAULT 0,
    "transaction_interval" smallint NOT NULL DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT 'now()'
);

CREATE TABLE "instrumentfilter"
(
    --"line" BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    --"path" varchar(256) NOT NULL,
    "product" varchar(3) NOT NULL,
    "exchange" varchar(6) NOT NULL,
    "instrument" varchar(32) NOT NULL,
    "underlying" varchar(32),
    "filter" int,
    "max_order_qty" decimal(24,9) DEFAULT 0,
    "max_order_price" decimal(24,9) DEFAULT 0,
    "max_order_value" decimal(24,9) DEFAULT 0,
    "buy_lmt" decimal(24,9) DEFAULT 0,
    "sell_lmt" decimal(24,9) DEFAULT 0,
    "exposure_factor" decimal(24,9) DEFAULT 0,
    "buy_exposure" decimal(24,9) DEFAULT 0,
    "sell_exposure" decimal(24,9) DEFAULT 0,
    "buy_qty_lmt" decimal(24,9) DEFAULT 0,
    "buy_qty_exposure" decimal(24,9) DEFAULT 0,
    "sell_qty_lmt" decimal(24,9) DEFAULT 0,
    "sell_qty_exposure" decimal(24,9) DEFAULT 0,
    "net_qty_lmt" decimal(24,9) DEFAULT 0,
    "net_qty_exposure" decimal(24,9) DEFAULT 0,
    "margin_type" char(1),
    "im" decimal(24,9) DEFAULT 0,
    "mm" decimal(24,9) DEFAULT 0,
    "marginable_ratio" decimal(24,9) DEFAULT 0,
    "margin_curr" varchar(3),
    "commission_type" char(1),
    "commission" decimal(24,9) DEFAULT 0,
    "min_commission" decimal(24,9) DEFAULT 0,
    "commission_curr" varchar(3),
    "authorized_order_type" bigint DEFAULT 0,
    "trading_curr" varchar(3),
    "price_code" varchar(30),
    "action_on_record" smallint NOT NULL,
    "ref_px" decimal(24,9) DEFAULT 0,
    "ref_px_type" smallint,
    "contract_size" decimal(24,9) DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    UNIQUE("product", "exchange", "instrument")
);

CREATE TABLE "currencies"
(
    "currency_pair" varchar(6) PRIMARY KEY NOT NULL,
    "bid" decimal(24,9) DEFAULT 0,
    "ask" decimal(24,9) DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz
);

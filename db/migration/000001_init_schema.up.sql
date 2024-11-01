CREATE TABLE "orders" (
    "order_id" varchar(20) NOT NULL,                    -- Unique order identifier e.g. TradingAcct+Exchange+Instrument+random_number
    "client_id" varchar(10) NOT NULL,                   -- client id
    "trading_acct" varchar(25) NOT NULL,                -- predefine set of trading_accts
    "product_code" varchar(3) NOT NULL,                 -- product code e.g. 'CS','OPT','CFD', 'FUT'
    "exchange_code" varchar(6) NOT NULL,                -- exchange code e.g. 'NASD', 'NYSE', 'AMEX', 'SEHK', 'SGX'
    "instrument_code" varchar(32) NOT NULL,             -- predefine set of instruments
    "order_side" int NOT NULL DEFAULT 1,                /* Side of the Order, depends on the value set:
                                                            1 - Long
                                                            2 - Short */
    "order_status" int NOT NULL DEFAULT 1,              /* Status of the position whether it is open, closed etc. Depending on the value set:
                                                            1 - open
                                                            2 - closed
                                                            3 - half-closed
                                                            4 - total (total position as per account and product) 
                                                            default is 1 */
    "order_type" int NOT NULL DEFAULT 1,                -- order type
    "order_qty" decimal(24,9) DEFAULT 0,                -- order Qty , DEFAULT value 0.0
    "order_price" decimal(24,9) DEFAULT 0,              -- order price, DEFAULT value 0.0
    "time_inforce" int DEFAULT 0,                       -- time in force , DEFAULT value 0(day)
    "trading_session_id" varchar(10),                   -- Used to identify which trading session is used in the trading session table
    "expire_date" date,                                 -- expire date, Local market date YYYYMMDD
    "exec_qty" decimal(24,9) DEFAULT 0,                 -- Executed Qty , DEFAULT value 0.0
    "exec_price" decimal(24,9) DEFAULT 0,               -- Executed price, DEFAULT value 0.0
    "cash" decimal(24,9) DEFAULT 0,                     -- Cash , DEFAULT value 0.0
    "Initial_margin" decimal(24,9) DEFAULT 0,           -- Initial margin for this instrument, DEFAULT value 0.0
    "exposure" decimal(24,9) DEFAULT 0,                 -- exposure , DEFAULT value 0.0
    "est_cost_to_close" decimal(24,9) DEFAULT 0,        -- Estimated cost to close the position, DEFAULT value 0.0
    "order_time" timestamptz,                           -- order time
    "underlying_code" varchar(32),                      -- Underlying code for the instrument
    "open_close" int NOT NULL DEFAULT 79,              -- OPT position with open or close 
    "put_or_call" int,                                  -- Determines whether to put or call this instrument, depending on mode:  0 - Put & 1 - Call
    "contract_size" decimal(24,9) DEFAULT 1,            -- Contract size for the instrument, DEFAULT value 1.0
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz
);

CREATE TABLE "positions"
(
    "position_id" varchar(20) NOT NULL,                 -- Unique position identifier e.g. TradingAcct+Exchange+Instrument+random_number
    "ref_position_id" varchar(20),                      -- can blank value
    "trading_acct" varchar(25) NOT NULL,                -- predefine set of trading_accts
    "client_id" varchar(10) NOT NULL,                   -- client id e.g. 'PSPL', 'PSHK'
    "product_code" varchar(3) NOT NULL,                 -- product code e.g. 'ES','OPT','CFD', 'FUT'
    "exchange_code" varchar(6) NOT NULL,                -- exchange code e.g. 'NASD', 'NYSE', 'AMEX', 'SEHK', 'SGX'
    "instrument_code" varchar(32) NOT NULL,             -- predefine set of instruments
    "position_side" int NOT NULL DEFAULT 1,             /* Side of the position, depends on the value set:
                                                            1 - Long
                                                            2 - Short */
    "position_status" int NOT NULL DEFAULT 1,           /* Status of the position whether it is open, closed etc. Depending on the value set:
                                                            1 - open
                                                            2 - closed
                                                            3 - half-closed
                                                            4 - total (total position as per account and product) 
                                                            default is 1 */
    "position_type" int NOT NULL DEFAULT 0,             /* 0 - history
                                                            1 - intraday
                                                            default is 0 */
    "settlement_status" int NOT NULL DEFAULT 0,         /*  To indicate whether the position is a settled position or still an outstanding position
                                                            0 = outstanding position
                                                            1 = settled position 
                                                            default is 0 */
    "calc_type" int NOT NULL DEFAULT 0,                 /*  0 - recalculate
                                                            1 - not recalculate
                                                            default is 0 */
    "open_qty" decimal(24,9) NOT NULL DEFAULT 0,        -- Quantity traded to open the position, DEFAULT value 0.0
    "open_px" decimal(24,9) NOT NULL DEFAULT 0,         -- Price traded to open the position, DEFAULT value 0.0
    "closed_qty" decimal(24,9) NOT NULL DEFAULT 0,      -- Quantity traded to close the position, DEFAULT value 0.0
    "closed_px" decimal(24,9) NOT NULL DEFAULT 0,       -- Price traded to close the position, DEFAULT value 0.0
    "trading_curr" varchar(3) NOT NULL,                 -- Trading currency e.g. USD, HKD, SGD
    "settle_curr" varchar(3) NOT NULL,                  -- Position settlement currency e.g. USD, HKD, SGD
    "mark_price" decimal(24,9) DEFAULT 0,               -- Mark to market price, DEFAULT value 0.0
    "mark_value" decimal(24,9) DEFAULT 0,               -- Mark to market value, DEFAULT value 0.0
    "exchange_rate" decimal(24,9) DEFAULT 0,            -- exchange rate , DEFAULT value 0.0
    "unrealized_pl" decimal(24,9) DEFAULT 0,            -- Unrealized profit and loss, DEFAULT value 0.0
    "realized_pl" decimal(24,9) DEFAULT 0,              -- Realized profit and loss, DEFAULT value 0.0
    "mm_and_cash" decimal(24,9) DEFAULT 0,              -- Amount charged for intraday transaction. Non margin instruments charge cash, marginable instruments charge the maintenance margin, DEFAULT value 0.0
    "collateral_haircut_ratio" decimal(24,9) DEFAULT 0, -- Percentage unavailable for margin trading, DEFAULT value 0.0
    "collateral_haircut" decimal(24,9) DEFAULT 0,       -- Amount unavailable for margin trading, DEFAULT value 0.0
    "open_date" timestamptz,                            -- Date of opening for the position
    "closed_date" timestamptz,                          -- Date of closing for the position
    "cost_to_close" decimal(24,9) DEFAULT 0,            -- Cost to close the position, DEFAULT value 0.0
    "trade_date" date,                                  -- Local market date, YYYYMMDD
    "trading_session_id" varchar(10),                   -- Used to identify which trading session is used in the trading session table, e.g. Exchange_Product
    "underlying_code" varchar(32),                      -- Underlying code for the instrument
    "open_close" int NOT NULL DEFAULT 79,              -- OPT position with open or close 
    "put_or_call" int,                              -- Determines whether to put or call this instrument, depending on mode: 0 - Put, 1 - Call 
    "contract_size" decimal(24,9) DEFAULT 1,            -- Contract size for the instrument, DEFAULT value 1.0
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz,
    UNIQUE("position_id", "trading_acct", "client_id")
);

CREATE TABLE "accountfilter"
(
    --"line" BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    --"parent_line" bigint NOT NULL,
    --"path" varchar(256) NOT NULL,
    "client_id" varchar(25),                        -- Client ID
    "trading_acct" varchar(25) PRIMARY KEY NOT NULL,-- Trading account name.
    "trader_id" varchar(25),                        -- Dealer id
    "cash_balance_bf" decimal(24,9) DEFAULT 0,      -- Cash balance available for this account before, DEFAULT value 0.0
    "cash_balance" decimal(24,9) DEFAULT 0,         -- Cash balance available for this account, DEFAULT value 0.0
    "est_trans_cost" decimal(24,9) DEFAULT 0,       -- estimated transaction cost for this account, DEFAULT value 0.0
    "equity_realized_pl" decimal(24,9) DEFAULT 0,   -- equity realized profit/loss for this account, DEFAULT value 0.0
    "equity_unrealized_pl" decimal(24,9) DEFAULT 0, -- equity unrealized profit/loss for this account, DEFAULT value 0.0
    "margin_realized_pl" decimal(24,9) DEFAULT 0,   -- margin realized profit/loss for this account, DEFAULT value 0.0
    "margin_unrealized_pl" decimal(24,9) DEFAULT 0, -- margin unrealized profit/loss for this account, DEFAULT value 0.0
    "est_closing_cost" decimal(24,9) DEFAULT 0,     -- Cash balance available for this account, DEFAULT value 0.0
    "transaction_not_booked" decimal(24,9) DEFAULT 0,-- estimated closing cost for this account, DEFAULT value 0.0
    "equity_order" decimal(24,9) DEFAULT 0,         -- equity order for this account, DEFAULT value 0.0
    "account_value" decimal(24,9) DEFAULT 0,        -- account value for this account, DEFAULT value 0.0
    "deposit_withdraw" decimal(24,9) DEFAULT 0,     -- Deposit/Withdraw, DEFAULT value 0.0
    "collateral_haircut" decimal(24,9) DEFAULT 0,   -- collateral haircut, DEFAULT value 0.0
    "equity_position_value" decimal(24,9) DEFAULT 0,-- equity position value, DEFAULT value 0.0
    "margin_requirement" decimal(24,9) DEFAULT 0,   -- margin requirement, DEFAULT value 0.0
    "currency" VARCHAR(3),                          -- Currency used for this account
    "buy_limit" decimal(30,9) DEFAULT 0,            -- Maximum buy limit, DEFAULT value 0.0
    "sell_limit" decimal(30,9) DEFAULT 0,           -- Maximum sell limit, DEFAULT value 0.0
    "buy_exposure" decimal(30,9) DEFAULT 0,         -- buy exposure, DEFAULT value 0.0
    "sell_exposure" decimal(30,9) DEFAULT 0,        -- sell exposure, DEFAULT value 0.0
    "loan_limit" decimal(30,9) DEFAULT 0,           -- Loan Limit, DEFAULT value 0.0
    "net_loss_limit" decimal(30,9) DEFAULT 0,       -- Maximum net loss limit, DEFAULT value 0.0
    "net_exposure_limit" decimal(30,9) DEFAULT 0,   -- Maximum net exposure limit, DEFAULT value 0.0
    "margin_factor" decimal(24,9) DEFAULT 0,        -- Margin adjustment factor. Default value is 1
    "authorized_order_type" bigint NOT NULL DEFAULT 0,       /* Order supported by exchange or counter party based on bit combination.Bit definition (0:disallow, 1: allow)
                                                        1 - Buy order 
                                                        2 - Sell order
                                                        3 - Short sell
                                                        4 - Market order
                                                        5 - Limit order
                                                        6 - Market to limit order
                                                        7 - GTD order
                                                        8 - GTC order
                                                        9- FOK order
                                                        10 - FAK order
                                                        11 - Stop Order(STP)
                                                        12 - Stop Limit Order(STP LMT)
                                                        13 - Market If Touch(MIT)
                                                        14 - Limit If Touch(LIT)
                                                        15 - Market On Close(MOC)
                                                        16 - Limit On Close(LOC)
                                                        17 - Market On Open(MOO)
                                                        18 - Limit On Open(LOO)
                                                        19 - Trailing Stop Order
                                                        20 - Trailing Stop Limit Order
                                                        21- OTO Order
                                                        22- OCO Order
                                                        23- Basket Order
                                                        24 = Buying-in
                                                        , DEFAULT value 0 */
    "calc_mode" int NOT NULL DEFAULT 0,                /* Calculation mode
                                                        0 = NONE, pass the order directly
                                                        1 = Cash A/C
                                                        2 = Margin A/C
                                                        , DEFAULT value 0*/
    "limit_mode" int NOT NULL DEFAULT 0,                     /* To indicate whether and how to check the limit (bit) (0=||, 1 = &)
                                                        1 =  check buy limit &/|| calc_mode
                                                        2 =  check sell limit &/|| calc_mode
                                                        3 =  check margin loan limit & calc_mode
                                                        4 - Enable net loss limit check
                                                        5 - Enable net exposure limit check
                                                        */
    "action_mode" int NOT NULL DEFAULT 0,                          /* Action taken when rules are breached for this account, depending on mode set. (mode 0, 1 , 2)
                                                        0 - Reject all
                                                        1 - accept all
                                                        2 - handle by dealer manually*/
    "netting_mode" int NOT NULL DEFAULT 0,    -- netting_mode, DEFAULT value 0
    "max_ord_qty" decimal(24,9) NOT NULL DEFAULT 0, -- max_ord_qty, DEFAULT value 0.0
    "max_ord_val" decimal(24,9) NOT NULL DEFAULT 0, -- max_ord_val, DEFAULT value 0.0
    "no_of_transaction" int NOT NULL DEFAULT 0,     -- no_of_transaction, DEFAULT value 0
    "transaction_interval" smallint NOT NULL DEFAULT 0,-- transaction_interval, DEFAULT value 0
    "created_at" timestamptz NOT NULL DEFAULT 'now()',   -- Current time
    "updated_at" timestamptz
);

CREATE TABLE "instrumentfilter"
(
    --"line" BIGINT GENERATED BY DEFAULT AS IDENTITY NOT NULL,
    --"path" varchar(256) NOT NULL,
    "product_code" varchar(3) NOT NULL,         -- product code e.g. 'ES','OPT','CFD', 'FUT', 'BO'
    "exchange_code" varchar(6) NOT NULL,        -- exchange code e.g. 'NASD', 'NYSE', 'AMEX', 'SEHK', 'SGX'
    "instrument_code" varchar(32) NOT NULL,     -- predefine set of instruments
    "underlying_code" varchar(32),              --  Underlying code for the instrument
    "filter_option" int NOT NULL DEFAULT 0,     /* Checking mode: (1 enabled, 0 disabled)
                                                    1 = Max order qty check
                                                    2 = Max order price
                                                    3 = Max order value
                                                    4 = Buy limit
                                                    5 = Sell limit
                                                    6 = Buy Qty Limit
                                                    7 = Sell Qty Limit
                                                    8 = Net Qty Limit
                                                    9 = order price range based on lower and higher price*/
    "max_order_qty" decimal(24,9) DEFAULT 0,    -- Maximum order quantity for the instrument, DEFAULT value 0.0
    "max_order_price" decimal(24,9) DEFAULT 0,  -- Maximum order price for the instrument, DEFAULT value 0.0
    "max_order_value" decimal(24,9) DEFAULT 0,  -- Maximum order value for single order, DEFAULT value 0.0
    "buy_limit" decimal(24,9) DEFAULT 0,        -- Maximum buying limit, DEFAULT value 0.0
    "sell_limit" decimal(24,9) DEFAULT 0,       -- Maximum selling limit, DEFAULT value 0.0
    "exposure_factor" decimal(24,9) DEFAULT 0,  -- Exposure factor used in instrument exposure calculation, DEFAULT value 0.0
    "buy_exposure" decimal(24,9) DEFAULT 0,     -- buy_exposure, DEFAULT value 0.0
    "sell_exposure" decimal(24,9) DEFAULT 0,    -- sell_exposure, DEFAULT value 0.0
    "buy_qty_limit" decimal(24,9) DEFAULT 0,    -- Maximum position and buy orders, DEFAULT value 0.0
    "buy_qty_exposure" decimal(24,9) DEFAULT 0, -- buy_qty_exposure DEFAULT value 0.0
    "sell_qty_limit" decimal(24,9) DEFAULT 0,   -- Maximum position and sell orders, DEFAULT value 0.0
    "sell_qty_exposure" decimal(24,9) DEFAULT 0,-- sell_qty_exposure, DEFAULT value 0.0
    "net_qty_limit" decimal(24,9) DEFAULT 0,    -- Sum of maximum net position quantity and orders quantity, DEFAULT value 0.0
    "net_qty_exposure" decimal(24,9) DEFAULT 0, -- net_qty_exposure , DEFAULT value 0.0
    "margin_type" int NOT NULL DEFAULT 0,       /* Indicates the margin type for this instrument, depending on mode set:
                                                    0 - non-margin product
                                                    1 - % per trades value
                                                    2 - value per lot
                                                    3 - % per trade size
                                                    used to indicate the product is margin or non-margin product. */
    "Initial_margin" decimal(24,9) DEFAULT 0,    -- Initial margin for this instrument, DEFAULT value 0.0
    "Maintenance_margin" decimal(24,9) DEFAULT 0,-- Maintenance margin for this instrument, DEFAULT value 0.0
    "marginable_ratio" decimal(24,9) DEFAULT 0, -- Marginable ratio for this instrument used only for margin account (accountfiler.calc_mode=2). Non-margin account not in use, DEFAULT value 0.0
    "margin_curr" varchar(3),                   -- Margin currency
    "commission_type" int NOT NULL DEFAULT 1,   /* Indicates the commission type for this instrument, depending on value set:
                                                    1 - % per trades value
                                                    2 - value per lot
                                                    3 - fixed value per trades
                                                    */
    "commission" decimal(24,9) DEFAULT 0,       -- Commission rate for this instrument, DEFAULT value 0.0
    "min_commission" decimal(24,9) DEFAULT 0,   -- Minimum commission charges for this instrument, DEFAULT value 0.0
    "commission_curr" varchar(3),               -- Currency used for commission
    "authorized_order_type" bigint NOT NULL DEFAULT 0,   /* Order supported by exchange or counter party based on bit combination.Bit definition (0:disallow, 1: allow)
                                                    1 - Buy order 
                                                    2 - Sell order
                                                    3 - Short sell
                                                    4 - Market order
                                                    5 - Limit order
                                                    6 - Market to limit order
                                                    7 - GTD order
                                                    8 - GTC order
                                                    9- FOK order
                                                    10 - FAK order
                                                    11 - Stop Order(STP)
                                                    12 - Stop Limit Order(STP LMT)
                                                    13 - Market If Touch(MIT)
                                                    14 - Limit If Touch(LIT)
                                                    15 - Market On Close(MOC)
                                                    16 - Limit On Close(LOC)
                                                    17 - Market On Open(MOO)
                                                    18 - Limit On Open(LOO)
                                                    19 - Trailing Stop Order
                                                    20 - Trailing Stop Limit Order
                                                    21- OTO Order
                                                    22- OCO Order
                                                    23- Basket Order
                                                    24 = Buying-in
                                                    */
    "trading_curr" varchar(3),                  -- Trading currency for the instrument e.g. USD, HKD, SGD
    "price_code" VARCHAR(20),                   -- Code used to receive price update from WDS or PMP (reserved column, not in use)
    "ref_px" decimal(24,9) NOT NULL DEFAULT 0,  -- referance price, DEFAULT value 0.0
    "ref_px_type" smallint NOT NULL DEFAULT 0,  /* Reference price type:
                                                    1 = Best bid price
                                                    2 = Last done price
                                                    3 = Best ask price
                                                    4 = Mid price
                                                    5 = Previous closing price
                                                    10 = Indicativeprice
                                                    */
    "contract_size" decimal(24,9) DEFAULT 1,    -- Contract size for the instrument, DEFAULT value 1.0
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz,
    UNIQUE("product_code", "exchange_code", "instrument_code") -- Assuming these three together form a unique identifier
);

CREATE TABLE "currencies"
(
    "currency_pair" varchar(6) PRIMARY KEY NOT NULL,
    "bid" decimal(24,9) DEFAULT 0,
    "ask" decimal(24,9) DEFAULT 0,
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz
);

CREATE TABLE "instruments" (
    product_code varchar(3) NOT NULL,       -- product code e.g. 'ES','OPT','CFD', 'FUT'
    exchange_code varchar(6) NOT NULL,      -- exchange code e.g. 'NASD', 'NYSE', 'AMEX', 'SEHK', 'SGX'
    instrument_code varchar(32) NOT NULL,   -- predefine set of instruments
    underlying_code VARCHAR(10),            -- Underlying code for the instrument
    instrument_desc VARCHAR(30),            -- Full name of the instrument
    maturity_monthyear VARCHAR(6),          -- Month and year for the maturity date of the instrument Format: yyyymm (reserved column, not in use)
    maturity_day VARCHAR(2),                -- Maturity day for the maturity date of the instrument. Format: DD(reserved column, not in use)
    put_or_call CHAR(1),                    /* Determines whether to put or call this instrument, depending on mode:
                                                0 - Put
                                                1 - Call
                                                (reserved column, not in use)*/
    strike_price decimal(24,9) DEFAULT 0,   -- Strike price for the instrument (reserved column, not in use)
    symbol VARCHAR(10),                     -- Exchange symbol
    symbol_sfx VARCHAR(15),                 -- Symbol suffix for use in US market
    cusip VARCHAR(15),                      -- CUSIP code for the instrument
    isin VARCHAR(15),                       -- ISIN code for the instrument
    ric VARCHAR(15),                        -- RIC code for the instrument
    sedol VARCHAR(15),                      -- Stock Exchange Daily Official List
    blpsyn VARCHAR(15),                     -- Bloomberg symbol for the instrument
    cficode VARCHAR(30),                    -- CFI code for the instrument
    price_code VARCHAR(20),                 -- Code used to receive price update from WDS or PMP (reserved column, not in use)
    effective_date TIMESTAMP,               -- Effective date for the instrument (reserved column, not in use)
    minqty int NOT NULL DEFAULT 0,          -- Minimum quantity for each order of the instrument
    sector VARCHAR(3),                      -- To which sector the instrument belongs (reserved column, not in use)
    contract_size decimal(24,9) NOT NULL DEFAULT 1,  -- Contract size for the instrument, DEFAULT value 1.0
    corporate_action VARCHAR(30),           -- Corporate actions, delimited by "," (reserved column, not in use)
    trading_session_id VARCHAR(10),         -- Trading session ID for this instrument, e.g. Exchange_Product
    trading_curr VARCHAR(3),                -- Trading currency for the instrument
    price_dec decimal(24,9) NOT NULL DEFAULT 0,      -- Number of decimal places for the price of the instrument (reserved column, not in use), DEFAULT value 0
    handl_inst CHAR(1),                     /* Handling instruction for the instrument, depending on mode set:
                                                0 - following the definition in exchange table, or else, take this value directly and insert into orders table
                                                1 - DMA, Automated execution order, private, no Broker intervention
                                                2 - Non-DMA, Automated execution order, public, Broker intervention OK. (Default)
                                                3 - Manual order, best execution*/
    tick_id int NOT NULL DEFAULT 0,         -- Tick ID for this instrument
    market_slippage decimal(24,9) DEFAULT 0,-- Market slippage percentage for market orders, DEFAULT value 0
    ref_px decimal(24,9) NOT NULL DEFAULT 0,-- Reference price , DEFAULT value 0.0
    ref_px_type smallint NOT NULL DEFAULT 0,     /* Reference price type:
                                                1 = Best bid price
                                                2 = Last done price
                                                3 = Best ask price
                                                4 = Mid price
                                                5 = Previous closing price
                                                10 = Indicativeprice
                                                , DEFAULT value 0 */
    tplusone int,                       -- T+1 restriction indicator
    "created_at" timestamptz NOT NULL DEFAULT 'now()',
    "updated_at" timestamptz,
    UNIQUE(product_code, exchange_code, instrument_code) -- Assuming these three together form a unique identifier
);
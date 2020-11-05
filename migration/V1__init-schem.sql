CREATE TABLE PRODUCTS
(
    product_id    SERIAL,
    product_name  VARCHAR(250) NOT NULL,
    product_price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    base_currency VARCHAR(3) NOT NULL default 'USD',
    product_title VARCHAR(250) NOT NULL,
    product_desc  TEXT NOT NULL,
    image_url   VARCHAR(250) NOT NULL
    CONSTRAINT products_pkey PRIMARY KEY (product_id)
);

CREATE TABLE USERS
(
    user_id        SERIAL,
    first_name     VARCHAR(250) NOT NULL,
    middle_name    VARCHAR(250) NOT NULL,
    last_name      VARCHAR(250) NOT NULL,
    user_email     VARCHAR(250) NOT NULL,
    password       VARCHAR(250) NOT NULL,
    phone_number   VARCHAR(15),
    date_create    DATE NOT NULL  default  CURRENT_DATE,
    date_updated   DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
);

CREATE TABLE shipping_address
(
    id             SERIAL,
    user_id        BIGINT NOT NULL,
    user_name      VARCHAR(250) NOT NULL,
    user_email     VARCHAR(250) NOT NULL,
    phone_number   VARCHAR(15) NOT NULL,
    pin_code       VARCHAR(15) NOT NULL,
    address1       VARCHAR(250) NOT NULL,
    address2       VARCHAR(250) NOT NULL,
    street         VARCHAR(250) NOT NULL,
    landmark       VARCHAR(250) NOT NULL,
    date_create    DATE NOT NULL  default  CURRENT_DATE,
    date_updated   DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT shipping_address_pkey PRIMARY KEY (id)
);

CREATE TABLE carts
(
    cart_id        SERIAL,
    cart_name      VARCHAR(250) NOT NULL,
    user_id        BIGINT NOT NULL,
    date_create    DATE NOT NULL  default  CURRENT_DATE,
    date_updated   DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT carts_pkey PRIMARY KEY (cart_id)
);

CREATE TABLE cart_items
(
    item_id        SERIAL,
    cart_id        BIGINT ,
    product_id     BIGINT ,
    product_name   VARCHAR(250) NOT NULL ,
    discount_id    INTEGER,
    date_create    DATE NOT NULL  default  CURRENT_DATE,
    date_updated   DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT cart_items_pkey PRIMARY KEY (item_id)
);

CREATE TABLE discount
(
    id             SERIAL,
    description    VARCHAR(250) NOT NULL ,
    type           VARCHAR(10) NOT NULL,
    rule_id        INTEGER,
    date_create    DATE NOT NULL  default  CURRENT_DATE,
    date_updated   DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT discount_pkey PRIMARY KEY (id)
);

CREATE TABLE discount_rules
(
    rule_id            SERIAL,
    description        VARCHAR(250) NOT NULL ,
    type               VARCHAR(10) NOT NULL,
    key                VARCHAR(250) NOT NULL,
    value              VARCHAR(250) NOT NULL,
    date_create        DATE NOT NULL  default  CURRENT_DATE,
    date_updated       DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT discount_rules_pkey PRIMARY KEY (rule_id)
);


CREATE TABLE discount_coupons
(
    coupon_id           SERIAL,
    description         VARCHAR(250) NOT NULL ,
    coupons_value       VARCHAR(10) NOT NULL,
    coupons_expiry      DATE NOT NULL  default  CURRENT_DATE,
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    active              BOOLEAN NOT NULL DEFAULT  true,
    CONSTRAINT discount_coupons_pkey PRIMARY KEY (coupon_id)
);


CREATE TABLE orders
(
    order_id            SERIAL,
    cart_id             BIGINT,
    user_id             BIGINT,
    amount              NUMERIC(10,2),
    active              BOOLEAN NOT NULL DEFAULT  true,
    order_data          json NOT NULL,
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT  orders_pkey PRIMARY KEY (order_id)
);

CREATE TABLE orders_items
(
    order_items_id      SERIAL,
    order_id            BIGINT,
    user_id             BIGINT,
    amount              NUMERIC(10,2),
    product_id         BIGINT ,
    product_name        VARCHAR(250) NOT NULL ,
    discount_id         INTEGER,
    order_data          json NOT NULL,
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT  orders_items_pkey PRIMARY KEY (order_id)
);

CREATE TABLE payments
(
    payment_id         SERIAL,
    user_id            BIGINT ,
    payment_mode       VARCHAR(10) NOT NULL,
    order_id           BIGINT,
    order_date         DATE,
    order_amount       NUMERIC(10,2),
    mode_of_payment    VARCHAR(10) NOT NULL,
    payment_status     VARCHAR(10) NOT NULL,
    name_on_card       VARCHAR(10) NOT NULL,
    card_tye           VARCHAR(10) NOT NULL,
    card_last_4_digit  VARCHAR(4) NOT NULL,
    order_data         json NOT NULL,
    date_create        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT transactions_pkey PRIMARY KEY (payment_id)
);

CREATE TABLE invoice
(
    invoice_id           SERIAL,
    user_id              BIGINT ,
    payment_id           BIGINT,
    shipping_address_id  BIGINT,
    order_id             BIGINT,
    user_name            VARCHAR(250) NOT NULL,
    payment_mode         VARCHAR(10) NOT NULL,
    order_date           DATE,
    order_amount         NUMERIC(10,2),
    mode_of_payment      VARCHAR(10) NOT NULL,
    payment_status       VARCHAR(10) NOT NULL,
    order_data           json NOT NULL,
    date_create          DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT invoice_pkey PRIMARY KEY (invoice_id)
);


CREATE TABLE inventory
(
    inventory_id        SERIAL,
    product_id          BIGINT NOT NULL,
    price               NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    quantity            INTEGER ,
    CONSTRAINT inventory_pkey PRIMARY KEY (inventory_id)
);
--DROP TABLE PRODUCTS CASCADE;
CREATE TABLE PRODUCTS
(
    product_id    BIGSERIAL,
    product_name  VARCHAR(250) NOT NULL,
    product_price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    base_currency VARCHAR(3) NOT NULL default 'USD',
    product_title VARCHAR(250) NOT NULL,
    product_desc  TEXT NOT NULL,
    image_url     VARCHAR(250) NOT NULL,
    CONSTRAINT products_pkey PRIMARY KEY (product_id)
);

--DROP TABLE USERS CASCADE;
CREATE TABLE USERS
(
    user_id        BIGSERIAL,
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
CREATE INDEX users_password_index ON USERS (password);
CREATE INDEX phone_number_index ON USERS (phone_number);

--DROP TABLE shipping_address CASCADE;
CREATE TABLE  shipping_address
(
    shipping_address_id             BIGSERIAL,
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
    CONSTRAINT  shipping_address_pkey PRIMARY KEY (shipping_address_id),
    CONSTRAINT  shipping_address_user_id_fk FOREIGN KEY(user_id) REFERENCES USERS(user_id)
);

--DROP TABLE offers CASCADE;
CREATE TYPE discount_types AS ENUM ('FLAT', 'PERCENTILE');
CREATE TYPE offers_type AS ENUM ('COMBO_OFFER', 'INDIVIDUAL_ITEM_OFFER','CHRISTMAS_OFFER', 'NEW_YEAR_OFFER','BLACK_FRIDAY'); --All type of discount
CREATE TABLE offers
(
    offers_id           BIGSERIAL,
    description         VARCHAR(250) NOT NULL ,
    discount            NUMERIC NOT NULL,
    discount_mode       discount_types NOT NULL,
    offers_type         offers_type NOT NULL  ,
    active              boolean NOT NULL,
    expire_date          DATE NOT NULL,
    date_created         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT offers_id_pkey PRIMARY KEY (offers_id)
);

--DROP TABLE offers_rules CASCADE;
CREATE TABLE offers_rules
(
    offers_rules_id      BIGSERIAL,
    offers_id            BIGINT,
    key                  VARCHAR(250) NOT NULL,
    value                VARCHAR(250) NOT NULL,
    operator             VARCHAR(10)  NOT NULL  default  '==',
    description          VARCHAR(250) NOT NULL ,
    date_created          DATE NOT NULL  default  CURRENT_DATE,
    date_updated         DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT offers_rules_pkey PRIMARY KEY (offers_rules_id),
    CONSTRAINT offers_rules_offers_id_fk FOREIGN KEY(offers_id) REFERENCES offers(offers_id)
);

--DROP TABLE offers_rules CASCADE;
CREATE TABLE product_offers
(
    product_offers_id    BIGSERIAL,
    offers_id            BIGINT,
    product_id           VARCHAR(250) NOT NULL,
    date_create          DATE NOT NULL  default  CURRENT_DATE,
    date_updated         DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT           product_offers_pkey            PRIMARY KEY (product_offers_id),
    CONSTRAINT           product_offers_offers_id_fk    FOREIGN KEY(offers_id) REFERENCES offers(offers_id),
    CONSTRAINT           product_offers_product_id_fk   FOREIGN KEY(product_id) REFERENCES products(product_id)
);


--DROP TABLE discount_coupons CASCADE;

CREATE TYPE discount_coupons_type AS ENUM ('ONE_TIME',  'MULTI_USE', 'NEW_YEAR_2020'); --All type of discount coupons
CREATE TABLE  discount_coupons
(
    discount_coupons_id   BIGSERIAL,
    description         VARCHAR(250) NOT NULL ,
    discount_coupon     VARCHAR(20) NOT NULL UNIQUE,
    discount            NUMERIC NOT NULL,
    discount_mode       discount_types NOT NULL,
    coupons_type        discount_coupons_type NOT NULL default  'ONE_TIME' ,
    expiry_date         DATE NOT NULL,
    date_created        DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    active              BOOLEAN NOT NULL DEFAULT  true,
    CONSTRAINT discount_coupons_id_pkey PRIMARY KEY (discount_coupons_id)
);

CREATE TABLE discount_coupons_rules
(
    discount_coupons_rules_id      BIGSERIAL,
    discount_coupons_id            BIGINT,
    key                  VARCHAR(250) NOT NULL,
    value                VARCHAR(250) NOT NULL,
    operator             VARCHAR(10)  NOT NULL  default  '==',
    description          VARCHAR(250) NOT NULL ,
    date_created          DATE NOT NULL  default  CURRENT_DATE,
    date_updated         DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT discount_coupons_rules_id_pkey PRIMARY KEY (discount_coupons_rules_id),
    CONSTRAINT discount_coupons_id_fk FOREIGN KEY(discount_coupons_id) REFERENCES discount_coupons(discount_coupons_id)
);

--DROP TABLE discount_coupons CASCADE;
CREATE TABLE  products_coupons_mapping
(
    products_coupons_mapping_id   BIGSERIAL,
    description                   VARCHAR(250) NOT NULL ,
    product_id                    BIGINT,
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    active              BOOLEAN NOT NULL DEFAULT  true,
    CONSTRAINT            products_coupons_mapping_id_pkey PRIMARY KEY (products_coupons_mapping_id),
    CONSTRAINT           products_coupons_mapping_discount_coupons_fk    FOREIGN KEY(product_id) REFERENCES discount_coupons(discount_coupons_id),
    CONSTRAINT           products_coupons_mapping_product_id_fk   FOREIGN KEY(product_id) REFERENCES products(product_id)
);



--DROP TABLE carts CASCADE;
CREATE TABLE carts
(
    cart_id             BIGSERIAL,
    cart_name           VARCHAR(250),
    user_id             BIGINT NOT NULL,
    --offers_id           BIGINT,
    discount_coupon     VARCHAR(20),
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT  carts_pkey PRIMARY KEY (cart_id),
    CONSTRAINT  carts_user_id_fk FOREIGN KEY(user_id) REFERENCES USERS(user_id),
    --CONSTRAINT  carts_offers_id_fk FOREIGN KEY(offers_id) REFERENCES offers(offers_id),
    CONSTRAINT  carts_discount_coupon_fk FOREIGN KEY(discount_coupon) REFERENCES discount_coupons(discount_coupon)
);

--DROP TABLE cart_items CASCADE;
CREATE TABLE cart_items
(
    cart_items_id       BIGSERIAL,
    cart_id             BIGINT ,
    product_id          BIGINT ,
    --product_name        VARCHAR(250) NOT NULL ,
    --product_price       NUMERIC(10,2) NOT NULL ,
    quantity            INTEGER NOT NULL DEFAULT 1,
    currency            VARCHAR(3)  default  'USD',
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT cart_items_pkey PRIMARY KEY (cart_items_id),
    CONSTRAINT cart_items_cart_id_fk FOREIGN KEY(cart_id) REFERENCES carts(cart_id),
    CONSTRAINT cart_items_product_id_fk FOREIGN KEY(product_id) REFERENCES products(product_id)
);


--DROP TABLE orders CASCADE;
CREATE TYPE order_status_type AS ENUM ('COMPLETED', 'CANCELED','PENDING_DELIVERY', 'DISPATCHED');
CREATE TABLE orders
(
    order_id            BIGSERIAL,
    cart_id             BIGINT,
    user_id             BIGINT,
    amount              NUMERIC(10,2),
    status              order_status_type,
    active              BOOLEAN NOT NULL DEFAULT  true,
    order_data          json NOT NULL,
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT  orders_pkey PRIMARY KEY (order_id)
);

--DROP TABLE orders_items CASCADE;
CREATE TABLE orders_items
(
    order_items_id      BIGSERIAL,
    order_id            BIGINT,
    product_id          BIGINT ,
    product_name        VARCHAR(250) NOT NULL ,
    product_price       NUMERIC(10,2),
    quantity            INTEGER,
    baseCurrency        VARCHAR(10) NOT NULL default  'USD',
    date_create         DATE NOT NULL  default  CURRENT_DATE,
    date_updated        DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT orders_items_pkey PRIMARY KEY (order_items_id),
    CONSTRAINT orders_items_order_id_fk FOREIGN KEY(order_id) REFERENCES orders(order_id)
);

--DROP TABLE payments CASCADE;
CREATE TABLE payments
(
    payment_id         BIGSERIAL,
    user_id            BIGINT ,
    payment_mode       VARCHAR(10) NOT NULL,
    order_id           BIGINT,
    order_date         DATE,
    order_amount       NUMERIC(10,2),
    mode_of_payment    VARCHAR(10) NOT NULL,
    payment_status     VARCHAR(10) NOT NULL,
    name_on_card       VARCHAR(10) NOT NULL,
    card_type           VARCHAR(10) NOT NULL,
    card_last_4_digit  VARCHAR(4) NOT NULL,
    order_data         json NOT NULL,
    date_create        DATE NOT NULL  default  CURRENT_DATE,
    date_updated       DATE NOT NULL  default  CURRENT_DATE,
    CONSTRAINT transactions_pkey PRIMARY KEY (payment_id)
    -- TODO add CONSTRAINT
);

--DROP TABLE invoice CASCADE;
CREATE TABLE invoice
(
    invoice_id           BIGSERIAL,
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
    -- TODO add CONSTRAINT
);

--DROP TABLE inventory CASCADE;
CREATE TABLE inventory
(
    inventory_id        BIGSERIAL,
    product_id          BIGINT NOT NULL,
    price               NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    quantity            INTEGER ,
    CONSTRAINT inventory_pkey PRIMARY KEY (inventory_id),
    CONSTRAINT inventory_product_id_fk FOREIGN KEY(product_id) REFERENCES products(product_id)
);
 
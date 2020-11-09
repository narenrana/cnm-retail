
--Add Sample Users
INSERT INTO users(user_id, first_name,middle_name,last_name, user_email, password, phone_number) VALUES (1201,'narender','','kumar',  'narenrana02@gmail.com', 'test@123','+6582924140');
INSERT INTO users(user_id, first_name,middle_name,last_name, user_email, password, phone_number) VALUES (1202,'Amit','','kumar',  'amit@gmail.com', 'test@123','+6582924140');
INSERT INTO users(user_id, first_name,middle_name,last_name, user_email, password, phone_number) VALUES (1203,'John','','',  'john@gmail.com', 'test@123','+6582924140');

--Add Sample Products

INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1001, 'Bananas', '10', 'USD','Banana' ,'','/static/images/products/fruit_25.jpg' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1002, 'Oranges', '5', 'USD','Oranges' ,'','/static/images/products/fruit_21.jpg' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1003, 'Pears', '10', 'USD','Pears' ,'','/static/images/products/fruit_37.jpeg' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1004, 'Grapes', '5', 'USD','Grapes' ,'','/static/images/products/fruit_36.jpg' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1005, 'Apples', '10', 'USD','Apples' ,'','/static/images/products/fruit_38.jpeg' );


INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1101, 'Tomato', '2.80', 'USD','Tomato' ,'','/static/images/products/product-0.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1102, 'Broccoli', '5.10', 'USD','Broccoli' ,'','/static/images/products/product-1.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1103, 'Mushroom', '6.20', 'USD','Mushroom' ,'','/static/images/products/product-2.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1104, 'Lettuce', '4.30', 'USD','Lettuce' ,'','/static/images/products/product-3.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1105, 'Radish', '9.40', 'USD','Radish' ,'','/static/images/products/product-4.png' );
--INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1106, 'Broccoli', '5.40', 'USD','Radish' ,'','/static/images/products/product-5.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1107, 'Cauliflower', '5.40', 'USD','Cauliflower' ,'','/static/images/products/product-6.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1108, 'Capcicum', '8.10', 'USD','Capcicum' ,'','/static/images/products/product-7.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1109, 'Cabbage', '2.90', 'USD','Cabbage' ,'','/static/images/products/product-8.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1110, 'Spinach', '1.50', 'USD','Spinach' ,'','/static/images/products/product-9.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1111, 'Eggplant', '3.30', 'USD','Eggplant' ,'','/static/images/products/product-10.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1112, 'Artichoke', '10.30', 'USD','Artichoke' ,'','/static/images/products/product-11.png' );

INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1113, 'Carrot', '10.30', 'USD','Carrot' ,'','/static/images/products/product-12.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1114, 'Purple Salad lettuce', '10.30', 'USD','Purple lettuce' ,'','/static/images/products/product-13.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1115, 'Cauliflower Ind', '13.30', 'USD','Cauliflower Ind' ,'','/static/images/products/product-14.png' );
INSERT INTO  PRODUCTS(product_id, product_name, product_price, base_currency, product_title, product_desc,image_url) VALUES (1116, 'Ooty Beans', '12.30', 'USD','Ooty Beans' ,'','/static/images/products/product-15.png' );




--Add discount coupons
INSERT INTO  discount_coupons(discount_coupons_id, description, discount_coupon, discount,discount_mode, expiry_date, active) VALUES (1, '30% Discount on Oranges', 'DIS237890WR',10.00,'PERCENTILE','2020-01-01', true);
INSERT INTO  discount_coupons(discount_coupons_id, description, discount_coupon, discount,discount_mode, expiry_date, active) VALUES (2, 'Discount Coupons', 'DIS237891WR',5.00,'PERCENTILE','2021-01-01', true);
INSERT INTO  discount_coupons(discount_coupons_id, description, discount_coupon, discount,discount_mode, expiry_date, active) VALUES (3, 'Discount Coupons', 'DIS237892WR',4.00,'PERCENTILE','2021-11-11', true);
INSERT INTO  discount_coupons(discount_coupons_id, description, discount_coupon, discount,discount_mode, expiry_date, active) VALUES (4, 'Discount Coupons', 'DIS237893WR',3.00,'PERCENTILE','2021-12-30', true);
INSERT INTO  discount_coupons(discount_coupons_id, description, discount_coupon, discount,discount_mode, expiry_date, active) VALUES (5, 'Discount Coupons', 'DIS237894WR',2.00,'PERCENTILE','2022-01-01', true);
INSERT INTO  discount_coupons_rules(discount_coupons_rules_id, discount_coupons_id, key, value, operator,description) VALUES (1, 1, 'Oranges', 0,'>=', '30% Discount on Oranges');


--Add offers
INSERT INTO  offers( offers_id, description, discount, discount_mode,offers_type,expire_date) VALUES (1, '4 Pears and 2 Banana Combo Offer- 30%', 30, 'PERCENTILE','COMBO_OFFER','2021-10-10');
INSERT INTO  offers_rules(offers_rules_id, offers_id, key, value, description) VALUES (1, 1, 'Pears', 4, '4 Pears and 2 Banana Combo Offer');
INSERT INTO  offers_rules(offers_rules_id, offers_id, key, value, description) VALUES (2, 1, 'Bananas', 2, '4 Pears and 2 Banana Combo Offer');

INSERT INTO  offers( offers_id, description, discount, discount_mode,offers_type,expire_date) VALUES (2, '7 or more apples - 10%', 10, 'PERCENTILE','INDIVIDUAL_ITEM_OFFER','2021-10-10');
INSERT INTO  offers_rules(offers_rules_id, offers_id, key, value, operator,description) VALUES (3, 2, 'Apples', 7,'>=', '7 or more Apples');


INSERT INTO  product_offers(product_offers_id, offers_id, product_id) VALUES (1, 2, 1002);


--increase sequence
alter sequence users_user_id_seq restart with 1300;
alter sequence offers_rules_offers_rules_id_seq restart with 1300;
alter sequence products_product_id_seq restart with 1300;
alter sequence discount_coupons_discount_coupons_id_seq restart with 1300;
alter sequence offers_offers_id_seq restart with 1300;
# Technologies and App Architecture

## Server Side

- Go Lang
- Gokit - Microservice friendly toolkit
- Goram - ORM
- Docker - App Containerization
- Database - Postgres
- Database migration - Flyway
- JWT base auth
- Microservice friendly structure

## Client Side

- React
- Redux
- Material UI

## How To run

Please make sure you have docker running in your machine and docker-compose command is working.

```
git clone https://github.com/narenrana/cnm-retail.git
cd cnm-retail
docker-compose build && docker-compose  up
```

Open `http://localhost:3500`

## Project Architecture

Each app can be independent and can be deployed and scaled individually .

![arch](./docs/arch.png)

## Database diagram

![db](./docs/db2.jpeg)

Offers are configurable . Admin can specify promotional offers in the database and what rules it holds to apply on buyings.

Coupons are also configurables .Admin can define coupon promotions and coupon rules in coupon rules tables.

### UI APP

Login Screen

![login](./docs/login.png)

Sign Up Page

![signup](./docs/signup.png)

Product Home Page

![products](./docs/products.png)

Add items in cart
![products](./docs/products-cart.png)

Checkout page

![checkout](./docs/checkout.png)

4 Pears and 2 banana scenario- 30 % discount on items
![checkout](./docs/checkout-1.png)

7 or more apple added in cart- 10% discount on apples
![checkout](./docs/checkout-2.png)

Apply a coupon DIS237890WR to get 30 % discount on oranges

![checkout](./docs/checkout-3.png)

For Demonstration purpose user can generate coupon valid for 10 or more seconds from
Coupon page

![coupon](./docs/tokens.png)

Fake Payment page
![pay](./docs/fake-pay.png)

Place order
![order](./docs/order.png)

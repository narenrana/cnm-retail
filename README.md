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
Open http://localhost:3500

## How To run as microservices

Please make sure you have minikube and kubectl are setup for locally. Please use kubeadmn for prod deployment.
Add below host entry in /etc/hosts file.
```
127.0.0.1 payment.retail.com
127.0.0.1 cnm.retail.com
```
Find your IP address using below command and update hardcoded host IP ( please use kubernetes configmap(https://kubernetes.io/docs/concepts/configuration/configmap/) and secrets(https://kubernetes.io/docs/concepts/configuration/secret/) services to load all environment variable time being these are hard coded. I am gona update this once i will get time. Please raise issue in case someone looking this on urgent basis, I will be happy to assist)

```
ifconfig |grep inet
```
Copy IP Address mentioned alongside your broadcast address.
Replace IP addr from  ```docker-compose-db.yml``` and ```connection-manager.go```

Run database and flyway migartion 
```
docker-compose -f docker-compose-db.yml up
```
Deploy pods in minikube 
```
sudo sh build.sh
```
Open `http://cnm.retail.com`

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

![products](./docs/products2.png)

Add items in cart
![products](./docs/products-cart2.png)

Checkout page

![checkout](./docs/checkout2-0.png)

4 Pears and 2 banana scenario- 30 % discount on items
![checkout](./docs/checkout2-1.png)

7 or more apple added in cart- 10% discount on apples
![checkout](./docs/checkout2-2.png)

Apply a coupon DIS237890WR to get 30 % discount on oranges

![checkout](./docs/checkout2-3.png)

For Demonstration purpose user can generate coupon valid for 10 or more seconds from
Coupon page

![coupon](./docs/tokens.png)

Fake Payment page
![pay](./docs/fake-pay.png)

Place order
![order](./docs/orders2.png)

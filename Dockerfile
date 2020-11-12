FROM golang:1.13

RUN mkdir -p /app

WORKDIR /app

#ADD . /app

COPY ./cnm-auth  ./cnm-auth
COPY ./cnm-carts ./cnm-carts
COPY ./cnm-core ./cnm-core
COPY ./cnm-coupons ./cnm-coupons
COPY ./cnm-offers ./cnm-offers
COPY ./cnm-orders ./cnm-orders
COPY ./cnm-payments ./cnm-payments
COPY ./cnm-products ./cnm-products
COPY ./cnm-users ./cnm-users
COPY ./config ./config
COPY ./migration ./migration
COPY ./go.mod ./
COPY ./main.go ./

RUN go install
RUN go build ./main.go

CMD ["./main"]
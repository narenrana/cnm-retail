package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/go-kit/kit/examples/shipping/routing"
	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
	"os"
	"os/signal"
	auth "shopping-cart/cnm-auth"
	carts "shopping-cart/cnm-carts"
	core "shopping-cart/cnm-core"
	coupons "shopping-cart/cnm-coupons"
	offers "shopping-cart/cnm-offers"
	products "shopping-cart/cnm-products"
	users "shopping-cart/cnm-users"
	"syscall"
)

const (
	defaultPort              = "8080"
	defaultRoutingServiceURL = "http://localhost:7878"
)

func main() {
	var (
		addr  = envString("PORT", defaultPort)
		rsurl = envString("ROUTINGSERVICE_URL", defaultRoutingServiceURL)

		httpAddr          = flag.String("http.addr", ":"+addr, "HTTP listen address")
		routingServiceURL = flag.String("service.routing", rsurl, "routing service URL")

		ctx = context.Background()
	)

	flag.Parse()

	var logger log.Logger
	logger = log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	logger = log.With(logger, "ts", log.DefaultTimestampUTC)



	fieldKeys := []string{"method"}

	var rs routing.Service
	rs = routing.NewProxyingMiddleware(ctx, *routingServiceURL)(rs)





	var authService auth.Service
	authService=auth.NewService();
	authService = auth.NewLoggingService(log.With(logger, "component", "auth"), authService)
	authService = auth.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "auth_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		authService,
	)

	var userService users.Service
	userService = users.NewService();
	userService = users.NewLoggingService(log.With(logger, "component", "users"), userService)
	userService = users.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "users_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "users_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		userService,
	)

	var productService products.Service
	productService = products.NewService();
	productService = products.NewLoggingService(log.With(logger, "component", "products"), productService)
	productService = products.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "products_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "products_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		productService,
	)

	var cartsService carts.Service
	cartsService = carts.NewService();
	cartsService = carts.NewLoggingService(log.With(logger, "component", "carts"), cartsService)
	cartsService = carts.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "carts_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "carts_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		cartsService,
	)

	var offersService offers.Service
	offersService = offers.NewService();
	offersService = offers.NewLoggingService(log.With(logger, "component", "offers"), offersService)
	offersService = offers.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "offers_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "offers_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		offersService,
	)

	var couponsService coupons.Service
	couponsService = coupons.NewService();
	couponsService = coupons.NewLoggingService(log.With(logger, "component", "coupons"), couponsService)
	couponsService = coupons.NewInstrumentingService(
		kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "coupons_service",
			Name:      "request_count",
			Help:      "Number of requests received.",
		}, fieldKeys),
		kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "coupons_service",
			Name:      "request_latency_microseconds",
			Help:      "Total duration of requests in microseconds.",
		}, fieldKeys),
		couponsService,
	)



	//ass = handling.NewService(handlingEvents, handlingEventFactory, handlingEventHandler)

	httpLogger := log.With(logger, "component", "http")

	mux := http.NewServeMux()

	//Initialize Database connection
	core.NewDatabaseManager();


	mux.Handle("/auth/v1/", auth.MakeHandler(authService, httpLogger))
	mux.Handle("/users/v1/", users.MakeHandler(userService, httpLogger))
	mux.Handle("/products/v1/", products.MakeHandler(productService, httpLogger))
	mux.Handle("/carts/v1/", carts.MakeHandler(cartsService, httpLogger))
	mux.Handle("/offers/v1/", offers.MakeHandler(offersService, httpLogger))
	mux.Handle("/coupons/v1/", coupons.MakeHandler(couponsService, httpLogger))



	http.Handle("/", accessControl(mux))
	http.Handle("/metrics", promhttp.Handler())

	errs := make(chan error, 2)
	go func() {
		logger.Log("transport", "http", "address", *httpAddr, "msg", "listening")
		errs <- http.ListenAndServe(*httpAddr, nil)
	}()
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	logger.Log("terminated", <-errs)
}

func accessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")

		if r.Method == "OPTIONS" {
			return
		}

		h.ServeHTTP(w, r)
	})
}

func envString(env, fallback string) string {
	e := os.Getenv(env)
	if e == "" {
		return fallback
	}
	return e
}



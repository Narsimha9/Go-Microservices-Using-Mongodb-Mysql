package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	httptransport "github.com/go-kit/kit/transport/http"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func main() {
	logger := log.NewLogfmtLogger(os.Stderr)
	r := mux.NewRouter()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here
	dbmname := "mysql"
	var svc AccountService
	svc = accountservice{}
	{
		if dbmname == "mongo" {
			db := GetMongoDB()
			repository, err := NewRepo(db, logger)
			if err != nil {
				level.Error(logger).Log("exit", err)
				os.Exit(-1)
			}
			svc = NewService(repository, logger)
		} else {
			sqldb := GetSqlDB()
			repository, err := NewSqlRepo(sqldb, logger)
			if err != nil {
				level.Error(logger).Log("exit", err)
				os.Exit(-1)
			}
			svc = NewService(repository, logger)
		}
	}

	svc = loggingMiddleware{logger, svc}
	svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}

	CreateAccountHandler := httptransport.NewServer(
		makeCreateCustomerEndpoint(svc),
		decodeCreateCustomerRequest,
		encodeResponse,
	)
	GetByIdHandler := httptransport.NewServer(
		makeGetCustomerByIdEndpoint(svc),
		decodeGetCustomerByIdRequest,
		encodeResponse,
	)
	GetAllCustomersHandler := httptransport.NewServer(
		makeGetAllCustomersEndpoint(svc),
		decodeGetAllCustomersRequest,
		encodeResponse,
	)
	DeleteCustomerHandler := httptransport.NewServer(
		makeDeleteCustomerEndpoint(svc),
		decodeDeleteCustomerRequest,
		encodeResponse,
	)
	UpdateCustomerHandler := httptransport.NewServer(
		makeUpdateCustomerEndpoint(svc),
		decodeUpdateCustomerRequest,
		encodeResponse,
	)
	http.Handle("/", r)
	http.Handle("/account", CreateAccountHandler)
	r.Handle("/account/{id}", GetByIdHandler).Methods("GET")
	r.Handle("/deletecustomer/{id}", DeleteCustomerHandler).Methods("GET")
	r.Handle("/updatecustomer/", UpdateCustomerHandler).Methods("PUT")
	r.Handle("/getall", GetAllCustomersHandler).Methods("GET")
	http.Handle("/metrics", promhttp.Handler())
	logger.Log("msg", "HTTP", "addr", ":8000")
	logger.Log("err", http.ListenAndServe(":8000", nil))
}

func Hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

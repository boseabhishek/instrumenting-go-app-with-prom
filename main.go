package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func oddOrEven(n int) string {
	if n%2 == 0 {
		evenNumbersProcessed.Inc()
		return "even"
	}

	oddNumbersProcessed.Inc()
	return "odd"
}

var (
	oddNumbersProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "odd_nums_processed_ops_total",
		Help: "The total number of odd nums processed events",
	})

	evenNumbersProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "even_nums_processed_ops_total",
		Help: "The total number of even nums processed events",
	})
)

func main() {

	for i := 1; i <= 100; i++ {
		oddOrEven(i)
	}

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}

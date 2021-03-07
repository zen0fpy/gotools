package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"math"
	"net/http"
	"time"
)

var (
	cpuTemp = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current Temperature of the CPU.",
	})

	hdFailures = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "hd_errors_total",
		Help: "Number of hard-disk errors.",
	},
		[]string{"device"},
	)

	reqCounter = prometheus.NewCounter(prometheus.CounterOpts{
		Name: "requests_total",
		Help: "The total number of requests served.",
	})

	httpReqs = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "http_requests_total",
		Help: "How many HTTP requests processed, paritioned by status code and HTTP method.",
	},
		[]string{"code", "method"},
	)

	tempHistogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:    "pond_temperature_celsius",
		Help:    "The temperature of the frog pond.",
		Buckets: prometheus.LinearBuckets(20, 5, 5),
	})
	tempSummary = prometheus.NewSummary(prometheus.SummaryOpts{
		Name:       "fenshan_temperature_celsius",
		Help:       "The temperature of fenshan.",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	})

	tempSummaryVec = prometheus.NewSummaryVec(prometheus.SummaryOpts{
		Name:       "fengshan_temperature_celsius",
		Help:       "The temperature of fengshang",
		Objectives: map[float64]float64{0.5: 0.05, 0.9: 0.01, 0.99: 0.001},
	}, []string{"species"})
)

func init() {
	prometheus.MustRegister(cpuTemp)
	prometheus.MustRegister(hdFailures)
	prometheus.MustRegister(httpReqs)
	prometheus.MustRegister(tempHistogram)
	prometheus.MustRegister(tempSummary)
	prometheus.MustRegister(tempSummaryVec)
	if err := prometheus.Register(reqCounter); err != nil {
		if are, ok := err.(prometheus.AlreadyRegisteredError); ok {
			reqCounter = are.ExistingCollector.(prometheus.Counter)
		} else {
			panic(err)
		}
	}
}

func main() {

	cpuTemp.Set(65.3)
	hdFailures.With(prometheus.Labels{"device": "/dev/sda"}).Inc()

	httpReqs.WithLabelValues("404", "POST").Inc()
	m := httpReqs.WithLabelValues("200", "GET")
	for i := 0; i < 1000000; i++ {
		m.Inc()
	}

	for i := 0; i < 1000; i++ {
		tempSummary.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)

	}

	for i := 0; i < 1000; i++ {
		tempSummaryVec.WithLabelValues("litoria-caerulea").Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
		tempSummaryVec.WithLabelValues("lithobates-catesbeianus").Observe(32 + math.Floor(100*math.Cos(float64(i)*0.11))/10)
	}

	//httpReqs.DeleteLabelValues("200", "GET")
	//httpReqs.Delete(prometheus.Labels{"method": "GET", "code": "200"})

	go func() {
		for {
			time.Sleep(5 * time.Second)
			for i := 0; i < 10000; i++ {
				tempHistogram.Observe(30 + math.Floor(120*math.Sin(float64(i)*0.1))/10)
			}
		}
	}()

	go func() {
		for {
			time.Sleep(2 * time.Second)
			reqCounter.Inc()
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9091", nil))

}

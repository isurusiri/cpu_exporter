package main

import (
	"fmt"
	"net/http"

	"github.com/isurusiri/cpu_exporter/cpustatcollector"
	"github.com/namsral/flag"
	"github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var log = logrus.New()

type config struct {
	port    uint
	verbose bool
}

func readAndValidateConfig() config {
	var result config

	flag.UintVar(&result.port, "port", 9581, "Port to expose scraping endpoint on")
	flag.BoolVar(&result.verbose, "verbose", false, "Enable verbose logging")
	flag.Parse()

	log.WithFields(logrus.Fields{
		"port":    result.port,
		"verbose": result.verbose,
	}).Infof("CPU Metrics exporter has been configured")

	return result
}

func configureRoutes() *http.ServeMux {
	var landingPage = []byte(`<html>
	<head><title>Azure ServiceBus exporter for Prometheus</title></head>
	<body>
	<h1>Azure ServiceBus exporter for Prometheus</h1>
	<p><a href='/metrics'>Metrics</a></p>
	</body>
	</html>
	`)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write(landingPage) // nolint: errcheck
	})
	mux.HandleFunc("/metrics", promhttp.Handler().ServeHTTP)

	return mux
}

func setupLogger(config config) {
	if config.verbose {
		log.Level = logrus.DebugLevel
	}
}

func startHTTPServer(config config, mux *http.ServeMux, collector *cpustatcollector.CPUCollector) {
	listenAddr := fmt.Sprintf(":%d", config.port)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
	log.Fatal(http.ListenAndServe(listenAddr, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "metrics" {
			collector.Collect()
		}
	})))
}

func main() {
	config := readAndValidateConfig()
	setupLogger(config)

	collector := cpustatcollector.New()
	collector.Init()

	mux := configureRoutes()
	startHTTPServer(config, mux, collector)
}

package main

import (
	"log"

	"github.com/mehmetolgundev/splitter/pkg/listener"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	memoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "splitter_memory_usage",
			Help: "Current memory usage of the application",
		},
		[]string{"instance"},
	)
)

func init() {
	prometheus.MustRegister(memoryUsage)
}
func main() {

	err := listener.NewTCPListener().Listen()
	if err != nil {
		log.Fatal(err)
	}
}

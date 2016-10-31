package main

import (
	"math/rand"

	"github.com/prometheus/client_golang/prometheus"
)

type randomMetrics struct {
	RandomNumbers float64
}

var (
	connections = prometheus.NewGauge(prometheus.GaugeOpts{
		Namespace: "random",
		Name:      "numbers",
		Help:      "Random collection",
	})
)

type Exporter struct {
}

// NewExporter instantiates a new NATS Exporter.
func NewExporter() *Exporter {
	return &Exporter{
	}
}

func (e *Exporter) Describe(ch chan<- *prometheus.Desc) {
	connections.Describe(ch)
}

func (e *Exporter) Collect(ch chan<- prometheus.Metric) {
	if err := e.collect(); err != nil {
		return
	}

	connections.Collect(ch)
}

func (e *Exporter) collect() error {
	var metrics randomMetrics

	metrics.RandomNumbers = rand.Float64()*5
	connections.Set(metrics.RandomNumbers)

	return nil
}

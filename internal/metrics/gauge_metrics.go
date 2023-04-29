package metrics

import "github.com/prometheus/client_golang/prometheus"

var RequestGauge = prometheus.NewGauge(prometheus.GaugeOpts{
	Namespace:   "",
	Subsystem:   "",
	Name:        "requests_count",
	Help:        "requests done since app started",
	ConstLabels: nil,
})

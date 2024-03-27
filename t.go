package prometheus

import "github.com/prometheus/client_golang/prometheus"

func foo() {
	_ = newStats(nil)
	_ = prometheus.NewGauge(nil)
}

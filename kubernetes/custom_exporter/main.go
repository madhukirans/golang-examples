
package main

import (
"net/http"
"time"

"github.com/prometheus/client_golang/prometheus"
"github.com/prometheus/client_golang/prometheus/promauto"
"github.com/prometheus/client_golang/prometheus/promhttp"
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			bridges_num_ports.Collect()
			cpuTemp.Set(1)
			time.Sleep(2 * time.Second)
		}
	}()
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})

	bridges_num_ports = promauto.NewGaugeVec(prometheus.GaugeOpts{
		//Namespace: namespace,
		Name:      "bridges_ports",
		Help:      "Number of ports attached to bridges",
	},
		[]string{"name"},
	)

	cpuTemp = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "cpu_temperature_celsius",
		Help: "Current temperature of the CPU.",
	},

	)
)

func main() {

	//cpuTemp.Set(1)
	recordMetrics()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":2112", nil)
}
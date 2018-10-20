package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"fmt"
	"github.com/prometheus/client_golang/prometheus/push"
)

func main() {
	gatewayUrl := "https://prometheus-gw.dev.jksc.sauron.us-phoenix-1.oracledx.com/metrics/job/k8s_cl1_federate_prometheus-metrics/instance"

	throughputGuage := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "throughput_temp",
		Help: "Throughput in Mbps",
	})
	throughputGuage.Set(800)

	err := push.New(gatewayUrl, "jksc-prometheus-pusher-955764f8b-r8gnr").BasicAuth("jksc","8QKRo7zv8P").Collector(throughputGuage).Push()


	//err := push.Collectors("throughput_job", push.HostnameGroupingKey(), gatewayUrl, throughputGuage)

	if err != nil {
		fmt.Println("Could not push completion time to Pushgateway:", err)
	}
}

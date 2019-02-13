
package main

import (
"net/http"

"github.com/prometheus/client_golang/prometheus"
//"github.com/prometheus/client_golang/prometheus/promauto"
"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/gorilla/mux"
	"fmt"
	"github.com/golang/glog"
	"time"
	"github.com/parnurzeal/gorequest"
	"github.com/fatih/color"
)

//func recordMetrics() {
//	go func() {
//		for {
//			opsProcessed.Inc()
//			bridges_num_ports.Collect()
//			cpuTemp.Set(1)
//			time.Sleep(2 * time.Second)
//		}
//	}()
//}
//
//var (
//	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
//		Name: "myapp_processed_ops_total",
//		Help: "The total number of processed events",
//	})
//
//	bridges_num_ports = promauto.NewGaugeVec(prometheus.GaugeOpts{
//		//Namespace: namespace,
//		Name:      "bridges_ports",
//		Help:      "Number of ports attached to bridges",
//	},
//		[]string{"name"},
//	)
//
//	cpuTemp = promauto.NewGauge(prometheus.GaugeOpts{
//		Name: "cpu_temperature_celsius",
//		Help: "Current temperature of the CPU.",
//	},
//
//	)
//)

func main() {

	//cpuTemp.Set(1)
//	recordMetrics()

	//http.Handle("/metrics", promhttp.Handler())
	//http.ListenAndServe(":2112", nil)

	go startServerandRegisterMetric()
	time.Sleep(2 * time.Second)
	addBackupFailureMetrics(1, "backupjob", "sauron-1", "data")
	addBackupFailureMetrics(1, "backupjob1", "sauron-1", "data")
	addBackupFailureMetrics(1, "backupjob2", "sauron-1", "data")
	addBackupFailureMetrics(1, "backupjob3", "sauron-1", "data")
	time.Sleep(10 * time.Second)
	backupMandos.Delete(prometheus.Labels{"job_name": "backupjob", "env_name": "sauron-1", "backup_type": "data"})
	time.Sleep(6000 * time.Second)


}

var backupMandos = prometheus.NewGaugeVec(
	prometheus.GaugeOpts{
		Namespace: "mandos",
		Name:      "backup",
		Help:      "value 0 means backup passed and value 1 means backup failed and value 2 means backup is still running ",
	},
	[]string{"job_name", "env_name", "backup_type"},
)

func func1 (){

		url := "http://localhost:9090/api/v1/series?match[]=" + monitor.Metric.String()
		fmt.Errorf("Prometheus get metric URL:>%s", url)
		request := gorequest.New()
		resp, body, errs := request.Get(url).End()
		if errs != nil {
			fmt.Println("Connecting prometheus error: ", errs)
		} else {
			color.Red(" %s Response Status: %d %s %s", monitor.Metric.EnvName, resp.Status, resp.Header, body)
		}
	}
}

func startServerandRegisterMetric() {
	port := 2112
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/federate", promhttp.Handler())

	// Register metrics
	prometheus.MustRegister(backupMandos)

	//Start the server
	http.ListenAndServe(fmt.Sprintf(":%d", port), router)
	glog.Fatalln("Quit unexpectedly: reason unknown")
}

func addBackupFailureMetrics(backupFailed int, backupJobLabel string, sauronName string, backupType string) {
	fmt.Println("111")
//	prometheus.MustRegister(backupMandos)
	// Convert it to float64 as Add method accepts float64
	i := float64(backupFailed)
	backupMandos.With(prometheus.Labels{"job_name": backupJobLabel, "env_name": sauronName, "backup_type": backupType}).Set(i)


}
package analytics

import "github.com/prometheus/client_golang/prometheus"

var (
	Anomalys = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "total_anomalys",
			Help: "Total number of anomalys",
		},
	)
)

func Init() {
	prometheus.MustRegister(Anomalys)
}

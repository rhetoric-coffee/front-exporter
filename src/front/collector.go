package front

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"time"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	up = prometheus.NewDesc(
		"front_up",
		"talking to front successfully",
		nil, nil,
	)
	client = &http.Client{
		Timeout: time.Second * 10,
	}
)

type FrontCollector struct{}

func (f FrontCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- up
}

func (f FrontCollector) Collect(ch chan<- prometheus.Metric) {
	front, err := New()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
	}
	teams, err := front.ListTeams()
	if err != nil {
		ch <- prometheus.MustNewConstMetric(up, prometheus.GaugeValue, 0)
		return
	}
	for _, t := range *teams {
		log.Println(t)
	}

}

func main() {
	log.Println("Starting the Front exporter...")
	f := FrontCollector{}
	prometheus.MustRegister(f)
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8000", nil))
}

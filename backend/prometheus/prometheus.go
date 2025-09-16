package prometheus

import (
	"statistics/statistics"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	visitors = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "statistics_traffic",
		Help: "Number of traffic",
	})
	active_users = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "statistics_active_users",
		Help: "Number of currently active users",
	})
	minuteSpentGauge = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "statistics_minute_spent",
		Help: "Number of minute spent",
	})
)

func RecordMetrics() {
	go func() {
		for {
			visitors.Set(float64(statistics.GetUsers(time.Now().Add(-5*time.Minute), time.Now(), "")))
			active_users.Set(float64(statistics.ActiveUsers("")))
			minuteSpentGauge.Set(float64(statistics.TimeOnSite("", time.Now().Add(-24*time.Hour), time.Now())))
			time.Sleep(2 * time.Second)
		}
	}()
}

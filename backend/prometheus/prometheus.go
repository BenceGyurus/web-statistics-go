package prometheus

import (
	"statistics/database"
	"statistics/statistics"
	"statistics/structs"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	visitorsBySite = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "statistics_traffic",
		Help: "Unique sessions (traffic) in last 5 minutes by site",
	}, []string{"site"})
	activeUsersBySite = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "statistics_active_users",
		Help: "Number of currently active users by site (last 5 minutes)",
	}, []string{"site"})
	minuteSpentBySite = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "statistics_minute_spent",
		Help: "Average minutes spent on site in last 24h by site",
	}, []string{"site"})
)

func RecordMetrics() {
	go func() {
		for {
			// Fetch distinct sites
			var sites []string
			_ = database.Session.Model(&structs.WebMetric{}).Distinct("site").Pluck("site", &sites).Error

			// Update metrics per site
			for _, site := range sites {
				if site == "" {
					continue
				}
				visitorsBySite.With(prometheus.Labels{"site": site}).Set(
					float64(statistics.GetUsers(time.Now().Add(-24*time.Minute), time.Now(), site)),
				)
				activeUsersBySite.With(prometheus.Labels{"site": site}).Set(
					float64(statistics.ActiveUsers(site)),
				)
				minuteSpentBySite.With(prometheus.Labels{"site": site}).Set(
					float64(statistics.TimeOnSite(site, time.Now().Add(-24*time.Hour), time.Now())),
				)
			}
			time.Sleep(2 * time.Second)
		}
	}()
}

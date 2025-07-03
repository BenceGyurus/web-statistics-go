package statistics

import (
	"statistics/database"
	"time"
)

func GetUsers(t1 time.Time, t2 time.Time, site string) int {
	var results int

	if site == "" {
		query := `SELECT COUNT (*) from (SELECT session_id FROM "web_metrics" WHERE "timestamp" >= ? AND "timestamp" <= ? GROUP BY session_id) as lamdba;`
		database.Session.Raw(query, t1, t2).Scan(&results)
	}
	if site != "" {
		query := `SELECT COUNT (*) from (SELECT session_id FROM "web_metrics" WHERE "timestamp" >= ? AND "timestamp" <= ? AND site = ? GROUP BY session_id) as lamdba;`
		database.Session.Raw(query, t1, t2, site).Scan(&results)
	}

	return results
}

type Result struct {
	URL            string `json:"url"`
	TotalTimeSpent int64  `json:"totalTimeSpent"`
}

func GetSpentTimeByPages(t1 time.Time, t2 time.Time) []Result {
	query := `WITH with_next AS (
            SELECT
                session_id,
                url,
                time,
                LEAD(time) OVER (PARTITION BY session_id ORDER BY time) AS next_time
            FROM traffic
        ),
        with_diffs AS (
            SELECT
                url,
                EXTRACT(EPOCH FROM (next_time - time)) * 1000 AS diff_ms
            FROM with_next
            WHERE next_time IS NOT NULL
        ),
        with_capped AS (
            SELECT
                url,
                LEAST(diff_ms, 300000) AS time_spent
            FROM with_diffs
        )
        SELECT
            url,
            SUM(time_spent)::BIGINT AS total_time_spent
        FROM with_capped
        GROUP BY url;`

	var results []Result
	database.Session.Raw(query, t1, t2).Scan(&results)

	return results
}

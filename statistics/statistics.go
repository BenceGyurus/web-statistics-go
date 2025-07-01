package statistics

import (
	"fmt"
	"statistics/database"
)

func GetUsers() { //from time.Time, to time.Time
	iter := database.Session.Query("SELECT * FROM statistics.traffic").Iter()
	data := iter.Scan(timestamp, page, &t.sessionId)
	fmt.Println("Data:", data)
}

package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"statistics/database"
	"statistics/statistics"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func userTraffic(c *gin.Context) {
	sessionId := c.Query("sessionId")
	if sessionId == "" {
		sessionId = uuid.New().String()
	}
	fmt.Println("Session ID:", sessionId)
	err := database.Session.Query(`
        INSERT INTO statistics.traffic (id, sessionId, timestamp, path, site) 
        VALUES (uuid(), ?, toTimestamp(now()), ?, ?)`,
		sessionId, c.Query("p"), c.Query("site"),
	).Exec()
	if err != nil {
		log.Println("Error inserting traffic data:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, sessionId)
}

func traffic(c *gin.Context) {
	/*from := c.Query("from")
	to := c.Query("to")
	if from == "" || to == "" {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	fromTime, err := time.Parse(time.RFC3339, from)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}
	toTime, err := time.Parse(time.RFC3339, to)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}*/

	statistics.GetUsers()
	c.String(http.StatusOK, "Traffic data processed successfully")
}

func Server() {
	router := gin.Default()
	port := os.Getenv("PORT")
	prefix := os.Getenv("PREFIX")
	if port == "" {
		port = "8080"
	}

	router.GET(prefix+"/put-traffic", userTraffic)

	router.POST(prefix + "/login")

	router.POST(prefix+"/traffic", traffic)

	log.Println("prefix", prefix)
	err := router.Run("localhost:" + port)
	if (err) == nil {
		log.Println("Failed to start server", "error", err)
		panic(err)
	}
}

package server

import (
	"log"
	"net/http"
	"os"
	"statistics/database"
	"statistics/statistics"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func userTraffic(c *gin.Context) {
	sessionId := c.Query("sessionId")
	if sessionId == "" {
		sessionId = uuid.New().String()
	}

	ip := c.Request.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = c.ClientIP()
	}

	record := database.WebMetric{
		SessionId: sessionId,
		Timestamp: time.Now(),
		Page:      c.Query("p"),
		Site:      c.Query("site"),
		Ip:        ip,
	}
	err := database.Session.Create(&record).Error
	if err != nil {
		log.Println("Error inserting traffic data:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.String(http.StatusOK, sessionId)
}

func traffic(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	page := c.Query("page")
	var fromTime, toTime time.Time
	var err error
	layout := "2006-01-02"
	if !(from == "" || to == "") {
		fromTime, err = time.Parse(layout, from)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		toTime, err = time.Parse(layout, to)
		if err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		fromTime = time.Now().Add(-24 * time.Hour)
		toTime = time.Now()
	}
	numberOfUsers := statistics.GetUsers(fromTime, toTime, page)
	c.JSON(http.StatusOK, gin.H{"traffic": numberOfUsers})
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

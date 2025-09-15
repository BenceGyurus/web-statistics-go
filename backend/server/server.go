package server

import (
	"log"
	"net/http"
	"os"
	"statistics/database"
	"statistics/statistics"
	"statistics/structs"
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
	record := structs.WebMetric{
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

func getSites(c *gin.Context) {
	var sites []string
	err := database.Session.Model(&structs.WebMetric{}).Distinct("site").Pluck("site", &sites).Error
	if err != nil {
		log.Println("Error fetching distinct sites:", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.JSON(http.StatusOK, gin.H{"sites": sites})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://irodalomerettsegi.hu")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Server() {
	router := gin.Default()
	port := os.Getenv("BACKEND_PORT")
	prefix := os.Getenv("PREFIX")
	if port == "" {
		port = "8080"
	}

	router.Use(CORSMiddleware())

	router.GET(prefix+"/put-traffic", userTraffic)

	router.POST(prefix+"/traffic", traffic)

	router.POST(prefix+"/sites", statistics.GetUsersByPages)

	router.POST(prefix+"/graph", statistics.GetTrafficStats)

	router.POST(prefix+"/active", statistics.GetActiveUsers)

	router.POST(prefix+"/time", statistics.GetTimeOnTheSite)

	router.POST(prefix+"/get-sites", getSites)

	// Health check endpoint
	router.GET(prefix+"/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "healthy"})
	})

	log.Println("prefix", prefix)
	log.Print("Starting server on port " + port)
	err := router.Run("0.0.0.0:" + port)
	if (err) == nil {
		log.Println("Failed to start server", "error", err)
		panic(err)
	}
}

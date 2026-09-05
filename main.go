package main

import (
	"math/rand/v2"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"go-prometheus-monitoring/middleware/metrics"
)

func main() {
	router := gin.Default()

	// Register /metrics before middleware
	router.GET("/metrics", PrometheusHandler())

	router.Use(metrics.RequestMetricsMiddleware())
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Up and running!",
		})
	})
	router.GET("/v1/users", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello from /v1/users",
		})
	})
	router.GET("/v1/random", RandomHandler)

	router.Run(":8000")
}

// Custom metrics handler with custom registry
func PrometheusHandler() gin.HandlerFunc {
	h := promhttp.HandlerFor(metrics.CustomRegistry, promhttp.HandlerOpts{})
	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func randomDelay(min, max int) int {
	return rand.IntN(max-min+1) + min
}

// function that will intentionally delay to simulate duration variation
func RandomHandler(c *gin.Context) {
	const (
		minDelay = 90
		maxDelay = 500
	)

	delayMs := rand.IntN(maxDelay-minDelay+1) + minDelay
	time.Sleep(time.Duration(delayMs) * time.Millisecond)

	// 80% success, 20% failure
	if rand.IntN(100) >= 80 {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":   "response delayed",
		"waited_ms": delayMs,
	})
}

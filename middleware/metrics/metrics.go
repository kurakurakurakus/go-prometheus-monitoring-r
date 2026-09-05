package metrics

import (
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

const cServiceName = "go-prometheus-monitoring-r"

// Custom registry (without default Go metrics)
var CustomRegistry = prometheus.NewRegistry()

// Define metrics
var (
	HttpRequestTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
		Name: "api_http_request_total",
		Help: "Total number of requests processed by the API",
	}, []string{"service", "path", "status"})

	// HttpRequestErrorTotal = prometheus.NewCounterVec(prometheus.CounterOpts{
	// 	Name: "api_http_request_error_total",
	// 	Help: "Total number of errors returned by the API",
	// }, []string{"path", "status"})

	HttpRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "api_http_request_duration_seconds",
			Help: "HTTP request duration in seconds",
		},
		[]string{"service", "path"},
	)
)

// Register metrics with custom registry
func init() {
	CustomRegistry.MustRegister(
		HttpRequestTotal,
		HttpRequestDuration,
	)
}

// Middleware to record incoming request metrics
func RequestMetricsMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		path := c.FullPath()
		status := c.Writer.Status()

		HttpRequestTotal.
			WithLabelValues(cServiceName, path, strconv.Itoa(status)).
			Inc()

		HttpRequestDuration.
			WithLabelValues(cServiceName, path).
			Observe(time.Since(start).Seconds())
	}
}

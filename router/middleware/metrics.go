package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"strconv"
	"strings"
	"time"
)

var httpRequestCount = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "http_request_count",
		Help: "http request count",
	},
	[]string{"endpoint", "code"},
)

var httpRequestDuration = prometheus.NewSummaryVec(
	prometheus.SummaryOpts{
		Name: "http_request_duration",
		Help: "http request duration",
	},
	[]string{"endpoint"},
)

func init() {
	prometheus.MustRegister(httpRequestCount)
	prometheus.MustRegister(httpRequestDuration)
}

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		status := c.Writer.Status()
		split := strings.Split(c.Request.RequestURI, "?")
		uri := split[0]
		httpRequestCount.WithLabelValues(uri, strconv.Itoa(status)).Inc()
		elapsed := (float64)(time.Since(start) / time.Millisecond)
		httpRequestDuration.WithLabelValues(uri).Observe(elapsed)
	}
}

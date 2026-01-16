package middleware

import (
	"go-admin/pkg/log"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Logger() gin.HandlerFunc {
	logger := log.Log()
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime) / time.Millisecond
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		header := c.Request.Header
		proto := c.Request.Proto
		statusCode := c.Writer.Status()
		clientIP := c.ClientIP()
		err := c.Err()
		body, _ := ioutil.ReadAll(c.Request.Body)
		logger.WithFields(logrus.Fields{
			"latency_time": latencyTime,
			"req_method":   reqMethod,
			"req_uri":      reqUri,
			"header":       header,
			"client_ip":    clientIP,
			"status_code":  statusCode,
			"proto":        proto,
			"err":          err,
			"body":         body,
		}).Info()
	}
}

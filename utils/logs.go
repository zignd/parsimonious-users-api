package utils

import (
	"bytes"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
)

func SetupLogrus() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
}

func GinJSONLogger() gin.HandlerFunc {
	l := log.New()
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		var buf bytes.Buffer
		l.Out = &buf
		l.SetFormatter(&logrus.JSONFormatter{})
		l.WithFields(log.Fields{
			"clientIP":     param.ClientIP,
			"method":       param.Method,
			"path":         param.Path,
			"protocol":     param.Request.Proto,
			"statusCode":   param.StatusCode,
			"latency":      param.Latency,
			"userAgent":    param.Request.UserAgent(),
			"errorMessage": param.ErrorMessage,
		}).Infof("%s %s (%d)", param.Method, param.Path, param.StatusCode)
		return buf.String()
	})
}

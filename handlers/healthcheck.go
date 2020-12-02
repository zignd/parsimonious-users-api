package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"github.com/zignd/parsimonious-users-api/healthcheck"
)

func GetHealthCheck(db *pg.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		isHealthy, err := healthcheck.Check(db)
		if err != nil {
			log.Error(fmt.Errorf("health check error: %w", err))
		}

		c.JSON(200, &gin.H{
			"isHealthy": isHealthy,
		})
	}
}

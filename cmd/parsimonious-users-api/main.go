package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"github.com/zignd/parsimonious-users-api/envs"
	"github.com/zignd/parsimonious-users-api/handlers"
	"github.com/zignd/parsimonious-users-api/utils"
)

func init() {
	if err := utils.ValidateEnvs(); err != nil {
		panic(fmt.Errorf("environment variables not properly set: %w", err))
	}
	utils.SetupLogrus()
	gin.SetMode(gin.ReleaseMode)
}

func main() {
	log.Info("Starting")
	db := utils.CreateDBConn(true)
	defer db.Close()

	r := gin.New()
	r.Use(utils.GinJSONLogger())
	r.GET("/health-check", handlers.GetHealthCheck(db))
	r.GET("/users/name/:name", handlers.GetUsersByName(db))
	r.GET("/users/username/:username", handlers.GetUsersByUsername(db))

	err := r.Run(fmt.Sprintf(":%s", os.Getenv(envs.HTTPServerPort)))
	if err != nil {
		panic(fmt.Errorf("application stopped unexpectedly: %w", err))
	}
}

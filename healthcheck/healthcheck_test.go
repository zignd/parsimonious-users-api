package healthcheck

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/require"
	"github.com/zignd/parsimonious-users-api/envs"
	"github.com/zignd/parsimonious-users-api/utils"
)

func TestCheckIsHealthy(t *testing.T) {
	err := utils.ValidateEnvs()
	require.Equal(t, err, nil, "environment variables should have been set for the test")

	db := utils.CreateDBConn(false)
	defer db.Close()

	isHealthy, err := Check(db)
	require.Equal(t, err, nil, "connection to the database should be okay")
	require.Equal(t, isHealthy, true)
}

func TestCheckIsNotHealthy(t *testing.T) {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(envs.PostgresAddr),
		User:     "invalid",
		Password: "invalid",
		Database: os.Getenv(envs.PostgresDatabase),
	})
	defer db.Close()

	isHealthy, err := Check(db)
	require.Equal(t, err.Error(), "failed to check if the database is accessible: FATAL #28P01 password authentication failed for user \"invalid\"")
	require.Equal(t, isHealthy, false)
}

package users

import (
	"os"
	"testing"

	"github.com/go-pg/pg/v10"
	"github.com/stretchr/testify/require"
	"github.com/zignd/parsimonious-users-api/envs"
	"github.com/zignd/parsimonious-users-api/models"
	"github.com/zignd/parsimonious-users-api/utils"
)

func TestGetUserByNameWithResults(t *testing.T) {
	db := utils.CreateDBConn(false)
	defer db.Close()

	usrs, err := GetUserByName(db, "Ta", 1, 2)
	require.Nil(t, err, "should not return an error")
	require.Len(t, usrs, 2, "length differs from the specified pageSize")
	require.Equal(t, usrs[0], &models.User{
		ID:       "092b7c45-cafd-4242-96e1-373c453961d8",
		Name:     "Taliana Weigert",
		Username: "taliana.weigert",
	})
	require.Equal(t, usrs[1], &models.User{
		ID:       "32caa8a3-6085-48d5-9eb5-39953de07b82",
		Name:     "Tatiane Eustaquio",
		Username: "tatiane.eustaquio",
	})
}

func TestGetUserByNameWithNoResults(t *testing.T) {
	db := utils.CreateDBConn(false)
	defer db.Close()

	usrs, err := GetUserByName(db, "Innexistent Name", 1, 5)
	require.Nil(t, err, "should not return an error")
	require.Len(t, usrs, 0, "should not return users for an innexistent name")
}

func TestGetUserByUsernameWithResults(t *testing.T) {
	db := utils.CreateDBConn(false)
	defer db.Close()

	usrs, err := GetUserByUsername(db, "ta", 1, 2)
	require.Nil(t, err, "should not return an error")
	require.Len(t, usrs, 2, "length differs from the specified pageSize")
	require.Equal(t, usrs[0], &models.User{
		ID:       "092b7c45-cafd-4242-96e1-373c453961d8",
		Name:     "Taliana Weigert",
		Username: "taliana.weigert",
	})
	require.Equal(t, usrs[1], &models.User{
		ID:       "32caa8a3-6085-48d5-9eb5-39953de07b82",
		Name:     "Tatiane Eustaquio",
		Username: "tatiane.eustaquio",
	})
}

func TestGetUserByUsernameWithNoResults(t *testing.T) {
	db := utils.CreateDBConn(false)
	defer db.Close()

	usrs, err := GetUserByUsername(db, "innexistent.username", 1, 5)
	require.Nil(t, err, "should not return an error")
	require.Len(t, usrs, 0, "should not return users for an innexistent username")
}

func TestGetUserByNameWithError(t *testing.T) {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(envs.PostgresAddr),
		User:     "invalid",
		Password: "invalid",
		Database: os.Getenv(envs.PostgresDatabase),
	})
	defer db.Close()

	usrs, err := GetUserByName(db, "", 1, 5)
	require.Equal(t, err.Error(), "failed to get users by name: failed to query for users: FATAL #28P01 password authentication failed for user \"invalid\"")
	require.Nil(t, usrs)
}

func TestGetUserByUsernameWithError(t *testing.T) {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(envs.PostgresAddr),
		User:     "invalid",
		Password: "invalid",
		Database: os.Getenv(envs.PostgresDatabase),
	})
	defer db.Close()

	usrs, err := GetUserByUsername(db, "", 1, 5)
	require.Equal(t, err.Error(), "failed to get users by username: failed to query for users: FATAL #28P01 password authentication failed for user \"invalid\"")
	require.Nil(t, usrs)
}

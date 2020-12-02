package utils

import (
	"context"
	"os"

	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"github.com/zignd/parsimonious-users-api/envs"
)

type DBLogger struct{}

func (d DBLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d DBLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	b, _ := q.FormattedQuery()
	log.Debug("Query executed: %s", b)
	return nil
}

func CreateDBConn(withQueryLogger bool) *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     os.Getenv(envs.PostgresAddr),
		User:     os.Getenv(envs.PostgresUser),
		Password: os.Getenv(envs.PostgresPassword),
		Database: os.Getenv(envs.PostgresDatabase),
	})
	if withQueryLogger {
		db.AddQueryHook(DBLogger{})
	}
	return db
}

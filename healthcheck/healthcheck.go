package healthcheck

import (
	"fmt"

	"github.com/go-pg/pg/v10"
)

type DBVersion struct {
	Version string
}

func Check(db *pg.DB) (bool, error) {
	var v DBVersion
	_, err := db.QueryOne(&v, `SELECT VERSION() as version`)
	if err != nil {
		return false, fmt.Errorf("failed to check if the database is accessible: %w", err)
	}
	return true, nil
}

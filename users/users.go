package users

import (
	"fmt"
	"strings"

	"github.com/go-pg/pg/v10"
	"github.com/zignd/parsimonious-users-api/models"
)

func getUsersBy(db *pg.DB, column, value string, page, pageSize int) ([]*models.User, error) {
	value = strings.Trim(value, " ")

	whereClause := fmt.Sprintf("\"user\".%s LIKE ?", column)
	whereArg := fmt.Sprintf("%%%s%%", value)

	users := make([]*models.User, 0)
	err := db.Model(&users).
		Join("LEFT JOIN relevance_1").
		JoinOn("relevance_1.id = \"user\".id").
		Join("LEFT JOIN relevance_2").
		JoinOn("relevance_2.id = \"user\".id").
		Where(whereClause, whereArg).
		Order("relevance_1.id").
		Order("relevance_2.id").
		Order("user.id").
		Offset((page - 1) * pageSize).
		Limit(pageSize).
		Select()
	if err != nil {
		return nil, fmt.Errorf("failed to query for users: %w", err)
	}

	return users, nil
}

func GetUserByName(db *pg.DB, name string, page, pageSize int) ([]*models.User, error) {
	users, err := getUsersBy(db, "name", name, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by name: %w", err)
	}
	return users, nil
}

func GetUserByUsername(db *pg.DB, username string, page, pageSize int) ([]*models.User, error) {
	users, err := getUsersBy(db, "username", username, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("failed to get users by username: %w", err)
	}
	return users, nil
}

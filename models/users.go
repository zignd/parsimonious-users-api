package models

type User struct {
	ID       string `pg:"type:uuid,pk"`
	Name     string `pg:",notnull"`
	Username string `pg:",notnull"`
}

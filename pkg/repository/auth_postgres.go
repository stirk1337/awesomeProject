package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/stirk1337/awesomeProject/pkg/user"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user user.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUserHashByUsername(username string) (user.User, error) {
	var usr user.User
	query := fmt.Sprintf("SELECT id, password_hash FROM %s WHERE username = $1", usersTable)
	err := r.db.Get(&usr, query, username)
	return usr, err
}

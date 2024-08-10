package repos

import (
	"fmt"
	"marketplace/internal/models"
)

/*
Ensures struct implements interface
*/
var _ models.UserRepo = (*UserRepo)(nil)

type UserRepo struct {
	db models.DB
}

func NewUserRepo(db models.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (u UserRepo) Create(user models.UserAttributes) error {
	statement := "INSERT INTO tcp_auth_db.user_tab(username, password_hash) VALUES ($1, $2)"
	stmt, err := u.db.Prepare(statement)
	if err != nil {
		return err
	}
	defer stmt.Close()
	fmt.Println(user)
	_, err = stmt.Exec(user.Username, user.PasswordHash)
	if err != nil {
		return err
	}
	return nil
}

func (u UserRepo) GetByUsername(username string) (*models.User, error) {
	statement := "SELECT * FROM tcp_auth_db.user_tab WHERE username = $1"
	stmt, err := u.db.Prepare(statement)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	var user models.User
	stmt.
}

func (u UserRepo) GetByID(userID uint64) (*models.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteByID(userID uint64) error {
	//TODO implement me
	panic("implement me")
}

func (u UserRepo) DeleteByUsername(username string) error {
	//TODO implement me
	panic("implement me")
}

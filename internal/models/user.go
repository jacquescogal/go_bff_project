package models

type User struct {
	UserID       string `json:"user_id"`
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserRepo interface {
	Create(user User) error
	ReadByUsername(username string) (User, error)
	ReadByID(userID string) (User, error)
	DeleteByID(userID string) error
	DeleteByUsername(username string) error
}

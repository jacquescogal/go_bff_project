package models

type User struct {
	UserID uint64 `json:"user_id"`
	UserAttributes
}

type UserAttributes struct {
	Username     string `json:"username"`
	PasswordHash string `json:"password_hash"`
}

type UserRepo interface {
	Create(user UserAttributes) error
	GetByUsername(username string) (*User, error)
	GetByID(userID uint64) (*User, error)
	DeleteByID(userID uint64) error
	DeleteByUsername(username string) error
}

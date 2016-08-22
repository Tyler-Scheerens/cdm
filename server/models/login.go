package models

type Login struct {
	UserId   int64  `json:"user_id" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"-" db:"password"`
	Salt     string `json:"-" db:"salt"`
  Email    string `json:"email" db:"email"`
}

package models

// User is
type User struct {
	ID       int64
	Name     string
	Role     string
	Username string
	Password []byte
}

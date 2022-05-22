package db

// User in DB
type User struct {
	ID    int    `json:"id" sql:"id"`
	Name  string `json:"name" sql:"name"`
	Email string `json:"email" sql:"email"`
}

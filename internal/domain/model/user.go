package model

type UserID int

type User struct {
	ID    UserID `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func NewUserFromRepository(id UserID, name string, email string) *User {
	return &User{
		ID:    id,
		Name:  name,
		Email: email,
	}
}

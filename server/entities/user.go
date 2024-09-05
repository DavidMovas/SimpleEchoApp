package entities

type User struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func NewUser(name string, age int, email string) *User {
	return &User{
		Name:  name,
		Age:   age,
		Email: email,
	}
}

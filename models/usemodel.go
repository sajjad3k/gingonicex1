package models

type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
	City  string `json:"city"`
}

//var rec []user

func (u *User) Tablename() string {
	return "user"
}

package models

import real "github.com/brianvoe/gofakeit/v6"

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Gender    string `json:"gender"`
}


func (u *User) NewUser() *User{
	return &User{
		FirstName: real.FirstName(),
		LastName: real.LastName(),
		Email: real.Email(),
		Phone: real.PhoneFormatted(),
		Gender: real.Gender(),
	}
}

func (u *User) GenUsers(amount int) []*User{
	var users []*User
	for i:=1; i <=amount; i++{
		users = append(users,u.NewUser())
	}
	return users
}


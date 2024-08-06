package models

type User struct {
	Vorname  string
	Nachname string
	Email    string
}

func (user User) NewUser(vorname, nachname, email string) User {
	return User{Vorname: vorname, Nachname: nachname, Email: email}
}

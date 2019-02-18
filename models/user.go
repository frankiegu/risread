package models

type User struct {
	Id       string
	Username string
	Password string
	Profile  Profile
}

type Profile struct {
	Gender  string
	Age     int
	Address string
	Email   string
}

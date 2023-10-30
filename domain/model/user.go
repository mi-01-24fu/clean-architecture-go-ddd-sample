package model

type User struct {
	ID          int
	UserName    UserName
	MailAddress MailAddress
	Password    Password
}

func NewSignupUser(userName UserName, mailAddress MailAddress, password Password) *User {
	return &User{UserName: userName, MailAddress: mailAddress, Password: password}
}

func NewUser(ID int, userName UserName, mailAddress MailAddress) *User {
	return &User{ID: ID, UserName: userName, MailAddress: mailAddress}
}

package model

type UserFactorier interface {
	Create(id int, userName, email string) (*User, error)
}

type UserFactory struct{}

func NewUserFactory() *UserFactory {
	return &UserFactory{}
}

func (u UserFactory) Create(id int, userName, email string) (*User, error) {
	name, err := NewUserName(userName)
	if err != nil {
		return nil, err
	}

	mailAddress, err := NewMailAddress(email)
	if err != nil {
		return nil, err
	}

	return NewUser(id, *name, *mailAddress), err

}

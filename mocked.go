package main

import "context"

type mockedDatabase struct{}

var (
	HomeStore  []Home
	UsersStore []User
)

func ConfigureMockedDB() database {
	return mockedDatabase{}
}

func (m mockedDatabase) GetUsers(ctx context.Context) ([]User, error) {
	return UsersStore, nil
}

func (m mockedDatabase) NewUser(input User) (User, error) {
	return input, nil
}

func (m mockedDatabase) NewHome(input Home) (Home, error) {
	return input, nil
}

package internal

import (
	"fmt"
)

type User struct {
	Id   int
	Name string
}

type UserManagement struct {
	data map[int]*User
}

func NewUserManagement() *UserManagement {
	return &UserManagement{
		data: make(map[int]*User),
	}
}

func (um *UserManagement) AddUser(id int, name string) {
	if _, exists := um.data[id]; exists {
		fmt.Printf("User with Id %d Already registered\n \n", id)
		return
	}
	um.data[id] = &User{Id: id, Name: name}
	fmt.Printf("User successfully added:\nId=%d\nName=%s\n \n", id, name)
}

func (um *UserManagement) GetUser(id int) (*User, error) {
	user, exists := um.data[id]
	if !exists {
		return nil, fmt.Errorf("User with Id %d not found\n \n", id)
	}
	return user, nil
}
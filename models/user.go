package models

import (
	"errors"
	"fmt"
)

type User struct {
	ID int
	FirstName string
	LastName string
}

var (
	users []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddUser(user User) (User, error) {
	if user.ID != 0 {
	return User{}, errors.New("new user must not include an ID or must be set to zero")
	}
	user.ID = nextID
	nextID++
	users = append(users, &user)
	return user, nil
}

func GetUserByID(id int)(User, error)  {
	for _, user := range users {
		if user.ID == id {
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("user with the ID '%v' not found", id)
}

func UpdateUser(user User)(User, error)  {
	for _, u := range users {
		if u.ID == user.ID {
			u = &user
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("user with the ID '%v' not found", user.ID)
}

func RemoveUserByID(id int)(error)  {
	for i, user := range users {
		if user.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("user with the ID '%v' not found", id)
}
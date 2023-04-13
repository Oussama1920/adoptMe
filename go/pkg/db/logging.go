package db

import (
	"context"
	"fmt"
)

func (service *Db) AddUser(ctx context.Context, user User) error {

	statement := `INSERT into users(name,firstname,address,dateOfBirth,phoneNumber) VALUES ($1, $2, $3, $4, $5)`

	_, err := service.handler.Exec(ctx, statement, user.Name, user.FirstName, user.Address, user.DateOfBirth, user.PhoneNumber)
	if err != nil {
		return fmt.Errorf("failed to insert new User - error:%#v", err)
	}

	return nil
}

type User struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	FirstName   string `json:"firstname"`
	Address     string `json:"address"`
	PhoneNumber string `json:"phonenumber"`
	DateOfBirth string `json:"dateOfBirth"`
}

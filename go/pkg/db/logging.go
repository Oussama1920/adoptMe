package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v4"
)

func (service *Db) AddUser(ctx context.Context, user User) error {

	statement := `INSERT into users(name,firstname,address,dateOfBirth,phoneNumber,email,password,provider,photo,role,createdAt,updatedAt,id,verificationcode) VALUES ($1, $2, $3, $4, $5,$6,$7,$8,$9,$10,$11,$12,$13,$14)`

	_, err := service.handler.Exec(ctx, statement, user.Name, user.FirstName, user.Address, user.DateOfBirth, user.PhoneNumber, user.Email, user.Password, user.Provider, user.Photo, user.Role, user.CreatedAt, user.UpdatedAt, user.ID, user.VerificationCode)
	if err != nil {
		return fmt.Errorf("failed to insert new User - error:%#v", err)
	}

	return nil
}
func (service *Db) Login(ctx context.Context, email string, password string) (*User, error) {
	var user User
	statement := `SELECT name,firstname from users where email=$1 AND password=$2`

	err := service.handler.QueryRow(ctx, statement, email, password).Scan(&user.Name, &user.FirstName)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new User - error:%#v", err)
	}
	return &user, nil
}
func (service *Db) GetVerificationCode(ctx context.Context, verificationCode string) (*User, error) {
	var user User
	statement := `SELECT id,name,firstname from users where verificationCode=$1`

	err := service.handler.QueryRow(ctx, statement, verificationCode).Scan(&user.ID, &user.Name, &user.FirstName)
	if err != nil {
		return nil, fmt.Errorf("failed to insert new User - error:%#v", err)
	}
	return &user, nil
}
func (service *Db) VerifyUser(ctx context.Context, user User) error {

	statement := `
	UPDATE users
	SET verificationCode = '',
	verified =  true
    WHERE id = $1
	`

	_, err := service.handler.Exec(ctx, statement, user.ID)
	if err != nil {
		return fmt.Errorf("failed to insert new User - error:%#v", err)
	}

	return nil
}
func (service *Db) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User
	statement := `SELECT id,name,firstname,password,verified,email from users where email=$1`

	err := service.handler.QueryRow(ctx, statement, email).Scan(&user.ID, &user.Name, &user.FirstName, &user.Password, &user.Verified, &user.Email)
	if err != nil && err != pgx.ErrNoRows {
		return nil, fmt.Errorf("failed to Get user by password - error:%#v", err)
	}
	return &user, nil
}
func (service *Db) GetUserById(ctx context.Context, id string) (*User, error) {
	var user User
	statement := `SELECT id,name,firstname,password,email,createdAt,address,phoneNumber,updatedAt,dateOfBirth from users where id=$1`

	err := service.handler.QueryRow(ctx, statement, id).Scan(&user.ID, &user.Name, &user.FirstName, &user.Password, &user.Email, &user.CreatedAt, &user.Address, &user.PhoneNumber, &user.UpdatedAt, &user.DateOfBirth)
	if err != nil {
		return nil, fmt.Errorf("failed to Get user by id - error:%#v", err)
	}
	return &user, nil
}
func (service *Db) UpdateUser(ctx context.Context, user User, newData User) error {

	statement := `
	UPDATE users
	SET 
	firstName =  $1,
	phoneNumber =  $2,
	dateOfBirth = $3,
	updatedAt= $4,
	photo = $5,
	address= $6,
	name=$7
    WHERE id = $8
	`
	dateOfBirth, err := time.Parse(time.RFC3339, newData.DateOfBirth)
	if err != nil {
		fmt.Println("Error parsing date:", err)
	}

	// Format the time to extract just the date portion
	dateOnly := dateOfBirth.Format("2006-01-02")

	_, err = service.handler.Exec(ctx, statement, newData.FirstName, newData.PhoneNumber, dateOnly, time.Now(), newData.Photo, newData.Address, newData.Name, user.ID)
	if err != nil {
		return fmt.Errorf("failed to update new User - error:%#v", err)
	}

	return nil
}

type User struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	FirstName        string `json:"firstname"`
	Address          string `json:"address"`
	PhoneNumber      string `json:"phonenumber"`
	DateOfBirth      string `json:"dateOfBirth"`
	Password         string `json:"password"`
	PasswordConfirm  string `json:"passwordConfirm"`
	Email            string `json:"email"`
	Provider         string `json:"provider"`
	Photo            string `json:"photo"`
	Role             string
	VerificationCode string
	Verified         bool
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}
type SignInInput struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}
type UserResponse struct {
	ID          string    `json:"id,omitempty"`
	Name        string    `json:"name,omitempty"`
	FirstName   string    `json:"firstname"`
	Address     string    `json:"address"`
	PhoneNumber string    `json:"phonenumber"`
	DateOfBirth string    `json:"dateOfBirth"`
	Email       string    `json:"email,omitempty"`
	Role        string    `json:"role,omitempty"`
	Photo       string    `json:"photo,omitempty"`
	Provider    string    `json:"provider"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

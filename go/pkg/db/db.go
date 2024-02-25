package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/sirupsen/logrus"
)

type PoolHandler interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Close()
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
}

type Db struct {
	dbc     DBConfig
	handler PoolHandler
	appLog  *logrus.Logger
}

type DbHandler interface {
	// Insert(ctx context.Context, queryString string, arguments ...interface{}) error
	Close(ctx context.Context) error
	IsDBUp(ctx context.Context) bool
	Connect(ctx context.Context) error
	AddUser(ctx context.Context, user User) error
	Login(ctx context.Context, email string, password string) (*User, error)
	GetVerificationCode(ctx context.Context, verificationCode string) (*User, error)
	VerifyUser(ctx context.Context, user User) error
	UpdateUser(ctx context.Context, user User, newData User) error
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	AddPet(ctx context.Context, pet Pet, userID string) (int, error)
	GetPet(ctx context.Context, id int) (*Pet, error)
	GetListPets(ctx context.Context, query string) ([]*Pet, error)
}

// NewDB establish a connection with the db and return a DbHandler
func NewDB(ctx context.Context, dbc DBConfig, appLog *logrus.Logger) (DbHandler, error) {
	cfg, err := dbc.toPgxConfig()
	if err != nil {
		return nil, err
	}

	// Connect to the "parameters" database.
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err != nil {
		appLog.Errorf("error connecting to the database: %v", err)
		return nil, fmt.Errorf("error connecting to the database: %+v", err)
	}

	return &Db{
		dbc:     dbc,
		handler: pool,
		appLog:  appLog,
	}, nil
}

// Insert : insert query entrypoint
func (service *Db) Insert(ctx context.Context, queryString string, arguments ...interface{}) error {
	_, err := service.handler.Exec(ctx, queryString, arguments...)
	if err != nil {
		return err
	}
	return nil
}

// Close : close db connection
func (service *Db) Close(ctx context.Context) error {
	service.handler.Close()
	return nil
}

// IsDBUp returns true if a connection with the DB can be established, false otherwise
func (service *Db) IsDBUp(ctx context.Context) bool {
	if service.handler == nil {
		return false
	}

	// Maybe this isn't the best way to do so, but it is the only one to work 100% of the times
	_, err := service.handler.Exec(ctx, "select pg_is_in_recovery()")

	return err == nil
}

func (service *Db) Connect(ctx context.Context) error {
	cfg, err := service.dbc.toPgxConfig()
	if err != nil {
		return err
	}
	pool, err := pgxpool.ConnectConfig(ctx, cfg)
	if err == nil {
		service.handler = pool
	}
	return err
}

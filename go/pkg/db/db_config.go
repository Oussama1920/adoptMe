package db

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DBConfig Database configuration
type DBConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	DBName       string
	PoolMaxConns int
	SSLMode      string
}

// toPgxConfig : build ConnConfig
func (dbc DBConfig) toPgxConfig() (*pgxpool.Config, error) {

	if err := dbc.check(); err != nil {
		return nil, err
	}

	// We need to temporarly disable the environmental variable PGSERVICE
	// because, if set, pgx will try to read the service from pg_service.conf
	// but pgx doesn't accept PGSYSCONFDIR and will look in $HOME/.pg_service.conf
	os.Setenv("PGSERVICE", "")

	return pgxpool.ParseConfig(dbc.toConfigString())
}

// check : check mandatories parameters
func (dbc DBConfig) check() error {
	var errMsgs []string

	if dbc.Host == "" {
		errMsgs = append(errMsgs, "Host can't be empty")
	}

	if _, err := net.LookupHost(dbc.Host); err != nil {
		errMsgs = append(errMsgs, "Not valid Host: "+dbc.Host)
	}

	if dbc.Port < 0 {
		errMsgs = append(errMsgs, fmt.Sprintf("Not valid port: %d", dbc.Port))
	}

	if dbc.User == "" {
		errMsgs = append(errMsgs, "User can't be empty")
	}

	if dbc.Password == "" {
		errMsgs = append(errMsgs, "Password can't be empty")
	}

	if dbc.DBName == "" {
		errMsgs = append(errMsgs, "DataBase Name can't be empty")
	}

	if dbc.PoolMaxConns < 0 {
		errMsgs = append(errMsgs, fmt.Sprintf("Not valid max pool connections: %d", dbc.PoolMaxConns))
	}

	if dbc.SSLMode == "" {
		errMsgs = append(errMsgs, "SSL Mode can't be empty")
	}

	if len(errMsgs) != 0 {
		return fmt.Errorf("wrong configuration for db connection: %s", strings.Join(errMsgs, ","))
	}

	return nil
}

// toConfigString : build string config from DBConfig
func (dbc DBConfig) toConfigString() string {
	var b strings.Builder
	appendString(&b, "host", dbc.Host)
	appendInt(&b, "port", dbc.Port)
	appendString(&b, "user", dbc.User)
	appendString(&b, "password", dbc.Password)
	appendString(&b, "dbname", dbc.DBName)
	appendInt(&b, "pool_max_conns", dbc.PoolMaxConns)
	appendString(&b, "sslmode", dbc.SSLMode)

	return b.String()
}

// appendString : add string to db connection string
func appendString(b *strings.Builder, label, value string) {
	if value == "" {
		return
	}
	fmt.Fprintf(b, "%s=%s ", label, value)
}

// appendInt : add int to db connection string
func appendInt(b *strings.Builder, label string, value int) {
	fmt.Fprintf(b, "%s=%d ", label, value)
}

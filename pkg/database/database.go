package database

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	Name     string
	Password string
}

type DbManager struct {
	conStr     string
	Connection *pgxpool.Pool
	context    context.Context
}

func NewDbMananger(connectionString string, ctx context.Context) (*DbManager, error) {
	if len(connectionString) == 0 {
		msg := "error creating new DbManager: empty connection string"
		return nil, errors.New(msg)
	}

	dbManager := DbManager{
		conStr:     connectionString,
		Connection: nil,
		context:    ctx,
	}

	err := dbManager.ConnectToDatabase()
	if err != nil {
		return nil, err
	}

	return &dbManager, nil
}

// Set connection string
func (d *DbManager) SetConnectionString(conStr string) {
	d.conStr = conStr
}

// Connect to database using set connection string
func (d *DbManager) ConnectToDatabase() error {
	if d.conStr == "" {
		errMsg := "connection string is not set"
		return errors.New(errMsg)
	}

	var err error
	d.Connection, err = pgxpool.New(context.Background(), d.conStr)

	return err
}

func (d *DbManager) GetUserByUserName(name string) (User, error) {
	user := User{}
	var err error
	if err = d.Connection.Ping(context.Background()); err != nil {
		return User{}, nil
	}

	query := "SELECT user_name, password_hash FROM users WHERE user_name = $1"
	err = d.Connection.QueryRow(d.context, query, name).Scan(&user.Name, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

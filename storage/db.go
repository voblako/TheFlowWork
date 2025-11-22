package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/voblako/TheFlowWork/internal/models"
)

type Config struct {
	DatabaseURL string
	//CacheURL
	//StorageURL
}

type DB struct {
	conn *pgx.Conn
}

func NewPostgressConnect(config Config) (DB, error) {
	conn, err := pgx.Connect(context.Background(), config.DatabaseURL)
	if err != nil {
		return DB{}, err
	}
	return DB{conn: conn}, nil
}

func (db *DB) Close() {
	db.conn.Close(context.Background())
}

func (db *DB) Version() (version string, err error) {
	version = ""
	err = db.conn.QueryRow(context.Background(), "SELECT version()").Scan(&version)
	if err != nil {
		return "", err
	}
	return version, nil
}

func (db *DB) Init() (postgresResp string, err error) {
	res := ""
	err = db.conn.QueryRow(context.Background(), "CREATE SCHEMA IF NOT EXISTS app;").Scan(&res)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("Can`t create service scheme: " + err.Error())
		}

	}

	err = db.conn.QueryRow(context.Background(), "CREATE TABLE IF NOT EXISTS app.users (id serial PRIMARY KEY,name VARCHAR(255) NOT NULL,surname VARCHAR(255) NOT NULL,third_name VARCHAR(255),email VARCHAR(255) NOT NULL UNIQUE,password_hash VARCHAR(255) NOT NULL,birthdate DATE);").Scan(&res)
	if err != nil {
		if !errors.Is(err, pgx.ErrNoRows) {
			return "", errors.New("Can`t create table app.users: " + err.Error())
		}
	}
	return res, nil
}

func (db *DB) CreateUser(user models.User) error {
	_, err := db.conn.Exec(context.Background(), "INSERT INTO app.users (name,surname,third_name,email,password_hash,birthdate) values ($1, $2, $3, $4, $5, $6)", user.Name, user.Surname, user.ThirdName, user.Email, user.PasswordHash, user.Birthdate)
	if err != nil {
		return errors.New("Can`t insert a new user:" + err.Error())
	}
	return nil
}

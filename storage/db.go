package storage

import (
	"context"

	"github.com/jackc/pgx/v5"
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
	if err!=nil{
		return "", err
	}
	return version, nil
}
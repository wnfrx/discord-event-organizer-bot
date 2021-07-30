package config

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func (c *Config) InitPostgres() (err error) {
	connString := `
		host=%s
		port=%s
        user=%s
        password=%s
        dbname=%s
        sslmode=%s
        application_name=%s
	`

	connString = fmt.Sprintf(connString,
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
		os.Getenv("APP_NAME"),
	)

	db, err := sqlx.Connect(os.Getenv("DB_ENGINE"), connString)
	if err != nil {
		log.Printf("[config][services] Failed while init database connection, %+v", err)
		return err
	}

	c.db = db

	return nil
}

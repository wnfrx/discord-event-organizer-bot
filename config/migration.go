package config

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Config) InitMigration() (err error) {
	// if c.db == nil {
	// 	return errors.New("database connection is not initialized yet")
	// }

	// driver, err := postgres.WithInstance(c.db.DB, &postgres.Config{})
	// if err != nil {
	// 	log.Fatalf("Failed to init migration postgres instance: %v", err)
	// 	return err
	// }

	// m, err := migrate.NewWithDatabaseInstance(
	// 	"file://db/migrations",
	// 	"postgres",
	// 	driver,
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to init migration with database instance: %v", err)
	// 	return err
	// }

	connString := "postgres://%s:%s@%s:%s/%s?sslmode=%s"
	connString = fmt.Sprintf(
		connString,
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_SSL_MODE"),
	)

	m, err := migrate.New(
		"file://db/migrations",
		connString,
	)
	if err != nil {
		log.Fatalf("Failed to init migration: %v", err)
	}

	if err = m.Force(1); err != nil {
		log.Fatalf("Failed to force migration version: %v", err)
		return err
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to do migration: %v", err)
		return err
	}

	return nil
}

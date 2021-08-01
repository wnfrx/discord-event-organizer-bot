package config

import (
	"errors"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func (c *Config) InitMigration() (err error) {
	if c.db == nil {
		return errors.New("database connection is not initialized yet")
	}

	driver, err := postgres.WithInstance(c.db.DB, &postgres.Config{})
	if err != nil {
		log.Fatalf("Failed to init migration postgres instance: %v", err)
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		"postgres",
		driver,
	)
	if err != nil {
		log.Fatalf("Failed to init migration with database instance: %v", err)
		return err
	}

	err = m.Run()
	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Failed to do migration: %v", err)
		return err
	}

	return nil
}

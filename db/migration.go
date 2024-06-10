package db

import (
	"database/sql"
	"errors"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func RunMigration(config DataSourceConfig) {
	db, err := sql.Open("postgres", config.GenerateDsn())
	if err != nil {
		panic(err)
	}

	driver, err := postgres.WithInstance(db, &postgres.Config{})

	m, err := migrate.NewWithDatabaseInstance("file://db/migrations", config.DB, driver)

	if err != nil {
		panic(err)
	}

	err2 := m.Up()
	if err2 != nil && !errors.Is(err2, migrate.ErrNoChange) {
		err := m.Down()
		if err != nil {
			panic(err2)
		}
		panic(err2)
	}
}

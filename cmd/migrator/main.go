package main

import (
	"errors"
	"flag"
	"fmt"

	// либа для миграций
	"github.com/golang-migrate/migrate/v4"
	// драйвер для sqllite
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	var storagePath, migrationPath, migrationTable string

	flag.StringVar(&storagePath, "storage-path", "", "path to storage")
	flag.StringVar(&migrationPath, "migration-path", "", "path to migration")
	flag.StringVar(&migrationTable, "migration-table", "migration", "table for migration")
	flag.Parse()

	if storagePath == "" {
		panic("storage path is empty")
	}

	if migrationPath == "" {
		panic("migration path is empty")
	}

	m, err := migrate.New(
		"file://"+migrationPath,
		fmt.Sprintf("sqlite3://%s?x-migrations-table=%s", storagePath, migrationTable),
	)
	if err != nil {
		panic(err)
	}

	if err := m.Up(); err != nil {
		if errors.Is(err, migrate.ErrNoChange) {
			fmt.Println("no migrations to apply")

			return
		}
		panic(err)
	}

	fmt.Println("migrations applied")

}

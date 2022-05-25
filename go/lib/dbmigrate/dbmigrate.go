package dbmigrate

import (
	"database/sql"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// An interface that can be used to allow swapping out different Migrators
type Migrator interface {
	Migrate(db *sql.DB) error
}

// This struct implements the Migrator interface, it makes use of "github.com/golang-migrate/migrate/v4"
// MigrateDir is the path to the directory containing the migration file(s), the directory URL should be
// formatted as follows: 'file://{path}' (see example for further details), the files inside the directories
// have to be of the form, {version/sequence number}_{title}.{[up|down]}.sql, example: 001254_create_user.up.sql
// for further information, see https://github.com/golang-migrate/migrate/blob/master/MIGRATIONS.md
// DbName is the name of the database
type GoMigrate struct {
	MigrateDirURL string
	DbName        string
}

func (gm GoMigrate) Migrate(db *sql.DB) error {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}
	migrator, err := migrate.NewWithDatabaseInstance(gm.MigrateDirURL, gm.DbName, driver)
	if err != nil {
		return err
	}
	err = migrator.Up()
	if err != nil {
		return err
	}
	return nil
}

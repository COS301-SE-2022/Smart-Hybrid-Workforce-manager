package dbmigrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"regexp"

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

// Manual Migration
type ManMigrate struct {
	// A prefix that is appended to each path in MigrateDirs
	PathPrefix string

	// MigrateDirs is an array of the directories containing the mirgration folders
	// the directories will be visited in the same sequence as migrateDirs, and thus
	// the order is important if there are dependencies
	MigrateDirs []string // Required

	// SeqRules contains a set of regex rules, the files in dirs of MigrateDirs will
	// be evaluated in the order of the SeqRules, see example for further details
	// If left out, all files will be evaluated and in any order
	SeqRules []string
}

func (mg ManMigrate) Migrate(db *sql.DB) error {
	paths, regexSeqs, err := mg.setup()
	if err != nil {
		return err
	}
	for _, dir := range paths {
		for _, regex := range regexSeqs {
			fmt.Println(dir)
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, file := range files {
				if file.IsDir() {
					continue // skip directories
				}
				if regex.MatchString(file.Name()) {
					fmt.Println("  ", file.Name())
				}
			}
		}
	}
	return nil
}

func (mg ManMigrate) setup() ([]string, []*regexp.Regexp, error) {
	regexSeqs := make([]*regexp.Regexp, len(mg.SeqRules))
	var err error
	for i, regexS := range mg.SeqRules {
		regexSeqs[i], err = regexp.Compile(regexS)
		if err != nil {
			return nil, nil, err
		}
	}
	paths := make([]string, len(mg.MigrateDirs))
	for i, path := range mg.MigrateDirs {
		paths[i] = mg.PathPrefix + path
	}
	return paths, regexSeqs, nil
}

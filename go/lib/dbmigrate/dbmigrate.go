package dbmigrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

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
	// A prefix that is prepended to each path in MigrateDirs
	PathPrefix string

	// MigrateDirs is an array of the directories containing the mirgration folders
	// the directories will be visited in the same sequence as migrateDirs, and thus
	// the order is important if there are dependencies
	MigrateDirs []string // Required

	// PathPatterns contains a set of shell path patterns the files in dirs of MigrateDirs will
	// be evaluated in the order of the PathPatterns, see example for further details
	// If left out, all files will be processed and in any order, no file will be processed
	// more than once.
	PathPatterns []string
}

func (mg ManMigrate) Migrate(db *sql.DB) error {
	fileProcessedMap := make(map[string]bool) // only used to check if key exists
	patterns, paths, err := mg.setup()
	if err != nil {
		return err
	}
	for _, dir := range paths {
		// fmt.Println(dir)
		for _, pathRule := range patterns {
			files, err := ioutil.ReadDir(dir)
			if err != nil {
				fmt.Println(err)
				return err
			}
			for _, file := range files {
				if file.IsDir() {
					continue // skip directories
				}
				matched, err := path.Match(pathRule, file.Name())
				if err != nil {
					return err
				}
				_, isPresent := fileProcessedMap[file.Name()]
				if matched && !isPresent {
					// fmt.Println("  ", file.Name(), "  ", dir, file.Name())
					fileProcessedMap[file.Name()] = true
					err = migrateFile(filepath.Join(dir, file.Name()), db)
					// err = migrateFile(dir+"/"+file.Name(), db)
					if err != nil {
						return err
					}
				}
			}
		}
	}
	return nil
}

func (mg ManMigrate) setup() (patterns []string, paths []string, err error) {
	// following two lines is to allow for easier change in future if needed
	patterns = make([]string, len(mg.PathPatterns))
	copy(patterns, mg.PathPatterns)
	paths = make([]string, len(mg.MigrateDirs))
	for i, path := range mg.MigrateDirs {
		paths[i] = filepath.Join(mg.PathPrefix, path)
	}
	return patterns, paths, nil
}

// assumes db is already open
func migrateFile(filepath string, db *sql.DB) error {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}
	_, err = db.Exec(string(file))
	return err
}

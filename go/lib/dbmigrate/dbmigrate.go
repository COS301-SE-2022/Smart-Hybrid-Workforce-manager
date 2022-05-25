package dbmigrate

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
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
	// A prefix that is prepended to each path in MigrateDirs
	PathPrefix string

	// MigrateDirs is an array of the directories containing the mirgration folders
	// the directories will be visited in the same sequence as migrateDirs, and thus
	// the order is important if there are dependencies
	MigrateDirs []string // Required

	// PathPatterns contains a set of shell path patterns, the files in dirs of MigrateDirs will
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
				_, isPresent := fileProcessedMap[filepath.Join(dir, file.Name())]
				if matched && !isPresent {
					// fmt.Println("  ", file.Name(), "  ", dir, file.Name())
					fileProcessedMap[filepath.Join(dir, file.Name())] = true
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

// var notGoAhead string = `^pq: (syntax error at or near .*$)`
// if the error matches this pattern, go ahead and try again later
var goAhead string = `^pq: (.* does not exist$)`

type AutoMigrate struct {
	// The root of the directory containing the relevant files,
	// this directory will be traversed recursively
	MigratePath string

	// PathPatterns contains a set of shell path patterns the, files will be processed
	// in the order of the PathPatterns, see example for further details
	// If left out, all files will be processed and in any order, no file will be processed
	// more than once.
	PathPatterns []string
}

func (am AutoMigrate) Migrate(db *sql.DB) error {
	processed := make(map[string]bool)
	allMigrated := false
	var err error
	for _, pattern := range am.PathPatterns {
		allMigrated = false
		for !allMigrated {
			allMigrated, err = recursiveMigrate(am.MigratePath, processed, pattern, db)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func recursiveMigrate(dir string, processed map[string]bool, pattern string, db *sql.DB) (bool, error) {
	// fmt.Println(dir)
	allMigrated := true
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return false, err
	}
	for _, file := range files {
		if file.IsDir() { // if file is a directory, recursively traverse
			dirMigrated, err := recursiveMigrate(filepath.Join(dir, file.Name()), processed, pattern, db)
			if err != nil {
				return false, err
			}
			allMigrated = allMigrated && dirMigrated
		} else {
			// fmt.Println("  ", file.Name())
			_, hasBeenProcessed := processed[filepath.Join(dir, file.Name())] // check if file has been succesfully processed already
			if !hasBeenProcessed {
				// fmt.Println("   ----------------")
				matched, err := path.Match(pattern, file.Name()) // check if it matches the current pattern
				if err != nil {
					return false, err
				}
				if matched {
					// fmt.Println("   *****************")
					fileBytes, err := os.ReadFile(filepath.Join(dir, file.Name()))
					if err != nil {
						return false, err
					}
					_, err = db.Exec(string(fileBytes))
					if err != nil {
						canGoAhead, regexErr := regexp.MatchString(goAhead, err.Error())
						if regexErr != nil {
							return false, regexErr
						}
						if canGoAhead {
							allMigrated = false
						} else {
							return false, err
						}
					} else {
						// fmt.Println("   +++++++++++++++")
						processed[filepath.Join(dir, file.Name())] = true
					}
				}
			}
		}
	}
	return allMigrated, nil // TODO
}

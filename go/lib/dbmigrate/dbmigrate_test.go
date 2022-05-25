package dbmigrate

import (
	. "lib/dockertest_db"
	tu "lib/testutils"
	"testing"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func TestMigrateExample(t *testing.T) {
	config := TestDbConfig{
		Verbose:  true,
		HostPort: "5433",
		HostAdrr: "127.0.0.1",
	}
	testdb, err := StartTestDb(config)
	defer StopTestDbWithTest(testdb, t, false)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
		t.FailNow()
	}
	driver, err := postgres.WithInstance(testdb.Db, &postgres.Config{})
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Error creating driver, err: %v", err))
		t.FailNow()
	}
	migrater, err := migrate.NewWithDatabaseInstance(
		"file://./migrate",
		"arche",
		driver,
	)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Error creating migrate, err: %v", err))
		t.FailNow()
	}
	err = migrater.Up()
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Error performing migration, err: %v", err))
		t.FailNow()
	}

	// for {
	// }
}

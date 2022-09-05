package test_setup

import (
	"api/db"
	"fmt"
	dbm "lib/dbmigrate"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	"testing"
)

func SetupTest(t *testing.T) dtdb.TestDb {
	// SETUP =============
	config := dtdb.TestDbConfig{
		Verbose:  true,
		HostPort: "5433",
		HostAdrr: "127.0.0.1",
	}
	testdb, err := dtdb.StartTestDb(config) // Start DB
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
		t.FailNow()
	}
	// Perform migration
	migrator := dbm.AutoMigrate{
		MigratePath: "../../../db/sql",
		PathPatterns: []string{
			`*.schema.*sql`,   // schema files first
			`*.function.*sql`, // function files second
			// `*mock.sql`,    // will use own mock data
		},
	}
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not performed migration, err: %v", err))
		t.FailNow()
	}
	// update env vars for db connection
	t.Setenv("DATABASE_DSN", testdb.Dsn)

	dbMaxIdleEnv := "5"
	t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", dbMaxIdleEnv)

	dbMaxOpenEnv := "5"
	t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", dbMaxOpenEnv)

	err = db.RegisterAccess()
	if err != nil {
		t.Logf(tu.Scolour(tu.PURPLE, "Error while connecting to DB: %v, skipping test"), err)
		t.FailNow()
	} else {
		fmt.Println(tu.Scolour(tu.GREEN, "DB connected"))
	}

	return testdb
	// ==================
}

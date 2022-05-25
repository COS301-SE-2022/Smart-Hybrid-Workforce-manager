package dbmigrate

import (
	"fmt"
	. "lib/dockertest_db"
	tu "lib/testutils"
	"testing"
)

func TestMigrateExample(t *testing.T) {
	// ====================================================
	// Start DB
	// ====================================================
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
	// ====================================================

	// ====================================================
	// Create and config migrator
	// ====================================================
	migrator := GoMigrate{MigrateDirURL: "file://../../../db/migrate", DbName: "arche"}

	// ====================================================
	// Perform migration
	// ====================================================
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not perform migration, err: ", err))
		t.FailNow()
	}

	// ====================================================
	// DB operations
	// ====================================================
	rows, err := testdb.Db.Query(`SELECT * FROM "user".identifier;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s1, s2, s3, s4, s5, s6, s7 string
	for rows.Next() {
		err := rows.Scan(&s1, &s2, &s3, &s4, &s5, &s6, &s7)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s %s %s %s %s %s", s1, s2, s3, s4, s5, s6, s7))
	}

}

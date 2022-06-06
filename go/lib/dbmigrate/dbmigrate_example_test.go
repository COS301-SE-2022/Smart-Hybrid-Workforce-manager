package dbmigrate

import (
	"fmt"
	. "lib/dockertest_db"
	tu "lib/testutils"
	"reflect"
	"testing"
)

func TestAutoMigrateExample(t *testing.T) {
	fmt.Println("Automatic Migration")
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
	migrator := AutoMigrate{
		// "../../../db/sql",
		"example-data/man",
		[]string{
			`*.schema.*sql`,   // schema files first
			`*.function.*sql`, // function files second
			`*mock.sql`,
			// `*`, 	       // can specify everything else like this
		},
	}
	// ====================================================

	// ====================================================
	// Perform migration
	// ====================================================
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not perform migration, err: %v", err))
		t.FailNow()
	}
	// ====================================================

	// ====================================================
	// DB operations
	// ====================================================
	// +++++++++
	// Check if tables in the pets schema were migrated
	table_rows, err := testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s string
	var tableNames []string
	var expectedTableNames = []string{"dog", "owns", "person"}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		tableNames = append(tableNames, s)
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if !reflect.DeepEqual(tableNames, expectedTableNames) {
		t.Error(tu.Scolourf(tu.RED, "Migration not performed correctly, actual created tables in the pets schema does not match expected"))
	}

	rows, err := testdb.Db.Query(`SELECT * FROM pets.get_ownership();`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s1, s2 string
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s", s1, s2))
	}
}

func TestManMigrateExample(t *testing.T) {
	fmt.Println("Manual Migration")
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
	migrator := ManMigrate{
		// "../../../db/sql",
		// []string{"user", "team", "role", "permission", "resource", "booking", "mock"},
		"example-data/man/",
		[]string{"dog", "person", "owns", "mock"},
		[]string{
			`*.schema.*sql`,   // schema files first
			`*.function.*sql`, // function files second
			`*mock.sql`,
			// `*`, 	       // can specify everything else like this
		},
	}
	// ====================================================

	// ====================================================
	// Perform migration
	// ====================================================
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not perform migration, err: %v", err))
		t.FailNow()
	}
	// ====================================================

	// ====================================================
	// DB operations
	// ====================================================
	// +++++++++
	// Check if tables in the pets schema were migrated
	table_rows, err := testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s string
	var tableNames []string
	var expectedTableNames = []string{"dog", "owns", "person"}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		tableNames = append(tableNames, s)
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if !reflect.DeepEqual(tableNames, expectedTableNames) {
		t.Error(tu.Scolourf(tu.RED, "Migration not performed correctly, actual created tables in the pets schema does not match expected"))
	}

	rows, err := testdb.Db.Query(`SELECT * FROM pets.get_ownership();`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s1, s2 string
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s", s1, s2))
	}
}

func TestMigrateExample(t *testing.T) {
	fmt.Println("Go Migrate")
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
	migrator := GoMigrate{MigrateDirURL: "file://./example-data/go", DbName: "arche"}

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
	// +++++++++
	// Check if tables in the pets schema were migrated
	table_rows, err := testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s string
	var tableNames []string
	var expectedTableNames = []string{"dog", "owns", "person"}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		tableNames = append(tableNames, s)
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if !reflect.DeepEqual(tableNames, expectedTableNames) {
		t.Error(tu.Scolourf(tu.RED, "Migration not performed correctly, actual created tables in the pets schema does not match expected"))
	}

	rows, err := testdb.Db.Query(`SELECT * FROM pets.get_ownership();`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s1, s2 string
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s", s1, s2))
	}
}

func TestClearDB(t *testing.T) {
	fmt.Println("Automatic Migration")
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
	migrator := AutoMigrate{
		// "../../../db/sql",
		"example-data/man",
		[]string{
			`*.schema.*sql`,   // schema files first
			`*.function.*sql`, // function files second
			`*mock.sql`,
			// `*`, 	       // can specify everything else like this
		},
	}
	// ====================================================

	// ====================================================
	// Perform migration
	// ====================================================
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not perform migration, err: %v", err))
		t.FailNow()
	}
	// ====================================================

	// ====================================================
	// DB operations
	// ====================================================
	// +++++++++
	// Check if tables in the pets schema were migrated
	table_rows, err := testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s string
	var tableNames []string
	var expectedTableNames = []string{"dog", "owns", "person"}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		tableNames = append(tableNames, s)
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if !reflect.DeepEqual(tableNames, expectedTableNames) {
		t.Error(tu.Scolourf(tu.RED, "Migration not performed correctly, actual created tables in the pets schema does not match expected"))
	}

	rows, err := testdb.Db.Query(`SELECT * FROM pets.get_ownership();`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	var s1, s2 string
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s", s1, s2))
	}

	// ====================================================
	// Clear DB
	// ===================================================
	fmt.Println(tu.Scolour(tu.BLUE, "Clearing DB"))
	err = ClearDB(testdb.Db)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "Could not clear db, err: %v", err))
	}
	// Print cleared db
	table_rows, err = testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	table_count := 0
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if table_count > 0 {
		t.Errorf("DB not cleared correctly, expected to get 0 tables, got %d", table_count)
		t.FailNow()
	}
	// ====================================================

	// ====================================================
	// Perform migration
	// ====================================================
	err = migrator.Migrate(testdb.Db)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not perform migration, err: %v", err))
		t.FailNow()
	}
	// ====================================================

	// ====================================================
	// DB operations
	// ====================================================
	// +++++++++
	// Check if tables in the pets schema were migrated
	tableNames = make([]string, 0)
	table_rows, err = testdb.Db.Query(
		`SELECT table_name FROM information_schema.tables WHERE table_schema = 'pets' ORDER BY table_name;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======= Tables ======="))
	for table_rows.Next() {
		err := table_rows.Scan(&s)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		tableNames = append(tableNames, s)
		fmt.Println(tu.Scolourf(tu.GREEN, "%s", s))
	}
	fmt.Println(tu.Scolour(tu.GREEN, "======================="))

	if !reflect.DeepEqual(tableNames, expectedTableNames) {
		t.Error(tu.Scolourf(tu.RED, "Migration not performed correctly, actual created tables in the pets schema does not match expected"))
	}

	rows, err = testdb.Db.Query(`SELECT * FROM pets.get_ownership();`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
		t.FailNow()
	}
	for rows.Next() {
		err := rows.Scan(&s1, &s2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "%s %s", s1, s2))
	}
}

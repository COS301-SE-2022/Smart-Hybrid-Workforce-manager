package testdb

import (
	"fmt"
	tu "lib/testutils"
	"testing"
)

// NB ==================================================================================
// 0. If the test fails in a way such that that StopDbTest is not called for the running
//    containers, the containers will have to be stopped manually. An example of this is
//    if execution is interrupted with with an interrupt such as ^C. I will look into
//    handling such interrupts at a later time to ensure that the resources are closed
// =====================================================================================

func TestExample(t *testing.T) {
	// ===================================================================================
	// 1. Create a configuration object, look at src code to see fields and other details
	// ===================================================================================
	config1 := TestDbConfig{
		Verbose:  true,
		HostPort: "5433",
		HostAdrr: "127.0.0.1",
	}

	// ================================================================================
	// 2. Start the container and mock db, this requires the config struct, if pool
	//    was included in config, that pool will be reused, a TestDb struct is returned
	//    see src for further details
	// ================================================================================
	testdb1, err := StartTestDb(config1)
	// ================================================================================
	// 3. Defer the stopping of the container
	// ================================================================================
	defer StopTestDbWithTest(testdb1, t, false)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
		t.FailNow()
	}
	fruitsExample(testdb1, t)

	// ================================================================================
	// 4. An example of a pool being reused
	// ================================================================================
	config2 := TestDbConfig{
		Verbose:  true,
		HostPort: "5434",
		HostAdrr: "127.0.0.1",
		Pool:     testdb1.pool, // using same pool
	}

	testdb2, err := StartTestDb(config2)
	defer StopTestDbWithTest(testdb2, t, false)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not start testDb, err: %v", err))
		t.FailNow()
	}
	dogsExample(testdb2, t)

	// Note ===========================================================================
	// 5. Containers can be stopped as follows if deferrig is not an option
	// ================================================================================
	// err = StopTestDb(testdb2)
	// if err != nil {
	// 	t.Error(tu.Scolourf(tu.RED, "Could not stop testDb, err: %v", err))
	// }
	// err = StopTestDb(testdb1)
	// if err != nil {
	// 	t.Error(tu.Scolourf(tu.RED, "Could not stop testDb, err: %v", err))
	// }
}

func runExample(create string, insert string, selectS string, testdb TestDb, t *testing.T) {
	query := create
	_, err := testdb.Db.Query(query)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
	}

	query = insert
	_, err = testdb.Db.Query(query)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
		t.FailNow()
	}

	query = selectS
	rows, err := testdb.Db.Query(query)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
		t.FailNow()
	}

	var id int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "ID: %d, Name: %s", id, name))
	}
}

func fruitsExample(testdb TestDb, t *testing.T) {
	runExample(
		`CREATE TABLE fruits(
			id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL
	 	);`,
		`INSERT INTO fruits(name)
		VALUES
			('Apple'),
			('Orange'),
			('Pear');`,
		`SELECT * FROM fruits;`,
		testdb,
		t,
	)
}

func dogsExample(testdb TestDb, t *testing.T) {
	runExample(
		`CREATE TABLE dogs(
			id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL
	 	);`,
		`INSERT INTO dogs(name)
		VALUES
			('Jenny'),
			('Lucky');`,
		`SELECT * FROM dogs;`,
		testdb,
		t,
	)
}

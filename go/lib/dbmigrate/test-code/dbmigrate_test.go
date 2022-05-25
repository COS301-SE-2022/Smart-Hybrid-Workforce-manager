package testdbmigrate

import (
	"fmt"
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
		"file://./psql",
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
	_, err = testdb.Db.Query(`
		INSERT INTO "users" (user_id, name, email) 
		VALUES (112, 'Ben', 'Smith');`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not insert, err: %v", err))
	}
	rows, err := testdb.Db.Query(`SELECT * FROM "users";`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
	}
	var id int
	var name, email string
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "ID: %d, Name: %s, Email: %s", id, name, email))
	}

	rows, err = testdb.Db.Query(`SELECT * FROM team;`)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "could not read rows, err: %v", err))
	}
	var id2 int
	for rows.Next() {
		err := rows.Scan(&id, &name, &id2)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
			t.FailNow()
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "ID: %d, Name: %s, team_id: %v", id, name, id2))
	}
}

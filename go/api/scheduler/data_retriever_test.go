package scheduler

import (
	"api/db"
	"encoding/json"
	"fmt"
	dbm "lib/dbmigrate"
	dtdb "lib/dockertest_db"
	tu "lib/testutils"
	"testing"
	"time"
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
			`*mock.sql`,
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

func TestGetSchedulerData(t *testing.T) {
	// t.Setenv("DATABASE_DSN", "host=127.0.0.1 port=5432 user=admin dbname=arche sslmode=disable")
	// t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	// t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "25")
	// t.Setenv("DATABASE_URL", "postgresql://admin:admin@localhost:5432/db?schema=public")
	testdb := SetupTest(t)
	defer dtdb.StopTestDbWithTest(testdb, t, false)
	// err := db.RegisterAccess()
	// if err != nil {
	// 	t.Fatal("Could not access db")
	// }
	// err := db.RegisterAccess()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// _, err = db.Open()
	// if err != nil {
	// 	t.Fatal(err)
	// }
	type args struct {
		from time.Time
		to   time.Time
	}
	tests := []struct {
		name    string
		args    args
		want    *SchedulerData
		wantErr bool
	}{
		{
			name: "Get data successfully",
			args: args{
				from: time.Date(1999, 1, 1, 1, 1, 1, 1, time.UTC),
				to:   time.Date(2023, 1, 1, 1, 1, 1, 1, time.UTC),
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSchedulerData(tt.args.from, tt.args.to)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSchedulerData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			b, err := json.MarshalIndent(got, "", "  ")
			if err != nil {
				t.Errorf("Could not unmarshall json")
			}
			fmt.Println(string(b))
		})
	}
}

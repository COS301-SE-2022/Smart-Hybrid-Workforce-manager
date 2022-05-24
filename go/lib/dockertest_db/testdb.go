package testdb

import (
	"database/sql"
	"fmt"
	tu "lib/testutils"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"

	_ "github.com/lib/pq"
)

type TestDbConfig struct {
	Verbose  bool   // indicates whether or not connection and other information should be output
	HostPort string // the host port to bind to
	HostAdrr string
	Pool     *dockertest.Pool // Optional
}

type TestDb struct {
	Db       *sql.DB              // db connection
	resource *dockertest.Resource // not export since it should be treated as final
	pool     *dockertest.Pool     // not export since it should be treated as final
	Dsn      string
}

// Common db config
var (
	user     = "admin"
	password = "admin"
	dbName   = "arche"
	dialect  = "postgres"
	dsn      = "host=%s port=%s user=%s dbname=%s sslmode=disable"
	idleConn = 5
	maxConn  = 5
	tag      = "14" // image tag
)

func StartTestDb(config TestDbConfig) (TestDb, error) {
	var db *sql.DB
	var pool *dockertest.Pool
	var err error = nil
	if config.Pool == nil { // Create a new pool if pool was not passed in
		pool, err = dockertest.NewPool("")
	} else {
		pool = config.Pool
	}
	if err != nil {
		if config.Verbose {
			fmt.Println(tu.Scolourf(tu.RED, "Could not connect to docker: %v", err))
		}
		return TestDb{}, err
	}

	opts := dockertest.RunOptions{
		Repository: "postgres",
		Tag:        tag,
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"POSTGRES_DB=" + dbName,
			"POSTGRES_HOST_AUTH_METHOD=trust",
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				{HostIP: config.HostAdrr, HostPort: config.HostPort},
			},
		},
	}

	resource, err := pool.RunWithOptions(
		// resource, err := pool.RunWithOptions(
		&opts,
		func(config *docker.HostConfig) {
			// set AutoRemove to true so that stopped container goes away by itself
			config.AutoRemove = true
			config.RestartPolicy = docker.RestartPolicy{
				Name: "no",
			}
		},
	)

	if err != nil {
		if config.Verbose {
			fmt.Println(tu.Scolourf(tu.RED, "Could not start resource: %v", err))
		}
		return TestDb{}, err
	}

	db_dsn := fmt.Sprintf(dsn, config.HostAdrr, config.HostPort, user, dbName)
	// Try connecting to db with exponential backoff
	if err = pool.Retry(func() error {
		if config.Verbose {
			fmt.Print(tu.Scolour(tu.BLUE, "RETRY "))
		}
		db, err = sql.Open(dialect, db_dsn)
		if err != nil {
			return err
		}
		err = db.Ping()
		if err != nil {
			return err
		}
		db.SetMaxIdleConns(idleConn)
		db.SetMaxOpenConns(maxConn)
		return nil
	}); err != nil {
		if config.Verbose {
			fmt.Println(tu.Scolourf(tu.RED, "Could not connect to Docker db, err: %v", err))
		}
		return TestDb{}, err
	}
	if config.Verbose {
		fmt.Println()
	}
	return TestDb{Db: db, resource: resource, Dsn: db_dsn, pool: pool}, nil
}

func StopTestDb(testDb TestDb) error {
	if testDb.Db != nil {
		testDb.Db.Close()
	}
	if err := testDb.pool.Purge(testDb.resource); err != nil {
		return err
	}
	return nil
}

// Mainly used to defer the closing and stopping of the container
func StopTestDbWithTest(testdb TestDb, t *testing.T, failNow bool) {
	err := StopTestDb(testdb)
	if err != nil {
		t.Errorf(tu.Scolourf(tu.RED, "Error stopping container or closing db, err: %v", err))
		if failNow {
			t.FailNow()
		}
	}
}

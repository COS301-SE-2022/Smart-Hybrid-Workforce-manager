package db

import (
	"database/sql"
	"fmt"
	tu "lib/testutils"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	user     = "postgres"
	password = "secret" // check if we use a password...
	dbS      = "postgres"
	port     = "5433"
	dialect  = "postgres"
	dsn      = "postgres://%s:%s@localhost:%s/%s?sslmode=disable"
	idleConn = 5
	maxConn  = 5
)

var (
	db *sql.DB
)

func TestDummy(t *testing.T) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		t.Errorf("Could not connect to docker: %s", err)
		t.FailNow()
	}
	// pool.Run("postgres", "14", []string{"POSTGRES_PASSWORD=secret"})
	// if err != nil {
	// 	t.Errorf("UGh Could not start resource: %s", err)
	// 	t.FailNow()
	// }

	opts := dockertest.RunOptions{
		Repository: "postgres", // image
		Tag:        "14",
		Env: []string{
			"POSTGRES_USER=" + user,
			"POSTGRES_PASSWORD=" + password,
			"listen_addresses = '*'",
			"POSTGRES_DB=" + dbS,
		},
		ExposedPorts: []string{"5432"},
		PortBindings: map[docker.Port][]docker.PortBinding{
			"5432": {
				// {HostIP: "0.0.0.0", HostPort: port},
				{HostIP: "127.0.0.1", HostPort: port},
			},
		},
	}

	resource, err := pool.RunWithOptions(
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
		t.Errorf("Could not start resource: %s", err.Error())
		t.FailNow()
	}
	fmt.Println(tu.Scolourf(tu.GREEN, "PORT %v", resource.GetPort("5432/tcp")))
	dsn = fmt.Sprintf(dsn, user, password, port, db)
	// Try connecting to db with exponential backoff
	if err = pool.Retry(func() error {
		fmt.Print(tu.Scolour(tu.BLUE, "IN RETRY"))
		db, err := sql.Open(dialect, "host=127.0.0.1 port=5433 user=postgres dbname=postgres password=secret sslmode=disable")
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
		t.Errorf(tu.Scolour(tu.RED, "Could not connect to Docker db, err: %v"), err)
	}

	defer func() {
		if db != nil {
			db.Close()
		}
	}()

	if err := pool.Purge(resource); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}

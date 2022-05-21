package testdb

import (
	"database/sql"
	"fmt"
	tu "lib/testutils"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type TestDbConfig struct {
	Verbose  bool   // indicates whether or not connection and other information should be output
	HostPort string // the host port to bind
	HostAdrr string
}

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

var db *sql.DB

func StartTestDb(config TestDbConfig) (*sql.DB, error) {
	pool, err := dockertest.NewPool("")
	if err != nil {
		if config.Verbose {
			fmt.Println(tu.Scolourf(tu.RED, "Could not connect to docker: %v", err))
		}
		return nil, err
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

	_, err = pool.RunWithOptions(
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
		return nil, err
	}

	dsn = fmt.Sprintf(dsn, config.HostAdrr, config.HostPort, user, dbName)
	// Try connecting to db with exponential backoff
	if err = pool.Retry(func() error {
		fmt.Print(tu.Scolour(tu.BLUE, "RETRY "))
		db, err = sql.Open(dialect, dsn)
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
		return nil, err
	}
	if config.Verbose {
		fmt.Println()
	}
	return db, nil
}

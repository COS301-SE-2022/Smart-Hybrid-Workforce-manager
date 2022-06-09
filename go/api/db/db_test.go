package db

/*import (
	"database/sql"
	"fmt"
	tu "lib/testutils"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

var (
	user     = "admin"
	password = "admin"
	dbName   = "postgres"
	port     = "5433" // host port
	dialect  = "postgres"
	// dsn      = "host=%s port=%s user=%s dbname=%s password=%s sslmode=disable"
	dsn      = "host=%s port=%s user=%s dbname=%s sslmode=disable"
	host     = "127.0.0.1"
	idleConn = 5
	maxConn  = 5
	tag      = "14" // image tag
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

	opts := dockertest.RunOptions{
		Repository: "postgres", // image
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
				{HostIP: host, HostPort: port},
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
	// dsn = fmt.Sprintf(dsn, host, port, user, dbName, password)
	dsn = fmt.Sprintf(dsn, host, port, user, dbName)
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
		t.Errorf(tu.Scolour(tu.RED, "Could not connect to Docker db, err: %v"), err)
	}
	fmt.Println()

	// DB operations below
	createQuery :=
		`CREATE TABLE fruits(
			id SERIAL PRIMARY KEY,
			name VARCHAR NOT NULL
		 );`
	_, err = db.Query(createQuery)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
	}

	insertQuery :=
		`INSERT INTO fruits(name)
		VALUES
			('Apple'),
			('Orange'),
			('Pear');`
	_, err = db.Query(insertQuery)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
	}

	selectQuery :=
		`SELECT * FROM fruits;`
	rows, err := db.Query(selectQuery)
	if err != nil {
		t.Error(tu.Scolourf(tu.RED, "Could not execute query, err: %v", err))
	}

	var id int
	var name string

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Error(tu.Scolourf(tu.RED, "Could not scan rows, err: %v", err))
		}
		fmt.Println(tu.Scolourf(tu.PURPLE, "ID: %d, Name: %s", id, name))
	}

	//defer func() {
	// if db != nil {

	// }
	//}()
	if db != nil {
		db.Close()
	}
	if err := pool.Purge(resource); err != nil {
		t.Errorf("Could not purge resource: %s", err)
	}
}*/

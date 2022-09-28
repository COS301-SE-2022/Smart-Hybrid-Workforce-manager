package test_setup

import (
	"api/db"
	"testing"
)

func ConnectDB(t *testing.T) error {
	t.Setenv("DATABASE_DSN", "host=127.0.0.1 port=5432 user=admin dbname=arche sslmode=disable")
	t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "25")
	t.Setenv("DATABASE_URL", "postgresql://admin:admin@localhost:5432/db?schema=public")
	// testdb := SetupTest(t)
	// defer dtdb.StopTestDbWithTest(testdb, t, false)
	err := db.RegisterAccess()
	if err != nil {
		t.Fatal(err)
		return err
	}
	_, err = db.Open()
	if err != nil {
		t.Fatal(err)
		return err
	}
	return nil
}

func DisconnectDB(t *testing.T) error {
	err := db.UnregisterAccess()
	if err != nil {
		t.Log(err)
		return err
	}
	return nil
}

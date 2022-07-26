//go:build integration
// +build integration

package scheduler

import (
	"api/db"
	"testing"
)

func Test_Integration_call(t *testing.T) {
	t.Setenv("DATABASE_DSN", "host=127.0.0.1 port=5432 user=admin dbname=arche sslmode=disable")
	t.Setenv("DATABASE_MAX_IDLE_CONNECTIONS", "5")
	t.Setenv("DATABASE_MAX_OPEN_CONNECTIONS", "25")
	t.Setenv("DATABASE_URL", "postgresql://admin:admin@localhost:5432/db?schema=public")
	err := db.RegisterAccess()
	if err != nil {
		t.Fatal(err)
	}
}
